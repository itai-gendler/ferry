---
references:
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#archive_volume
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#restore_volume
---

# Backup

An archive of the Server's data, at a location the Deployer names: messages, uploaded files, joined devices and the access setting.

## Information kept

- **Archive** — The chosen archive location.
- **Contents** — A consistent snapshot of Ferry database files and uploaded blobs.
