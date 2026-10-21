---
kind: validation
routes:
  devices: Web Devices
steps:
- text: The User saves an oversized password
  kind: actor
  entities: []
  actor: user
  contexts:
    devices:
      place: web::connected::devices
- text: The Product reports invalid input
  kind: product
  entities:
  - entity: access-settings
    effect: reads
    facts:
    - Password required
  actor: user
  contexts:
    devices:
      place: web::connected::devices
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: internal/ferry/server.go#updateAccessSettings
- kind: code
  role: implementation
  target: internal/ferry/store.go#setAccessPassword
---

# Reject a password longer than 256 bytes

## Trigger

The User saves a password longer than 256 UTF-8 bytes.

## Outcome

The error is shown in Devices and the current setting stays as it was.
