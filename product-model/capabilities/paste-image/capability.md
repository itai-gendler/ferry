---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#pastedImage
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageInput.swift#PastedImage
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#pasteContent
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/ContentStore.kt#snapshot
domain: timeline
---

# Paste an image

Paste an image from the clipboard into the message field, where it becomes the attachment.
