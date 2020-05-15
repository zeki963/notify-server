# Notify-Server
## lazy like me (っ´ω`c)

![Logo](/logo-m.png)

微服務-通知訊息伺服器

Notify-Server by microservices

透過gRPC同時對多個應用程式發出通知

Support application
- Discord 
- Line

### Todo
- [x] Welcome msg
- [x] GRPC
- [x] Log
- [ ] Admin page 

### Application
- [x] Discord
- [ ] Slack
- [x] Line
- [ ] Telegram

Work Explanation
---
---
[![](https://mermaid.ink/img/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cdFNlcnZlci0-PitOb3RpZnlfU2VydmVyOiBQb3N0IE1zZ1xuXHROb3RpZnlfU2VydmVyLS0-PithcHAxX3NlcnZlcjogUG9zdCBFdmVudFxuXHRhcHAxX3NlcnZlci0tPj4rTm90aWZ5X1NlcnZlcjogUmVzcCBzdGF0dXNcbiAgICBOb3RpZnlfU2VydmVyLS0-PithcHAyX3NlcnZlcjogUG9zdCBFdmVudFxuXHRhcHAyX3NlcnZlci0tPj4rTm90aWZ5X1NlcnZlcjogUmVzcCBzdGF0dXNcblx0Tm90aWZ5X1NlcnZlci0tPj4rU2VydmVyOiByZXR1cm4gc3RhdHVzIiwibWVybWFpZCI6eyJ0aGVtZSI6ImZvcmVzdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cdFNlcnZlci0-PitOb3RpZnlfU2VydmVyOiBQb3N0IE1zZ1xuXHROb3RpZnlfU2VydmVyLS0-PithcHAxX3NlcnZlcjogUG9zdCBFdmVudFxuXHRhcHAxX3NlcnZlci0tPj4rTm90aWZ5X1NlcnZlcjogUmVzcCBzdGF0dXNcbiAgICBOb3RpZnlfU2VydmVyLS0-PithcHAyX3NlcnZlcjogUG9zdCBFdmVudFxuXHRhcHAyX3NlcnZlci0tPj4rTm90aWZ5X1NlcnZlcjogUmVzcCBzdGF0dXNcblx0Tm90aWZ5X1NlcnZlci0tPj4rU2VydmVyOiByZXR1cm4gc3RhdHVzIiwibWVybWFpZCI6eyJ0aGVtZSI6ImZvcmVzdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)

QuickStart
---
---

```bash
vim config/demo.toml
go run notify-server.go
```