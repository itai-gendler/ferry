---
kind: validation
routes:
  api: HTTP API
steps:
- text: The API client sends a file larger than 64 MiB
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product rejects the file
  kind: product
  entities: []
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

# Reject a file beyond the size limit

## Trigger

The file turns out to be larger than 67,108,864 bytes while it is received.

## Outcome

No file message is published and the partial stored bytes are removed.
