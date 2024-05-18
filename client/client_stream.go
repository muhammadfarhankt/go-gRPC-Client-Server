// package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
// )

// func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
// 	log.Printf("client streaming starrted")
// 	stream, errr := client.SayHelloClientStream(context.Background())
// 	if errr != nil {
// 		log.Fatalf("error while calling SayHelloClientStream: %v", errr)
// 	}
// 	for _, name := range names.Names {
// 		req := &pb.HelloRequest{
// 			Name: name,
// 		}
// 		if err := stream.Send(req); err != nil {
// 			log.Fatalf("error while sending request: %v", err)
// 		}
// 		log.Printf("client sent: %v", name)
// 		time.Sleep(time.Second * 2)
// 	}

// 	res, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalf("error while receiving response: %v", err)
// 	}
// 	log.Printf("response from server: %v", res.Messages)
// }
