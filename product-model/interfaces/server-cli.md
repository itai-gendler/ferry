---
type: cli
actors:
- deployer
references:
- kind: code
  role: implementation
  target: cmd/ferry/main.go
- kind: code
  role: implementation
  target: scripts/ferry-data.sh
- kind: code
  role: implementation
  target: docker-entrypoint.sh
---

# Server CLI

How the Deployer runs Ferry: starting and stopping the Server with its listener, data location and optional proxy origin, and backing up or restoring its data with the supplied utility.
