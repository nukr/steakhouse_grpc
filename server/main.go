package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/nukr/steakhouse_grpc/pb"
	"google.golang.org/grpc"
	// pq ...
	_ "github.com/lib/pq"
	"github.com/nukr/checkerrors"
)

var (
	db     *sql.DB
	dbAddr string
)

func init() {
	var err error
	dbAddr = os.Getenv("DB_ADDR")
	db, err = sql.Open("postgres", dbAddr)
	checkerrors.Do(err, "fatal")
}

// Server ...
type Server struct{}

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 2323))
	s := grpc.NewServer()
	pb.RegisterSteakhouseServer(s, &Server{})
	log.Println("server listening on port 2323")
	s.Serve(lis)
}
