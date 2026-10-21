---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User chooses Send with text in the composer
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
- text: The Product creates a text message
  kind: product
  entities:
  - entity: text-message
    effect: creates
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
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

# Send text

## Trigger

The User has typed or pasted text and chooses Send.

## Outcome

The text appears in every joined device's timeline under this device's name, and the sent text leaves the composer.

## Edge cases

- Whitespace around nonblank text is kept as sent.
- On iOS and Android, text edited while the send was in progress stays in the composer.
