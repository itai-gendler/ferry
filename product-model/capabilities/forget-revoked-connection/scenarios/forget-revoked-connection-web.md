---
kind: alternate
routes:
  web: Web
steps:
- text: The credential this browser presents is refused
  kind: condition
  entities: []
  unattended: true
  contexts:
    web:
      place: web::connected::timeline
- text: The Product clears the composer and returns the connection to setup
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Status
    - Polling
    from: Connected
    to: Setup
  contexts:
    web:
      place: web::connected::timeline
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#showAccess
---

# Leave the Web after the device is revoked

## Trigger

The Server refuses the browser's credential, because another device revoked this device.

## Outcome

The Web shows that this device is no longer connected, clears the unsent text and attachment, and returns to Connect, where the browser can join again as a new device.
