---
kind: primary
routes:
  public-api: HTTP API
steps:
- text: The API client supplies a name and, if one is set, the access password
  kind: actor
  entities: []
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
- text: The Product creates the joined identity
  kind: product
  entities:
  - entity: device
    effect: creates
    facts:
    - Name
    - Kind
    - Joined at
    - Current device
    - Other devices remain
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
---

# Join through the HTTP API

## Trigger

A client sends a device name and, if one is set, the access password.

## Outcome

The new device and its credential are returned; the credential is shown only in this response.

## Edge cases

- Names are trimmed and must be 1–64 UTF-8 bytes without control characters.
- A client may state its device kind; otherwise the kind is inferred from the name, with browser as the fallback.
