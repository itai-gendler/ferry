---
appliesTo:
- type: entity
  id: file-message
references:
- kind: code
  role: implementation
  target: internal/ferry/model.go#normalizeFileName
- kind: code
  role: implementation
  target: internal/ferry/model.go#normalizeMediaType
---

# Stored filenames are normalized display names

The final path component becomes the display name. Names must be nonempty, valid UTF-8, at most 255 UTF-8 bytes and without control characters. Invalid media types become generic binary content.
