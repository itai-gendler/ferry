---
appliesTo:
- type: entity
  id: draft
- type: entity
  id: attachment
- type: entity
  id: text-message
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageInput.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
---

# Clipboard changes require an explicit paste or copy

Ferry does not automatically read the clipboard, upload clipboard content or replace the clipboard when messages arrive.
