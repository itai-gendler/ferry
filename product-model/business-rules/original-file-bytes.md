---
appliesTo:
- type: entity
  id: file-message
- type: entity
  id: file-export
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#OpenFile
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift
- kind: code
  role: implementation
  target: android/app/src/main/java/com/max1874/ferry/FerryClient.kt#download
---

# Downloads preserve the uploaded bytes

Preview decoding does not rewrite the received file. A paste may prepare a portable encoding before upload; saving the received content uses those uploaded bytes.
