---
kind: alternate
routes:
  web-access: Web
steps:
- text: The access setting requires a password
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
- text: The User supplies the shared password and submits Connect
  kind: actor
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
- text: The Product starts the connection attempt
  kind: product
  actor: user
  entities:
  - entity: connection
    effect: changes
    from: Setup
    to: Connecting
    facts:
    - Server address
    - Device name
    - Status
  contexts: &id001
    web-access:
      place: web::access::connect
- text: The Product creates the named device
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
  actor: user
  contexts: *id001
- text: The Product completes the connection
  kind: product
  actor: user
  entities:
  - entity: connection
    effect: changes
    from: Connecting
    to: Connected
    facts:
    - Status
    - Polling
  contexts: *id001
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#joinAccess
---

# Join from the browser password prompt

## Trigger

A new browser opens the Server address while an access password is set.

## Outcome

The browser joins under the chosen name, remembers its identity and opens the shared timeline.
