---
appliesTo:
- type: entity
  id: text-message
references:
- kind: code
  role: implementation
  target: internal/ferry/model.go#validateText
---

# Text is nonblank and at most 64 KB (65,536 UTF-8 bytes)

Whitespace around nonblank text is preserved; all-whitespace input is rejected.
