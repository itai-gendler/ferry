---
appliesTo:
- type: entity
  id: connection
- type: entity
  id: draft
- type: entity
  id: attachment
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#resetSession
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#invalidateSession
---

# Old connection results do not replace a new connection

After leaving or replacing a session, delayed callbacks cannot overwrite current native client state. Text and attachment edits during native sends remain newer work.
