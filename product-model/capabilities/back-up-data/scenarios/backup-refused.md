---
kind: validation
routes:
  cli: Server CLI
steps:
- text: The Deployer selects an existing archive name
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product reports the destination error
  kind: product
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
references:
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#archive_volume
---

# Refuse an unusable backup destination

## Trigger

The archive already exists or its parent directory is absent.

## Outcome

The prior archive and data are preserved.
