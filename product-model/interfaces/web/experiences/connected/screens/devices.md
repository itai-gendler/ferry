---
entities:
- entity: device
  shows:
  - Name
  - Kind
  - Joined at
  - Current device
- entity: access-settings
  shows:
  - Password required
  collects:
  - Password
- entity: connection
  shows:
  - Server address
  - Device name
  - Status
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
---

# Devices

Connecting another device and looking after the joined ones: the Server address with its QR code, the joined devices with Revoke, and the access password.
