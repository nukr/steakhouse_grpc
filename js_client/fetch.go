package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/nukr/steakhouse_grpc/pb"
)

func main() {
	buf := yield()
	client := http.Client{}
	req := makeReq(buf)
	client.Do(req)
}

func yield() *bytes.Buffer {
	cuisine := pb.Cuisine{
		Id:          "1234",
		Title:       "美沙",
		Price:       580,
		Description: "ggggggggg",
	}
	j, _ := json.Marshal(cuisine)
	s := fmt.Sprintf("%s\r\n%s\r\n", string(j), string(j))
	buf := bytes.NewBuffer([]byte(s))
	return buf
}

func makeReq(buf *bytes.Buffer) *http.Request {
	req, err := http.NewRequest("POST", "http://localhost:9091/v1/cuisine", buf)
	req.Header.Add("Transfer-Encoding", "chunked")

	if err != nil {
		log.Println(err)
	}
	return req
}
