package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-server/pb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	SUCCES = "SUCCESS"
	FAILED = "FAILED"
	PORT   = 8080
)

type server struct {
	pb.AttendanceServiceServer
}

func (s *server) CheckIn(ctx context.Context, req *pb.CheckInRequest) (*pb.CheckInResponse, error) {
	var resp pb.CheckInResponse
	log.Printf("Request check in %s", req.Username)
	dateTime, err := time.Parse("2006-01-02 15:04", req.Datetime)
	if err != nil {
		return nil, errors.New("format dateTime must be 'YYYY-MM-DD HH:MM'")
	}
	checkInTime, _, _ := dateTime.Clock()
	fmt.Println("checkInTime: ", checkInTime)
	if checkInTime > 9 {
		resp.Status = FAILED
		resp.Description = "late check in"
		return &resp, nil
	}

	resp.Status = SUCCES
	resp.Description = fmt.Sprintf("%s succesfull check in", req.Username)
	return &resp, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		panic(err)
	}
	log.Println("GRPC running on port", PORT)

	s := grpc.NewServer()
	pb.RegisterAttendanceServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
