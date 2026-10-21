---
kind: edge
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: A received image cannot be fetched or decoded
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
- text: The Product replaces the received image with the file card
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
    - Size
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
  target: internal/webui/assets/app.js#loadMessageImage
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#loadImage
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#loadImage
---

# Fall back when an image cannot be displayed

## Trigger

The image cannot be fetched or decoded.

## Outcome

The message shows the file card instead, so the file can still be saved. The Web also falls back when loading takes more than 60 seconds; the iOS and Android apps do not try that image again in the same session.
