---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#downloadMessageFile
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#openViewer
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift#SaveControl
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#download
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
domain: timeline
---

# Save a file

Save a received file, or an opened image, to this device through the browser's downloads or the system's save flow, keeping the original bytes.
