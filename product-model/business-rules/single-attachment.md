---
appliesTo:
- type: entity
  id: attachment
- type: entity
  id: draft
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#updateComposer
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#send
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#send
---

# A composer holds one attachment and preserves its text draft

A selected file pauses text editing. Sending or removing it leaves existing text available; file and text are separate sends.
