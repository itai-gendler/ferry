---
appliesTo:
- type: entity
  id: download-link
  effect: creates
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#createDownloadLink
permits:
- actors:
  - user
  when:
  - entity: user
    fact: Device access
    is: Valid
- actors:
  - api-client
  when:
  - entity: api-client
    fact: Device access
    is: Valid
---

# Joined clients create download links

Creating a saving address requires a joined device and an existing stored file.
