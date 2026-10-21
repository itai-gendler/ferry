---
kind: alternate
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User chooses another local file
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
- text: The Product replaces the attachment
  kind: product
  entities:
  - entity: attachment
    effect: changes
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
  target: internal/webui/assets/app.js#selectedAttachment
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#selectFile
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#selectContent
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
---

# Replace the attachment

## Trigger

An attachment is already in the composer and the User picks another file.

## Outcome

Only the newest choice remains the attachment.
