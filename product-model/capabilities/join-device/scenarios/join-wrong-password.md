---
kind: validation
routes:
  web-access: Web
  ios-access: iOS
  android-access: Android
steps:
- text: The User supplies a name and a password and chooses Connect
  kind: actor
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
- text: The password does not match the access setting
  kind: condition
  entities:
  - entity: access-settings
    effect: reads
    facts:
    - Password required
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
- text: The Product refuses the join and reports that the password is incorrect
  kind: product
  entities: []
  actor: user
  contexts:
    web-access:
      place: web::access::connect
    ios-access:
      place: ios::access::connect
    android-access:
      place: android::access::connect
references:
- kind: code
  role: implementation
  target: internal/ferry/server.go#join
- kind: code
  role: implementation
  target: internal/ferry/auth.go#verifyAccessPassword
- kind: code
  role: implementation
  target: internal/webui/assets/app.js#showJoinFailure
---

# Refuse a wrong access password

## Trigger

An access password is set and the User supplies a different one.

## Outcome

No device joins. Connect stays open with the password field, so the User can try again.
