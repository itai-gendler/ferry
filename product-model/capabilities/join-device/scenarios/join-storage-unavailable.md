---
kind: edge
routes:
  web-access: Web
steps:
- text: The User opens Ferry
  kind: actor
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
- text: The Product detects unavailable site storage
  kind: condition
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
- text: The Product requests that site storage be enabled
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

# Stop joining when browser storage is unavailable

## Trigger

The browser cannot store anything for this site, so it could not remember a new identity.

## Outcome

No join request is sent, and Connect asks the User to enable site storage.
