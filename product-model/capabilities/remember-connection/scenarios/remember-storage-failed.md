---
kind: edge
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The local credential write fails after a successful join
  kind: condition
  entities: []
  unattended: true
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product retains the issued identity in the current remembered connection with a storage warning
  kind: product
  entities:
  - entity: remembered-connection
    effect: creates
    facts:
    - Server address
    - Device identity
    - Storage warning
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#storeToken
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#connect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#connect
---

# Keep the connection with a storage warning

## Trigger

The client cannot save the new credential on this device.

## Outcome

The client stays connected and shows a warning that stays visible; after reopening, it may have to join again.
