// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shared/model/routes/routes.proto

package routes // import "github.com/rhino1998/programme/backend/lib/model/routes"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

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
	Walking   TravelMethod = 0
	Biking    TravelMethod = 1
	Rideshare TravelMethod = 2
	Car       TravelMethod = 3
)

var TravelMethod_name = map[int32]string{
	0: "Walking",
	1: "Biking",
	2: "Rideshare",
	3: "Car",
}
var TravelMethod_value = map[string]int32{
	"Walking":   0,
	"Biking":    1,
	"Rideshare": 2,
	"Car":       3,
}

func (TravelMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_routes_f6d447a3dfa46f4f, []int{0}
}

type Location struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Location) Reset()      { *m = Location{} }
func (*Location) ProtoMessage() {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_routes_f6d447a3dfa46f4f, []int{0}
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

type Trip struct {
	Start    *Location    `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End      *Location    `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	Method   TravelMethod `protobuf:"varint,3,opt,name=method,proto3,enum=shared.model.routes.TravelMethod" json:"method,omitempty"`
	Duration int64        `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *Trip) Reset()      { *m = Trip{} }
func (*Trip) ProtoMessage() {}
func (*Trip) Descriptor() ([]byte, []int) {
	return fileDescriptor_routes_f6d447a3dfa46f4f, []int{1}
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
	return Walking
}

func (m *Trip) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*Location)(nil), "shared.model.routes.Location")
	proto.RegisterType((*Trip)(nil), "shared.model.routes.Trip")
	proto.RegisterEnum("shared.model.routes.TravelMethod", TravelMethod_name, TravelMethod_value)
}
func (x TravelMethod) String() string {
	s, ok := TravelMethod_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
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
func (this *Location) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&routes.Location{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
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
func valueToGoStringRoutes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
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
		n1, err := m.Start.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.End != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoutes(dAtA, i, uint64(m.End.Size()))
		n2, err := m.End.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
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

func encodeVarintRoutes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
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
func (this *Location) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Location{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
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
func valueToStringRoutes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
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
	proto.RegisterFile("shared/model/routes/routes.proto", fileDescriptor_routes_f6d447a3dfa46f4f)
}

var fileDescriptor_routes_f6d447a3dfa46f4f = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0x2a, 0x41,
	0x14, 0xc6, 0x77, 0x58, 0x2e, 0x7f, 0x0e, 0xf7, 0xde, 0x90, 0xb9, 0xcd, 0xe6, 0x26, 0x9e, 0xac,
	0x54, 0xc4, 0x62, 0x37, 0x42, 0xa1, 0x54, 0x26, 0xd8, 0x6a, 0xb3, 0x21, 0x31, 0xb1, 0x9b, 0x65,
	0x27, 0xec, 0x84, 0x9d, 0x1d, 0x32, 0x0c, 0xd6, 0x3e, 0x82, 0x8f, 0xe1, 0x7b, 0xd8, 0x58, 0x52,
	0x52, 0xca, 0xd0, 0x58, 0xf2, 0x08, 0xc6, 0x01, 0x8d, 0x26, 0x14, 0x56, 0x73, 0x26, 0xe7, 0xf7,
	0xfb, 0x4e, 0xf1, 0x41, 0x38, 0xcf, 0x99, 0xe6, 0x59, 0x2c, 0x55, 0xc6, 0x8b, 0x58, 0xab, 0x85,
	0xe1, 0xf3, 0xfd, 0x13, 0xcd, 0xb4, 0x32, 0x8a, 0xfe, 0xdb, 0x11, 0x91, 0x23, 0xa2, 0xdd, 0xaa,
	0x83, 0xd0, 0xb8, 0x52, 0x63, 0x66, 0x84, 0x2a, 0x29, 0x85, 0x6a, 0xc9, 0x24, 0x0f, 0x48, 0x48,
	0xba, 0xcd, 0xc4, 0xcd, 0x9d, 0x27, 0x02, 0xd5, 0x91, 0x16, 0x33, 0xda, 0x87, 0x5f, 0x73, 0xc3,
	0xb4, 0x71, 0xdb, 0x56, 0xef, 0x28, 0x3a, 0x90, 0x16, 0x7d, 0x44, 0x25, 0x3b, 0x96, 0xc6, 0xe0,
	0xf3, 0x32, 0x0b, 0x2a, 0x3f, 0x51, 0xde, 0x49, 0x3a, 0x80, 0x9a, 0xe4, 0x26, 0x57, 0x59, 0xe0,
	0x87, 0xa4, 0xfb, 0xb7, 0x77, 0x7c, 0xd0, 0x19, 0x69, 0x76, 0xc7, 0x8b, 0x6b, 0x07, 0x26, 0x7b,
	0x81, 0xfe, 0x87, 0x46, 0xb6, 0xd0, 0x2e, 0x2b, 0xa8, 0x86, 0xa4, 0xeb, 0x27, 0x9f, 0xff, 0x93,
	0x0b, 0xf8, 0xfd, 0xd5, 0xa1, 0x2d, 0xa8, 0xdf, 0xb0, 0x62, 0x2a, 0xca, 0x49, 0xdb, 0xa3, 0x00,
	0xb5, 0xa1, 0x70, 0x33, 0xa1, 0x7f, 0xa0, 0x99, 0x88, 0x8c, 0xbb, 0xa3, 0xed, 0x0a, 0xad, 0x83,
	0x7f, 0xc9, 0x74, 0xdb, 0x1f, 0xca, 0xe5, 0x1a, 0xbd, 0xd5, 0x1a, 0xbd, 0xed, 0x1a, 0xc9, 0xbd,
	0x45, 0xf2, 0x68, 0x91, 0x3c, 0x5b, 0x24, 0x4b, 0x8b, 0xe4, 0xc5, 0x22, 0x79, 0xb5, 0xe8, 0x6d,
	0x2d, 0x92, 0x87, 0x0d, 0x7a, 0xcb, 0x0d, 0x7a, 0xab, 0x0d, 0x7a, 0xb7, 0x67, 0x13, 0x61, 0xf2,
	0x45, 0x1a, 0x8d, 0x95, 0x8c, 0x75, 0x2e, 0x4a, 0x75, 0x3a, 0x18, 0x9c, 0xc7, 0x33, 0xad, 0x26,
	0x9a, 0x49, 0xc9, 0xe3, 0x94, 0x8d, 0xa7, 0xbc, 0xcc, 0xe2, 0x42, 0xa4, 0xdf, 0x7a, 0x4b, 0x6b,
	0xae, 0xb1, 0xfe, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x83, 0x19, 0xf5, 0xd5, 0x01, 0x00,
	0x00,
}