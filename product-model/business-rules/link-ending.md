---
appliesTo:
- type: entity
  id: download-link
  effect: changes
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go
permits:
- unattended: true
---

# A download link lasts five minutes, ending early if its device is revoked or the Server restarts

Following a link does not use it up: it can be followed again within its lifetime. The file itself stays in the timeline, and a new link can be issued for it.
