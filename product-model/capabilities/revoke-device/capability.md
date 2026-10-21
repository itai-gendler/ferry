---
availability:
- place: web::connected
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#loadDevices
- kind: code
  role: implementation
  target: internal/ferry/server.go#deleteDevice
- kind: code
  role: implementation
  target: internal/ferry/store.go#DeleteDevice
domain: devices
---

# Revoke a device

Remove another device from the joined devices, so its credential stops working.
