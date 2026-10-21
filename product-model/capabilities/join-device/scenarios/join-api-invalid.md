---
kind: validation
routes:
  public-api: HTTP API
steps:
- text: The API client supplies an invalid name or kind
  kind: actor
  entities: []
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
- text: The Product returns the validation error
  kind: product
  entities: []
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
---

# Reject an invalid join request

## Trigger

A join request has an unusable name, an unknown device kind or a malformed body.

## Outcome

No device identity is created.
