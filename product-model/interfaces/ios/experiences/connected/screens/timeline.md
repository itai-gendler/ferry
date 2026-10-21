---
entities:
- entity: connection
  shows:
  - Device name
  - Status
- entity: text-message
  shows:
  - Text
  - Sender name
- entity: file-message
  shows:
  - Name
  - Size
  - Content
  - Sender name
- entity: draft
  collects:
  - Text
- entity: attachment
  shows:
  - Name
  - Size
- entity: upload
  shows:
  - Error
- entity: remembered-connection
  shows:
  - Storage warning
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift
---

# Timeline

The shared timeline of text and files from every joined device, with this device's name, the connection status and the composer for sending.
