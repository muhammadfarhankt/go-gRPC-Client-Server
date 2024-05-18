package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Muhammad", "Farhan", "KT"},
	}

	callSayHello(client)

	callSayHelloServerStream(client, names)

	callSayHelloClientStream(client, names)

	//callSayHelloBidirectionalStream(client, names)
}

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}
	log.Printf("response from server: %v", res.Message)
}

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("calling SayHelloServerStream : Streaming started...")

	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("failed to call SayHelloServerStream: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}
		log.Printf("response from server: %v \n", res.Message)
	}
	log.Println("Streaming finished...")
}

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("client streaming starrted")
	stream, errr := client.SayHelloClientStream(context.Background())
	if errr != nil {
		log.Fatalf("error while calling SayHelloClientStream: %v", errr)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request: %v", err)
		}
		log.Printf("client sent: %v", name)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}
	log.Printf("response from server: %v", res.Messages)
}

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("calling SayHelloBidirectionalStream : Streaming started...")
	stream, err := client.SayHelloBiDirectionalStream(context.Background())
	fmt.Print("stream : ", stream)
	if err != nil {
		log.Fatalf("failed to call SayHelloBidirectionalStream: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send request: %v", err)
		}
		time.Sleep(time.Second * 2)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("failed receive response: %v", err)
			}
			log.Printf("response from server: %v", message)
		}
		close(waitc)
	}()

	stream.CloseSend()
	<-waitc
	log.Println("Bidirectional Streaming finished...")
}
