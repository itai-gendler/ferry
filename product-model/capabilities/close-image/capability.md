---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#openViewer
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
domain: timeline
---

# Close an image

Close an opened image and return to the timeline.
