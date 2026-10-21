---
id: ferry
summary: Self-hosted text and file sharing across your devices on a private network.
category: productivity
tags:
- self-hosted
- file-sharing
- clipboard
license: MIT
authors:
- name: max1874
languages:
- en
limitations:
- Ferry serves devices on a trusted private network.
- Content is sent deliberately; Ferry does not automatically synchronize clipboards.
- The shared space has device identities rather than user accounts or social groups.
- Files are kept and served as uploaded, up to 64 MB each, without re-encoding.
- Native apps check for new messages only while they are open; there is no background transfer.
references:
- kind: doc
  role: intent
  target: README.md
- kind: spec
  role: intent
  target: docs/product-core.md
- kind: doc
  role: context
  target: CONTRIBUTING.md
---

# Ferry

Ferry ferries text and files between devices on one trusted private network through a shared, persistent timeline. A self-hosted Server serves the Web app; optional iOS and Android apps use the same service.

## Intent

“Your devices, one timeline, on your own network.” Ferry is “a self-hosted clipboard and file ferry for the devices on one trusted private network”: everything you send lands in a single timeline that every joined device can read, “so moving a link from your phone to your laptop is a paste and a copy, not an email to yourself.” Its core actions are deliberate: the user sends a piece of text or a file, other connected devices see it in the same timeline, and the user copies the text, or previews and downloads the file. Nothing leaves the network you run it on, and there is no account to create.

## Existing specifications

[Product Core](../docs/product-core.md) defines product boundaries, [OpenAPI](../api/openapi.yaml) defines the HTTP contract, and [Contributing](../CONTRIBUTING.md) defines contributor expectations. This model complements them and does not replace them: it gathers the behavior of the Web, iOS and Android clients, the HTTP API and the Server tools in one place, using their terms.
