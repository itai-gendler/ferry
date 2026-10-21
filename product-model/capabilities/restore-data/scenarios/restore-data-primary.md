---
kind: primary
routes:
  cli: Server CLI
steps:
- text: The Deployer supplies the source archive and a new safety snapshot location
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product validates the archive and creates the safety backup
  kind: product
  entities:
  - entity: backup
    effect: creates
    facts:
    - Archive
    - Contents
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product replaces the Server data with the selected archive contents
  kind: product
  entities:
  - entity: backup
    effect: reads
    facts:
    - Contents
  - entity: server
    effect: changes
    facts:
    - Data location
  actor: deployer
  contexts:
    cli:
      place: server-cli
references:
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#restore_volume
---

# Restore data

## Trigger

The Deployer runs the restore with an archive and a new path for the safety backup.

## Outcome

The archive's database and uploaded files replace the current data together. A Server that was running is started again only after a successful restore.
