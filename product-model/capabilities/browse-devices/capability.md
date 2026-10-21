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
  target: internal/ferry/server.go#listDevices
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListDevices
domain: devices
---

# Browse devices

See every joined device, with its name, kind and when it joined, and which one is this device.
