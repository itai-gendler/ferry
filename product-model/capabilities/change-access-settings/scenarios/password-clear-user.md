---
kind: alternate
routes:
  devices: Web Devices
steps:
- text: The User saves the access password
  kind: actor
  entities: []
  actor: user
  contexts:
    devices:
      place: web::connected::devices
- text: The Product saves the access requirement
  kind: product
  entities:
  - entity: access-settings
    effect: changes
    facts:
    - Password
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

# Clear the access password

## Trigger

The User saves Devices → Access password with the field empty.

## Outcome

Devices reports that new devices join directly; devices already joined stay joined.
