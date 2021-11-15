package main

import (
	Node "DisysExercise2/NodePack/proto"
	"fmt"

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

	clientConfig(conn)

	

	//message stream
	
	fmt.Println("Client Node: ", CritAccess)
}
//Need to update clientConfig so it doesnt get random numbers.
 
//sets name for client and status
func clientConfig(conn *grpc.ClientConn) {

	//reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to Chitty Chat! \n")
	fmt.Printf("Enter your id : ")
	var input int
	_, err := fmt.Scan(&input)
	fmt.Printf("Your id is: %v", input)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}

	
	//call ChatService to create a stream
	clientObj := client {clientid: int64(input), nodeClient: Node.NewNodeClient(conn)}
	

	clientObj.AskForPermission()
	//ch.clientName = strings.Trim(name, "\r\n")
	//ch.SendStatus()
}
type client struct{
	clientid int64
	nodeClient Node.NodeClient
}

func (client *client) AskForPermission() {
	
	ServerPermission, err := client.nodeClient.Permission(context.Background(), &Node.RequestPermission{Nodeid: int32(client.clientid)})
	log.Printf("Client %v has asked for permission \n", client.clientid)
	
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}
	
	if(ServerPermission != nil){
		AccesCrit, err := client.nodeClient.AccesCrit(context.Background(), &Node.GoIntoCrit{Nodeid: int32(client.clientid)})
		log.Printf("Client %v has accessed the crit section \n", client.clientid)
		CritAccess++
		if err != nil {
			log.Fatalf("Failed to call ChatService :: %v", err)
		}
		if(AccesCrit != nil){
			ExitCrit, err := client.nodeClient.ExitCrit(context.Background(), &Node.ReleaseToken{Nodeid: int32(client.clientid)})
			log.Printf("Client %v has exited the crit section \n", client.clientid)
			if err != nil {
				log.Fatalf("Failed to call ChatService :: %v", err)
			}
			if (ExitCrit != nil){
				client.AskForPermission()
			}
		}
	}
}
// func (client *client) EnterCrit(){
// 	AccesCrit, err := client.nodeClient.AccesCrit(context.Background(), &Node.GoIntoCrit{GoIntoCrit: "Client going into crit"})
// 		log.Printf("Client %v has accessed the crit section \n", client.clientid)
// 		CritAccess++
// 		if err != nil {
// 			log.Fatalf("Failed to call ChatService :: %v", err)
// 		}
// }

// func (client *client) ExitCrit(){
// 	ExitCrit, err := client.nodeClient.ExitCrit(context.Background(), &Node.ReleaseToken{Nodeid: int32(client.clientid)})
// 			log.Printf("Client %v has exited the crit section \n", client.clientid)
// 			if err != nil {
// 				log.Fatalf("Failed to call ChatService :: %v", err)
// 			}
// }
