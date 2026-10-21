---
kind: edge
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User sends the selected file
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
- text: The Product reports the upload failure
  kind: product
  entities:
  - entity: upload
    effect: changes
    facts:
    - Error
    from: Sending
    to: Failed
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
  target: internal/ferry/server.go#createFile
- kind: code
  role: implementation
  target: internal/ferry/store.go#createFile
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#sendFile
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#send
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#send
---

# Report a failed file send

## Trigger

Reading, transferring or storing the attached file fails.

## Outcome

The error is visible, no file message appears, and the attachment stays in the composer to send again.
