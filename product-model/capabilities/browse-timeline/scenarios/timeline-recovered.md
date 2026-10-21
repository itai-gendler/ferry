---
kind: alternate
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: A later scheduled poll succeeds
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
- text: The Product presents the returned history and clears the connection failure
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
  - entity: connection
    effect: changes
    facts:
    - Status
    from: Offline
    to: Connected
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

# Recover the timeline after a temporary failure

## Trigger

A check for new messages succeeds after a failure.

## Outcome

History resumes without repeating already displayed entries.
