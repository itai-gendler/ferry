---
kind: validation
routes:
  cli: Server CLI
steps:
- text: The Deployer supplies an invalid archive
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product reports that the archive is not accepted
  kind: product
  entities:
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
  target: scripts/ferry-data.sh#restore_volume
---

# Reject an invalid restore archive

## Trigger

The archive is not a Ferry data backup, or the safety backup path already exists.

## Outcome

The existing Server data is not replaced.
