---
kind: primary
routes:
  web-access: Web
  ios-access: iOS
  android-access: Android
steps:
- text: The User opens the client again
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
- text: The Product reads the saved identity
  kind: product
  entities:
  - entity: remembered-connection
    effect: reads
    facts:
    - Server address
    - Device identity
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
- text: The Product checks the current identity and restores the connection
  kind: product
  entities:
  - entity: device
    effect: reads
    facts:
    - Name
  - entity: connection
    effect: changes
    facts:
    - Server address
    - Device name
    from: Setup
    to: Connecting
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
- text: The Product completes the connection
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Status
    - Polling
    from: Connecting
    to: Connected
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

# Reconnect with the remembered identity

## Trigger

The client is reopened with a saved credential for its Server address.

## Outcome

The client reconnects as the same device and loads the timeline.
