---
availability:
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#resetSession
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#setActive
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#invalidateSession
domain: timeline
---

# Cancel an upload

Stop a native file send that is in progress when the User leaves its connection; on iOS, also when the app goes to the background.
