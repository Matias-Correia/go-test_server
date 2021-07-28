package main

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	pb "github.com/Matias-Correia/go-test_server/server/protologs"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedLogTestDataServer
	database sql.DB
}

// LogTestData implements protologs.LogTestData
func (s *server) SendLogs(ctx context.Context, in *pb.Log) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetBlockID())
	log.Printf("Was delivered?: %v", in.GetBlockDelivered())
	log.Printf("Request Delay: %v", in.GetRequestDelay())
	log.Printf("Block Delay: %v", in.GetBlockDelay())

	recordReceivedLogs(s.database, in.GetBlockID(), in.GetRequestDelay(), in.GetBlockDelay(), in.GetBlockDelivered())

	return &(pb.Empty{}), nil
}

// Save info in the BD
func recordReceivedLogs( dbCon sql.DB, blockId string, requestDelay int64, blockDelay int64, blockDelivered bool){
	query := fmt.Sprintf("INSERT INTO test_logs (BlcokID,RequestDelay,BlockDelay,BlockDelivered) VALUES('%s',%v,%v,%t)",blockId,requestDelay,blockDelay,blockDelivered)
	fmt.Printf(query)
	insert, err := dbCon.Query(query)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }

    defer insert.Close()

}

func openDataBaseConnection() sql.DB{
	db, err := sql.Open("mysql", "admin:raiz@(db:3306)/logs")
	if err != nil {
  		panic(err)
	}
	return *db
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db := openDataBaseConnection()
	s := grpc.NewServer()
	pb.RegisterLogTestDataServer(s, &server{database: db})
	if err := s.Serve(lis); err != nil {
		defer db.Close()
		log.Fatalf("failed to serve: %v", err)
	}
}
