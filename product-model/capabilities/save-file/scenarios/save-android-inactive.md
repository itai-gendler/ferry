---
kind: alternate
routes:
  android: Android
steps:
- text: The User chooses Save on a received file and picks a destination
  kind: actor
  actor: user
  entities: []
  contexts:
    android:
      place: android::connected::timeline
- text: The app goes to the background while the file is still being saved
  kind: condition
  entities: []
  contexts:
    android:
      place: android::connected::timeline
- text: The Product finishes writing the file and reports it saved
  kind: product
  actor: user
  entities:
  - entity: file-export
    effect: changes
    from: Saving
    to: Saved
    facts:
    - Content
  contexts:
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#setActive
- kind: code
  role: context
  target: android/app/src/test/java/com/max1874/ferry/FerryViewModelTest.kt#userDownloadContinuesWhileTheActivityIsTemporarilyStopped
---

# Keep saving while the Android app is briefly in the background

## Trigger

A save the User started is still running when the app goes to the background.

## Outcome

The save finishes and reports Saved, while checking for new messages stays paused until the app returns. Leaving the connection still stops the save; nothing else transfers in the background.
