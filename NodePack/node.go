package NodePack

import (
	context "context"
	"log"
	pb "DisysExercise2/NodePack/proto"
	//"fmt"
	// "log"
	// "net"
	// "os"
	//"google.golang.org/grpc"
	"sync"
	"time"
	
	
)

var MeNode node

type Nodeserver struct{
	pb.UnimplementedNodeServer
}

	
type CriticalSection struct {
	CritSection int32 
	mu sync.Mutex
}
var CritObject = CriticalSection{}

// func MakeNode(id int, totalNodes int) node{
// 	MeNode = node{
// 		id: id,
// 		time: 0,
// 		state: 0,
// 		numberTotalNode: totalNodes, 
// 	}
// 	return MeNode
	
// }

func (c *Nodeserver) Permission(ctx context.Context, in *pb.RequestPermission) (*pb.GivePermission, error) {
	
	//c.AddClient(client{ClientUniqueCode: in.Nodeid})

	
	// if(len(clientObject.CQue)==0){
		log.Printf("Node %v has asked for permission \n", in.Nodeid)
		CritObject.mu.Lock()
		log.Printf("The CritSection is locked and Node %v has permission \n", in.Nodeid)
		return &pb.GivePermission{GivePermission: "You have permission"}, nil
		
	// 	fmt.Println("Crit Locked - Node Crit: ", CritObject.CritSection);
	// } else {
	// 	clientObject.CQue = append(clientObject.CQue, client{ClientUniqueCode: in.Nodeid})
	// }
}

func (c *Nodeserver) AccesCrit(ctx context.Context, in *pb.GoIntoCrit) (*pb.ServerDoneInCrit, error) {
	
	CritObject.CritSection++
	log.Printf("The CritSection has been accessed by %v and incremented \n", in.Nodeid)
	return &pb.ServerDoneInCrit{ServerDoneInCrit: "The Server is done"}, nil	
}
//not used
func (c *Nodeserver) ExitCrit(ctx context.Context, in *pb.ReleaseToken) (*pb.Empty, error) {
	time.Sleep(5 * time.Second)
	CritObject.mu.Unlock()
	log.Printf("The CritSection has been unlocked \n")
	return &pb.Empty{}, nil	
}


func (is *Nodeserver) AddClient(client client) {
		clientObject.mu.Lock()
		clientObject.CQue = append(clientObject.CQue, client)
		clientObject.mu.Unlock()	
}


//Structs
type client struct {
	ClientUniqueCode int32
}

type que struct {
	CQue []client
	mu   sync.Mutex
}

var clientObject = que{}

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


/*
Client sender request
Server checker om CS er i brug
Hvis ikke i brug, giv permission
Hvis i brug, send request i køen
Når client exiter, siger vi til server at CS ikke er i brug
Og server vælger den næste i køen og giver permission
*/