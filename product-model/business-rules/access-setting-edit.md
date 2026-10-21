---
appliesTo:
- type: entity
  id: access-settings
  effect: changes
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#updateAccessSettings
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

# Any joined device changes the access password

Ferry has no administrator: every joined device may set, replace or clear the shared password. A new password applies to devices that join later; devices already joined stay joined.
