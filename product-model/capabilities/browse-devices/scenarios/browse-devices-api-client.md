---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client requests the joined devices
  kind: actor
  entities:
  - entity: device
    effect: reads
    facts: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product returns the connected identities
  kind: product
  entities:
  - entity: device
    effect: reads
    facts:
    - Name
    - Kind
    - Joined at
    - Current device
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#listDevices
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListDevices
---

# List joined devices through the HTTP API

## Trigger

A client requests the joined devices.

## Outcome

Every joined device is returned with its name, kind and join time, oldest first.
