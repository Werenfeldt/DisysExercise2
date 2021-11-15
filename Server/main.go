package main

import (
	"fmt"
	"log"
	"net"
	"os"

	Node "DisysExercise2/NodePack"
	proto "DisysExercise2/NodePack/proto"

	"google.golang.org/grpc"
)

func main() {
	//Init log file
	
	LOG_FILE := "../Node_log"

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	//log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	//Init port
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080" //default Port set to 5000 if PORT is not set in env
	}

	//Init listener
	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", Port, err)
	}
	log.Println("Listening @ : " + Port)
	fmt.Println("Listening @ : " + Port)

	//gRPC server instance
	grpcserver := grpc.NewServer()

	//register ChatService
	cs := Node.Nodeserver{}
	proto.RegisterNodeServer(grpcserver, &cs)

	//grpc listen and serve
	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	} else {

	}
}