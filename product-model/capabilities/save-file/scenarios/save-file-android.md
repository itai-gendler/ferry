---
kind: primary
routes:
  android: Android
  android-image: Android image
steps:
- text: The User chooses Save
  kind: actor
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
  actor: user
  contexts:
    android:
      place: android::connected::timeline
    android-image:
      place: android::connected::image
- text: The Product starts saving the received bytes
  kind: product
  entities:
  - entity: file-export
    effect: creates
    facts:
    - Name
    - Destination
    - Content
    - Error
    to: Saving
  - entity: file-message
    effect: reads
    facts:
    - Content
    - Name
  actor: user
  contexts:
    android:
      place: android::connected::timeline
    android-image:
      place: android::connected::image
- text: The Product completes the export
  kind: product
  entities:
  - entity: file-export
    effect: changes
    facts:
    - Content
    from: Saving
    to: Saved
  actor: user
  contexts:
    android:
      place: android::connected::timeline
    android-image:
      place: android::connected::image
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#download
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
---

# Save a file on Android

## Trigger

The User chooses Save and picks a destination in the system picker.

## Outcome

The destination holds the original uploaded bytes, checked against the file's size, and the app reports Saved.
