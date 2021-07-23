package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "github.com/Matias-Correia/go-test_server/protologs"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedLogTestDataServer
}

// LogTestData implements protologs.LogTestData
func (s *server) SendLogs(ctx context.Context, in *pb.Log) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetBlockID())
	log.Printf("Was delivered?: %v", in.GetBlockDelivered())
	log.Printf("Request Delay: %v", in.GetRequestDelay())
	log.Printf("Block Delay: %v", in.GetBlockDelay())

	RecordReceivedLogs(in.GetBlockID(), time.Duration(in.GetRequestDelay()) * time.Millisecond, in.GetBlockDelivered())

	return &(pb.Empty{}), nil
}

// Save info in the BD
func RecordReceivedLogs(blockId string, delay time.Duration, blockDelivered bool){
	
}



func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogTestDataServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
