package cmd

import (
	"log"
	"net"
	"time"

	"github.com/zorhayashi/notify-server/config"
	pb "github.com/zorhayashi/notify-server/grpc"
	"github.com/zorhayashi/notify-server/supapp"
	"github.com/zorhayashi/notify-server/util"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//PostMsgServer type
type PostMsgServer struct{}

//PostMsg 傳送訊息
func (e *PostMsgServer) PostMsg(ctx context.Context, req *pb.MsgRequest) (resp *pb.MsgReply, err error) {
	util.Info("receive client request, client send: " + req.Msg)

	supapp.DiscordPost(req.Msg)

	return &pb.MsgReply{
		Msg:      req.Msg,
		Unixtime: time.Now().Unix(),
	}, nil
}

//GrpcServer 主Server
func GrpcServer() {
	apiListener, err := net.Listen("tcp", ":9999"+config.Global.po)
	if err != nil {
		log.Println(err)
		return
	}
	// 註冊 grpc
	es := &PostMsgServer{}
	grpc := grpc.NewServer()
	//pb.Re(grpc, es)
	pb.RegisterPostMsgServer(grpc, es)
	reflection.Register(grpc)
	if err := grpc.Serve(apiListener); err != nil {
		log.Fatal(" grpc.Serve Error: ", err)
		return
	}
}
