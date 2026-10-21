---
appliesTo:
- type: entity
  id: draft
- type: entity
  id: attachment
- type: entity
  id: upload
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#showAccess
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#resetSession
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#resetSession
---

# Unsent text and attachments stay on the device until sent

A draft and a selected attachment belong to the client composing them: nothing reaches the Server or another device until the User chooses Send. Leaving the connection discards them.
