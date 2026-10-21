---
appliesTo:
- type: entity
  id: device
  effect: creates
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
permits:
- actors:
  - user
  - api-client
  when:
  - entity: access-settings
    fact: Password required
    is: false
- actors:
  - user
  when:
  - entity: user
    fact: Admission
    is: Accepted
- actors:
  - api-client
  when:
  - entity: api-client
    fact: Admission
    is: Accepted
---

# Devices join under the current access requirement

A new identity is admitted when password access is disabled or the exact supplied password satisfies the current setting. The requirement is checked when the join is processed.
