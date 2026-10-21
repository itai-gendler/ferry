---
availability:
- place: web::connected
- place: ios::connected
- place: android::connected
domain: timeline
references:
- kind: spec
  role: intent
  target: docs/product-core.md
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageInput.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt#Composer
---

# Paste text

Paste clipboard text into the message field with the system's Paste, to review before sending.
