---
appliesTo:
- type: entity
  id: file-message
  effect: removes
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#newHandler
- kind: code
  role: implementation
  target: internal/ferry/store.go#DeleteDevice
- kind: code
  role: implementation
  target: internal/ferry/store.go#OpenStore
permits: []
---

# Sent file messages remain in history

The product offers no per-message deletion or age-based expiry. Restarting against the same data location and revoking a sender preserve its sent history.
