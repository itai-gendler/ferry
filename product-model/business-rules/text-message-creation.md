---
appliesTo:
- type: entity
  id: text-message
  effect: creates
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#createText
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

# Only joined devices send text, under their own name

The Server takes the sender's identity, name and kind from the joined device making the request, and checks it is still joined when the text is published.
