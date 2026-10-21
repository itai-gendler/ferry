---
actors:
- user
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#openViewer
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryApp.kt
---

# Save a received image

## Goal

Keep a local copy of an image opened from the timeline.

## Success criterion

The opened image’s original uploaded bytes reach the platform save flow.
