package main

import (
	"fmt"
	"io"

	"golang.org/x/net/context"

	pb "github.com/nukr/steakhouse_grpc/pb"
)

// GetCuisine ...
func (s *Server) GetCuisine(
	ctx context.Context,
	in *pb.Query,
) (cuisine *pb.Cuisine, err error) {
	return queryCuisineByID(in.Id)
}

// GetCuisines ...
func (s *Server) GetCuisines(
	in *pb.Query,
	stream pb.Steakhouse_GetCuisinesServer,
) error {
	eachCuisines(func(cuisine *pb.Cuisine) {
		stream.Send(cuisine)
	})
	return nil
}

// CreateCuisine ...
func (s *Server) CreateCuisine(stream pb.Steakhouse_CreateCuisineServer) error {
	for {
		in, err := stream.Recv()
		fmt.Println("create cuisine", in)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		createCuisine(in)
		stream.Send(&pb.Status{
			Code: 0,
			Msg:  "done",
		})
	}
}

// UpdateCuisine ...
func (s *Server) UpdateCuisine(
	ctx context.Context,
	in *pb.Cuisine,
) (*pb.Status, error) {
	err := updateCuisine(in)
	if err != nil {
		return &pb.Status{
			Code: 0,
			Msg:  "done",
		}, nil
	}
	return &pb.Status{
		Code: 999,
		Msg:  "no new row",
	}, nil
}

// DeleteCuisine ...
func (s *Server) DeleteCuisine(
	ctx context.Context,
	in *pb.Query,
) (*pb.Status, error) {
	err := deleteCuisine(in.Id)
	if err != nil {
		return &pb.Status{
			Code: 999,
			Msg:  "gg",
		}, nil
	}
	return &pb.Status{
		Code: 0,
		Msg:  "ok",
	}, nil
}
