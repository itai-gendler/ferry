---
kind: primary
routes:
  devices: Web Devices
steps:
- text: The User opens the joined devices
  kind: actor
  entities:
  - entity: device
    effect: reads
    facts: []
  actor: user
  contexts:
    devices:
      place: web::connected::devices
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
  target: internal/ferry/server.go#listDevices
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListDevices
---

# List joined devices

## Trigger

The User opens Devices.

## Outcome

Every joined device is listed with its name, kind and join date; this device is marked as This device and every other one offers Revoke.
