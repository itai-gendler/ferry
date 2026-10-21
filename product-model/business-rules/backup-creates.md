---
appliesTo:
- type: entity
  id: backup
  effect: creates
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

# Deployers create backups

The deployer operates the locally hosted service and its data utility.
