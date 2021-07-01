package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"dlpay.club/services/idgen/internal/pb"
	"dlpay.club/services/idgen/internal/snowflake"
)

const APIVersion = "v1"

type server struct {
	pb.UnimplementedIDGeneratorServer
}

var _ pb.IDGeneratorServer = server{}

func (s server) Get(_ context.Context, r *pb.IDGeneratorRequest) (*pb.IDGeneratorResponse, error) {
	if r.Version != APIVersion {
		return &pb.IDGeneratorResponse{
			Version: APIVersion,
			Id:      "",
		}, status.Error(codes.InvalidArgument, "不受支持的版本")
	}

	nextID, err := snowflake.Generator.NextID()
	for err != nil {
		nextID, err = snowflake.Generator.NextID()
	}

	return &pb.IDGeneratorResponse{Version: APIVersion, Id: strconv.FormatUint(nextID, 10)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	srv := server{}
	pb.RegisterIDGeneratorServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
