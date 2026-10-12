# iOS 与 Android 应用的图片粘贴

> [English](native-image-paste.md) | 简体中文

## 问题

- **Observed（Max，2026-10-11，FerryDrop build 4）**：在 iPhone 上复制的截图粘贴不进 FerryDrop 的消息框。图片粘贴此前只有 Web 端做了（2026-09-01，`web-attachment-picker.md`），那次任务的范围排除了原生应用；这是一次任务的边界，不是产品决定。
- **Decided（Max，2026-10-11）**：粘贴图片是基本能力，iOS、Android、Web 都要支持。
- 完成的标准：系统剪贴板里只有一张图片时，在消息框里用系统的「粘贴」，它会成为输入区唯一的附件；发送后得到一条各端都能预览的图片消息；粘贴文字仍然是文字。

## 旧输入框为什么做不到

- iOS：SwiftUI `TextField` 绑定的是 `String`。剪贴板里只有图片时，粘贴什么也不会插入（2026-10-11 在模拟器上用 build 4 复现）。
- Android：Compose 的 `contentReceiver` 只对 `TextFieldState` 版输入框生效，按值绑定的 `OutlinedTextField` 收不到非文本内容。
- iOS 上传一律写 `Content-Type: application/octet-stream`。Server 把这一段的 Content-Type 原样存为消息的媒体类型，各端只预览 `image/*`，所以连 iOS 选取的照片到了别的设备上也只是文件卡片。现在选取的文件按文件扩展名发送对应类型。

## 机制

- iOS：消息框改为 `UITextView`（`MessageInput.swift`）。`paste(_:)` 从剪贴板取第一张图片；PNG、JPEG、GIF、WebP 保留原字节，其他格式（HEIC、TIFF）有透明通道时重新编码为 PNG，否则为 JPEG。字节放在选择项里（`SelectedFile.Contents.pasted`），因为没有文件承载它们。只在用户自己的「粘贴」命令里读取剪贴板，因此 iOS 不会弹出粘贴提示，Ferry 也仍然不会主动读取剪贴板。
- Android：输入框改为带 `contentReceiver` 的 `TextFieldState` 版 `OutlinedTextField`。图片剪贴内容的第一个 content URI 交给 `FerryViewModel.pasteContent`，它立刻读出字节（`ContentStore.snapshot`）：只要再复制任何东西，系统就会收回应用对剪贴板 URI 的读取授权，所以字节不能等到发送时再读。和图片一起到来的文字会被丢弃，与 Web 一致。输入框拥有自己的文字并把编辑同步给 model；model 只在发送成功或会话重置后写回，用来清空它。
- 两个平台都沿用现有的单附件位、64 MB 上限、移除按钮和文件上传。

## 验证中遇到的坑

- Xcode 27 的 Device Hub 在使用其粘贴菜单时会把 Mac 剪贴板复制进模拟器，覆盖用 `simctl` 放进去的内容。图片必须在按下「粘贴」前一刻放进模拟器剪贴板。
- Android 模拟器里的 Chrome 对注入的触摸不弹长按菜单，系统分享面板也没有「复制」，所以 Android 的验证用一个一次性的辅助应用，通过它自己的 ContentProvider 把 `content://` 图片放进剪贴板，与截图工具的做法相同。

## 验收（2026-10-11）

- iOS 模拟器，iOS 26.5，本地 Server：剪贴板里只有 `public.png` 时，「粘贴」得到附件 `image.png 438 KB`；发送后存为 `file|image.png|image/png|438423`（与原图字节数一致），时间线和 Web 端都显示为图片预览。同样步骤下 build 4 什么也没有得到。粘贴六行文字后输入框长到五行，之后滚动。
- Android 模拟器（API 37）：在图片剪贴内容上「粘贴」得到 `Screenshot_20261011.png 428.1 KB`；随后把剪贴板换成文字，发送仍存为 `image/png`、438423 字节，并在时间线中预览。粘贴文字并发送后，文字已存储，输入框清空。
- 按审查意见修正后（2026-10-12）两个应用重新构建，再走一遍图片粘贴 → 发送：iOS 存为第 2 条消息，Android 存为第 3 条，都是 `image/png`、438423 字节。iOS 在选中附件期间输入框对外报告为不可用，移除或发送后恢复可编辑。
- 这里没有验证：剪贴板里只有图片时，iOS Safari 是否会在 Web 应用的 `<textarea>` 里提供「粘贴」。由 Max 在手机上确认。
