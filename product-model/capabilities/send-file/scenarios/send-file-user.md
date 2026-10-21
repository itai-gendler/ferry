---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User sends the attached file
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
- text: The Product begins the upload
  kind: product
  entities:
  - entity: upload
    effect: creates
    facts:
    - Name
    - Error
    to: Sending
  actor: user
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product creates the file message after storing its bytes
  kind: product
  entities:
  - entity: file-message
    effect: creates
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
    - Name
    - Media type
    - Size
    - Content
    - Download address
  - entity: upload
    effect: changes
    facts:
    - Error
    from: Sending
    to: Sent
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

# Send a file

## Trigger

The User has attached a file and chooses Send.

## Outcome

The file appears in every joined device's timeline with its name and size, and images preview inline. The attachment leaves the composer; any text draft stays.
