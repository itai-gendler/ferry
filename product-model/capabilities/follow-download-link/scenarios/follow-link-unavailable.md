---
kind: validation
routes:
  public-api: HTTP API
steps:
- text: The API client follows an expired address
  kind: actor
  entities: []
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
- text: The Product reports that the link is invalid or expired
  kind: product
  entities: []
  actor: api-client
  contexts:
    public-api:
      place: http-api::access
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#downloadByLink
---

# Report an unavailable link

## Trigger

The link has expired, has ended or never existed.

## Outcome

The link is reported as invalid or expired and no file is returned.
