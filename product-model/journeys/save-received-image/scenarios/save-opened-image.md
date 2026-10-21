---
kind: primary
result: achieved
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
  capability: open-image
  actor: user
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product opens the received image for saving
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  capability: open-image
  actor: user
  contexts:
    web:
      place: web::connected::image
    ios:
      place: ios::connected::image
    android:
      place: android::connected::image
- text: The User chooses Save
  kind: actor
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  capability: save-file
  actor: user
  contexts:
    web:
      place: web::connected::image
    ios:
      place: ios::connected::image
    android:
      place: android::connected::image
- text: The Product hands the original received bytes to the platform save flow
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  capability: save-file
  actor: user
  contexts:
    web:
      place: web::connected::image
    ios:
      place: ios::connected::image
    android:
      place: android::connected::image
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

# Save an opened image

## Trigger

The User opens a received image from the timeline.

## Outcome

The original uploaded bytes are offered to the platform save flow.
