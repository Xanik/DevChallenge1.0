package storage

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/Xanik/DevChallenge1.0/GeneratedProtobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"sync"
)


func newStorageService(server pb.StorageServiceServer) {
	port := ":3030"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterStorageServiceServer(s, server)

	log.Println("Starting Server On Port:" + port)

	e := s.Serve(lis)
	if e != nil {
		log.Fatalf("failed to Serve: %v", e)
	}
}

type storage struct {
	mutex *sync.RWMutex
	state map[int64]interface{}
}

type response struct {
	Message string
	Value   float32
}

func newStorage() *storage {
	return &storage{mutex: &sync.RWMutex{}, state: make(map[int64]interface{})}
}

func (s storage) generateID() int64 {
	return int64(len(s.state) + 1)
}
func (s storage) Store(ctx context.Context, m *pb.StorageRequest) (*pb.StorageResponse, error) {

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata Received: %v\n", md)
	}

	key := s.generateID()

	s.mutex.Lock()
	s.state[key] = response{
		Message: m.Message,
		Value:   m.Value,
	}
	s.mutex.Unlock()

	return &pb.StorageResponse{
		Id:      key,
		Message: m.Message,
		Value:   m.Value,
	}, nil
}

func (s storage) Read(ctx context.Context, m *pb.GetByID) (*pb.StorageResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata Received: %v\n", md)
	}

	s.mutex.RLock()
	fmt.Println(s.state[m.Id])
	data, ok := s.state[m.Id]
	s.mutex.RUnlock()

	if !ok {
		return nil, errors.New("Data not found")
	}

	fmt.Printf("%T here", data)
	d, ok := data.(response)

	if !ok {
		return nil, errors.New("Data is not of type response")
	}

	return &pb.StorageResponse{
		Id:      m.Id,
		Message: d.Message,
		Value:   d.Value,
	}, nil
}

func (s storage) GetAll(ctx context.Context,m *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	s.mutex.RLock()

	d := s.state

	s.mutex.RUnlock()

	var data []*pb.StorageResponse

	for key, value := range d {
		v := value.(response)
		s := &pb.StorageResponse{
			Id:      key,
			Message: v.Message,
			Value:   v.Value,
		}
		data = append(data, s)
	}
	return &pb.GetAllResponse{
		Responses:            data,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

func (s storage) Update(ctx context.Context, m *pb.UpdateRequest) (*pb.StorageResponse, error) {

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata Received: %v\n", md)
	}

	s.mutex.RLock()
	_, ok := s.state[m.Id]
	s.mutex.RUnlock()

	if !ok {
		return nil, errors.New("Data not found")
	}

	s.mutex.Lock()
	s.state[m.Id] = response{
		Message: m.Message,
		Value:   m.Value,
	}
	s.mutex.Unlock()

	s.mutex.RLock()
	data, ok := s.state[m.Id]
	s.mutex.RUnlock()

	d, ok := data.(response)

	if !ok {
		return nil, errors.New("Data is not of type response")
	}

	return &pb.StorageResponse{
		Id:      m.Id,
		Message: d.Message,
		Value:   d.Value,
	}, nil
}

func (s storage) Delete(ctx context.Context, m *pb.GetByID) (*pb.StorageResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata Received: %v\n", md)
	}

	s.mutex.RLock()
	_, ok := s.state[m.Id]
	s.mutex.RUnlock()

	if !ok {
		return nil, errors.New("Data not found")
	}

	s.mutex.Lock()

	delete(s.state, m.Id)
	s.mutex.Unlock()
	return &pb.StorageResponse{
		Id: m.Id,
	}, nil
}
