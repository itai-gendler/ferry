---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#copyControl
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift#copy
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt#CopyButton
domain: timeline
---

# Copy text

Put one text message on this device's clipboard with its Copy control.
