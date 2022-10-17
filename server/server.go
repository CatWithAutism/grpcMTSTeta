package main

import (
	"context"
	"flag"
	"log"
	"net"
	"sync"

	protoFlight "flightService/proto"
	"google.golang.org/grpc"
)

type flightService struct {
	customers map[int]*protoFlight.Person
	id        int
	sync.Mutex
}

func (cs *flightService) MustEmbedUnimplementedFlightServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewFlightService() *flightService {
	cs := flightService{
		id:        1,
		customers: make(map[int]*protoFlight.Person),
	}
	return &cs
}

func (cs *flightService) CancelBooking(ctx context.Context, request *protoFlight.BookingCancellationRequest) (*protoFlight.BookingCancellationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *flightService) BookTicket(ctx context.Context, request *protoFlight.BookTicketRequest) (*protoFlight.BookTicketResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *flightService) ChangeBooking(ctx context.Context, request *protoFlight.ChangeBookingRequest) (*protoFlight.ChangeBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	protoFlight.RegisterFlightServiceServer(server, NewFlightService())
	server.Serve(lis)
}
