---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: Joining has succeeded on this client
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
- text: The Product stores the remembered connection
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

# Remember a connection

## Trigger

Joining succeeds.

## Outcome

Reopening the client on this device reconnects as the same device without joining again.
