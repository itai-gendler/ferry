---
kind: alternate
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
- text: The access setting requires no password
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
- text: The Product creates the default-named device
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

# Join automatically from a new browser

## Trigger

A new browser opens the Server address while no access password is set.

## Outcome

The browser joins under its default device name, remembers its identity and opens the shared timeline.
