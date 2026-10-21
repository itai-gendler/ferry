---
kind: primary
routes:
  ios-access: iOS
  android-access: Android
steps:
- text: The User supplies a name and, if one is set, the access password
  kind: actor
  entities: []
  actor: user
  contexts:
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
- text: The Product starts the connection attempt
  kind: product
  actor: user
  entities:
  - entity: connection
    effect: changes
    from: Setup
    to: Connecting
    facts:
    - Server address
    - Device name
    - Status
  contexts: &id001
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
- text: The Product creates the joined identity
  kind: product
  entities:
  - entity: device
    effect: creates
    facts:
    - Name
    - Kind
    - Joined at
    - Current device
    - Other devices remain
  actor: user
  contexts: *id001
- text: The Product completes the connection
  kind: product
  actor: user
  entities:
  - entity: connection
    effect: changes
    from: Connecting
    to: Connected
    facts:
    - Status
    - Polling
  contexts: *id001
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#connect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#connect
---

# Join from a native app

## Trigger

The User enters the Server address, a device name and, if one is set, the access password, and chooses Connect.

## Outcome

The Server joins a new device under that name and the app opens the shared timeline.

## Edge cases

- Names are trimmed and must be 1–64 UTF-8 bytes without control characters.
- A connection failure stays visible on Connect; the app does not show a completed join.
