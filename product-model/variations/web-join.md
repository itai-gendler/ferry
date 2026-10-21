---
kind: configuration
of: capability-scenario
settings:
- entity: access-settings
  fact: Password required
takesEffect: The browser reads the current access setting before its join attempt. The Server checks the
  requirement again when processing the attempt.
stability: Existing joined devices retain their identities when the setting changes. If password access
  turns on between status and join, the browser reveals the password prompt.
alternatives:
- id: join-web-direct
  selectedWhen: Password required is false; a fresh Server starts in this state.
- id: join-web-password
  selectedWhen: Password required is true.
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#beginAccess
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
---

# Web join

The access-password setting selects whether the browser asks the User to submit a password.
