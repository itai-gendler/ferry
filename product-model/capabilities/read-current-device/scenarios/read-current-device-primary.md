---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client requests its current identity
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product returns the current device
  kind: product
  entities:
  - entity: device
    effect: reads
    facts:
    - Name
    - Kind
    - Joined at
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#session
---

# Read the current device

## Trigger

A client asks which device its credential belongs to.

## Outcome

The device's name, kind and join time are returned; a credential that is not a joined device's is refused.
