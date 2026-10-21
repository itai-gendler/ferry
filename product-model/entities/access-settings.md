---
singleton: true
domain: devices
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#updateAccessSettings
- kind: code
  role: implementation
  target: internal/ferry/store.go#AccessPassword
---

# Access settings

The Server-wide access-password setting.

## Information kept

- **Password required** — Whether new devices must satisfy the shared password.
- **Password** — The newly supplied value when saving; empty disables password access. The saved setting does not disclose the value.
