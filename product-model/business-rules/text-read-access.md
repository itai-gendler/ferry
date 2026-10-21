---
appliesTo:
- type: entity
  id: text-message
  effect: reads
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#requireDevice
- kind: code
  role: implementation
  target: internal/ferry/store.go#ListMessages
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
- unattended: true
---

# Every joined device reads the whole timeline

There is one shared timeline: every joined device reads all of its text, whichever device sent it. Clients keep it current under their existing connection.
