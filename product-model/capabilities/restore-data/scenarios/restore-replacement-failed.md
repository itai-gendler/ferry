---
kind: edge
routes:
  cli: Server CLI
steps:
- text: The Deployer starts a restore
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product creates the safety backup
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
- text: The Product reports the replacement failure and leaves the Server stopped
  kind: product
  entities:
  - entity: server
    effect: changes
    facts:
    - Health
    from: Running
    to: Stopped
  actor: deployer
  contexts:
    cli:
      place: server-cli
references:
- kind: code
  role: implementation
  target: scripts/ferry-data.sh#restore_volume
---

# Leave the Server stopped when replacing the data fails

## Trigger

Unpacking the archive fails after the safety backup was written.

## Outcome

The safety backup stays available, and the Server is not started on partly restored data.
