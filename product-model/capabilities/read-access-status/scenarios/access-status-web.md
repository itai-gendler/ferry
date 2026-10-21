---
kind: primary
routes:
  web-access: Web Connect
  devices: Web Devices
steps:
- text: The User opens the access information
  kind: actor
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    devices:
      place: web::connected::devices
- text: The Product presents the current requirement
  kind: product
  entities:
  - entity: access-settings
    effect: reads
    facts:
    - Password required
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    devices:
      place: web::connected::devices
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#accessStatus
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#beginAccess
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#openDevices
---

# Read the join requirement

## Trigger

A new browser opens Ferry, or the User opens Devices.

## Outcome

Connect asks for the password only when one is required, and Devices says whether a password is enabled; the password itself is never shown.
