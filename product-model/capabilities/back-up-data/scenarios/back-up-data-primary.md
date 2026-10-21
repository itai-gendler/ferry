---
kind: primary
routes:
  cli: Server CLI
steps:
- text: The Deployer chooses a new archive location
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product snapshots the stopped service data and creates a backup
  kind: product
  entities:
  - entity: backup
    effect: creates
    facts:
    - Archive
    - Contents
  - entity: server
    effect: reads
    facts:
    - Data location
  actor: deployer
  contexts:
    cli:
      place: server-cli
references:
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#archive_volume
---

# Back up data

## Trigger

The Deployer runs the backup with a new archive path.

## Outcome

A validated archive of the database and uploaded files is written. A Server that was running is started again; one that was stopped stays stopped.
