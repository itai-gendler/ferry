---
kind: primary
routes:
  devices: Web Devices
steps:
- text: The User chooses Copy
  kind: actor
  entities:
  - entity: connection
    effect: reads
    facts:
    - Server address
  actor: user
  contexts:
    devices:
      place: web::connected::devices
- text: The Product writes the chosen address to the clipboard
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
  target: internal/webui/assets/app.js#copyToClipboard
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#renderConnectSection
---

# Copy the Server address

## Trigger

The User chooses Copy next to the Server address in Devices.

## Outcome

Devices reports that the address was copied; if the browser refuses, it selects the address for the User to copy.
