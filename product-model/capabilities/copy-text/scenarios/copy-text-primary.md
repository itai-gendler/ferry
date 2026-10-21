---
kind: primary
routes:
  web: Web
  ios: iOS
  android: Android
steps:
- text: The User chooses Copy on a text message
  kind: actor
  entities:
  - entity: text-message
    effect: reads
    facts:
    - Text
  actor: user
  contexts:
    web:
      place: web::connected::timeline
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product writes that text to the local clipboard
  kind: product
  entities:
  - entity: text-message
    effect: reads
    facts:
    - Text
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
  target: internal/webui/assets/app.js#copyControl
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift#copy
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt#CopyButton
---

# Copy a text message

## Trigger

The User chooses Copy on a text message.

## Outcome

The message text is on this device's clipboard and Copy briefly shows Copied. If the browser refuses, the Web asks the User to press Ctrl+C.
