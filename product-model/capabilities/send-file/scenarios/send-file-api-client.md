---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client sends one file
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product begins the upload
  kind: product
  entities:
  - entity: upload
    effect: creates
    facts:
    - Name
    - Error
    to: Sending
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product creates the file message after storing its bytes
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
  - entity: upload
    effect: changes
    facts:
    - Error
    from: Sending
    to: Sent
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

# Send a file through the HTTP API

## Trigger

A client sends one file in a multipart request.

## Outcome

Once its bytes are stored, the file message is added to the timeline under the requesting device and returned with its name, type and size.
