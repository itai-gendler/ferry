---
kind: edge
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: A timeline poll fails temporarily
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
- text: The Product reports the connection problem
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Status
    from: Connected
    to: Offline
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
  target: internal/webui/assets/app.js#loadMessages
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#refreshMessages
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#refresh
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListMessages
---

# Keep displayed history through a temporary failure

## Trigger

A check for new messages fails, for example because the Server cannot be reached.

## Outcome

The messages already loaded stay visible and the connection shows Offline. A later successful check clears the error.
