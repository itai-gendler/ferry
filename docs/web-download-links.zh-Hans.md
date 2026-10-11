# Web 下载链接

> [English](web-download-links.md) | 简体中文

## 问题

- **Observed（Max，2026-10-11）**：在一台安卓手机上，用某个 App 的内置浏览器打开 Ferry Web，文件列表能显示，但点文件只弹出系统提示「没有应用可执行此操作。」，什么也没存下来。
- **Observed**：Web 原来的下载方式是带 Bearer 头取回文件，把字节包成 `blob:` URL，再点一个指向它的 `<a download>`。安卓的内置 WebView 常常自己存不了 `blob:` URL，会把它交给系统，而系统里没有能处理这个协议的应用。这条路径还会在保存前把整个文件（最大 64 MiB）放在页面内存里。
- **Inferred**：报告里的内置浏览器就是这类 WebView；没有在那台手机上复现，结论来自提示文字和代码路径。

## 机制

**Recommended**，还不是 Max 的决定。

1. Web 带着 Bearer token 请求 `POST /api/v1/files/{message_id}/link`。
2. Server 返回同源路径 `/api/v1/downloads/{ticket}`：32 字节随机数，内存里只存它的 SHA-256，连同消息 ID、签发设备和五分钟后的过期时间。
3. Web 用普通的 `<a download>` 跳到这个路径。浏览器用自己的下载器保存；WebView 则可以把这个普通的 `http://` URL 交给系统浏览器，系统浏览器照样能打开。
4. `GET /api/v1/downloads/{ticket}` 不需要 Bearer 头。格式不对、不存在、已过期，或签发设备已被吊销的 ticket，一律返回同样的 `404 not_found`。

带 Bearer 的 `GET /api/v1/files/{message_id}` 不变，原生 App 和 Web 的图片预览仍然用它。

## 取舍

URL 本身就是凭证：五分钟内，能访问到 Server 所在网络、又拿到这个 URL 的人，可以下载这一个文件。它不给其他文件、时间线或任何写权限。它不落盘，Server 重启即失效，签发设备被吊销即失效；Ferry 已经发送 `Referrer-Policy: no-referrer`，所以它不会作为 Referer 被带到别处。它可能出现在那个浏览器的下载记录里。

## 否决的方案

- **用 Cookie 会话下载**：Ferry 早就对所有凭证否决了 Cookie，因为 Cookie 不按端口隔离，同一台主机上的其他服务也会收到（见 `docs/lan-auth.zh-Hans.md`）。
- **用 Service Worker 给下载跳转加 Bearer 头**：Service Worker 只在安全上下文里存在，而 Ferry 在局域网里走明文 HTTP。
- **一次性 ticket**：WebView 可能先请求一次 URL，再交给系统浏览器重新请求；第一次请求就作废的 ticket 恰好会弄坏这次要修的场景。

## 验收

- Go 测试 `TestDownloadLinkGrantsOneFileWithoutBearerUntilExpiryOrRevocation`：签发需要 Bearer；文本消息返回 404；链接不带 Bearer 也能以附件形式下载到原始字节，有效期内可重复使用；到期、签发设备被吊销、ticket 格式不对或不存在时返回 404。
- 尚未完成：在运行这版代码的 Server 上，用安卓内置浏览器和手机系统浏览器各真实下载一次。
