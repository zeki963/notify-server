package main

import (
	"context"
	"log"

	pb "github.com/zorhayashi/notify-server/grpc"
	"google.golang.org/grpc"
)

func test() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	c := pb.NewPostMsgClient(conn)

	r, err := c.PostMsg(context.Background(), &pb.MsgRequest{Msg: "HI,我是GRPC "})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%s , 時間:%d", r.Msg, r.Unixtime)

}
