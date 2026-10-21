---
kind: primary
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

# Set the access password

## Trigger

The User enters a password in Devices → Access password and saves it.

## Outcome

Devices reports that the password is enabled for new devices. Later joins need this exact password; devices already joined stay joined.
