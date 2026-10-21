---
kind: edge
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User sends the current text
  kind: actor
  entities: []
  actor: user
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product reports the send error and keeps the draft
  kind: product
  entities:
  - entity: draft
    effect: reads
    facts:
    - Text
  actor: user
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
  target: internal/ferry/server.go#createText
- kind: code
  role: implementation
  target: internal/ferry/store.go#createText
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#sendText
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#send
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#send
---

# Keep unsent text after a send failure

## Trigger

Sending fails before the Server accepts the text.

## Outcome

The error stays visible and the text stays in the composer to send again; a later successful timeline check does not hide the error.
