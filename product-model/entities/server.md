---
singleton: true
references:
- kind: code
  role: implementation
  target: cmd/ferry/main.go
---

# Server

The self-hosted Ferry service and its selected data location.

## Information kept

- **Listen address** — Loopback by default; private-network or all-interface listening requires explicit selection.
- **Published address** — The address reported for opening Ferry from a browser.
- **Trusted origin** — The optional origin of the deployer’s reverse proxy.
- **Data location** — The persistent location for history, file content, devices and access settings.
- **Health** — Whether the service answers its health check.

## States

### Stopped

The process is not serving requests.

### Running

The service is serving the Web app and HTTP contract.
