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


type session struct {
	conn *grpc.ClientConn `json:"conn"`
	stream pb.Add_AddNumbersClient `json:"stream"`
	err error `json:"err"`
}

var s session

func main(){
	// Set up a connection to the server.
	ip := os.Getenv("SERVER_IP")
	fmt.Println("Connecting to "+ip+":50051 ... ")
	_conn, err := grpc.Dial(ip+":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Connection failed!")
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Connection was successfull")
	defer _conn.Close()
	s.conn = _conn

	// Spawn the client
	fmt.Println("Spawning client....")
	c := pb.NewAddClient(s.conn)
	_stream, err := c.AddNumbers(context.Background())
	s.stream = _stream
	s.err = err
	if err != nil{
		log.Fatalf("Failed to spawn the client")
	}
	fmt.Println("Client spawned successfully")
	
	// Test the server
	_stream = s.stream
	_err := s.err
	var sum int64
	for i:=0;i<5;i++{
		time.Sleep(time.Second*1)
		fmt.Println("Sending:",i)
		sum += int64(i)
		fmt.Println("Expected:",sum)
		_err = _stream.Send(&pb.Request{int64(i)})
		if _err != nil {
			log.Fatalf("could not send data to Addition server: %v", _err)
		}
		ret,_ := _stream.Recv()
		fmt.Println("Recieved Sum:",ret.Sum)
		if sum != ret.Sum{
			log.Fatal("Incorrect response")
		}
	}

fmt.Println("Server test passed")
}

