package main

import (
	"fmt"
	"log"
	"net"

	pb "GoWebApi/rpcdatacontract"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":60000"
)

// server is used to implement rpcdatacontract.AIServer.
type server struct{}

// SayHello implements rpcdatacontract.AIServer
func (s *server) Process(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	// cell1 := pb.Cell{Value: pb.Cell_ValInteger{ValInteger: 10.0}}
	col := pb.Column{
		EnglishName: "en_name_num",
		ChineseName: "zh_name_num",
		Type:        "number",
	}

	col2 := pb.Column{
		EnglishName: "en_name",
		ChineseName: "zh_name",
		Type:        "string",
	}

	cell := pb.Cell{Value: &pb.Cell_ValInteger{ValInteger: 10}}
	cell1 := pb.Cell{Value: &pb.Cell_ValString{ValString: "king"}}

	row := pb.Row{Cells: []*pb.Cell{&cell, &cell1}}
	table := pb.Table{
		Columns: []*pb.Column{&col, &col2},
		Rows:    []*pb.Row{&row},
	}

	// row := pb.Row
	return &pb.Reply{Tables: []*pb.Table{&table}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Listening at port %s", port)
	s := grpc.NewServer()
	pb.RegisterAIServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
