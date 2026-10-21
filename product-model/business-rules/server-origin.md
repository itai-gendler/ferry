---
appliesTo:
- type: entity
  id: connection
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/ServerEndpoint.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/ServerEndpoint.kt
---

# Native connections name one HTTP or HTTPS origin

The entered Server address is normalized by case and default port. The connection uses a single origin, without an appended application path, query or fragment.
