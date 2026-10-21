---
availability:
- place: web::access
- place: http-api::access
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#downloadByLink
domain: timeline
---

# Follow a download link

Open a download link to receive its file while the link is still usable.
