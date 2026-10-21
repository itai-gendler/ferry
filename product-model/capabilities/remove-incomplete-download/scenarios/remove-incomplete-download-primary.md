---
kind: primary
routes:
  android: Android
steps:
- text: Saving failed before completion
  kind: condition
  entities: []
  unattended: true
  contexts:
    android:
      place: android::connected::timeline
- text: The Product removes the incomplete file export
  kind: product
  entities:
  - entity: file-export
    effect: removes
    from: Incomplete
  contexts:
    android:
      place: android::connected::timeline
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#download
---

# Remove an incomplete download

## Trigger

Saving a file fails before it completes.

## Outcome

The unfinished destination file is removed. If the system does not allow that, the error asks the User to remove it.
