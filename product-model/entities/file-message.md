---
domain: timeline
references:
- kind: code
  role: implementation
  target: internal/ferry/store.go#createFile
- kind: code
  role: implementation
  target: internal/ferry/store.go#OpenFile
relations:
- entity: download-link
  verb: has
  cardinality: one-to-many
---

# File message

One uploaded file in the shared timeline; images are files with an image media type.

## Information kept

- **Sequence** — Server order shared by text and files.
- **Sender name** — The sending device name captured at send time.
- **Sender kind** — The sending device category.
- **Sent at** — The Server timestamp.
- **Name** — The normalized filename.
- **Media type** — The uploaded media type, with a generic binary fallback.
- **Size** — The actual uploaded byte count.
- **Content** — The original uploaded bytes.
- **Download address** — The address used to retrieve this file.
