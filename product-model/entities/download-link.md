---
domain: timeline
references:
- kind: code
  role: implementation
  target: internal/ferry/downloads.go
---

# Download link

A short-lived address that lets a browser save one file without the device credential, for browsers that can only save a file by opening its address.

## Information kept

- **Address** — The temporary download address.
- **Expires at** — Five minutes after issuance.
- **Usable** — Whether the link is within its lifetime, its issuing device is still joined and the Server has not restarted since it was issued.

## States

### Available

The link can serve its one file repeatedly.

### Expired

Its five-minute lifetime ended.

### Ended

The issuing device was revoked or the Server restarted.
