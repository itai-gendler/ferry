---
kind: primary
routes:
  cli: Server CLI
steps:
- text: The Deployer supplies deployment settings
  kind: actor
  entities: []
  actor: deployer
  contexts:
    cli:
      place: server-cli
- text: The Product validates the Server settings
  kind: product
  entities:
  - entity: server
    effect: changes
    facts:
    - Listen address
    - Published address
    - Trusted origin
    - Data location
  actor: deployer
  contexts:
    cli:
      place: server-cli
references:
- kind: code
  role: implementation
  target: cmd/ferry/main.go#parseConfig
- kind: code
  role: implementation
  target: docker-entrypoint.sh
---

# Change Server settings

## Trigger

The Deployer sets the data location, listen address and, optionally, the published host address and the proxy origin before starting the Server.

## Outcome

Valid settings are ready for startup. Invalid addresses, ports or an empty data location stop startup with a visible error.
