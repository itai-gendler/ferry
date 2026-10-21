---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: A received image comes into view
  kind: condition
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Media type
    - Name
  unattended: true
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product fetches the received image
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product shows the decoded image in its message
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
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

# Show a received image inline

## Trigger

An image file comes into view in the timeline.

## Outcome

The image is shown in its message and can be opened.
