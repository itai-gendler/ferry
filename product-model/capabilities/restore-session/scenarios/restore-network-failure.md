---
kind: edge
routes:
  web-access: Web
  ios-access: iOS
  android-access: Android
steps:
- text: The User opens a remembered client
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
- text: The Product reports the failed connection and keeps the saved identity
  kind: product
  entities:
  - entity: remembered-connection
    effect: reads
    facts:
    - Device identity
  - entity: connection
    effect: changes
    facts:
    - Status
    from: Connecting
    to: Offline
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
  target: internal/webui/assets/app.js#loadSession
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#start
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#start
---

# Keep the remembered identity through a temporary outage

## Trigger

The client is reopened while its Server cannot be reached.

## Outcome

The failure is visible and the saved credential is kept. The Web and Android keep trying to reconnect; iOS shows its timeline offline with Change Server.
