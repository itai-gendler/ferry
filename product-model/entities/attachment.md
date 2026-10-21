---
domain: timeline
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#selectedAttachment
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#selectFile
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#selectContent
---

# Attachment

The single file selected for the next send. Removing it does not remove a sent file.

## Information kept

- **Name** — Selected filename.
- **Size** — Known byte count, or unknown until read on Android.
- **Media type** — The file type used for preview and upload.
- **Content** — The picked file or the bytes captured by an explicit image paste.
