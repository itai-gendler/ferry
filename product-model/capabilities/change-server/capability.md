---
availability:
- place: ios::connected
- place: android::connected
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#disconnect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#disconnect
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/TimelineView.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
domain: devices
---

# Change Server

Leave the current Server in a native app and return to Connect. Android offers Change at any time; iOS offers Change Server when the Server cannot be reached.
