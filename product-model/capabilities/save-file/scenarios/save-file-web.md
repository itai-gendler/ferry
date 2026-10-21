---
kind: primary
routes:
  web: Web
steps:
- text: The User chooses to save a received file
  kind: actor
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
  actor: user
  contexts:
    web:
      place: web::connected::timeline
- text: The Product obtains a temporary address and hands it to the downloader
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Download address
  - entity: download-link
    effect: creates
    facts:
    - Address
    - Expires at
    - Usable
    to: Available
  actor: user
  contexts:
    web:
      place: web::connected::timeline
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#downloadMessageFile
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#openViewer
---

# Save a file from the Web

## Trigger

The User chooses a file card in the timeline.

## Outcome

The browser opens a short-lived download link and saves the file itself. If the link cannot be obtained, the Web shows the error.
