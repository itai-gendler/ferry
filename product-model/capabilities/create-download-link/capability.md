---
availability:
- place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#createDownloadLink
domain: timeline
---

# Create a download link

Get a short-lived address for one file that a browser can open without the device credential.
