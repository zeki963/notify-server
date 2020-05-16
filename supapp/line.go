package supapp

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/zorhayashi/notify-server/config"
	"github.com/zorhayashi/notify-server/util"
)

var (
	iD = "123"
)

var messages []linebot.SendingMessage

//LinePost 傳送訊息
func LinePost() {

	bot, err := linebot.New(config.Global.Line.ChannelSecret, config.Global.Line.ChannelAccessToken)
	_, err = bot.PushMessage(iD, messages...).Do()
	if err != nil {
		//Do something when some bad happened
		util.Error(err.Error())
	}
}
