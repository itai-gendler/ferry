---
kind: primary
routes:
  cli: Server CLI
steps:
- text: The Deployer stops the service
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product shuts down the Server
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
  target: cmd/ferry/main.go#run
- kind: code
  role: implementation
  target: cmd/ferry/main.go#main
---

# Stop the Server

## Trigger

The Deployer stops the Server.

## Outcome

The service stops; stored messages, uploaded files, devices and access settings remain available for the next start.
