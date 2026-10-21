---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#createFile
- kind: code
  role: implementation
  target: internal/ferry/store.go#createFile
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#sendFile
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#send
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#send
domain: timeline
---

# Send a file

Send the attached file to the shared timeline; it appears once all its bytes are stored.
