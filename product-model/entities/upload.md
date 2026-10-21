---
domain: timeline
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#sendFile
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#send
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#send
---

# Upload

One send of a selected file, from the moment it starts until it is accepted, fails or is cancelled.

## Information kept

- **Name** — The file being sent.
- **Error** — The visible send failure, if any.

## States

### Sending

A selected file is being sent.

### Sent

A file message has been accepted.

### Failed

Sending ended with a visible error.

### Cancelled

The client ended this transfer when leaving its connection or activity.
