---
kind: primary
routes:
  public-api: HTTP API
steps:
- text: The API client requests the service health
  kind: actor
  entities: []
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
- text: The Product returns the Server health response
  kind: product
  entities:
  - entity: server
    effect: reads
    facts:
    - Health
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#health
---

# Check Server health

## Trigger

A monitor or the Deployer asks the Server whether it is up.

## Outcome

A running Server answers ok. This shows the service responds, not that devices can transfer content.
