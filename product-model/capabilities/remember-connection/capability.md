---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#storeToken
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#connect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#connect
domain: devices
---

# Remember a connection

Keep the identity a client received when it joined, so reopening it reconnects as the same device.
