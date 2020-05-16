## Notify-Server
### lazy like me (っ´ω`c)

[![Go Report Card](https://goreportcard.com/badge/github.com/zorhayashi/notify-server)](https://goreportcard.com/report/github.com/zorhayashi/notify-server)

![Logo](/logo-m.png)

微服務-通知訊息伺服器

Notify-Server by microservices

透過gRPC同時對多個應用程式發出通知

### Todo
- [x] Welcome msg
- [x] GRPC
- [x] Log
- [ ] Admin page 

### Support application
- [x] Discord
- [ ] Slack
- [x] Line
- [ ] Telegram

Work Explanation
---
---
[![](https://mermaid.ink/img/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cdHJlY3QgcmdiKDIxMCwgMjU1LCAyNDApXG5cdE5vdGUgcmlnaHQgb2YgU2VydmVyOiBSZWdpc3RlcmVkXG5cdFNlcnZlci0-PitOb3RpZnlfU2VydmVyOiBBcHBsaWNhdGlvblxuXHROb3RpZnlfU2VydmVyLS0-PitTZXJ2ZXI6IFRva2VuXG5cdGVuZFxuXHRcblx0U2VydmVyLT4-K05vdGlmeV9TZXJ2ZXI6IFBvc3QgTXNnXG5cdE5vdGlmeV9TZXJ2ZXItLT4-K2FwcDFfc2VydmVyOiBQb3N0IEV2ZW50XG5cdGFwcDFfc2VydmVyLS0-PitOb3RpZnlfU2VydmVyOiBSZXNwIHN0YXR1c1xuICAgIE5vdGlmeV9TZXJ2ZXItLT4-K2FwcDJfc2VydmVyOiBQb3N0IEV2ZW50XG5cdGFwcDJfc2VydmVyLS0-PitOb3RpZnlfU2VydmVyOiBSZXNwIHN0YXR1c1xuXHROb3RpZnlfU2VydmVyLS0-PitTZXJ2ZXI6IHJldHVybiBzdGF0dXMiLCJtZXJtYWlkIjp7InRoZW1lIjoiZm9yZXN0In0sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cdHJlY3QgcmdiKDIxMCwgMjU1LCAyNDApXG5cdE5vdGUgcmlnaHQgb2YgU2VydmVyOiBSZWdpc3RlcmVkXG5cdFNlcnZlci0-PitOb3RpZnlfU2VydmVyOiBBcHBsaWNhdGlvblxuXHROb3RpZnlfU2VydmVyLS0-PitTZXJ2ZXI6IFRva2VuXG5cdGVuZFxuXHRcblx0U2VydmVyLT4-K05vdGlmeV9TZXJ2ZXI6IFBvc3QgTXNnXG5cdE5vdGlmeV9TZXJ2ZXItLT4-K2FwcDFfc2VydmVyOiBQb3N0IEV2ZW50XG5cdGFwcDFfc2VydmVyLS0-PitOb3RpZnlfU2VydmVyOiBSZXNwIHN0YXR1c1xuICAgIE5vdGlmeV9TZXJ2ZXItLT4-K2FwcDJfc2VydmVyOiBQb3N0IEV2ZW50XG5cdGFwcDJfc2VydmVyLS0-PitOb3RpZnlfU2VydmVyOiBSZXNwIHN0YXR1c1xuXHROb3RpZnlfU2VydmVyLS0-PitTZXJ2ZXI6IHJldHVybiBzdGF0dXMiLCJtZXJtYWlkIjp7InRoZW1lIjoiZm9yZXN0In0sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)

QuickStart
---
---

 須先申請Bot相關帳號資料
- [Discord](https://discord.com/developers/docs/intro)
- [Line](https://account.line.biz/login?scope=line&redirectUri=https%3A%2F%2Fdevelopers.line.biz%2Fconsole%2F)

Downland file

```bash
vim config/demo.toml
go run notify-server.go
```

Licenses
---

This program is under the terms of the MIT License. See [LICENSE](/LICENSE) for the full license text.