---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#loadMessages
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#refreshMessages
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#refresh
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListMessages
domain: timeline
---

# Browse the timeline

Read the shared timeline of text and files in Server order, and keep receiving new messages while it is open.
