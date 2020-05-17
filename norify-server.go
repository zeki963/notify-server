package main

import (
	"github.com/zorhayashi/notify-server/cmd"
)

func main() {
	//export LC_CTYPE="en_US.UTF-8"
	cmd.GrpcServer()
}
