package main 

import (
    "fmt"
    "log"
    "net"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/ychen11/godrone/godrone"
)

const (
  port = ":40001"
)

type server struct{}

// Implement Put 
func (s *server) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutResponse, error) {
  cachePutOne (in.Key, in.Value)
  return &pb.PutResponse{Key: in.Key, Success: true}, nil
}

// Implement GetOne
func (s *server) GetOne(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
  value, isSucceedded := cacheGetOne (in.Key)
  if isSucceedded == false {
    return &pb.GetResponse{Key: in.Key, Value: ""}, nil
  } else {
    return &pb.GetResponse{Key: in.Key, Value: value}, nil
  }
}


func main() {
  initCache()
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterCacheOpsServer(s, &server{})
  s.Serve(lis)
  fmt.Printf("This is a main file. \n")
}