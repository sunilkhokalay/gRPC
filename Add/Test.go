package main

import (
	"log"
	"time"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRPC/Add/proto"
	"os"
)


func main(){
	// Set up a connection to the server.
	ip := os.Getenv("SERVER_IP")
	fmt.Println("Connecting to "+ip+":50051 ... ")
	conn, err := grpc.Dial(ip+":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Connection failed!")
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Connection was successfull")
	defer conn.Close()


	// Spawn the client
	fmt.Println("Spawning client....")
	c := pb.NewAddClient(conn)
	stream, err := c.AddNumbers(context.Background())
	if err != nil{
		log.Fatalf("Failed to spawn the client")
	}
	fmt.Println("Client spawned successfully")
	
	// Test the server
	var sum int64
	for i:=0;i<5;i++{
		time.Sleep(time.Second*1)
		fmt.Println("Sending:",i)
		sum += int64(i)
		fmt.Println("Expected:",sum)
		err = stream.Send(&pb.Request{int64(i)})
		if err != nil {
			log.Fatalf("could not send data to Addition server: %v", err)
		}
		ret,_ := stream.Recv()
		fmt.Println("Recieved Sum:",ret.Sum)
		if sum != ret.Sum{
			log.Fatal("Incorrect response")
		}
	}

fmt.Println("Server test passed")
}
