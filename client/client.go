package main

import (
	protoFlight "flightService/proto"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var address = "127.0.0.1:11111"

func changeBooking(bookingId int32, flightNumber int32) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := protoFlight.NewFlightServiceClient(conn)

	req := &protoFlight.ChangeBookingRequest{
		BookingId:    bookingId,
		FlightNumber: flightNumber,
	}

	response, err := client.ChangeBooking(context.Background(), req)
	if err != nil {
		return err
	}

	fmt.Printf("Is success: %t", response.IsSuccess)
	return nil
}

func BookTicket(flightNumber int32, person *protoFlight.Person) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := protoFlight.NewFlightServiceClient(conn)

	req := &protoFlight.BookTicketRequest{
		FlightNumber: flightNumber,
		Person:       person,
	}

	response, err := client.BookTicket(context.Background(), req)
	if err != nil {
		return err
	}

	fmt.Printf("Is success: %t. Booking id: %d", response.IsSuccess, response.BookingId)
	return nil
}

func CancelBooking(bookingId int32) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := protoFlight.NewFlightServiceClient(conn)

	req := &protoFlight.BookingCancellationRequest{
		BookingId: bookingId,
	}

	response, err := client.CancelBooking(context.Background(), req)
	if err != nil {
		return err
	}

	fmt.Printf("Is success: %t.", response.IsSuccess)
	return nil
}

func main() {

}
