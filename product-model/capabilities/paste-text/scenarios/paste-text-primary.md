---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User explicitly pastes clipboard text
  kind: actor
  actor: user
  entities: []
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product keeps the pasted text in the draft
  kind: product
  actor: user
  entities:
  - entity: draft
    effect: creates
    facts:
    - Text
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
references:
- kind: spec
  role: intent
  target: docs/product-core.md
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageInput.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt#Composer
---

# Paste text into an empty composer

## Trigger

The User chooses the system Paste action with clipboard text while the composer is editable.

## Outcome

Text is present in the unsent draft. Nothing is sent until the User chooses Send. Pasting into an existing draft edits its text using the system field’s selection behavior.
