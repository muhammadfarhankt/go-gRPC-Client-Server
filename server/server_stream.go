// package main

// import (
// 	"log"
// 	"time"

// 	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
// )

// func (s *helloServer) SayHelloServerStream(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
// 	log.Printf("server received name requests: %v", req.Names)
// 	for _, name := range req.Names {
// 		res := &pb.HelloResponse{
// 			Message: "Hello, " + name,
// 		}
// 		if err := stream.Send(res); err != nil {
// 			return err
// 		}
// 		time.Sleep(time.Second * 2)
// 	}
// 	return nil
// }
