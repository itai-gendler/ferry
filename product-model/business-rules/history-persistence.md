---
appliesTo:
- type: entity
  id: server
- type: entity
  id: text-message
- type: entity
  id: file-message
- type: entity
  id: device
- type: entity
  id: access-settings
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#OpenStore
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#resetSession
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#resetSession
---

# Server data persists across restarts at the same data location

Messages, uploaded files, device identities and access settings persist. Native timeline copies are session-local; reconnecting reloads from the Server.
