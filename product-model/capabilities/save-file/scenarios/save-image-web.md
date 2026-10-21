---
kind: alternate
routes:
  web-image: Web image
steps:
- text: The User chooses Save on the received image
  kind: actor
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  actor: user
  contexts:
    web-image:
      place: web::connected::image
- text: The Product hands the already fetched bytes to the browser
  kind: product
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Content
    - Name
  actor: user
  contexts:
    web-image:
      place: web::connected::image
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#downloadMessageFile
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#openViewer
---

# Save an opened image from the Web

## Trigger

The User chooses Save on an opened image.

## Outcome

The browser saves the image bytes it already loaded, under the file's name.
