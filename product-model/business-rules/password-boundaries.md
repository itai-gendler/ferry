---
appliesTo:
- type: entity
  id: access-settings
references:
- kind: code
  role: implementation
  target: internal/ferry/auth.go#validAccessPassword
- kind: code
  role: implementation
  target: internal/ferry/server.go#updateAccessSettings
---

# Access passwords use exact UTF-8 values

A nonempty password contains at most 256 UTF-8 bytes. Empty disables the setting; leading and trailing spaces in a nonempty password are meaningful. The stored password is never returned by access status.
