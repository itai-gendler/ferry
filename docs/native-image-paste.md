# Image paste in the iOS and Android apps

> English | [简体中文](native-image-paste.zh-Hans.md)

## Problem

- **Observed (Max, 2026-10-11, FerryDrop build 4)**: a screenshot copied on the iPhone could not be pasted into the FerryDrop message field. Image paste existed only in the Web app (2026-09-01, `web-attachment-picker.md`), whose scope excluded the native apps; that was the boundary of one task, not a product decision.
- **Decided (Max, 2026-10-11)**: pasting an image is basic behaviour; iOS, Android and the Web must all support it.
- Done means: with only an image on the system clipboard, the system Paste command in the message field turns it into the composer's single attachment, sending it produces an image message that every client previews, and pasted text still pastes as text.

## Why the old fields could not do it

- iOS: SwiftUI `TextField` is bound to a `String`. With only an image on the pasteboard, Paste inserted nothing (reproduced on build 4 in the Simulator, 2026-10-11).
- Android: Compose `contentReceiver` only works on the `TextFieldState` text field; the value-based `OutlinedTextField` cannot receive non-text content.
- iOS uploads always said `Content-Type: application/octet-stream`. The Server stores the part's Content-Type as the message's media type and every client previews only `image/*`, so even a picked photo from iOS arrived as a file card. Picked files now send the type of their filename extension.

## Mechanism

- iOS: the message field is a `UITextView` (`MessageInput.swift`). `paste(_:)` takes the first image off the pasteboard; PNG, JPEG, GIF and WebP keep their bytes, anything else (HEIC, TIFF) is re-encoded as PNG when it has alpha and JPEG otherwise. The bytes live in the selection (`SelectedFile.Contents.pasted`) because no file owns them. The pasteboard is read only inside the user's own Paste command, so iOS shows no paste prompt and Ferry still never reads the clipboard on its own.
- Android: the field is the `TextFieldState` `OutlinedTextField` with `contentReceiver`. The first content URI of an image clip goes to `FerryViewModel.pasteContent`, which reads its bytes at once (`ContentStore.snapshot`): the system revokes the app's read grant on a clipboard URI as soon as anything else is copied, so the bytes cannot wait for Send. Text that arrives together with an image is dropped, as on the Web. The field owns its text and reports edits to the model; the model writes back only to empty it after a send or a session reset.
- Both platforms reuse the existing single attachment slot, size limit (64 MB), remove control and file upload.

## Pitfalls met during verification

- Xcode 27 Device Hub copies the Mac clipboard into the Simulator when its Paste menu is used, replacing whatever was put there with `simctl`. The image has to be placed on the Simulator pasteboard immediately before Paste is pressed.
- The Android emulator's Chrome did not open its long-press menu for injected touches, and the system share sheet had no Copy action, so the Android check put a `content://` image on the clipboard from a throwaway helper app with its own ContentProvider, the way a screenshot tool does.

## Acceptance (2026-10-11)

- iOS Simulator, iOS 26.5, local Server: with only `public.png` on the pasteboard, Paste produced the `image.png 438 KB` attachment; Send stored `file|image.png|image/png|438423` (the seed's exact size) and the timeline previewed it, as did the Web app. Build 4 under the same steps produced nothing. Six pasted lines grew the field to five lines and then scrolled.
- Android emulator (API 37): Paste on an image clip produced `Screenshot_20261011.png 428.1 KB`; the clipboard was then replaced with text, and Send still stored `image/png`, 438423 bytes, previewed in the timeline. Pasting text and sending it stored the text and emptied the field.
- After the review fixes (2026-10-12) both apps were rebuilt and the image paste → send path repeated: iOS stored message 2 and Android message 3, each `image/png`, 438423 bytes. On iOS the field reports itself disabled while an attachment is selected and becomes editable again after Remove or Send.
- Not verified here: whether iOS Safari offers Paste in the Web app's `<textarea>` when the pasteboard holds only an image. Max will check it on a phone.
