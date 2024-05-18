// package main

// import (
// 	"context"
// 	"io"
// 	"log"
// 	"time"

// 	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
// )

// func callHelloBidirectionalStream(client pb.GreetServiceClient, names pb.NameList) {
// 	log.Printf("calling SayHelloBidirectionalStream : Streaming started...")
// 	stream, err := client.SayHelloBiDirectionalStream(context.Background())
// 	if err != nil {
// 		log.Fatalf("failed to call SayHelloBidirectionalStream: %v", err)
// 	}

// 	waitc := make(chan struct{})

// 	go func() {
// 		for {
// 			message, err := stream.Recv()
// 			if err != nil {
// 				if err == io.EOF {
// 					break
// 				}
// 				log.Fatalf("failed to receive response: %v", err)
// 			}
// 			log.Printf("response from server: %v", message)
// 		}
// 		close(waitc)
// 	}()

// 	for _, name := range names.Names {
// 		req := &pb.HelloRequest{
// 			Name: name,
// 		}
// 		if err := stream.Send(req); err != nil {
// 			log.Fatalf("failed to send request: %v", err)
// 		}
// 		time.Sleep(time.Second * 2)
// 	}
// 	stream.CloseSend()
// 	<-waitc
// 	log.Println("Bidirectional Streaming finished...")
// }
