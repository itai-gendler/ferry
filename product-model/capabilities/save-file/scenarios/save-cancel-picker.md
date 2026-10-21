---
kind: alternate
routes:
  ios: iOS
  android: Android
steps:
- text: The User cancels destination selection
  kind: actor
  entities: []
  actor: user
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product leaves the received file unchanged
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
  actor: user
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift#SaveControl
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#download
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
---

# Cancel choosing where to save

## Trigger

The User dismisses the system save picker or exporter.

## Outcome

Nothing is saved.
