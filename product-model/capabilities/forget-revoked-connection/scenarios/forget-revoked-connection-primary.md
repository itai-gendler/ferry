---
kind: primary
routes:
  ios: iOS
  android: Android
steps:
- text: The credential this app presents is refused
  kind: condition
  entities: []
  unattended: true
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
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product clears the timeline and returns the connection to setup
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Status
    - Polling
    from: Connected
    to: Setup
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#handle
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#handle
---

# Leave a native app after the device is revoked

## Trigger

The Server refuses the app's credential, because another device revoked this device.

## Outcome

The app shows that this device is no longer connected, forgets the saved credential for that Server, clears the timeline and returns to Connect.
