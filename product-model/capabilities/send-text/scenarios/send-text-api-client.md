---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client sends text
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product creates a text message
  kind: product
  entities:
  - entity: text-message
    effect: creates
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
    - Text
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

# Send text through the HTTP API

## Trigger

A client sends a piece of text.

## Outcome

The text message is added to the timeline under the requesting device and returned.

## Edge cases

- Whitespace around nonblank text is kept as sent.
