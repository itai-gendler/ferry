---
appliesTo:
- type: entity
  id: file-message
  effect: reads
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#requireDevice
- kind: code
  role: implementation
  target: internal/ferry/downloads.go#downloadByLink
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
- actors:
  - user
  when:
  - entity: user
    fact: Download link usable
    is: true
- actors:
  - api-client
  when:
  - entity: api-client
    fact: Download link usable
    is: true
- unattended: true
---

# Joined devices read every file; a download link reads only its own

Every joined device retrieves any file in the timeline with its credential. Someone following a usable download link retrieves only the file it was issued for. Image previews load under the client's existing connection.
