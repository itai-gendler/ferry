---
availability:
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#setActive
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#setActive
domain: timeline
---

# Pause timeline updates

Stop checking for new messages while the native app is in the background.
