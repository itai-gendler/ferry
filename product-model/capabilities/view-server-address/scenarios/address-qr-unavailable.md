---
kind: edge
routes:
  devices: Web Devices
steps:
- text: The User opens the sharing address
  kind: actor
  entities: []
  actor: user
  contexts:
    devices:
      place: web::connected::devices
- text: The Product presents the origin and a text fallback
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

# Keep the address when QR generation fails

## Trigger

The QR code cannot be generated.

## Outcome

Devices still shows the address and asks the User to type or copy it on the other device.
