---
kind: alternate
routes:
  android: Android
steps:
- text: The User chooses Change
  kind: actor
  entities: []
  actor: user
  contexts:
    android:
      place: android::connected::timeline
- text: The Product ends the local connection
  kind: product
  entities:
  - entity: connection
    effect: changes
    facts:
    - Status
    - Polling
    from: Connected
    to: Setup
  actor: user
  contexts:
    android:
      place: android::connected::timeline
- text: The Product forgets the remembered connection for this address
  kind: product
  entities:
  - entity: remembered-connection
    effect: removes
  actor: user
  contexts:
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#disconnect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
---

# Change Server from a connected Android app

## Trigger

The User chooses Change while the Android app is connected.

## Outcome

Connect opens, ready for another Server address; the device stays joined on the old Server.
