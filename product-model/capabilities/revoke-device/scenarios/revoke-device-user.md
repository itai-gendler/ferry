---
kind: primary
routes:
  devices: Web Devices
steps:
- text: The User chooses Revoke on another device
  kind: actor
  entities:
  - entity: device
    effect: reads
    facts: []
  actor: user
  contexts:
    devices:
      place: web::connected::devices
- text: The Product removes that device
  kind: product
  entities:
  - entity: device
    effect: removes
  actor: user
  contexts:
    devices:
      place: web::connected::devices
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#loadDevices
- kind: code
  role: implementation
  target: internal/ferry/server.go#deleteDevice
- kind: code
  role: implementation
  target: internal/ferry/store.go#DeleteDevice
---

# Revoke another joined device

## Trigger

The User chooses Revoke on another device in Devices.

## Outcome

The device disappears from the list and its later requests are refused, so it returns to joining. The text and files it sent stay in the timeline.

## Edge cases

- This device has no Revoke control.
