---
kind: primary
routes:
  ios: iOS
  android: Android
steps:
- text: The current connection is left while sending
  kind: condition
  entities:
  - entity: connection
    effect: reads
    facts: []
  unattended: true
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product cancels the upload
  kind: product
  entities:
  - entity: upload
    effect: changes
    facts: []
    from: Sending
    to: Cancelled
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#resetSession
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#setActive
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#invalidateSession
---

# Cancel an upload

## Trigger

The User leaves the connection while a file is being sent; on iOS, the app also goes to the background.

## Outcome

The send stops and reports nothing to the client's new state. If the Server had already accepted the file, it stays in the timeline.
