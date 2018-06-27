package main

import (
	"fmt"
	"log"
	"os"
	"time"

	//pb "github.com/backendservice/example-service/helloworld"
	"github.com/backendservice/car-sharing"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	//	address     = "ec2-18-191-204-27.us-east-2.compute.amazonaws.com:50051"
	address = "localhost:50051"
	//defaultName = "CarSharing"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := CarSharing.NewCarSharingClient(conn)

	// Contact the server and print out its response.
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	if command == "register" {
		// Owner Model Place StartDate EndDate
		message, err := RegisterCar(c, context.Background(), os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			log.Fatalf("could not register car: %v", err)
		}
		log.Printf("Register car success: %s", *message)

	} else if command == "search" {
		// Place StartDate EndDate, Place can be "Any"
		message, err := SearchCars(c, context.Background(), os.Args[2], os.Args[3], os.Args[4])
		if err != nil {
			log.Fatalf("could not search cars: %v", err)
		}
		log.Printf("search cars success: %s", *message)

	} else if command == "use" {
		// CarId UserName Model Place StartDate EndDate
		message, err := UseCar(c, context.Background(), os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7])
		if err != nil {
			log.Fatalf("could not use the car: %v", err)
		}
		log.Printf("use the car success: %s", *message)
	} else if command == "return" {
		// CarId UserName Place Date
		message, err := UseCar(c, context.Background(), os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7])
		if err != nil {
			log.Fatalf("could not return the car: %v", err)
		}
		log.Printf("return the car success: %s", *message)
	}
}

func RegisterCar(c CarSharing.CarSharingClient, ctx context.Context, owner string, model string, place string, start string, end string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.RegisterCar(ctx, &CarSharing.RegisterCarRequest{OwnerName: owner, Model: model, Place: place, StartDate: start, EndDate: end})
	if err != nil {
		return nil, err
	}
	return &r.Status, nil
}

// SearchCars - search cars for the condition
func SearchCars(c CarSharing.CarSharingClient, ctx context.Context, place string, start string, end string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.SearchCars(ctx, &CarSharing.SearchCarsRequest{Place: place, StartDate: start, EndDate: end})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Available %d cars\n", len(r.Cars))
	for _, aCar := range r.Cars {
		fmt.Printf("Id %s, Place %s, Model %s, Date %s - %s\n", aCar.CarId, aCar.Place, aCar.Model, aCar.ServiceStartDate, aCar.ServiceEndDate)
	}
	return &r.Status, nil
}

func UseCar(c CarSharing.CarSharingClient, ctx context.Context, id string, user string, model string, pickplace string, start string, end string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.UseCar(ctx, &CarSharing.UseCarRequest{CarId: id, UserName: user, PickupPlace: pickplace, StartDate: start, EndDate: end})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Use car: %s\n", r.CarId)
	return &r.Status, nil
}

func ReturnCar(c CarSharing.CarSharingClient, ctx context.Context, id string, user string, place string, return_date string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.ReturnCar(ctx, &CarSharing.ReturnCarRequest{CarId: id, UserName: user, Place: place, ReturnDate: return_date})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Return cars: %s\n", r.Status)
	return &r.Status, nil
}

/*
func SayHello(c pb.GreeterClient, ctx context.Context, timeout time.Duration, name string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return &r.Message, nil
}
*/
