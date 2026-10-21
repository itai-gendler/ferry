---
kind: edge
routes:
  android: Android
steps:
- text: The User picks a destination for a second file while a save is running
  kind: actor
  entities: []
  actor: user
  contexts:
    android:
      place: android::connected::timeline
- text: The Product reports the busy save and removes the unused destination when possible
  kind: product
  entities:
  - entity: file-export
    effect: reads
    facts:
    - Destination
  actor: user
  contexts:
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#download
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
---

# Allow one Android save at a time

## Trigger

A file is being saved.

## Outcome

Save is unavailable on other files until the running save finishes. A destination picked for a second file anyway is refused with a request to wait, and its empty file is removed, or the User is asked to remove it.
