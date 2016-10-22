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

// func startLocalCache () {
// 	startCache()
// }

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutResponse, error) {
	cachePutOne (in.Key, in.Value)
	return &pb.PutResponse{Key: in.Key, Success: true}, nil
}

func main() {
	//start()
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