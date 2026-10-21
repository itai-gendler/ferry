---
appliesTo:
- type: entity
  id: backup
- type: entity
  id: server
references:
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#restore_volume
---

# Restore preserves a safety snapshot before replacing data

The source is validated before current data is replaced. The safety archive must have a new path. Failed replacement leaves the service stopped and the safety archive available.
