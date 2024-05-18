// package main

// import (
// 	"io"
// 	"log"
// 	"time"

// 	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
// )

// func (s *helloServer) SayHelloBidirectionalStream(stream pb.GreetService_SayHelloBiDirectionalStreamServer) error {
// 	log.Printf("server starting SayHelloBidirectionalStream RPC...")
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		log.Printf("server received name request: %v", req.Name)
// 		res := &pb.HelloResponse{
// 			Message: "Hello, " + req.Name,
// 		}
// 		if err := stream.Send(res); err != nil {
// 			return err
// 		}
// 		time.Sleep(time.Second * 2)
// 	}
// }
