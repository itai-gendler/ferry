---
availability:
- place: android::connected
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#invalidateSession
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
domain: timeline
---

# Cancel a download

Stop an Android file save that is in progress when the User leaves its connection.
