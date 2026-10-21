---
appliesTo:
- type: entity
  id: access-settings
  effect: reads
  facts:
  - Password
permits: []
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#accessStatus
---

# Stored passwords are never disclosed

Access status returns only whether a password is required. Saving accepts the new password as input; it is never returned as a setting value.
