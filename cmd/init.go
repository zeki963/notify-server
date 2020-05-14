package cmd

import (
	"fmt"

	util "github.com/zorhayashi/notify-server/util"
)

var (
	version = "0.0.1"
)

func init() {

}

//Print init
func Print() {
	xMsg := util.GetLineMsg()
	logo := fmt.Sprintf(` 	     _   _  __                                          
 _ __   ___ | |_(_)/ _|_   _       ___  ___ _ ____   _____ _ __ 
| '_ \ / _ \| __| | |_| | | |_____/ __|/ _ \ '__\ \ / / _ \ '__|
| | | | (_) | |_| |  _| |_| |_____\__ \  __/ |   \ V /  __/ |   
|_| |_|\___/ \__|_|_|  \__, |     |___/\___|_|    \_/ \___|_|   
             	        |___/               - version : %s -
----------------------------------------------------------------
%s`, version, xMsg)
	fmt.Println(logo)
	fmt.Println()
}
