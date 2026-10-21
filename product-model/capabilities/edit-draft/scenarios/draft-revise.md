---
kind: alternate
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User changes unsent text
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
- text: The Product replaces the draft text
  kind: product
  entities:
  - entity: draft
    effect: changes
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

# Revise an existing draft

## Trigger

The message field already holds unsent text and the User changes it.

## Outcome

The composer holds the changed text.
