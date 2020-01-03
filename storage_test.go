package storage

import (
	"context"
	pb "github.com/Xanik/DevChallenge1.0/GeneratedProtobuf"
	"google.golang.org/grpc"
	"log"
	"testing"
)

// Initialize Server Service
func init() {
	go newStorageService(newStorage())

}

// New Storage Client Connection Created
func newClient() pb.StorageServiceClient {
	port := ":3030"

	conn, err := grpc.Dial("localhost" + port,  grpc.WithInsecure())


	if err != nil {
		log.Fatalf("failed to Serve: %v", err)
	}


	client := pb.NewStorageServiceClient(conn)
	return client

}

// TestStore Function
func TestStore(t *testing.T) {

	res, err := newClient().Store(context.Background(), &pb.StorageRequest{
		Message:              "Hello From Client",
		Value:                12.50,
	})
	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}

// TestRead Function
func TestRead(t *testing.T) {

	res, err := newClient().Read(context.Background(), &pb.GetByID{
		Id:                   1,
	})

	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}

// TestGetAll Function
func TestGetAll(t *testing.T) {
	res, err := newClient().GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}

// TestUpdate Function
func TestUpdate(t *testing.T) {
	res, err := newClient().Update(context.Background(), &pb.UpdateRequest{
		Id:                   1,
		Message:              "Client Loading",
		Value:                5.34,
	})
	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}

// TestDelete Function
func TestDelete(t *testing.T) {
	res, err := newClient().Delete(context.Background(), &pb.GetByID{
		Id:                   1,
	})
	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}