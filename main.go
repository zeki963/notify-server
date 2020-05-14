package main

import (
	"github.com/zorhayashi/notify-server/cmd"
	//"github.com/zorhayashi/notify-server/discord"
	//"github.com/zorhayashi/notify-server/util"
)

func main() {
	//xMsg := util.GetLineMsg()
	//discord.Post(xMsg)
	cmd.Print()
	cmd.GrpcServer()
}
