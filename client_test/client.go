package main

import (
	"context"
	"log"
	"fmt"
	pb "github.com/zorhayashi/notify-server/grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

    fmt.Print("Enter text: ")
    var postmsg string
    fmt.Scanln(&postmsg)

	c := pb.NewPostMsgClient(conn)

	r, err := c.PostMsg(context.Background(), &pb.MsgRequest{Msg: postmsg})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%s , 時間:%d", r.Msg, r.Unixtime)

}
