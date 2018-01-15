package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRPC/Add/proto"
	"fmt"
	"time"
)



func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddClient(conn)
	stream, err := c.AddNumbers(context.Background())
	for i:=0;i<100;i++{
		time.Sleep(time.Second*1)
		fmt.Println("Sending:",i)
		err = stream.Send(&pb.Request{int64(i)})
		if err != nil {
			log.Fatalf("could not send data to Addition server: %v", err)
		}
		ret,_ := stream.Recv()
		fmt.Println("Recieved Sum:",ret.Sum)
	}
}
