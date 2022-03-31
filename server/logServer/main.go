package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Matias-Correia/go-test_server/server/protologs"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	port = ":50051"
)

// server is used to implement protologs.server.
type server struct {
	pb.UnimplementedLogTestDataServer
	database sql.DB
}

// SendLogs implements protologs.SendLogs
func (s *server) SendLogs(ctx context.Context, in *pb.Log) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetBlockID())
	log.Printf("Sent By: %v", in.GetLocalpeer())
	log.Printf("Received By: %v", in.GetRemotepeer())
	log.Printf("SentAt: %v", in.GetSentAt())
	log.Printf("ReceivedAt: %v", in.GetReceivedAt())
	log.Printf("BlockRequestedAt: %v", in.GetBlockRequestedAt())
	log.Printf("Duplicate: %v", in.GetDuplicate())

	recordReceivedLogs(s.database, in.GetBlockID(), in.GetLocalpeer(), in.GetRemotepeer(), in.GetSentAt(), in.GetReceivedAt(), in.GetBlockRequestedAt(), in.GetDuplicate())

	return &(pb.Empty{}), nil
}

// Save the received logs in the BD
func recordReceivedLogs(dbCon sql.DB, blockId string, localpeer string, remotepeer string, sentAt *timestamppb.Timestamp, receivedAt *timestamppb.Timestamp, blockRequestedAt *timestamppb.Timestamp, duplicate bool) {

	//convert timestamps to DB format with milliseconds
	sentAtAsTime := fmt.Sprintf("%v", sentAt.AsTime().Format("2006-01-02 15:04:05.000"))
	receivedAtAsTime := fmt.Sprintf("%v", receivedAt.AsTime().Format("2006-01-02 15:04:05.000"))
	blockRequestedAtAsTime := fmt.Sprintf("%v", blockRequestedAt.AsTime().Format("2006-01-02 15:04:05.000"))

	query := fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,SentAt,ReceivedAt,BlockRequestedAt,Duplicate) VALUES('%s','%s','%s','%v','%v','%v',%t)", blockId, localpeer, remotepeer, sentAtAsTime, receivedAt, blockRequestedAt, duplicate)

	if sentAt == nil {
		if receivedAt == nil {
			if blockRequestedAt == nil {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,Duplicate) VALUES('%s','%s','%s',%t)", blockId, localpeer, remotepeer, duplicate)
			} else {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,BlockRequestedAt,Duplicate) VALUES('%s','%s','%s','%v',%t)", blockId, localpeer, remotepeer, blockRequestedAtAsTime, duplicate)
			}
		} else {
			if blockRequestedAt == nil {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,ReceivedAt,Duplicate) VALUES('%s','%s','%s','%v',%t)", blockId, localpeer, remotepeer, receivedAtAsTime, duplicate)
			} else {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,ReceivedAt,BlockRequestedAt,Duplicate) VALUES('%s','%s','%s','%v','%v',%t)", blockId, localpeer, remotepeer, receivedAtAsTime, blockRequestedAtAsTime, duplicate)
			}
		}
	} else {
		if receivedAt == nil {
			if blockRequestedAt == nil {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,SentAt,Duplicate) VALUES('%s','%s','%s','%v',%t)", blockId, localpeer, remotepeer, sentAtAsTime, duplicate)
			} else {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,SentAt,BlockRequestedAt,Duplicate) VALUES('%s','%s','%s','%v','%v',%t)", blockId, localpeer, remotepeer, sentAtAsTime, blockRequestedAtAsTime, duplicate)
			}
		} else {
			if blockRequestedAt == nil {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,SentAt,ReceivedAt,Duplicate) VALUES('%s','%s','%s','%v','%v',%t)", blockId, localpeer, remotepeer, sentAtAsTime, receivedAtAsTime, duplicate)
			} else {
				query = fmt.Sprintf("INSERT INTO test_logs (BlockID,LocalPeer,RemotePeer,SentAt,ReceivedAt,BlockRequestedAt,Duplicate) VALUES('%s','%s','%s','%v','%v','%v',%t)", blockId, localpeer, remotepeer, sentAtAsTime, receivedAtAsTime, blockRequestedAtAsTime, duplicate)
			}
		}
	}

	fmt.Printf(query)
	insert, err := dbCon.Query(query)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func openDataBaseConnection() sql.DB {
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
	db.SetMaxIdleConns(0)
	s := grpc.NewServer()
	pb.RegisterLogTestDataServer(s, &server{database: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer db.Close()
}
