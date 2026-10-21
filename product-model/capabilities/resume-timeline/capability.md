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

# Resume timeline updates

Start checking for new messages again when the native app returns to the foreground.
