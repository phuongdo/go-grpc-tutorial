package main

import (
	"flag"

	"github.com/golang/glog"
	pb "github.com/go-grpc-tutorial/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

// Implements of EchoServiceServer

type echoServer struct{}

func newEchoServer() pb.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	glog.Info(msg)
	return msg, nil
}

func Run() error {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, newEchoServer())
	server.Serve(listen)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}