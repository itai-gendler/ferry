---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#createText
- kind: code
  role: implementation
  target: internal/ferry/store.go#createText
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#sendText
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#send
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#send
domain: timeline
---

# Send text

Send the text in the composer to the shared timeline.
