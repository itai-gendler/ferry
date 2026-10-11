package ferry

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestDownloadLinkGrantsOneFileWithoutBearerUntilExpiryOrRevocation(t *testing.T) {
	store := openTestStore(t)
	now := time.Date(2026, 10, 11, 8, 0, 0, 0, time.UTC)
	handler := newHandler(store, HandlerOptions{Logger: log.New(io.Discard, "", 0)}, func() time.Time { return now })
	join := func(name string) claimResponse {
		t.Helper()
		response := requestJSON(t, handler, http.MethodPost, "/api/v1/access/join", `{"device_name":"`+name+`","password":""}`)
		assertStatus(t, response, http.StatusCreated)
		return decodeClaim(t, response)
	}
	owner := join("Owner")
	phone := join("Phone")
	file, err := store.CreateFileForDevice(t.Context(), owner.Device, "notes.txt", "text/plain", strings.NewReader("ferry file\n"))
	if err != nil {
		t.Fatal(err)
	}
	text, err := store.CreateTextForDevice(t.Context(), owner.Device, "hello")
	if err != nil {
		t.Fatal(err)
	}
	type linkResponse struct {
		URL       string    `json:"url"`
		ExpiresAt time.Time `json:"expires_at"`
	}
	linkPath := "/api/v1/files/" + file.ID + "/link"
	createLink := func() linkResponse {
		t.Helper()
		response := authenticatedRequest(t, handler, http.MethodPost, linkPath, "", phone.Token)
		assertStatus(t, response, http.StatusCreated)
		if response.Header().Get("Cache-Control") != "no-store" {
			t.Fatal("link response was cacheable")
		}
		assertJSONKeys(t, response.Body.Bytes(), "download link", []string{"expires_at", "url"})
		var link linkResponse
		decodeResponse(t, response, &link)
		if !strings.HasPrefix(link.URL, "/api/v1/downloads/") || !link.ExpiresAt.Equal(now.Add(downloadLinkLifetime)) {
			t.Fatalf("link = %#v", link)
		}
		return link
	}

	assertStatus(t, authenticatedRequest(t, handler, http.MethodPost, linkPath, "", ""), http.StatusUnauthorized)
	textLink := authenticatedRequest(t, handler, http.MethodPost, "/api/v1/files/"+text.ID+"/link", "", phone.Token)
	if textLink.Code != http.StatusNotFound || errorCode(t, textLink) != "not_found" {
		t.Fatalf("text message link status = %d, body = %s", textLink.Code, textLink.Body.String())
	}

	link := createLink()
	download := authenticatedRequest(t, handler, http.MethodGet, link.URL, "", "")
	assertStatus(t, download, http.StatusOK)
	if download.Body.String() != "ferry file\n" {
		t.Fatalf("download body = %q", download.Body.String())
	}
	if disposition := download.Header().Get("Content-Disposition"); !strings.HasPrefix(disposition, "attachment") || !strings.Contains(disposition, "notes.txt") {
		t.Fatalf("Content-Disposition = %q", disposition)
	}
	if download.Header().Get("Cache-Control") != "no-store" {
		t.Fatal("download response was cacheable")
	}

	now = now.Add(downloadLinkLifetime - time.Second)
	assertStatus(t, authenticatedRequest(t, handler, http.MethodGet, link.URL, "", ""), http.StatusOK)
	now = now.Add(time.Second)
	expired := authenticatedRequest(t, handler, http.MethodGet, link.URL, "", "")
	if expired.Code != http.StatusNotFound || errorCode(t, expired) != "not_found" {
		t.Fatalf("expired link status = %d, body = %s", expired.Code, expired.Body.String())
	}

	revocable := createLink()
	assertStatus(t, authenticatedRequest(t, handler, http.MethodDelete, "/api/v1/devices/"+phone.Device.ID, "", owner.Token), http.StatusNoContent)
	assertStatus(t, authenticatedRequest(t, handler, http.MethodGet, revocable.URL, "", ""), http.StatusNotFound)

	for _, path := range []string{"/api/v1/downloads/not-a-ticket", "/api/v1/downloads/" + strings.Repeat("A", 43)} {
		assertStatus(t, authenticatedRequest(t, handler, http.MethodGet, path, "", ""), http.StatusNotFound)
	}
}
