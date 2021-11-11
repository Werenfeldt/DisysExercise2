package main

import (
	context "context"

	"DisysExercise2/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

//var MeNode node

const (
	name = "node"
	port = ":8080"
)

type Node struct {
	api.UnimplementedNodeServer
	name string
}

func (n Node) Message(ctx context.Context, in *api.Request) (*api.Reply, error) {
	return &api.Reply{Reply: "Response right here"}, nil
}

func main() {

	log.Printf("Server started")

}

func NewNodeServer(name string) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", port)
	grpcServer := grpc.NewServer()
	NodeServer, err := NewNodeServer(name)
	if err != nil {
		return nil, err
	}
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	api.RegisterNodeServer(grpcServer, *NodeServer)

	return grpcServer, nil
}

// type InitNode struct {

// }

// func MakeNode(id int, totalNodes int) node{
// 	MeNode = node{
// 		id: id,
// 		time: 0,
// 		state: 0,
// 		numberTotalNode: totalNodes,
// 	}
// 	return MeNode

// }

// func InitServer(){
// 	Port := os.Getenv("PORT")
// 	if Port == "" {
// 		Port = "8080" //default Port set to 8080 if PORT is not set in env
// 	}

// 	//init listener
// 	listen, err := net.Listen("tcp", ":"+Port)
// 	if err != nil {
// 		log.Fatalf("Could not listen @ %v :: %v", Port, err)
// 	}
// 	grpcserver := grpc.NewServer()
// 	log.Printf("new grpc server - listen on: %v", listen )

// 	RegisterNodeServer(grpcserver, &InitNode{})
// 	fmt.Println("InitNode")

// 	// err = grpcserver.Serve(listen)
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to start gRPC Server :: %v", err)
// 	// }
// 	fmt.Println("Starting Client")
// 	InitClient()

// 	if err := grpcserver.Serve(listen); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// 	fmt.Println("Listening")
// }

// func InitClient(){
// 	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

// 	if err != nil {
// 		log.Fatalf("Failed to connect to gRPC server :: %v", err)
// 	} else {
// 		log.Println("Dialed localhost:8080")
// 	}
// 	defer conn.Close()

// 	nodeClient := NewNodeClient(conn)

// 	ctx := context.Background()
// 	//defer cancel()
// 	reply, err := nodeClient.Message(ctx, &Request{Nodeid: 1, Time: 1})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}

// 	log.Println(reply)
// }

// type node struct{
// 	id 		int
// 	time 	int
// 	state 	State
// 	numberTotalNode int
// 	//partici []node
// }

// type State int

// const (
// 	Released  State = iota 	//Released har index 0
// 	Wanted 					//Wanted har index 1
// 	Held					//Held har index 2
// )
