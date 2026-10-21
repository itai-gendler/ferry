---
appliesTo:
- type: entity
  id: server
  effect: changes
references:
- kind: code
  role: implementation
  target: cmd/ferry/main.go
- kind: code
  role: implementation
  target: scripts/ferry-data.sh
permits:
- actors:
  - deployer
---

# Deployers configure, start and stop the Server

The deployer operates the locally hosted service and its data utility.
