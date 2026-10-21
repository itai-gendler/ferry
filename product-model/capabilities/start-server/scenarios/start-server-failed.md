---
kind: edge
routes:
  cli: Server CLI
steps:
- text: The Deployer starts the service
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product reports the startup failure
  kind: product
  entities:
  - entity: server
    effect: reads
    facts:
    - Listen address
    - Data location
  actor: deployer
  contexts:
    cli:
      place: server-cli
references:
- kind: code
  role: implementation
  target: cmd/ferry/main.go#run
- kind: code
  role: implementation
  target: internal/ferry/store.go#OpenStore
---

# Report a startup failure

## Trigger

The listener is busy or the data location cannot be opened.

## Outcome

No running service is announced.
