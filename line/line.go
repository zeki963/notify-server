package line

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	channelSecret      = "123"
	channelAccessToken = "foo"
)

func init() {
	bot, err := linebot.New(channelSecret, channelAccessToken)
}

var messages []linebot.SendingMessage

//PostLine 傳送訊息
func PostLine() {
	_, err := bot.PushMessage(ID, messages...).Do()
	if err != nil {
		//Do something when some bad happened
		fmt.Println(error(err))
	}
}
