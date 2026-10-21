---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User opens the timeline
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
- text: The Product returns the next ordered messages
  kind: product
  entities:
  - entity: text-message
    effect: reads
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
    - Text
  - entity: file-message
    effect: reads
    facts:
    - Sequence
    - Sender name
    - Sender kind
    - Sent at
    - Name
    - Media type
    - Size
    - Download address
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

# Read the shared timeline

## Trigger

The User opens the timeline; while it stays open, the client keeps checking for new messages.

## Outcome

Text and file messages from every joined device appear in Server order, each once, with the sending device and time.

## Edge cases

- An empty timeline invites the User to send something; the Web also offers Connect another device.
