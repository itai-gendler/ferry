---
domain: timeline
references:
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#download
---

# File export

A file the Android app is saving to the document the User chose in the system picker. The Web hands files to the browser and the iOS app hands them to the system exporter, so neither keeps one.

## Information kept

- **Name** — The original received filename.
- **Destination** — The location selected through the platform exporter.
- **Content** — The downloaded original bytes.
- **Error** — The saving or cleanup failure.

## States

### Saving

The client is transferring bytes for export.

### Saved

Export finished successfully.

### Incomplete

Saving failed and its destination needs cleanup.

### Cancelled

The current connection ended the transfer.
