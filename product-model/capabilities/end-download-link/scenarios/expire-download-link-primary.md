---
kind: primary
routes:
  web: Web
  api: HTTP API
steps:
- text: The saving address reaches its expiry
  kind: condition
  entities:
  - entity: download-link
    effect: reads
    facts:
    - Expires at
  unattended: true
  contexts:
    web:
      place: web::connected::timeline
    api:
      place: http-api::connected
- text: The Product expires the download link
  kind: product
  entities:
  - entity: download-link
    effect: changes
    facts:
    - Usable
    from: Available
    to: Expired
  contexts:
    web:
      place: web::connected::timeline
    api:
      place: http-api::connected
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#pruneLocked
---

# Expire a download link

## Trigger

Five minutes pass after the link was issued.

## Outcome

Subsequent use reports the link as unavailable; the stored file remains.
