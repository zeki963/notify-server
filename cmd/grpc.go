package cmd

import (
	"log"
	"net"
	"time"

	pb "github.com/zorhayashi/notify-server/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//PostMsgServer 12
type PostMsgServer struct{}

//PostMsg 12
func (e *PostMsgServer) PostMsg(ctx context.Context, req *pb.MsgRequest) (resp *pb.MsgReply, err error) {

	log.Printf("receive client request, client send:%s\n", req.Msg)
	return &pb.MsgReply{
		Msg:      req.Msg,
		Unixtime: time.Now().Unix(),
	}, nil

}

//GrpcServer 12
func GrpcServer() {
	apiListener, err := net.Listen("tcp", ":9999")
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
