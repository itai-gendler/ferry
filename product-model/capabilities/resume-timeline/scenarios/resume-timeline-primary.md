---
kind: primary
routes:
  ios: iOS
  android: Android
steps:
- text: The native app becomes active again
  kind: condition
  entities: []
  unattended: true
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product resumes polling for the connection
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

# Resume timeline updates

## Trigger

The app comes back to the foreground.

## Outcome

A connected or offline app starts checking for new messages again and catches up. The Android app also resumes reconnecting if that had not finished.
