---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#clearAttachment
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#clearSelectedFile
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#clearSelectedFile
domain: timeline
---

# Remove an attachment

Remove the attachment from the composer before sending.
