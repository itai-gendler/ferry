---
appliesTo:
- type: entity
  id: file-message
- type: entity
  id: upload
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#validMultipartShape
- kind: code
  role: implementation
  target: internal/ferry/store.go#createFile
---

# A file is at most 64 MB (67,108,864 bytes)

The actual byte count is enforced during upload. A valid zero-byte file is accepted. Each upload contains one file, with no text caption.
