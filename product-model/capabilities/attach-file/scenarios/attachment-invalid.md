---
kind: validation
routes:
  ios: iOS
  android: Android
steps:
- text: The User chooses a file larger than 64 MiB
  kind: actor
  entities: []
  actor: user
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
- text: The Product reports the selection error
  kind: product
  entities: []
  actor: user
  contexts:
    ios:
      place: ios::connected::timeline
    android:
      place: android::connected::timeline
references:
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

# Reject an oversized native selection

## Trigger

The chosen file is known to be larger than 64 MB.

## Outcome

The app reports that files must be 64 MB or smaller, and the choice does not become the attachment.
