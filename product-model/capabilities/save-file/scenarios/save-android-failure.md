---
kind: edge
routes:
  android: Android
steps:
- text: The User chooses Save on a received file and picks a destination
  kind: actor
  entities: []
  actor: user
  contexts:
    android:
      place: android::connected::timeline
- text: The Product reports the failed save
  kind: product
  entities:
  - entity: file-export
    effect: changes
    facts:
    - Error
    from: Saving
    to: Incomplete
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

# Report a failed Android save

## Trigger

The transfer fails, or the bytes received do not match the file's size.

## Outcome

The error is shown and no Saved is reported; the unfinished destination file is then removed.
