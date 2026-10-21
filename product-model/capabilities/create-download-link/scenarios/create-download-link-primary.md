---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client requests a saving link
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product creates the download link
  kind: product
  entities:
  - entity: download-link
    effect: creates
    facts:
    - Address
    - Expires at
    - Usable
    to: Available
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#createDownloadLink
---

# Create a download link

## Trigger

A joined client asks for a download link for one file message.

## Outcome

The link's address and its expiry, five minutes later, are returned. A missing file, or a text message, is reported as not found.
