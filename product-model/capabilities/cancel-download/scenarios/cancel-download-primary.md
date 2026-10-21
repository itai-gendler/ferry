---
kind: primary
routes:
  android: Android
steps:
- text: The current connection is left while saving
  kind: condition
  entities:
  - entity: connection
    effect: reads
    facts: []
  unattended: true
  contexts:
    android:
      place: android::connected::timeline
- text: The Product cancels the file export transfer
  kind: product
  entities:
  - entity: file-export
    effect: changes
    facts: []
    from: Saving
    to: Cancelled
  contexts:
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#invalidateSession
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
---

# Cancel a download

## Trigger

The User leaves the connection, for example with Change, while a file is being saved.

## Outcome

The save stops without reporting a result to the new connection, and its unfinished destination file is removed.
