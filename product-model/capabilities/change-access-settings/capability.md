---
availability:
- place: web::connected
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: internal/ferry/server.go#updateAccessSettings
- kind: code
  role: implementation
  target: internal/ferry/store.go#setAccessPassword
domain: devices
---

# Change access settings

Set, replace or clear the shared access password that new devices must give to join.
