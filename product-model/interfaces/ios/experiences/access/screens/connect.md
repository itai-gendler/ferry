---
entities:
- entity: connection
  shows:
  - Status
  collects:
  - Device name
  - Server address
- entity: access-settings
  collects:
  - Password
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/AppModel.swift
---

# Connect

Where the iOS app is given a Server address, a device name and, if one is set, the access password, and shows why connecting failed.
