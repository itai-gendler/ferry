---
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#loadSession
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt
---

# Connection

A client's current connection to one Server. While it is Offline, the timeline already loaded stays visible.

## Information kept

- **Server address** — The origin the client is using.
- **Device name** — The chosen or restored local device name.
- **Status** — The currently visible connection or failure message.
- **Polling** — Running while the native client is active, suspended while inactive.

## States

### Setup

A connection must be established.

### Connecting

Joining or restoring is in progress.

### Connected

The current device is connected.

### Offline

A temporary failure interrupted access to a remembered connection.
