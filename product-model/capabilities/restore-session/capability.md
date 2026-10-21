---
availability:
- place: web::access
- place: ios::access
- place: android::access
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#loadSession
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#start
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#start
domain: devices
---

# Restore a session

Reconnect as the same device when the client is reopened, using the identity it remembered for that Server.
