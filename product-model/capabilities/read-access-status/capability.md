---
availability:
- place: web::access
- place: web::connected
- place: http-api::access
- place: http-api::connected
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
domain: devices
---

# Read access status

Find out whether new devices must give a password; the Web also shows it in Devices.
