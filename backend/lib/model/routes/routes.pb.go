// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shared/model/routes/routes.proto

package routes // import "github.com/rhino1998/programme/backend/lib/model/routes"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TravelMethod int32

const (
	Unknown   TravelMethod = 0
	Walking   TravelMethod = 1
	Biking    TravelMethod = 2
	Rideshare TravelMethod = 3
	Car       TravelMethod = 4
)

var TravelMethod_name = map[int32]string{
	0: "Unknown",
	1: "Walking",
	2: "Biking",
	3: "Rideshare",
	4: "Car",
}
var TravelMethod_value = map[string]int32{
	"Unknown":   0,
	"Walking":   1,
	"Biking":    2,
	"Rideshare": 3,
	"Car":       4,
}

func (TravelMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_routes_4e57bbe1d56c3914, []int{0}
}

type Coords struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
}

func (m *Coords) Reset()      { *m = Coords{} }
func (*Coords) ProtoMessage() {}
func (*Coords) Descriptor() ([]byte, []int) {
	return fileDescriptor_routes_4e57bbe1d56c3914, []int{0}
}
func (m *Coords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Coords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Coords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Coords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coords.Merge(dst, src)
}
func (m *Coords) XXX_Size() int {
	return m.Size()
}
func (m *Coords) XXX_DiscardUnknown() {
	xxx_messageInfo_Coords.DiscardUnknown(m)
}

var xxx_messageInfo_Coords proto.InternalMessageInfo

func (m *Coords) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Coords) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Location struct {
	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Coords *Coords `protobuf:"bytes,2,opt,name=coords" json:"coords,omitempty"`
}

func (m *Location) Reset()      { *m = Location{} }
func (*Location) ProtoMessage() {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_routes_4e57bbe1d56c3914, []int{1}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Location.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return m.Size()
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetCoords() *Coords {
	if m != nil {
		return m.Coords
	}
	return nil
}

type Trip struct {
	Start    *Location    `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End      *Location    `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	Method   TravelMethod `protobuf:"varint,3,opt,name=method,proto3,enum=shared.model.routes.TravelMethod" json:"method,omitempty"`
	Duration int64        `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *Trip) Reset()      { *m = Trip{} }
func (*Trip) ProtoMessage() {}
func (*Trip) Descriptor() ([]byte, []int) {
	return fileDescriptor_routes_4e57bbe1d56c3914, []int{2}
}
func (m *Trip) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trip.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Trip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trip.Merge(dst, src)
}
func (m *Trip) XXX_Size() int {
	return m.Size()
}
func (m *Trip) XXX_DiscardUnknown() {
	xxx_messageInfo_Trip.DiscardUnknown(m)
}

var xxx_messageInfo_Trip proto.InternalMessageInfo

func (m *Trip) GetStart() *Location {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Trip) GetEnd() *Location {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Trip) GetMethod() TravelMethod {
	if m != nil {
		return m.Method
	}
	return Unknown
}

func (m *Trip) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type Trips struct {
	Trips []*Trip `protobuf:"bytes,1,rep,name=trips" json:"trips,omitempty"`
}

func (m *Trips) Reset()      { *m = Trips{} }
func (*Trips) ProtoMessage() {}
func (*Trips) Descriptor() ([]byte, []int) {
	return fileDescriptor_routes_4e57bbe1d56c3914, []int{3}
}
func (m *Trips) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trips) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trips.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Trips) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trips.Merge(dst, src)
}
func (m *Trips) XXX_Size() int {
	return m.Size()
}
func (m *Trips) XXX_DiscardUnknown() {
	xxx_messageInfo_Trips.DiscardUnknown(m)
}

var xxx_messageInfo_Trips proto.InternalMessageInfo

func (m *Trips) GetTrips() []*Trip {
	if m != nil {
		return m.Trips
	}
	return nil
}

type Locations struct {
	Locations []*Location  `protobuf:"bytes,1,rep,name=locations" json:"locations,omitempty"`
	Method    TravelMethod `protobuf:"varint,2,opt,name=method,proto3,enum=shared.model.routes.TravelMethod" json:"method,omitempty"`
}

func (m *Locations) Reset()      { *m = Locations{} }
func (*Locations) ProtoMessage() {}
func (*Locations) Descriptor() ([]byte, []int) {
	return fileDescriptor_routes_4e57bbe1d56c3914, []int{4}
}
func (m *Locations) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Locations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Locations.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Locations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Locations.Merge(dst, src)
}
func (m *Locations) XXX_Size() int {
	return m.Size()
}
func (m *Locations) XXX_DiscardUnknown() {
	xxx_messageInfo_Locations.DiscardUnknown(m)
}

var xxx_messageInfo_Locations proto.InternalMessageInfo

func (m *Locations) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *Locations) GetMethod() TravelMethod {
	if m != nil {
		return m.Method
	}
	return Unknown
}

func init() {
	proto.RegisterType((*Coords)(nil), "shared.model.routes.Coords")
	proto.RegisterType((*Location)(nil), "shared.model.routes.Location")
	proto.RegisterType((*Trip)(nil), "shared.model.routes.Trip")
	proto.RegisterType((*Trips)(nil), "shared.model.routes.Trips")
	proto.RegisterType((*Locations)(nil), "shared.model.routes.Locations")
	proto.RegisterEnum("shared.model.routes.TravelMethod", TravelMethod_name, TravelMethod_value)
}
func (x TravelMethod) String() string {
	s, ok := TravelMethod_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Coords) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Coords)
	if !ok {
		that2, ok := that.(Coords)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Latitude != that1.Latitude {
		return false
	}
	if this.Longitude != that1.Longitude {
		return false
	}
	return true
}
func (this *Location) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Location)
	if !ok {
		that2, ok := that.(Location)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Coords.Equal(that1.Coords) {
		return false
	}
	return true
}
func (this *Trip) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Trip)
	if !ok {
		that2, ok := that.(Trip)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Start.Equal(that1.Start) {
		return false
	}
	if !this.End.Equal(that1.End) {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	return true
}
func (this *Trips) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Trips)
	if !ok {
		that2, ok := that.(Trips)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Trips) != len(that1.Trips) {
		return false
	}
	for i := range this.Trips {
		if !this.Trips[i].Equal(that1.Trips[i]) {
			return false
		}
	}
	return true
}
func (this *Locations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Locations)
	if !ok {
		that2, ok := that.(Locations)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Locations) != len(that1.Locations) {
		return false
	}
	for i := range this.Locations {
		if !this.Locations[i].Equal(that1.Locations[i]) {
			return false
		}
	}
	if this.Method != that1.Method {
		return false
	}
	return true
}
func (this *Coords) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&routes.Coords{")
	s = append(s, "Latitude: "+fmt.Sprintf("%#v", this.Latitude)+",\n")
	s = append(s, "Longitude: "+fmt.Sprintf("%#v", this.Longitude)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Location) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&routes.Location{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	if this.Coords != nil {
		s = append(s, "Coords: "+fmt.Sprintf("%#v", this.Coords)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Trip) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&routes.Trip{")
	if this.Start != nil {
		s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	}
	if this.End != nil {
		s = append(s, "End: "+fmt.Sprintf("%#v", this.End)+",\n")
	}
	s = append(s, "Method: "+fmt.Sprintf("%#v", this.Method)+",\n")
	s = append(s, "Duration: "+fmt.Sprintf("%#v", this.Duration)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Trips) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&routes.Trips{")
	if this.Trips != nil {
		s = append(s, "Trips: "+fmt.Sprintf("%#v", this.Trips)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Locations) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&routes.Locations{")
	if this.Locations != nil {
		s = append(s, "Locations: "+fmt.Sprintf("%#v", this.Locations)+",\n")
	}
	s = append(s, "Method: "+fmt.Sprintf("%#v", this.Method)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRoutes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoutesAPIClient is the client API for RoutesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutesAPIClient interface {
	CalcTravelTime(ctx context.Context, in *Locations, opts ...grpc.CallOption) (*Trips, error)
}

type routesAPIClient struct {
	cc *grpc.ClientConn
}

func NewRoutesAPIClient(cc *grpc.ClientConn) RoutesAPIClient {
	return &routesAPIClient{cc}
}

func (c *routesAPIClient) CalcTravelTime(ctx context.Context, in *Locations, opts ...grpc.CallOption) (*Trips, error) {
	out := new(Trips)
	err := c.cc.Invoke(ctx, "/shared.model.routes.RoutesAPI/CalcTravelTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutesAPIServer is the server API for RoutesAPI service.
type RoutesAPIServer interface {
	CalcTravelTime(context.Context, *Locations) (*Trips, error)
}

func RegisterRoutesAPIServer(s *grpc.Server, srv RoutesAPIServer) {
	s.RegisterService(&_RoutesAPI_serviceDesc, srv)
}

func _RoutesAPI_CalcTravelTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Locations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesAPIServer).CalcTravelTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.model.routes.RoutesAPI/CalcTravelTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesAPIServer).CalcTravelTime(ctx, req.(*Locations))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoutesAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shared.model.routes.RoutesAPI",
	HandlerType: (*RoutesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalcTravelTime",
			Handler:    _RoutesAPI_CalcTravelTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shared/model/routes/routes.proto",
}

func (m *Coords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coords) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Latitude != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Latitude))))
		i += 8
	}
	if m.Longitude != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Longitude))))
		i += 8
	}
	return i, nil
}

func (m *Location) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Location) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Coords != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(m.Coords.Size()))
		n1, err := m.Coords.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Trip) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trip) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(m.Start.Size()))
		n2, err := m.Start.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.End != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(m.End.Size()))
		n3, err := m.End.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Method != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(m.Method))
	}
	if m.Duration != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(m.Duration))
	}
	return i, nil
}

func (m *Trips) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trips) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Trips) > 0 {
		for _, msg := range m.Trips {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRoutes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Locations) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Locations) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Locations) > 0 {
		for _, msg := range m.Locations {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRoutes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Method != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(m.Method))
	}
	return i, nil
}

func encodeVarintRoutes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Coords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Latitude != 0 {
		n += 9
	}
	if m.Longitude != 0 {
		n += 9
	}
	return n
}

func (m *Location) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRoutes(uint64(l))
	}
	if m.Coords != nil {
		l = m.Coords.Size()
		n += 1 + l + sovRoutes(uint64(l))
	}
	return n
}

func (m *Trip) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != nil {
		l = m.Start.Size()
		n += 1 + l + sovRoutes(uint64(l))
	}
	if m.End != nil {
		l = m.End.Size()
		n += 1 + l + sovRoutes(uint64(l))
	}
	if m.Method != 0 {
		n += 1 + sovRoutes(uint64(m.Method))
	}
	if m.Duration != 0 {
		n += 1 + sovRoutes(uint64(m.Duration))
	}
	return n
}

func (m *Trips) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Trips) > 0 {
		for _, e := range m.Trips {
			l = e.Size()
			n += 1 + l + sovRoutes(uint64(l))
		}
	}
	return n
}

func (m *Locations) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Locations) > 0 {
		for _, e := range m.Locations {
			l = e.Size()
			n += 1 + l + sovRoutes(uint64(l))
		}
	}
	if m.Method != 0 {
		n += 1 + sovRoutes(uint64(m.Method))
	}
	return n
}

func sovRoutes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoutes(x uint64) (n int) {
	return sovRoutes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Coords) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Coords{`,
		`Latitude:` + fmt.Sprintf("%v", this.Latitude) + `,`,
		`Longitude:` + fmt.Sprintf("%v", this.Longitude) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Location) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Location{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Coords:` + strings.Replace(fmt.Sprintf("%v", this.Coords), "Coords", "Coords", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Trip) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Trip{`,
		`Start:` + strings.Replace(fmt.Sprintf("%v", this.Start), "Location", "Location", 1) + `,`,
		`End:` + strings.Replace(fmt.Sprintf("%v", this.End), "Location", "Location", 1) + `,`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`Duration:` + fmt.Sprintf("%v", this.Duration) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Trips) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Trips{`,
		`Trips:` + strings.Replace(fmt.Sprintf("%v", this.Trips), "Trip", "Trip", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Locations) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Locations{`,
		`Locations:` + strings.Replace(fmt.Sprintf("%v", this.Locations), "Location", "Location", 1) + `,`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRoutes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Coords) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoutes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Coords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coords: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Latitude = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Longitude = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipRoutes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoutes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Location) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoutes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Location: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Location: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Coords == nil {
				m.Coords = &Coords{}
			}
			if err := m.Coords.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoutes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoutes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Trip) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoutes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Trip: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trip: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Start == nil {
				m.Start = &Location{}
			}
			if err := m.Start.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.End == nil {
				m.End = &Location{}
			}
			if err := m.End.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= (TravelMethod(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoutes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoutes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Trips) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoutes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Trips: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trips: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trips", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trips = append(m.Trips, &Trip{})
			if err := m.Trips[len(m.Trips)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoutes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoutes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Locations) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoutes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Locations: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Locations: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoutes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locations = append(m.Locations, &Location{})
			if err := m.Locations[len(m.Locations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= (TravelMethod(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoutes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoutes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRoutes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoutes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoutes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRoutes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoutes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRoutes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRoutes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoutes   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("shared/model/routes/routes.proto", fileDescriptor_routes_4e57bbe1d56c3914)
}

var fileDescriptor_routes_4e57bbe1d56c3914 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0xf6, 0xc6, 0x49, 0xee, 0x32, 0x81, 0x53, 0xb4, 0x34, 0x21, 0xc0, 0xca, 0xb8, 0x8a, 0x28,
	0x6c, 0xe1, 0x14, 0x5c, 0x44, 0x45, 0x52, 0x21, 0x19, 0x09, 0x99, 0x20, 0x04, 0xdd, 0xc6, 0x5e,
	0x25, 0xab, 0xd8, 0xbb, 0xd1, 0x7a, 0x03, 0x2d, 0x12, 0x2f, 0xc0, 0x63, 0xf0, 0x1e, 0x34, 0x94,
	0x29, 0xaf, 0x24, 0x4e, 0x43, 0x79, 0x8f, 0x80, 0xbc, 0xf6, 0xfd, 0x20, 0xe5, 0x38, 0x5d, 0xe5,
	0x19, 0xcf, 0xf7, 0x33, 0xdf, 0x48, 0x0b, 0x4e, 0xbe, 0xa4, 0x8a, 0x25, 0x7e, 0x26, 0x13, 0x96,
	0xfa, 0x4a, 0x6e, 0x34, 0xcb, 0xeb, 0x8f, 0xb7, 0x56, 0x52, 0x4b, 0xfc, 0xa0, 0x42, 0x78, 0x06,
	0xe1, 0x55, 0x23, 0x77, 0x02, 0xed, 0xa9, 0x94, 0x2a, 0xc9, 0xf1, 0x00, 0x8e, 0x43, 0xaa, 0xb9,
	0xde, 0x24, 0xac, 0x8f, 0x1c, 0x34, 0x44, 0xd1, 0x65, 0x8f, 0x1f, 0x43, 0x27, 0x94, 0x62, 0x51,
	0x0d, 0x1b, 0x66, 0x78, 0xf5, 0xc3, 0x7d, 0x07, 0xc7, 0xa1, 0x8c, 0xa9, 0xe6, 0x52, 0x60, 0x0c,
	0x4d, 0x41, 0xb3, 0x4a, 0xa1, 0x13, 0x99, 0x1a, 0x8f, 0xa0, 0x1d, 0x1b, 0x0f, 0x43, 0xed, 0x06,
	0x8f, 0xbc, 0x03, 0x9b, 0x78, 0xd5, 0x1a, 0x51, 0x0d, 0x75, 0x7f, 0x22, 0x68, 0xce, 0x14, 0x5f,
	0xe3, 0x11, 0xb4, 0x72, 0x4d, 0x95, 0x36, 0x92, 0xdd, 0xe0, 0xc9, 0x41, 0xf2, 0x85, 0x7f, 0x54,
	0x61, 0xb1, 0x0f, 0x36, 0x13, 0x49, 0xed, 0x77, 0x0b, 0xa5, 0x44, 0xe2, 0x31, 0xb4, 0x33, 0xa6,
	0x97, 0x32, 0xe9, 0xdb, 0x0e, 0x1a, 0x9e, 0x04, 0x4f, 0x0f, 0x72, 0x66, 0x8a, 0x7e, 0x66, 0xe9,
	0x1b, 0x03, 0x8c, 0x6a, 0x42, 0x79, 0xb8, 0x64, 0xa3, 0x8c, 0x56, 0xbf, 0xe9, 0xa0, 0xa1, 0x1d,
	0x5d, 0xf6, 0xee, 0x29, 0xb4, 0xca, 0x10, 0x39, 0xf6, 0xa1, 0xa5, 0xcb, 0xa2, 0x8f, 0x1c, 0x7b,
	0xd8, 0x0d, 0x1e, 0xde, 0x20, 0xcf, 0xd7, 0x51, 0x85, 0x73, 0xbf, 0xa1, 0xf2, 0xe6, 0xd5, 0x8a,
	0x39, 0x7e, 0x09, 0x9d, 0xf4, 0xa2, 0xa9, 0x25, 0x6e, 0x49, 0x75, 0x85, 0xbf, 0x96, 0xad, 0x71,
	0xc7, 0x6c, 0xcf, 0x42, 0xb8, 0x77, 0xfd, 0x3f, 0xee, 0xc2, 0xd1, 0x7b, 0xb1, 0x12, 0xf2, 0x8b,
	0xe8, 0x59, 0x65, 0xf3, 0x81, 0xa6, 0x2b, 0x2e, 0x16, 0x3d, 0x84, 0x01, 0xda, 0x13, 0x6e, 0xea,
	0x06, 0xbe, 0x0f, 0x9d, 0x88, 0x27, 0xcc, 0xb8, 0xf4, 0x6c, 0x7c, 0x04, 0xf6, 0x94, 0xaa, 0x5e,
	0x33, 0xf8, 0x08, 0x9d, 0xc8, 0x98, 0xbd, 0x7a, 0xfb, 0x1a, 0x87, 0x70, 0x32, 0xa5, 0x69, 0x5c,
	0xc9, 0xcf, 0x78, 0xc6, 0x30, 0xf9, 0x6f, 0xa2, 0x7c, 0x30, 0xb8, 0xf1, 0x68, 0xf9, 0x24, 0xdb,
	0xee, 0x88, 0x75, 0xb6, 0x23, 0xd6, 0xf9, 0x8e, 0xa0, 0xaf, 0x05, 0x41, 0x3f, 0x0a, 0x82, 0x7e,
	0x15, 0x04, 0x6d, 0x0b, 0x82, 0x7e, 0x17, 0x04, 0xfd, 0x29, 0x88, 0x75, 0x5e, 0x10, 0xf4, 0x7d,
	0x4f, 0xac, 0xed, 0x9e, 0x58, 0x67, 0x7b, 0x62, 0x7d, 0x7a, 0xb1, 0xe0, 0x7a, 0xb9, 0x99, 0x7b,
	0xb1, 0xcc, 0x7c, 0xb5, 0xe4, 0x42, 0x3e, 0x1f, 0x8f, 0x4f, 0xfd, 0xb5, 0x92, 0x0b, 0x45, 0xb3,
	0x8c, 0xf9, 0x73, 0x1a, 0xaf, 0x98, 0x48, 0xfc, 0x94, 0xcf, 0xff, 0x79, 0x58, 0xf3, 0xb6, 0x79,
	0x52, 0xa3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xcb, 0xc0, 0xc8, 0x76, 0x03, 0x00, 0x00,
}
