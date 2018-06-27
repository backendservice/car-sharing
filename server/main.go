package main

import ( // 패키지 디렉토리

	"fmt"
	"log"
	"net"
	"strconv"

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
	fmt.Printf("RegisterCar: owner %s, model %s, place %s\n", in.OwnerName, in.Model, in.Place)
	aCar := car{ID: carID, InUse: false, OwnerName: in.OwnerName, Model: in.Model, Place: in.Place, ServiceStartDate: in.StartDate, ServiceEndDate: in.EndDate}
	cars[carID] = aCar
	//str := fmt.Sprintf("%d", carID)
	str := strconv.Itoa(carID)
	carID++ // 증가하는 유일한 값
	return &CarSharing.RegisterCarReply{Status: "Success", CarId: str}, nil
}

func (s *server) SearchCars(ctx context.Context, in *CarSharing.SearchCarsRequest) (*CarSharing.SearchCarsReply, error) {
	//search cars for meeting Place, StartData, EndDate
	//var Cars []*SearchCarsResult
	//a := make([]int, 5) // len(a)=5
	//myCarIds := make([]string, 50)
	fmt.Printf("SearchCars: place %s, date %s - %s\n", in.Place, in.StartDate, in.EndDate)
	rsltCars := []*CarSharing.SearchCarsResult{}
	for _, aCar := range cars {
		fmt.Printf("SearchCars: checking a car, owner %s, model %s\n", aCar.OwnerName, aCar.Model)
		if aCar.Place == in.Place || in.Place == "Any" {
			if aCar.ServiceStartDate <= in.StartDate && in.EndDate <= aCar.ServiceEndDate {
				fmt.Printf("SearchCars: place & dates are ok!,  date %s - %s\n", aCar.ServiceStartDate, aCar.ServiceEndDate)
				if !aCar.InUse || aCar.EndDate < in.StartDate {
					// found a car, add a car to pointer array of cars
					fmt.Printf("SearchCars: adding a car\n")
					rslt := &CarSharing.SearchCarsResult{CarId: strconv.Itoa(aCar.ID), Model: aCar.Model, Place: aCar.Place, ServiceStartDate: aCar.ServiceStartDate, ServiceEndDate: aCar.ServiceEndDate}
					rsltCars = append(rsltCars, rslt)
				}
			}
		}
	}
	if len(rsltCars) > 0 {
		return &CarSharing.SearchCarsReply{Status: "Success", Cars: rsltCars}, nil
	} else {
		return &CarSharing.SearchCarsReply{Status: "Fail"}, nil
	}
}

func (s *server) UseCar(ctx context.Context, in *CarSharing.UseCarRequest) (*CarSharing.UseCarReply, error) {
	id, _ := strconv.Atoi(in.CarId)
	aCar, exists := cars[id]
	if !exists {
		return &CarSharing.UseCarReply{Status: "Fail", CarId: "Any"}, nil
	} else if aCar.InUse == true {
		return &CarSharing.UseCarReply{Status: "Fail", CarId: "Any"}, nil
	} else {
		aCar.InUse = true
		aCar.UserName = in.UserName
		aCar.StartDate = in.StartDate
		aCar.EndDate = in.EndDate
		return &CarSharing.UseCarReply{Status: "Success", CarId: "Any"}, nil
	}
}

func (s *server) ReturnCar(ctx context.Context, in *CarSharing.ReturnCarRequest) (*CarSharing.ReturnCarReply, error) {
	id, _ := strconv.Atoi(in.CarId)
	aCar, exists := cars[id]
	if !exists {
		return &CarSharing.ReturnCarReply{Status: "Fail"}, nil
	} else if aCar.InUse == false {
		return &CarSharing.ReturnCarReply{Status: "Fail"}, nil
	} else {
		aCar.InUse = false
		aCar.UserName = ""
		aCar.StartDate = ""
		aCar.EndDate = ""
		return &CarSharing.ReturnCarReply{Status: "Success"}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("CarSharing service started")
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
	ID               int    // generated number string
	OwnerName        string // 소유자 이름
	Model            string // model name
	Place            string // 장소
	ServiceStartDate string // 서비스 시작일
	ServiceEndDate   string // 서비스 종료일
	InUse            bool   // 임대중인지, "True" or "False"
	UserName         string // 임대인 이름
	StartDate        string // 'YYYYMMDDHH' 임대 시작일
	EndDate          string // 임대 종료일
}

var cars = map[int]car{}
var carID = int(0) // 증가하는 유일한 값
