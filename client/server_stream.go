// package main

// import (
// 	"context"
// 	"io"
// 	"log"

// 	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
// )

// func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
// 	log.Printf("calling SayHelloServerStream : Streaming started...")

// 	stream, err := client.SayHelloServerStream(context.Background(), names)
// 	if err != nil {
// 		log.Fatalf("failed to call SayHelloServerStream: %v", err)
// 	}

// 	for {
// 		res, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("failed to receive response: %v", err)
// 		}
// 		log.Printf("response from server: %v \n", res.Message)
// 	}
// }
