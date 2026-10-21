---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User chooses Remove
  kind: actor
  entities:
  - entity: attachment
    effect: reads
    facts:
    - Name
  actor: user
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product removes the attachment
  kind: product
  entities:
  - entity: attachment
    effect: removes
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
  target: internal/webui/assets/app.js#clearAttachment
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#clearSelectedFile
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#clearSelectedFile
---

# Remove an attachment

## Trigger

The User chooses Remove on the attachment.

## Outcome

The attachment leaves the composer, text editing resumes, and any text draft is still there.
