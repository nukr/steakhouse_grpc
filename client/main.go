package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/nukr/checkerrors"
	pb "github.com/nukr/steakhouse_grpc/pb"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:2323", grpc.WithInsecure())
	checkerrors.Do(err, "fatal")
	defer conn.Close()
	client := pb.NewSteakhouseClient(conn)
	createCuisine(client)
}

func getCuisine(client pb.SteakhouseClient) {
	cuisine, err := client.GetCuisine(context.Background(), &pb.Query{Id: "123123"})
	checkerrors.Do(err, "fatal")
	fmt.Println(cuisine)
}

func getCuisines(client pb.SteakhouseClient) {
	stream, err := client.GetCuisines(context.Background(), &pb.Query{Id: "123123"})
	checkerrors.Do(err, "fatal")
	value, err := stream.Recv()
	checkerrors.Do(err, "fatal")
	fmt.Println("cuisine from stream", value)
	value, err = stream.Recv()
	checkerrors.Do(err, "fatal")
	fmt.Println("cuisine from stream", value)
	value, err = stream.Recv()
	checkerrors.Do(err, "fatal")
	fmt.Println("cuisine from stream", value)
}

func createCuisine(client pb.SteakhouseClient) {
	start := time.Now()
	stream, err := client.CreateCuisine(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			checkerrors.Do(err, "fatal")
			log.Printf("Got status code: %d, msg: %s", in.Code, in.Msg)
		}
	}()
	for i := 0; i < 100; i++ {
		cuisine := pb.Cuisine{
			Id:          uuid.NewV4().String(),
			Name:        "test",
			Description: "aa",
			Price:       12.3,
		}
		checkerrors.Do(err, "fatal")
		time.Sleep(2 * time.Second)
		stream.Send(&cuisine)
		fmt.Println("sent")
	}
	errStreamCloseSend := stream.CloseSend()
	checkerrors.Do(errStreamCloseSend, "fatal")
	<-waitc
	fmt.Println(time.Since(start))
}
