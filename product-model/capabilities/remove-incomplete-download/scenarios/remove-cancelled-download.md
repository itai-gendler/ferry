---
kind: alternate
routes:
  android: Android
steps:
- text: A transfer was cancelled
  kind: condition
  entities: []
  unattended: true
  contexts:
    android:
      place: android::connected::timeline
- text: The Product removes the cancelled file export
  kind: product
  entities:
  - entity: file-export
    effect: removes
    from: Cancelled
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

# Remove the destination of a stopped save

## Trigger

The User left the connection while a file was being saved.

## Outcome

The unfinished destination file is removed when the system allows it.
