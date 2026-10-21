---
kind: primary
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
- text: The Product ends the local connection
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Status
    - Polling
    from: Offline
    to: Setup
  actor: user
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product forgets the remembered connection for this address
  kind: product
  entities:
  - entity: remembered-connection
    effect: removes
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

# Change Server while offline

## Trigger

The app cannot reach its Server and the User chooses Change Server (iOS) or Change (Android).

## Outcome

Connect opens, ready for another Server address. The device stays joined on the old Server and the messages it sent remain.
