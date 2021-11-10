package main

import (
	context "context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var MeNode node

func main(){
	fmt.Println("Starting Server")
	InitServer()
	// fmt.Println("Starting Client")
	// InitClient()
}

type InitNode struct {
	
}

func MakeNode(id int, totalNodes int) node{
	MeNode = node{
		id: id,
		time: 0,
		state: 0,
		numberTotalNode: totalNodes, 
	}
	return MeNode
	
}

func InitServer(){
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080" //default Port set to 8080 if PORT is not set in env
	}

	//init listener
	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", Port, err)
	}
	grpcserver := grpc.NewServer()
	log.Printf("new grpc server - listen on: %v", listen )

	RegisterNodeServer(grpcserver, &InitNode{})
	fmt.Println("InitNode")

	// err = grpcserver.Serve(listen)
	// if err != nil {
	// 	log.Fatalf("Failed to start gRPC Server :: %v", err)
	// }
	fmt.Println("Starting Client")
	InitClient()
	
	if err := grpcserver.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("Listening")
}


func InitClient(){
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to gRPC server :: %v", err)
	} else {
		log.Println("Dialed localhost:8080")
	}
	defer conn.Close()

	nodeClient := NewNodeClient(conn)


	ctx := context.Background()
	//defer cancel()
	reply, err := nodeClient.Message(ctx, &Request{Nodeid: 1, Time: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	
	log.Println(reply)
}


func (c *InitNode) Message(ctx context.Context, in *Request) (*Reply, error) {
	return &Reply{Reply: "Response right here"}, nil
}

type node struct{
	id 		int 
	time 	int 	
	state 	State	
	numberTotalNode int 
	//partici []node
}

type State int 

const (
	Released  State = iota 	//Released har index 0
	Wanted 					//Wanted har index 1
	Held					//Held har index 2
)


