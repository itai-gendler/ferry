---
kind: primary
routes:
  web-download: Web
steps:
- text: The User's browser follows the download link
  kind: actor
  entities:
  - entity: download-link
    effect: reads
    facts:
    - Address
  actor: user
  contexts:
    web-download:
      place: web::access
- text: The Product serves the linked file
  kind: product
  entities:
  - entity: download-link
    effect: reads
    facts:
    - Usable
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
    - Media type
    - Size
  actor: user
  contexts:
    web-download:
      place: web::access
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#downloadByLink
---

# Follow a download link

## Trigger

The browser opens the download link the Web just obtained, or the system browser it hands the link to.

## Outcome

The browser receives the file under its stored name and saves it. The link is not used up: it can be opened again within its lifetime.
