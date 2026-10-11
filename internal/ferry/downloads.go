package ferry

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// downloadLinkLifetime bounds how long a download URL works. A link is reused
// within it rather than spent on first use: an in-app WebView may probe the
// URL and then hand it to the system browser, which requests it again.
const downloadLinkLifetime = 5 * time.Minute

// downloadLinks holds short-lived URLs that each grant one file without the
// Bearer header, for browsers that can only save a file by navigating to it.
// They live in memory only: a restart ends them, and nothing reaches disk.
type downloadLinks struct {
	mu      sync.Mutex
	entries map[[32]byte]downloadLink
}

type downloadLink struct {
	messageID string
	deviceID  string
	expiresAt time.Time
}

func (l *downloadLinks) add(hash [32]byte, link downloadLink, now time.Time) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.pruneLocked(now)
	l.entries[hash] = link
}

func (l *downloadLinks) lookup(hash [32]byte, now time.Time) (downloadLink, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.pruneLocked(now)
	link, ok := l.entries[hash]
	return link, ok
}

func (l *downloadLinks) remove(hash [32]byte) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.entries, hash)
}

func (l *downloadLinks) pruneLocked(now time.Time) {
	for hash, link := range l.entries {
		if !now.Before(link.expiresAt) {
			delete(l.entries, hash)
		}
	}
}

func (s *server) createDownloadLink(w http.ResponseWriter, r *http.Request) {
	message, file, err := s.store.OpenFile(r.Context(), r.PathValue("message_id"))
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "not_found", "file message was not found")
		return
	}
	if err != nil {
		s.internalError(w, err)
		return
	}
	file.Close()
	// A ticket has the same canonical 256-bit encoding as a device token, and
	// like a device token only its hash is kept.
	ticket, hash, err := newDeviceToken()
	if err != nil {
		s.internalError(w, fmt.Errorf("generate download link: %w", err))
		return
	}
	now := s.now()
	expiresAt := now.Add(downloadLinkLifetime)
	s.links.add(hash, downloadLink{messageID: message.ID, deviceID: currentDevice(r).ID, expiresAt: expiresAt}, now)
	w.Header().Set("Cache-Control", "no-store")
	writeJSON(w, http.StatusCreated, map[string]string{
		"url":        "/api/v1/downloads/" + ticket,
		"expires_at": expiresAt.UTC().Format(time.RFC3339),
	})
}

// downloadByLink sits outside the Bearer middleware. A malformed, unknown or
// expired ticket and one whose issuing device has been revoked all answer the
// same 404, so the response says nothing about which tickets exist.
func (s *server) downloadByLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	hash, ok := hashDeviceToken(r.PathValue("ticket"))
	if !ok {
		s.downloadLinkNotFound(w)
		return
	}
	link, ok := s.links.lookup(hash, s.now())
	if !ok {
		s.downloadLinkNotFound(w)
		return
	}
	exists, err := s.store.DeviceExists(r.Context(), link.deviceID)
	if err != nil {
		s.internalError(w, err)
		return
	}
	if !exists {
		s.links.remove(hash)
		s.downloadLinkNotFound(w)
		return
	}
	s.serveFile(w, r, link.messageID)
}

func (s *server) downloadLinkNotFound(w http.ResponseWriter) {
	writeError(w, http.StatusNotFound, "not_found", "download link is invalid or has expired")
}
