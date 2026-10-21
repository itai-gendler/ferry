---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User opens a received image
  kind: actor
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  actor: user
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product opens the received image
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
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
  target: internal/webui/assets/app.js#openViewer
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
---

# Open an image

## Trigger

The User taps or clicks a received image in the timeline.

## Outcome

The Image view shows the received image and offers Save.
