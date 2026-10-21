---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client requests a file
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product serves original bytes and the stored filename
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
    - Media type
    - Size
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#serveFile
- kind: code
  role: implementation
  target: internal/ferry/store.go#OpenFile
---

# Download a file

## Trigger

A joined client requests a file message's file.

## Outcome

The original bytes are returned as a download under the stored file name, with its media type and size.
