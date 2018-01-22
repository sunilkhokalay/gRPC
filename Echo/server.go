package main

import (
	"log"
	"net"
	"fmt"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRPC/Echo/proto"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)


type server struct{}


func (s *server) EchoBack(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("Message recieved from client:"+in.MsgRequest)
	fmt.Println("Reply to client: "+in.MsgRequest)
	return &pb.EchoResponse{MsgReponse: in.MsgRequest}, nil
}

func main() {
	host,_ := os.Hostname()
	addrs,_ := net.LookupIP(host)
	for _,addr := range addrs{
	if ipv4 := addr.To4(); ipv4 != nil{
	fmt.Println("IPv4: ",ipv4)
	}}

	fmt.Print("Starting Echo Server on port 50051...\n")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
