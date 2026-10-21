---
availability:
- place: server-cli
references:
- kind: code
  role: implementation
  target: cmd/ferry/main.go#parseConfig
- kind: code
  role: implementation
  target: docker-entrypoint.sh
---

# Change Server settings

Choose the data location, the listen address, an optional published host address and an optional reverse-proxy origin before starting the Server.
