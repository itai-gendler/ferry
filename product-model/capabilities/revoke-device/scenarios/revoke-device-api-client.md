---
kind: primary
routes:
  api: HTTP API
steps:
- text: The API client asks to revoke another device
  kind: actor
  entities:
  - entity: device
    effect: reads
    facts: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product removes that device
  kind: product
  entities:
  - entity: device
    effect: removes
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#deleteDevice
- kind: code
  role: implementation
  target: internal/ferry/store.go#DeleteDevice
---

# Revoke another device through the HTTP API

## Trigger

A client asks to revoke another joined device.

## Outcome

The device is no longer joined and its later requests are refused. The text and files it sent stay in the timeline.

## Edge cases

- A device that is no longer joined is reported as not found.
- Revoking the requesting device itself, or the only joined device, is refused.
