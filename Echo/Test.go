package main

import (
	"log"
	"os"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRPC/Echo/proto"
)


func main(){

	ip := os.Getenv("SERVER_IP")
	address := ip+":50051"
	fmt.Println("Dialing to server "+address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to dial to the server")
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()
	fmt.Println("Dialing succeeded to server")

	c := pb.NewEchoClient(conn)

	// Contact the server and print out its response.
	requestMsg := "Hi"
	fmt.Println("Message sent to server: "+requestMsg)
	r, err := c.EchoBack(context.Background(), &pb.EchoRequest{MsgRequest: requestMsg})
		if err != nil {
			log.Fatalf("could not send message to server: %v", err)
		}
	

	log.Printf("Server response :%s", r.MsgReponse)
	

	if r.MsgReponse != requestMsg{
		log.Fatal("Server Test failed")
	}
fmt.Println("Server test passed")
}

