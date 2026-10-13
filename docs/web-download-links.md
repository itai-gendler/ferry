# Web download links

> English | [简体中文](web-download-links.zh-Hans.md)

## Problem

- **Observed (Max, 2026-10-11)**: on an Android phone, Ferry's Web app opened inside an app's built-in browser listed the files, but tapping one only showed the system message "没有应用可执行此操作。" ("No app can perform this action."); nothing was saved.
- **Observed**: the Web app downloaded a file by fetching it with the Bearer header, wrapping the bytes in a `blob:` URL and clicking an `<a download>` for it. Built-in Android WebViews commonly cannot save a `blob:` URL themselves and pass it to the system, where no app handles that scheme. The same flow also held the whole file, up to 64 MiB, in page memory before saving it.
- **Inferred**: the in-app browser in the report is a WebView of this kind; the failure was not reproduced on that phone, only read from the message and the code path.

## Mechanism

**Recommended**, not yet a decision by Max.

1. The Web app asks `POST /api/v1/files/{message_id}/link` with its Bearer token.
2. The Server returns a same-origin path `/api/v1/downloads/{ticket}`: 32 random bytes, kept in memory as a SHA-256 hash together with the message ID, the issuing device and an expiry five minutes ahead.
3. The Web app navigates to that path with a plain `<a download>`. The browser's own downloader saves the file, or a WebView hands the ordinary `http://` URL to the system browser, which can still open it.
4. `GET /api/v1/downloads/{ticket}` needs no Bearer header. A malformed, unknown or expired ticket, or one whose issuing device has been revoked, gets the same `404 not_found`.

`GET /api/v1/files/{message_id}` with Bearer is unchanged; native apps and the Web image previews still use it.

## Trade-off

The URL itself is the credential: for five minutes anyone who has it can download that one file from inside the network the Server is reachable from. It does not grant other files, the timeline or any write. It never touches disk, ends on Server restart, ends when the issuing device is revoked, and is not sent onward as a Referer because Ferry already sends `Referrer-Policy: no-referrer`. It can show up in that browser's download history.

## Rejected alternatives

- **Cookie session for downloads**: rejected earlier for all Ferry credentials, because cookies are not isolated by port and other services on the same host would receive them (see `docs/lan-auth.md`).
- **A Service Worker that adds the Bearer header to a download navigation**: Service Workers exist only in secure contexts, and Ferry serves plain HTTP on the LAN.
- **A single-use ticket**: a WebView may request the URL and then hand it to the system browser, which requests it again; a ticket spent on the first request breaks exactly the case this change is for.

## Acceptance

- Go test `TestDownloadLinkGrantsOneFileWithoutBearerUntilExpiryOrRevocation`: issuing needs Bearer; a text message gets 404; the link downloads the bytes as an attachment without Bearer and is reusable within its lifetime; it is 404 at expiry, after the issuing device is revoked, and for malformed or unknown tickets.
- Desktop Chrome (2026-10-13), image built from `8371f22` on loopback: after a 3 MiB random file was sent, clicking its card made one `POST …/link` and saved the file through the browser's downloader with the original SHA-256.
- Not verified: a real download from an Android in-app browser and from the phone's system browser. Max waived that check for the 1.0.2 release (2026-10-13).
