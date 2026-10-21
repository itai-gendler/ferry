---
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#storeToken
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/CredentialStore.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/CredentialStore.kt
---

# Remembered connection

The joined identity a client keeps on this device for one Server address, so reopening the client connects again without joining. Forgetting it does not revoke the device on the Server.

## Information kept

- **Server address** — The origin to which this saved identity belongs.
- **Device identity** — The credential issued at join, used to restore the same identity.
- **Storage warning** — An error remembering or removing the local credential.
