---
kind: primary
routes:
  ios: iOS
  android: Android
steps:
- text: The native app becomes inactive
  kind: condition
  entities: []
  unattended: true
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product pauses polling for the connection
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Polling
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#setActive
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#setActive
---

# Pause timeline updates

## Trigger

The app goes to the background.

## Outcome

The app stops checking for new messages. The timeline already loaded stays, and being in the background is not shown as a failure.
