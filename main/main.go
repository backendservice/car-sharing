package main

import ( // 패키지 디렉토리
	"fmt"
	"log"
	"net"

	"github.com/backendservice/car-sharing"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// $GOPATH/src/github.com/backendservice/CarSharing
)

/*
struct
func RaiseError() (string, error) {
	return "", &NotFound{Message:"Can't find", Code: int32(404))}
}

func main() {
	if _, err := RaiseError(); err != nil {
		if _, ok := err.(*NotFound); ok {
			fmt.Print("Got Not Found")
		}
		else {
			fmt.Print("Got A Generic Error")
		}
	}
}
*/

const (
	port = ":50051"
)

// server is used to implement CarSharing.CarSharingServer.
type server struct{}

// RegisterCar implements CarSharing.CarSharingServer
func (s *server) RegisterCar(ctx context.Context, in *CarSharing.RegisterCarRequest) (*CarSharing.RegisterCarReply, error) {
	aCar := car{ID: carID, OwnerName: in.OwnerName, Model: in.Model, Place: in.Place, StartDate: in.StartDate, EndDate: in.EndDate}
	cars[carID] = aCar
	str := fmt.Sprintf("%d", carID)
	return &CarSharing.RegisterCarReply{Status: "Success", CarId: str}, nil
}

func (s *server) SearchCars(ctx context.Context, in *CarSharing.SearchCarsRequest) (*CarSharing.SearchCarsReply, error) {
	return &CarSharing.SearchCarsReply{Status: "Success"}, nil //CarId: "Any"}, nil
}

func (s *server) UseCar(ctx context.Context, in *CarSharing.UseCarRequest) (*CarSharing.UseCarReply, error) {
	return &CarSharing.UseCarReply{Status: "Success", CarId: "Any"}, nil
}

func (s *server) ReturnCar(ctx context.Context, in *CarSharing.ReturnCarRequest) (*CarSharing.ReturnCarReply, error) {
	return &CarSharing.ReturnCarReply{Status: "Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	CarSharing.RegisterCarSharingServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

/*
func main() {
	now := time.Now()
	formattedTime := now.Format(time.RFC3339)
	fmt.Println(formattedTime)
	//parsed, _ = time.Parse(formattedTime, time.RFC3339)
	req := CarSharing.RegisterCarRequest{"Owner1", "그랜저", "서초", "20180626", "20181231"}
	fmt.Print(req.OwnerName)
}
*/

// car type
type car struct {
	ID        int32  // generated number string
	OwnerName string // 소유자 이름
	Model     string // model name
	Place     string // 장소 or "InUse" 장소를 이용하여 사용중인지 판단
	//InUse bool // "True" or "False"
	UserName  string // 임대인 이름
	StartDate string // 'YYYYMMDD' 임대중
	EndDate   string
}

var cars map[int32]car
var carID = int32(0)
