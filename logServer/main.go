package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "github.com/Matias-Correia/go-test_server/protologs"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedLogTestDataServer
}

// LogTestData implements protologs.LogTestData
func (s *server) LogTestData(ctx context.Context, in *pb.Log) (*pb.google.protobuf.Empty, error) {
	log.Printf("Received: %v", in.GetBlockID())
	log.Printf("Was delivered?: %v", in.GetBlockDelivered())
	log.Printf("With Delay: %v", in.GetRequestDelay())
	
	RecordReceivedLogs(in.GetBlockID, in.GetRequestDelay, in.GetBlockDelivered)

	return *emptypb.Empty, nil
}

// Save info in the BD
func RecordReceivedLogs(blockId string, delay time.Duration, bool blockDelivered){
	
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
