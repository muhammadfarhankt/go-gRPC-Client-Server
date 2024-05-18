// package main

// import (
// 	"io"
// 	"log"

// 	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
// )

// func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
// 	log.Printf("server received client stream requests")
// 	var names []string
// 	for {
// 		req, err := stream.Recv()
// 		if err != nil {
// 			if err == io.EOF {
// 				return stream.SendAndClose(&pb.MessagesList{Messages: names})
// 			}
// 			return err
// 		}
// 		log.Printf("server received name request: %v", req.Names)
// 		//names = append(names, "Hello", req.Names...)
// 		names = append(names, "Hello", req.Name)
// 	}
// }
