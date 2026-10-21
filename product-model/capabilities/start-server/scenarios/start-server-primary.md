---
kind: primary
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
- text: The Product starts the Server
  kind: product
  entities:
  - entity: server
    effect: changes
    facts:
    - Health
    - Published address
    from: Stopped
    to: Running
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

# Start the Server

## Trigger

The Deployer starts the Server.

## Outcome

The service reports an address to open and serves existing data from the selected location; a fresh location starts an empty space.
