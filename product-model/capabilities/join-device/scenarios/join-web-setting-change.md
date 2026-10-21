---
kind: edge
routes:
  web-access: Web
steps:
- text: The User opens Ferry while no password is required
  kind: actor
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
- text: The access requirement changes before joining
  kind: condition
  entities:
  - entity: access-settings
    effect: reads
    facts:
    - Password required
  actor: user
  contexts:
    web-access:
      place: web::access::connect
- text: The Product presents the password prompt and preserves the failed join as an error
  kind: product
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#joinAccess
---

# Reveal the password prompt after the setting changes

## Trigger

An access password is set after the browser read that none was required, before it joined.

## Outcome

The User can supply the required password without reloading.
