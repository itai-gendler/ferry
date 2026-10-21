---
kind: alternate
routes:
  api: HTTP API
steps:
- text: The API client saves the access setting
  kind: actor
  entities: []
  actor: api-client
  contexts:
    api:
      place: http-api::connected
- text: The Product saves the access requirement
  kind: product
  entities:
  - entity: access-settings
    effect: changes
    facts:
    - Password
    - Password required
  actor: api-client
  contexts:
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#updateAccessSettings
- kind: code
  role: implementation
  target: internal/ferry/store.go#setAccessPassword
---

# Clear the access password through the HTTP API

## Trigger

A client saves an empty password.

## Outcome

New devices join directly; existing identities remain valid.
