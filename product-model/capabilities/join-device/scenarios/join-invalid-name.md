---
kind: validation
routes:
  web-access: Web
  ios-access: iOS
  android-access: Android
steps:
- text: The User supplies an empty or oversized name
  kind: actor
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
- text: The Product reports invalid input without creating an identity
  kind: product
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#joinAccess
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#connect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#connect
---

# Reject an unusable device name

## Trigger

The submitted device name is empty after trimming, longer than 64 UTF-8 bytes or contains control characters.

## Outcome

The User can correct the name and reconnect.
