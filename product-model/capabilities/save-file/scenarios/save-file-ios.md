---
kind: primary
routes:
  ios: iOS
  ios-image: iOS image
steps:
- text: The User chooses Save
  kind: actor
  actor: user
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
  contexts:
    ios:
      place: ios::connected::timeline
    ios-image:
      place: ios::connected::image
- text: The Product fetches original bytes for the system exporter
  kind: product
  actor: user
  entities:
  - entity: file-message
    effect: reads
    facts:
    - Name
    - Content
  contexts:
    ios:
      place: ios::connected::timeline
    ios-image:
      place: ios::connected::image
references:
- kind: code
  role: implementation
  target: ios/Ferry/Ferry/MessageRow.swift#SaveControl
---

# Save a file on iOS

## Trigger

The User chooses Save on a file card or on an opened image.

## Outcome

The app fetches the original uploaded bytes and hands them, with the file name, to the system exporter, where the User picks the destination. The bytes are released when the exporter closes.
