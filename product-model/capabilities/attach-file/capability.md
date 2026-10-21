---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#selectedAttachment
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#selectFile
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#selectContent
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
domain: timeline
---

# Attach a file

Choose one file to send with the composer; the Web and Android also offer a photo picker.
