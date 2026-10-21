---
appliesTo:
- type: entity
  id: device
- type: entity
  id: text-message
- type: entity
  id: file-message
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#listMessages
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#renderMessage
---

# Messages are attributed to the sending device, not its name

Each message keeps the name and kind of the device that sent it. The Web tells this device's own messages apart by device identity, so two devices with the same name stay distinct, and a revoked device's messages keep their sender.
