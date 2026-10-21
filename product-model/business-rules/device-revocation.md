---
appliesTo:
- type: entity
  id: device
  effect: removes
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#DeleteDevice
permits:
- actors:
  - user
  when:
  - entity: user
    fact: Device access
    is: Valid
  - fact: Current device
    is: false
  - fact: Other devices remain
    is: true
- actors:
  - api-client
  when:
  - entity: api-client
    fact: Device access
    is: Valid
  - fact: Current device
    is: false
  - fact: Other devices remain
    is: true
---

# A joined device revokes any other device, never itself or the last one

Revoking removes only that device's identity: the text and files it sent stay in the timeline, and its later requests are refused.
