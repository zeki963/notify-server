package discord

import (
	"net/http"
	"strings"

	"github.com/zorhayashi/notify-server/config"
)

//Post discord webhooks
func Post(msg string) {
	resp, err := http.Post(config.Global.Discord.WebhookLink,
		"application/x-www-form-urlencoded",
		strings.NewReader("content="+msg))

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}
