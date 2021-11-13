package main

import (
	Node "DisysExercise2/NodePack"
	"fmt"
	"math/rand"

	//"bufio"
	"context"
	//"fmt"
	"log"

	"google.golang.org/grpc"
)

var CritAccess int32
	
func main() {
	//connect to grpc server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to gRPC server :: %v", err)
	}
	defer conn.Close()

	//call ChatService to create a stream
	client := Node.NewNodeClient(conn)

	AskForPermission(client)

	//message stream
	
	fmt.Println("Client Node: ", CritAccess)
}
//Need to update clientConfig so it doesnt get random numbers.
 
func clientConfig(){

}

func AskForPermission(client Node.NodeClient) {
	nodeid := rand.Intn(1e3)
	ServerPermission, err := client.Permission(context.Background(), &Node.RequestPermission{Nodeid: int32(nodeid)})
	log.Printf("Client %v has asked for permission \n", nodeid)
	
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}
	
	if(ServerPermission != nil){
		AccesCrit, err := client.AccesCrit(context.Background(), &Node.GoIntoCrit{GoIntoCrit: "Client going into crit"})
		log.Printf("Client %v has accessed the crit section \n", nodeid)
		CritAccess++
		if err != nil {
			log.Fatalf("Failed to call ChatService :: %v", err)
		}
		if(AccesCrit != nil){
			ExitCrit, err := client.ExitCrit(context.Background(), &Node.ReleaseToken{Nodeid: int32(nodeid)})
			log.Printf("Client %v has exited the crit section \n", nodeid)
			if err != nil {
				log.Fatalf("Failed to call ChatService :: %v", err)
			}
			if (ExitCrit != nil){
				AskForPermission(client)
			}
		}
	}
}