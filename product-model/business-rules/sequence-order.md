---
appliesTo:
- type: entity
  id: text-message
- type: entity
  id: file-message
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListMessages
- kind: code
  role: implementation
  target: internal/ferry/server.go#listMessages
---

# One Server sequence orders all timeline entries

Text and file messages share ascending Server order. Pagination continues strictly after the cursor; timestamps and duplicate display names do not establish identity or order.
