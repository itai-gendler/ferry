---
entities:
- entity: connection
  shows:
  - Status
  collects:
  - Device name
  - Server address
- entity: access-settings
  collects:
  - Password
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt
---

# Connect

Where the Android app is given a Server address, a device name and, if one is set, the access password, and shows why connecting failed.
