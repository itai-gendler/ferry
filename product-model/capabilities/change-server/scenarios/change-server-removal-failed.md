---
kind: edge
routes:
  ios: iOS
  android: Android
steps:
- text: The User chooses to leave the current service
  kind: actor
  entities: []
  actor: user
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product returns to setup and reports a storage error
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Status
    from: Offline
    to: Setup
  - entity: remembered-connection
    effect: changes
    facts:
    - Storage warning
  actor: user
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
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
---

# Report that the saved credential could not be forgotten

## Trigger

The app cannot remove the saved credential while leaving the Server.

## Outcome

The app still returns to Connect, and shows that the saved credential could not be removed.
