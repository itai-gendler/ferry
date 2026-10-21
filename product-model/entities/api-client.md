---
acts: external
kind: system
references:
- kind: spec
  role: intent
  target: api/openapi.yaml
---

# API client

A program using Ferry's published HTTP contract on behalf of a device. The facts below describe the credential or link it presents with a request.

## Information kept

- **Device access** — Valid while the device credential this client presents belongs to a joined device.
- **Admission** — Accepted when the password supplied with a join satisfies the current access setting.
- **Download link usable** — Whether the download link this client follows is still usable for its file.
