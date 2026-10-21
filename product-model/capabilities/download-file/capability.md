---
availability:
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#serveFile
- kind: code
  role: implementation
  target: internal/ferry/store.go#OpenFile
domain: timeline
---

# Download a file

Retrieve a file's original bytes with a joined device's credential.
