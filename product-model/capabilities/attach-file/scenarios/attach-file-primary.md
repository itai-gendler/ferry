---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User chooses a local file
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
- text: The Product keeps the selected attachment
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

# Attach a file

## Trigger

The User chooses to add an attachment and picks a file; the Web and Android also offer Photos beside Files.

## Outcome

The attachment shows its name, and on iOS and Android its size; text editing pauses until the attachment is sent or removed.
