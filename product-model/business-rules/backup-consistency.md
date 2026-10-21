---
appliesTo:
- type: entity
  id: backup
- type: entity
  id: server
references:
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#archive_volume
---

# Backups capture database and files together

The utility stops a running service before taking the snapshot, validates the archive, then resumes only a previously running service. Existing archive paths are never overwritten.
