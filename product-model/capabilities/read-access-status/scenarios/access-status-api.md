---
kind: primary
routes:
  public-api: HTTP API, before joining
  api: HTTP API, joined
steps:
- text: The API client requests the access requirement
  kind: actor
  entities: []
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
    api:
      place: http-api::connected
- text: The Product returns the requirement
  kind: product
  entities:
  - entity: access-settings
    effect: reads
    facts:
    - Password required
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#accessStatus
---

# Read the join requirement through the HTTP API

## Trigger

A client asks whether a password is required to join.

## Outcome

Whether a password is required is returned; the password itself never is.
