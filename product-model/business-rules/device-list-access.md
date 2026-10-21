---
appliesTo:
- type: entity
  id: device
  effect: reads
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#requireDevice
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

# Only joined devices see the joined devices

The current device and the list of joined devices are read with a joined device's credential.
