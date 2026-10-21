---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#showAccess
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#handle
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#handle
domain: devices
---

# Forget a revoked connection

Return to Connect when the Server no longer accepts this device, because it was revoked.
