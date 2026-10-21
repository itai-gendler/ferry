---
kind: primary
routes:
  devices: Web Devices
steps:
- text: The User opens the sharing information
  kind: actor
  entities: []
  actor: user
  contexts:
    devices:
      place: web::connected::devices
- text: The Product presents the current origin
  kind: product
  entities:
  - entity: connection
    effect: reads
    facts:
    - Server address
  actor: user
  contexts:
    devices:
      place: web::connected::devices
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#renderConnectSection
---

# View the Server address

## Trigger

The User opens Devices, or chooses Connect another device on an empty timeline.

## Outcome

Devices shows the address this browser is using and a QR code of it, made in the browser. At an address only this computer can reach, it explains that a LAN address is needed instead of showing a code.
