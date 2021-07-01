package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"dlpay.club/services/idgen/internal/pb"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("请指定服务器连接地址")
	}
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewIDGeneratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.Get(ctx, &pb.IDGeneratorRequest{Version: "v1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", resp.Id)
}
