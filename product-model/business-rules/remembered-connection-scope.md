---
appliesTo:
- type: entity
  id: remembered-connection
- type: entity
  id: device
references:
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#storeToken
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/CredentialStore.swift#KeychainCredentialStore
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/CredentialStore.kt
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift#disconnect
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryViewModel.kt#disconnect
---

# A remembered connection belongs to one device and one Server address

Each client keeps the identity it was given when it joined, for the Server address it joined, and reconnects with it when reopened. Forgetting it on the device, through Change Server, does not revoke the device on the Server: it stays among the joined devices until another device revokes it.
