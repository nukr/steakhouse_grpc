package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/nukr/steakhouse_grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// pq ...
	_ "github.com/lib/pq"
	"github.com/nukr/checkerrors"
)

var (
	db     *sql.DB
	dbaddr = "postgres://postgres:postgres@localhost:5432/steakhouse?sslmode=disable"
)

func init() {
	d, err := sql.Open("postgres", dbaddr)
	checkerrors.Do(err, "fatal")
	db = d
}

// Server ...
type Server struct{}

// GetCuisine ...
func (s *Server) GetCuisine(ctx context.Context, in *pb.Query) (*pb.Cuisine, error) {
	db.Exec(`
  SELECT * from cuisine
  `)
	return &pb.Cuisine{
		Name:        "aa",
		Description: "bb",
	}, nil
}

// GetCuisines ...
func (s *Server) GetCuisines(in *pb.Query, stream pb.Steakhouse_GetCuisinesServer) error {
	stream.Send(&pb.Cuisine{
		Name:        "aa",
		Description: "bb",
	})
	stream.Send(&pb.Cuisine{
		Name:        "bb",
		Description: "cc",
	})
	return nil
}

// CreateCuisine ...
func (s *Server) CreateCuisine(stream pb.Steakhouse_CreateCuisineServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		checkerrors.Do(err, "fatal")

		db.Exec(`
    INSERT INTO cuisine (id, name, description, price)
    VALUES ($1, $2, $3, $4)
    `, in.Id, in.Name, in.Description, in.Price)

		stream.Send(&pb.Status{
			Code: 0,
			Msg:  "done",
		})
	}
}

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 2323))
	s := grpc.NewServer()
	pb.RegisterSteakhouseServer(s, &Server{})
	log.Println("server listening on port 2323")
	s.Serve(lis)
}
