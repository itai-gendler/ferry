---
availability:
- place: web::access
- place: ios::access
- place: android::access
- place: http-api::access
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#joinAccess
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#connect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#connect
domain: devices
---

# Join a device

Join this device to the Server under a name, giving the access password when one is set, and open the shared timeline.
