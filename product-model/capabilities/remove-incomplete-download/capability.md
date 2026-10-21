---
availability:
- place: android::connected
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#download
domain: timeline
---

# Remove an incomplete download

Remove the unfinished destination file when an Android save fails or is stopped.
