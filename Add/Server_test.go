package main

import (
	"testing"
	"log"
	"time"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRPC/Add/proto"
)


type session struct {
	conn *grpc.ClientConn
	stream pb.Add_AddNumbersClient
	err error
}

var s session

func TestConnection(t *testing.T){
	fmt.Println("--------------------------------")
	fmt.Println(s)
	fmt.Println("--------------------------------")
	// Set up a connection to the server.
  fmt.Println("Enter server IP:")
	reader := bufio.NewReader()
	ip,_ := reader.ReadString('\n')
	ip = ip[:len(ip)-1]
	fmt.Println("Connecting to "+ip+":50051 ... ")
	conn, err := grpc.Dial(ip+":50051", grpc.WithInsecure())
	if err != nil {
		t.Fail()
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("Connection was successfull")
	defer conn.Close()
	s.conn = conn
	fmt.Println("===========================")
	fmt.Println(s)
	fmt.Println("===========================")
}

func TestSpawnNewClient(t *testing.T)  {
	fmt.Println("--------------------------------")
	fmt.Println(s)
	fmt.Println("--------------------------------")
	fmt.Println("Spawning client....")
	c := pb.NewAddClient(s.conn)
	_stream, err := c.AddNumbers(context.Background())
	s.stream = _stream
	s.err = err
	fmt.Println("Client spawned successfully")
	fmt.Println("===========================")
	fmt.Println(s)
	fmt.Println("===========================")
}

func TestServer(t *testing.T) {
	fmt.Println("--------------------------------")
	fmt.Println(s)
	fmt.Println("--------------------------------")
	_stream := s.stream
	_err := s.err
	for i:=0;i<100;i++{
		time.Sleep(time.Second*1)
		fmt.Println("Sending:",i)
		_err = _stream.Send(&pb.Request{int64(i)})
		if _err != nil {
			log.Fatalf("could not send data to Addition server: %v", _err)
		}
		ret,_ := _stream.Recv()
		fmt.Println("Recieved Sum:",ret.Sum)
	}


	}

