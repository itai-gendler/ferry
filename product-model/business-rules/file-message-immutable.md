---
appliesTo:
- type: entity
  id: file-message
  effect: changes
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#newHandler
permits: []
---

# Sent file messages are immutable

The shared timeline offers no operation to edit a sent entry.
