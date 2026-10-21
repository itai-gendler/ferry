---
kind: alternate
routes:
  api: HTTP API
steps:
- text: The API client sends an empty file
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product creates the file message with a zero size
  kind: product
  entities:
  - entity: file-message
    effect: creates
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
    - Name
    - Media type
    - Size
    - Content
    - Download address
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#createFile
- kind: code
  role: implementation
  target: internal/ferry/store.go#createFile
---

# Send a zero-byte file

## Trigger

A client sends a file with a valid name and no bytes.

## Outcome

The zero-byte file is valid and downloadable.
