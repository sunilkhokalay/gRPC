package main

import (
	pb "gRPC/Add/proto"
	"google.golang.org/grpc"
	"fmt"
	"io"
	"flag"
	"net"
	"os"
)

type addServer struct{
}

func (s *addServer)AddNumbers(stream pb.Add_AddNumbersServer)  error{
	fmt.Println("Entering Add numbers method")
	var sum int64
	defer func() {fmt.Println("Leaving AddNumbers method")}()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Got EOF from client .. Closing the connection..")
			break
		} else if err != nil {
			return err
		}
		fmt.Printf("Received: %d\n", req.Num)
		sum += req.Num
		stream.Send(&pb.Response{sum})


	}


	return nil
}

func newAddServer()  pb.AddServer{
	return &addServer{}
}

func main() {
	host,_ := os.Hostname()
	addrs,_ := net.LookupIP(host)
	for _,addr := range addrs{
	if ipv4 := addr.To4(); ipv4 != nil{
	fmt.Println("IPv4: ",ipv4)
	}}
	port := flag.Int("port", 50051, "Port for the server to run on")
	flag.Parse()
	fmt.Printf("Starting Addition Server on port %d...\n", *port)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAddServer(grpcServer, newAddServer())
	grpcServer.Serve(conn)
}
