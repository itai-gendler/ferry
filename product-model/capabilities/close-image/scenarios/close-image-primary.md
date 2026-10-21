---
kind: primary
routes:
  web-image: Web image
  ios-image: iOS image
  android-image: Android image
steps:
- text: The User closes the received image
  kind: actor
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  actor: user
  contexts:
    web-image:
      place: web::connected::image
    ios-image:
      place: ios::connected::image
    android-image:
      place: android::connected::image
- text: The Product returns the received image to the timeline
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  actor: user
  contexts:
    web-image:
      place: web::connected::image
    ios-image:
      place: ios::connected::image
    android-image:
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

# Close an image

## Trigger

The User dismisses an opened image.

## Outcome

The timeline is shown again as it was.
