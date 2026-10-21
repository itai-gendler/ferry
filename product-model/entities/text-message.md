---
domain: timeline
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#createText
---

# Text message

A deliberately sent piece of text in the shared timeline.

## Information kept

- **Sequence** — Server order shared by text and files.
- **Sender name** — The sending device name captured at send time.
- **Sender kind** — The sending device category.
- **Sent at** — The Server timestamp.
- **Text** — The original nonblank UTF-8 text.
