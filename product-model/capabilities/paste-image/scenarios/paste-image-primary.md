---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User explicitly pastes an image
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
- text: The Product keeps the pasted attachment
  kind: product
  entities:
  - entity: attachment
    effect: creates
    facts:
    - Name
    - Size
    - Media type
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
  target: internal/webui/assets/app.js#pastedImage
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageInput.swift#PastedImage
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#pasteContent
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/ContentStore.kt#snapshot
---

# Paste an image

## Trigger

The User chooses the system Paste in the message field while the clipboard holds an image.

## Outcome

The image becomes the attachment, ready for Send; nothing is sent until then. The iOS app keeps PNG, JPEG, GIF and WebP images as they are and converts other images to PNG or JPEG; the Android app copies the image at once, because its permission to read the clipboard item does not last.
