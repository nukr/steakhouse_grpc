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
	q := db.QueryRow(`
  SELECT id, name, description, price, is_deleted from cuisine
  WHERE id = $1
  `, in.Id)

	var id, name, description string
	var price float32
	var isDeleted bool

	err := q.Scan(&id, &name, &description, &price, &isDeleted)
	checkerrors.Do(err, "fatal")
	cuisine := &pb.Cuisine{
		Id:          id,
		Name:        name,
		Description: description,
		Price:       price,
		IsDeleted:   isDeleted,
	}

	return cuisine, nil
}

// GetCuisines ...
func (s *Server) GetCuisines(in *pb.Query, stream pb.Steakhouse_GetCuisinesServer) error {
	rows, err := db.Query(`
  SELECT id, name, description, price, is_deleted from cuisine
  `)

	checkerrors.Do(err, "fatal")

	defer rows.Close()
	for rows.Next() {
		var id, name, description string
		var price float32
		var isDeleted bool
		err := rows.Scan(&id, &name, &description, &price, &isDeleted)
		checkerrors.Do(err, "fatal")
		stream.Send(&pb.Cuisine{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			IsDeleted:   isDeleted,
		})
	}
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

		result, err := db.Exec(`
    INSERT INTO cuisine (id, name, description, price)
    VALUES ($1, $2, $3, $4)
    `, in.Id, in.Name, in.Description, in.Price)
		checkerrors.Do(err, "fatal")

		affected, err := result.RowsAffected()
		checkerrors.Do(err, "fatal")

		if affected > 0 {
			stream.Send(&pb.Status{
				Code: 0,
				Msg:  "done",
			})
		} else {
			stream.Send(&pb.Status{
				Code: 999,
				Msg:  "no new row",
			})
		}
	}
}

// UpdateCuisine ...
func (s *Server) UpdateCuisine(ctx context.Context, in *pb.Cuisine) (*pb.Status, error) {
	result, err := db.Exec(`
  UPDATE cuisine
  SET name = $1, description = $2, price = $3
  WHERE id = $4
  `, in.Name, in.Description, in.Price, in.Id)
	checkerrors.Do(err, "fatal")
	affected, err := result.RowsAffected()
	checkerrors.Do(err, "fatal")
	if affected > 0 {
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
func (s *Server) DeleteCuisine(ctx context.Context, in *pb.Query) (*pb.Status, error) {
	result, err := db.Exec(`
  DELETE FROM cuisine
  WHERE id = $1
  `, in.Id)
	checkerrors.Do(err, "fatal")
	affected, err := result.RowsAffected()
	checkerrors.Do(err, "fatal")
	if affected > 0 {
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

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 2323))
	s := grpc.NewServer()
	pb.RegisterSteakhouseServer(s, &Server{})
	log.Println("server listening on port 2323")
	s.Serve(lis)
}
