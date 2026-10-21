---
entities:
- entity: connection
  shows:
  - Status
  collects:
  - Device name
- entity: access-settings
  collects:
  - Password
  shows:
  - Password required
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js
---

# Connect

Where a browser joins Ferry: it asks for a device name and, when one is set, the access password, and shows why joining failed.
