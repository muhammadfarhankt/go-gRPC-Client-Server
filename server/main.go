package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	pb "github.com/muhammadfarhankt/go-gRPC-Client-Server/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *helloServer) SayHelloServerStream(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("server received name requests: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello, " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

// unary RPC
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, from the server",
	}, nil
}

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	log.Printf("server received client stream requests")
	var names []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.MessagesList{Messages: names})
			}
			return err
		}
		log.Printf("server received name request: %v", req.Name)
		//names = append(names, "Hello", req.Names...)
		names = append(names, "Hello", req.Name)
	}
}

// func (s *helloServer) SayHelloBidirectionalStream(ctx context.Context, stream pb.GreetService_SayHelloBiDirectionalStreamServer) error {
// 	log.Printf("server starting SayHelloBidirectionalStream RPC...")

// 	//log.Printf("stream : %v", stream)
// 	if stream == nil {
// 		return errors.New("stream is empty")
// 	}
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			log.Println("err == io.EOF")
// 			return nil
// 		}
// 		if err != nil {
// 			log.Println("err != nil")
// 			return err
// 		}
// 		if req == nil {
// 			log.Println("request nil")
// 			return errors.New("request is empty")
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

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamServer) error {
	log.Println("server starting SayHelloBidirectionalStream RPC...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
