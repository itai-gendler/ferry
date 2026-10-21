---
kind: primary
routes:
  public-api: HTTP API
steps:
- text: The API client follows the download link
  kind: actor
  entities:
  - entity: download-link
    effect: reads
    facts:
    - Address
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
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
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#downloadByLink
---

# Follow a download link through the HTTP API

## Trigger

A client requests a download link's address without a device credential.

## Outcome

The linked file is delivered. Following it does not consume the link; another request within its lifetime can succeed.
