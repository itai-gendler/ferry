---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client requests the next messages
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product returns the next ordered messages
  kind: product
  entities:
  - entity: text-message
    effect: reads
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
    - Text
  - entity: file-message
    effect: reads
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
    - Name
    - Media type
    - Size
    - Download address
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListMessages
---

# Read the timeline through the HTTP API

## Trigger

A client requests the messages after the last sequence it has.

## Outcome

Text and file messages are returned in increasing Server sequence, with the cursor to continue from.

## Edge cases

- A page holds 100 messages unless the client asks for 1–200.
- Each message says whether it was sent by the requesting device.
