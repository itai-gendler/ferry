---
entities:
- entity: connection
  shows:
  - Status
- entity: text-message
  shows:
  - Text
  - Sender name
  - Sender kind
  - Sent at
- entity: file-message
  shows:
  - Name
  - Size
  - Content
  - Sender name
  - Sender kind
  - Sent at
- entity: draft
  collects:
  - Text
- entity: attachment
  shows:
  - Name
- entity: upload
  shows:
  - Error
- entity: remembered-connection
  shows:
  - Storage warning
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
---

# Timeline

The shared timeline of text and files from every joined device, with the connection status and the composer for sending. An empty timeline offers Connect another device.
