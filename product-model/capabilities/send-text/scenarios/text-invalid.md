---
kind: validation
routes:
  api: HTTP API
steps:
- text: The API client submits blank or oversized text
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product returns an input error
  kind: product
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#createText
- kind: code
  role: implementation
  target: internal/ferry/store.go#createText
---

# Reject blank or oversized text

## Trigger

The text is only whitespace or longer than 65,536 UTF-8 bytes.

## Outcome

No text message is created.
