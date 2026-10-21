---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#draft
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#updateDraft
domain: timeline
---

# Edit a draft

Type or revise text in the composer before sending it.
