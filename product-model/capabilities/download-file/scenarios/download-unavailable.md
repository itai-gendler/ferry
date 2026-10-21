---
kind: edge
routes:
  api: HTTP API
steps:
- text: The API client requests a missing file or a text entry as a file
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product reports not found
  kind: product
  entities: []
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

# Report an unavailable file

## Trigger

The requested message does not exist or is a text message.

## Outcome

Not found is reported and no file is returned.
