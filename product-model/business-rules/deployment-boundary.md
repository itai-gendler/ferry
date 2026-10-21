---
appliesTo:
- type: entity
  id: server
references:
- kind: code
  role: implementation
  target: cmd/ferry/main.go#parseConfig
- kind: code
  role: implementation
  target: cmd/ferry/main.go#browserAddressMessage
---

# The service uses an explicitly selected private deployment

Loopback is the default listener. LAN or all-interface listening is an explicit opt-in. A deployer’s reverse proxy is configured as one trusted origin, and startup reports the browser address to use.
