---
kind: primary
routes:
  web: Web
  api: HTTP API
steps:
- text: The issuer is revoked or the service restarts
  kind: condition
  entities: []
  unattended: true
  contexts:
    web:
      place: web::connected::timeline
    api:
      place: http-api::connected
- text: The Product ends the download link
  kind: product
  entities:
  - entity: download-link
    effect: changes
    facts:
    - Usable
    from: Available
    to: Ended
  contexts:
    web:
      place: web::connected::timeline
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#downloadByLink
- kind: code
  role: implementation
  target: internal/ferry/server.go#newHandler
---

# End a download link early

## Trigger

The device that asked for the link is revoked, or the Server restarts, before the link expires.

## Outcome

The old address no longer downloads content; a new address can be requested for the retained file.
