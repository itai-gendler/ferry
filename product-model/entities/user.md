---
acts: external
kind: person
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt
---

# User

A person ferrying text and files between their own devices. Ferry keeps no account for them: the facts below describe the device they are using right now.

## Information kept

- **Device access** — Valid while the device credential this client presents belongs to a joined device.
- **Admission** — Accepted when the password supplied with a join satisfies the current access setting.
- **Download link usable** — Whether the download link this client follows is still usable for its file.
