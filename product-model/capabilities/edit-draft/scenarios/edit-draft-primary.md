---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User enters or revises text
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
- text: The Product keeps the current draft
  kind: product
  entities:
  - entity: draft
    effect: creates
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
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#draft
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#updateDraft
---

# Edit a draft

## Trigger

The User types in the empty message field.

## Outcome

The text waits in the composer until the User sends it. Attaching and sending a file leaves it there.
