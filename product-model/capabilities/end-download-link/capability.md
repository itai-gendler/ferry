---
availability:
- place: web::connected
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#downloadByLink
- kind: code
  role: implementation
  target: internal/ferry/server.go#newHandler
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#pruneLocked
domain: timeline
---

# End a download link

A download link stops working five minutes after it was issued, or earlier if its device is revoked or the Server restarts.
