---
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
- kind: code
  role: implementation
  target: internal/ferry/store.go#DeleteDevice
domain: devices
relations:
- entity: text-message
  verb: sent
  cardinality: one-to-many
- entity: file-message
  verb: sent
  cardinality: one-to-many
- entity: download-link
  verb: issued
  cardinality: one-to-many
---

# Device

A device that has joined Ferry, told apart by its identity rather than its name: two devices may share a name. Revoking it removes it from the joined devices; what it sent stays in the timeline.

## Information kept

- **Name** — The chosen display name, trimmed at join.
- **Kind** — iPhone, iPad, Mac, Android, Windows or browser.
- **Joined at** — When this identity joined.
- **Current device** — Whether this is the device making the request.
- **Other devices remain** — Whether at least one other device would stay joined if this one were revoked.
