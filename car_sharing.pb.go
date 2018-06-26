// Code generated by protoc-gen-go. DO NOT EDIT.
// source: car_sharing.proto

package CarSharing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// request
type RegisterCarRequest struct {
	OwnerName            string   `protobuf:"bytes,1,opt,name=OwnerName,proto3" json:"OwnerName,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=Model,proto3" json:"Model,omitempty"`
	Place                string   `protobuf:"bytes,3,opt,name=Place,proto3" json:"Place,omitempty"`
	StartDate            string   `protobuf:"bytes,4,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              string   `protobuf:"bytes,5,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterCarRequest) Reset()         { *m = RegisterCarRequest{} }
func (m *RegisterCarRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterCarRequest) ProtoMessage()    {}
func (*RegisterCarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{0}
}
func (m *RegisterCarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterCarRequest.Unmarshal(m, b)
}
func (m *RegisterCarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterCarRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterCarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCarRequest.Merge(dst, src)
}
func (m *RegisterCarRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterCarRequest.Size(m)
}
func (m *RegisterCarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCarRequest proto.InternalMessageInfo

func (m *RegisterCarRequest) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *RegisterCarRequest) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *RegisterCarRequest) GetPlace() string {
	if m != nil {
		return m.Place
	}
	return ""
}

func (m *RegisterCarRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *RegisterCarRequest) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

// response
type RegisterCarReply struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	CarId                string   `protobuf:"bytes,2,opt,name=CarId,proto3" json:"CarId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterCarReply) Reset()         { *m = RegisterCarReply{} }
func (m *RegisterCarReply) String() string { return proto.CompactTextString(m) }
func (*RegisterCarReply) ProtoMessage()    {}
func (*RegisterCarReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{1}
}
func (m *RegisterCarReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterCarReply.Unmarshal(m, b)
}
func (m *RegisterCarReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterCarReply.Marshal(b, m, deterministic)
}
func (dst *RegisterCarReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCarReply.Merge(dst, src)
}
func (m *RegisterCarReply) XXX_Size() int {
	return xxx_messageInfo_RegisterCarReply.Size(m)
}
func (m *RegisterCarReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCarReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCarReply proto.InternalMessageInfo

func (m *RegisterCarReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RegisterCarReply) GetCarId() string {
	if m != nil {
		return m.CarId
	}
	return ""
}

// request
type SearchCarsRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=Model,proto3" json:"Model,omitempty"`
	Place                string   `protobuf:"bytes,3,opt,name=Place,proto3" json:"Place,omitempty"`
	StartDate            string   `protobuf:"bytes,4,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              string   `protobuf:"bytes,5,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCarsRequest) Reset()         { *m = SearchCarsRequest{} }
func (m *SearchCarsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchCarsRequest) ProtoMessage()    {}
func (*SearchCarsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{2}
}
func (m *SearchCarsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCarsRequest.Unmarshal(m, b)
}
func (m *SearchCarsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCarsRequest.Marshal(b, m, deterministic)
}
func (dst *SearchCarsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCarsRequest.Merge(dst, src)
}
func (m *SearchCarsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchCarsRequest.Size(m)
}
func (m *SearchCarsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCarsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCarsRequest proto.InternalMessageInfo

func (m *SearchCarsRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SearchCarsRequest) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *SearchCarsRequest) GetPlace() string {
	if m != nil {
		return m.Place
	}
	return ""
}

func (m *SearchCarsRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *SearchCarsRequest) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

// response
type SearchCarsReply struct {
	Status               string              `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Cars                 []*SearchCarsResult `protobuf:"bytes,2,rep,name=Cars,proto3" json:"Cars,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SearchCarsReply) Reset()         { *m = SearchCarsReply{} }
func (m *SearchCarsReply) String() string { return proto.CompactTextString(m) }
func (*SearchCarsReply) ProtoMessage()    {}
func (*SearchCarsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{3}
}
func (m *SearchCarsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCarsReply.Unmarshal(m, b)
}
func (m *SearchCarsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCarsReply.Marshal(b, m, deterministic)
}
func (dst *SearchCarsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCarsReply.Merge(dst, src)
}
func (m *SearchCarsReply) XXX_Size() int {
	return xxx_messageInfo_SearchCarsReply.Size(m)
}
func (m *SearchCarsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCarsReply.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCarsReply proto.InternalMessageInfo

func (m *SearchCarsReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SearchCarsReply) GetCars() []*SearchCarsResult {
	if m != nil {
		return m.Cars
	}
	return nil
}

type SearchCarsResult struct {
	CarId                string   `protobuf:"bytes,1,opt,name=CarId,proto3" json:"CarId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCarsResult) Reset()         { *m = SearchCarsResult{} }
func (m *SearchCarsResult) String() string { return proto.CompactTextString(m) }
func (*SearchCarsResult) ProtoMessage()    {}
func (*SearchCarsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{4}
}
func (m *SearchCarsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCarsResult.Unmarshal(m, b)
}
func (m *SearchCarsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCarsResult.Marshal(b, m, deterministic)
}
func (dst *SearchCarsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCarsResult.Merge(dst, src)
}
func (m *SearchCarsResult) XXX_Size() int {
	return xxx_messageInfo_SearchCarsResult.Size(m)
}
func (m *SearchCarsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCarsResult.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCarsResult proto.InternalMessageInfo

func (m *SearchCarsResult) GetCarId() string {
	if m != nil {
		return m.CarId
	}
	return ""
}

type UseCarRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=Model,proto3" json:"Model,omitempty"`
	PickupPlace          string   `protobuf:"bytes,3,opt,name=PickupPlace,proto3" json:"PickupPlace,omitempty"`
	DropPlace            string   `protobuf:"bytes,4,opt,name=DropPlace,proto3" json:"DropPlace,omitempty"`
	StartDate            string   `protobuf:"bytes,5,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              string   `protobuf:"bytes,6,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseCarRequest) Reset()         { *m = UseCarRequest{} }
func (m *UseCarRequest) String() string { return proto.CompactTextString(m) }
func (*UseCarRequest) ProtoMessage()    {}
func (*UseCarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{5}
}
func (m *UseCarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseCarRequest.Unmarshal(m, b)
}
func (m *UseCarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseCarRequest.Marshal(b, m, deterministic)
}
func (dst *UseCarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseCarRequest.Merge(dst, src)
}
func (m *UseCarRequest) XXX_Size() int {
	return xxx_messageInfo_UseCarRequest.Size(m)
}
func (m *UseCarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UseCarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UseCarRequest proto.InternalMessageInfo

func (m *UseCarRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UseCarRequest) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *UseCarRequest) GetPickupPlace() string {
	if m != nil {
		return m.PickupPlace
	}
	return ""
}

func (m *UseCarRequest) GetDropPlace() string {
	if m != nil {
		return m.DropPlace
	}
	return ""
}

func (m *UseCarRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *UseCarRequest) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

type UseCarReply struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	CarId                string   `protobuf:"bytes,2,opt,name=CarId,proto3" json:"CarId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseCarReply) Reset()         { *m = UseCarReply{} }
func (m *UseCarReply) String() string { return proto.CompactTextString(m) }
func (*UseCarReply) ProtoMessage()    {}
func (*UseCarReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{6}
}
func (m *UseCarReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseCarReply.Unmarshal(m, b)
}
func (m *UseCarReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseCarReply.Marshal(b, m, deterministic)
}
func (dst *UseCarReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseCarReply.Merge(dst, src)
}
func (m *UseCarReply) XXX_Size() int {
	return xxx_messageInfo_UseCarReply.Size(m)
}
func (m *UseCarReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UseCarReply.DiscardUnknown(m)
}

var xxx_messageInfo_UseCarReply proto.InternalMessageInfo

func (m *UseCarReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UseCarReply) GetCarId() string {
	if m != nil {
		return m.CarId
	}
	return ""
}

type ReturnCarRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	CarId                string   `protobuf:"bytes,2,opt,name=CarId,proto3" json:"CarId,omitempty"`
	Place                string   `protobuf:"bytes,3,opt,name=Place,proto3" json:"Place,omitempty"`
	ReturnDate           string   `protobuf:"bytes,4,opt,name=ReturnDate,proto3" json:"ReturnDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnCarRequest) Reset()         { *m = ReturnCarRequest{} }
func (m *ReturnCarRequest) String() string { return proto.CompactTextString(m) }
func (*ReturnCarRequest) ProtoMessage()    {}
func (*ReturnCarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{7}
}
func (m *ReturnCarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnCarRequest.Unmarshal(m, b)
}
func (m *ReturnCarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnCarRequest.Marshal(b, m, deterministic)
}
func (dst *ReturnCarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnCarRequest.Merge(dst, src)
}
func (m *ReturnCarRequest) XXX_Size() int {
	return xxx_messageInfo_ReturnCarRequest.Size(m)
}
func (m *ReturnCarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnCarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnCarRequest proto.InternalMessageInfo

func (m *ReturnCarRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ReturnCarRequest) GetCarId() string {
	if m != nil {
		return m.CarId
	}
	return ""
}

func (m *ReturnCarRequest) GetPlace() string {
	if m != nil {
		return m.Place
	}
	return ""
}

func (m *ReturnCarRequest) GetReturnDate() string {
	if m != nil {
		return m.ReturnDate
	}
	return ""
}

type ReturnCarReply struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnCarReply) Reset()         { *m = ReturnCarReply{} }
func (m *ReturnCarReply) String() string { return proto.CompactTextString(m) }
func (*ReturnCarReply) ProtoMessage()    {}
func (*ReturnCarReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_sharing_37b954a60e11ab7c, []int{8}
}
func (m *ReturnCarReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnCarReply.Unmarshal(m, b)
}
func (m *ReturnCarReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnCarReply.Marshal(b, m, deterministic)
}
func (dst *ReturnCarReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnCarReply.Merge(dst, src)
}
func (m *ReturnCarReply) XXX_Size() int {
	return xxx_messageInfo_ReturnCarReply.Size(m)
}
func (m *ReturnCarReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnCarReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnCarReply proto.InternalMessageInfo

func (m *ReturnCarReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterCarRequest)(nil), "CarSharing.RegisterCarRequest")
	proto.RegisterType((*RegisterCarReply)(nil), "CarSharing.RegisterCarReply")
	proto.RegisterType((*SearchCarsRequest)(nil), "CarSharing.SearchCarsRequest")
	proto.RegisterType((*SearchCarsReply)(nil), "CarSharing.SearchCarsReply")
	proto.RegisterType((*SearchCarsResult)(nil), "CarSharing.SearchCarsResult")
	proto.RegisterType((*UseCarRequest)(nil), "CarSharing.UseCarRequest")
	proto.RegisterType((*UseCarReply)(nil), "CarSharing.UseCarReply")
	proto.RegisterType((*ReturnCarRequest)(nil), "CarSharing.ReturnCarRequest")
	proto.RegisterType((*ReturnCarReply)(nil), "CarSharing.ReturnCarReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CarSharingClient is the client API for CarSharing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarSharingClient interface {
	RegisterCar(ctx context.Context, in *RegisterCarRequest, opts ...grpc.CallOption) (*RegisterCarReply, error)
	SearchCars(ctx context.Context, in *SearchCarsRequest, opts ...grpc.CallOption) (*SearchCarsReply, error)
	UseCar(ctx context.Context, in *UseCarRequest, opts ...grpc.CallOption) (*UseCarReply, error)
	ReturnCar(ctx context.Context, in *ReturnCarRequest, opts ...grpc.CallOption) (*ReturnCarReply, error)
}

type carSharingClient struct {
	cc *grpc.ClientConn
}

func NewCarSharingClient(cc *grpc.ClientConn) CarSharingClient {
	return &carSharingClient{cc}
}

func (c *carSharingClient) RegisterCar(ctx context.Context, in *RegisterCarRequest, opts ...grpc.CallOption) (*RegisterCarReply, error) {
	out := new(RegisterCarReply)
	err := c.cc.Invoke(ctx, "/CarSharing.CarSharing/RegisterCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carSharingClient) SearchCars(ctx context.Context, in *SearchCarsRequest, opts ...grpc.CallOption) (*SearchCarsReply, error) {
	out := new(SearchCarsReply)
	err := c.cc.Invoke(ctx, "/CarSharing.CarSharing/SearchCars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carSharingClient) UseCar(ctx context.Context, in *UseCarRequest, opts ...grpc.CallOption) (*UseCarReply, error) {
	out := new(UseCarReply)
	err := c.cc.Invoke(ctx, "/CarSharing.CarSharing/UseCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carSharingClient) ReturnCar(ctx context.Context, in *ReturnCarRequest, opts ...grpc.CallOption) (*ReturnCarReply, error) {
	out := new(ReturnCarReply)
	err := c.cc.Invoke(ctx, "/CarSharing.CarSharing/ReturnCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarSharingServer is the server API for CarSharing service.
type CarSharingServer interface {
	RegisterCar(context.Context, *RegisterCarRequest) (*RegisterCarReply, error)
	SearchCars(context.Context, *SearchCarsRequest) (*SearchCarsReply, error)
	UseCar(context.Context, *UseCarRequest) (*UseCarReply, error)
	ReturnCar(context.Context, *ReturnCarRequest) (*ReturnCarReply, error)
}

func RegisterCarSharingServer(s *grpc.Server, srv CarSharingServer) {
	s.RegisterService(&_CarSharing_serviceDesc, srv)
}

func _CarSharing_RegisterCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarSharingServer).RegisterCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CarSharing.CarSharing/RegisterCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarSharingServer).RegisterCar(ctx, req.(*RegisterCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarSharing_SearchCars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarSharingServer).SearchCars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CarSharing.CarSharing/SearchCars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarSharingServer).SearchCars(ctx, req.(*SearchCarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarSharing_UseCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarSharingServer).UseCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CarSharing.CarSharing/UseCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarSharingServer).UseCar(ctx, req.(*UseCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarSharing_ReturnCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarSharingServer).ReturnCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CarSharing.CarSharing/ReturnCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarSharingServer).ReturnCar(ctx, req.(*ReturnCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarSharing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CarSharing.CarSharing",
	HandlerType: (*CarSharingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCar",
			Handler:    _CarSharing_RegisterCar_Handler,
		},
		{
			MethodName: "SearchCars",
			Handler:    _CarSharing_SearchCars_Handler,
		},
		{
			MethodName: "UseCar",
			Handler:    _CarSharing_UseCar_Handler,
		},
		{
			MethodName: "ReturnCar",
			Handler:    _CarSharing_ReturnCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "car_sharing.proto",
}

func init() { proto.RegisterFile("car_sharing.proto", fileDescriptor_car_sharing_37b954a60e11ab7c) }

var fileDescriptor_car_sharing_37b954a60e11ab7c = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x8f, 0x93, 0x40,
	0x14, 0x95, 0xee, 0x16, 0xed, 0x25, 0xba, 0xbb, 0x13, 0xa3, 0x88, 0x75, 0xd3, 0x10, 0x1f, 0x78,
	0x22, 0x66, 0x7d, 0x32, 0x26, 0xc6, 0x2c, 0x6b, 0x8c, 0x26, 0xab, 0x04, 0xb2, 0x4f, 0x3e, 0x98,
	0x29, 0x4c, 0x5a, 0x52, 0x0a, 0x38, 0x33, 0x68, 0xea, 0xdf, 0x30, 0xf1, 0x77, 0xf8, 0xea, 0xbf,
	0x33, 0xcc, 0xf0, 0x31, 0xb4, 0xd2, 0xd8, 0x17, 0xdf, 0xb8, 0xe7, 0x0c, 0x87, 0x7b, 0xcf, 0x3d,
	0x03, 0x9c, 0x45, 0x98, 0x7e, 0x66, 0x4b, 0x4c, 0x93, 0x6c, 0xe1, 0x16, 0x34, 0xe7, 0x39, 0x02,
	0x0f, 0xd3, 0x50, 0x22, 0xf6, 0x4f, 0x0d, 0x50, 0x40, 0x16, 0x09, 0xe3, 0x84, 0x7a, 0x98, 0x06,
	0xe4, 0x4b, 0x49, 0x18, 0x47, 0x53, 0x98, 0x7c, 0xfc, 0x96, 0x11, 0xfa, 0x01, 0xaf, 0x89, 0xa9,
	0xcd, 0x34, 0x67, 0x12, 0x74, 0x00, 0xba, 0x0f, 0xe3, 0xeb, 0x3c, 0x26, 0xa9, 0x39, 0x12, 0x8c,
	0x2c, 0x2a, 0xd4, 0x4f, 0x71, 0x44, 0xcc, 0x23, 0x89, 0x8a, 0xa2, 0x52, 0x0a, 0x39, 0xa6, 0xfc,
	0x0a, 0x73, 0x62, 0x1e, 0x4b, 0xa5, 0x16, 0x40, 0x26, 0xdc, 0x7e, 0x93, 0xc5, 0x82, 0x1b, 0x0b,
	0xae, 0x29, 0xed, 0xd7, 0x70, 0xda, 0xeb, 0xab, 0x48, 0x37, 0xe8, 0x01, 0xe8, 0x21, 0xc7, 0xbc,
	0x64, 0x75, 0x4b, 0x75, 0x55, 0x7d, 0xd9, 0xc3, 0xf4, 0x5d, 0xdc, 0xf4, 0x23, 0x0a, 0xfb, 0x87,
	0x06, 0x67, 0x21, 0xc1, 0x34, 0x5a, 0x7a, 0x98, 0xb2, 0x66, 0x32, 0x0b, 0xee, 0xdc, 0xb0, 0xde,
	0x60, 0x6d, 0xfd, 0x5f, 0xe6, 0xfa, 0x04, 0x27, 0x6a, 0x53, 0xfb, 0xc6, 0x7a, 0x06, 0xc7, 0xd5,
	0x21, 0x73, 0x34, 0x3b, 0x72, 0x8c, 0x8b, 0xa9, 0xdb, 0xad, 0xcd, 0x55, 0x25, 0x58, 0x99, 0xf2,
	0x40, 0x9c, 0xb4, 0x1d, 0x38, 0xdd, 0x66, 0x3a, 0x73, 0x34, 0xd5, 0x9c, 0xdf, 0x1a, 0xdc, 0xbd,
	0x61, 0x44, 0x59, 0xf9, 0xe1, 0xc6, 0xcc, 0xc0, 0xf0, 0x93, 0x68, 0x55, 0x16, 0xaa, 0x3d, 0x2a,
	0x54, 0x99, 0x74, 0x45, 0xf3, 0x9a, 0xaf, 0x4d, 0x6a, 0x81, 0xbe, 0x85, 0xe3, 0x3d, 0x16, 0xea,
	0x7d, 0x0b, 0x5f, 0x82, 0xd1, 0xb4, 0x7e, 0x78, 0x2a, 0xbe, 0x57, 0xb9, 0xe2, 0x25, 0xcd, 0xfe,
	0x7d, 0xf4, 0x5d, 0x95, 0x81, 0x4c, 0x9c, 0x03, 0x48, 0x6d, 0x25, 0x14, 0x0a, 0x62, 0x3b, 0x70,
	0x4f, 0xf9, 0xf6, 0x9e, 0xde, 0x2f, 0x7e, 0x8d, 0x40, 0xb9, 0xa5, 0xe8, 0x1a, 0x0c, 0xe5, 0x32,
	0xa0, 0x73, 0x35, 0x0a, 0xbb, 0xb7, 0xd7, 0x9a, 0x0e, 0xf2, 0x45, 0xba, 0xb1, 0x6f, 0xa1, 0xf7,
	0x00, 0x5d, 0x4c, 0xd0, 0x93, 0xa1, 0x60, 0x49, 0xb1, 0xc7, 0x43, 0xb4, 0xd4, 0x7a, 0x05, 0xba,
	0x5c, 0x06, 0x7a, 0xa4, 0x1e, 0xec, 0x65, 0xcb, 0x7a, 0xf8, 0x37, 0x4a, 0xbe, 0xff, 0x16, 0x26,
	0xad, 0x27, 0x68, 0xab, 0xf1, 0xfe, 0x9a, 0x2c, 0x6b, 0x80, 0x15, 0x42, 0x97, 0x2f, 0xe0, 0x69,
	0x94, 0xaf, 0xdd, 0x45, 0xc2, 0x97, 0xe5, 0xdc, 0x9d, 0xe3, 0x68, 0x45, 0xb2, 0x98, 0x11, 0xfa,
	0x35, 0x89, 0x88, 0xf2, 0xe2, 0xe5, 0x49, 0xf7, 0xec, 0x57, 0xbf, 0x43, 0x5f, 0x9b, 0xeb, 0xe2,
	0xbf, 0xf8, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xc1, 0x06, 0xf9, 0x2c, 0x05, 0x00,
	0x00,
}
