// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shared/model/schedule/schedule.proto

package schedule // import "github.com/rhino1998/programme/backend/lib/model/schedule"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import routes "github.com/rhino1998/programme/backend/lib/model/routes"

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

type Day struct {
	Datetime int64   `protobuf:"varint,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Tasks    []*Task `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
	// Types that are valid to be assigned to Weather:
	//	*Day_WeatherNull
	Weather isDay_Weather `protobuf_oneof:"weather"`
	// Types that are valid to be assigned to Traffic:
	//	*Day_TrafficNull
	Traffic isDay_Traffic `protobuf_oneof:"traffic"`
}

func (m *Day) Reset()      { *m = Day{} }
func (*Day) ProtoMessage() {}
func (*Day) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_a5fb929fc46108ca, []int{0}
}
func (m *Day) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Day) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Day.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Day) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Day.Merge(dst, src)
}
func (m *Day) XXX_Size() int {
	return m.Size()
}
func (m *Day) XXX_DiscardUnknown() {
	xxx_messageInfo_Day.DiscardUnknown(m)
}

var xxx_messageInfo_Day proto.InternalMessageInfo

type isDay_Weather interface {
	isDay_Weather()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}
type isDay_Traffic interface {
	isDay_Traffic()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Day_WeatherNull struct {
	WeatherNull bool `protobuf:"varint,12,opt,name=weather_null,json=weatherNull,proto3,oneof"`
}
type Day_TrafficNull struct {
	TrafficNull bool `protobuf:"varint,14,opt,name=traffic_null,json=trafficNull,proto3,oneof"`
}

func (*Day_WeatherNull) isDay_Weather() {}
func (*Day_TrafficNull) isDay_Traffic() {}

func (m *Day) GetWeather() isDay_Weather {
	if m != nil {
		return m.Weather
	}
	return nil
}
func (m *Day) GetTraffic() isDay_Traffic {
	if m != nil {
		return m.Traffic
	}
	return nil
}

func (m *Day) GetDatetime() int64 {
	if m != nil {
		return m.Datetime
	}
	return 0
}

func (m *Day) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Day) GetWeatherNull() bool {
	if x, ok := m.GetWeather().(*Day_WeatherNull); ok {
		return x.WeatherNull
	}
	return false
}

func (m *Day) GetTrafficNull() bool {
	if x, ok := m.GetTraffic().(*Day_TrafficNull); ok {
		return x.TrafficNull
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Day) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Day_OneofMarshaler, _Day_OneofUnmarshaler, _Day_OneofSizer, []interface{}{
		(*Day_WeatherNull)(nil),
		(*Day_TrafficNull)(nil),
	}
}

func _Day_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Day)
	// weather
	switch x := m.Weather.(type) {
	case *Day_WeatherNull:
		t := uint64(0)
		if x.WeatherNull {
			t = 1
		}
		_ = b.EncodeVarint(12<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Day.Weather has unexpected type %T", x)
	}
	// traffic
	switch x := m.Traffic.(type) {
	case *Day_TrafficNull:
		t := uint64(0)
		if x.TrafficNull {
			t = 1
		}
		_ = b.EncodeVarint(14<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Day.Traffic has unexpected type %T", x)
	}
	return nil
}

func _Day_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Day)
	switch tag {
	case 12: // weather.weather_null
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Weather = &Day_WeatherNull{x != 0}
		return true, err
	case 14: // traffic.traffic_null
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Traffic = &Day_TrafficNull{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _Day_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Day)
	// weather
	switch x := m.Weather.(type) {
	case *Day_WeatherNull:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// traffic
	switch x := m.Traffic.(type) {
	case *Day_TrafficNull:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Task struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Duration int64  `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Stress   int64  `protobuf:"varint,3,opt,name=stress,proto3" json:"stress,omitempty"`
	// Types that are valid to be assigned to Location:
	//	*Task_LocationNull
	//	*Task_LocationValue
	Location isTask_Location `protobuf_oneof:"location"`
}

func (m *Task) Reset()      { *m = Task{} }
func (*Task) ProtoMessage() {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_a5fb929fc46108ca, []int{1}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Task.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return m.Size()
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

type isTask_Location interface {
	isTask_Location()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Task_LocationNull struct {
	LocationNull bool `protobuf:"varint,14,opt,name=location_null,json=locationNull,proto3,oneof"`
}
type Task_LocationValue struct {
	LocationValue *routes.Location `protobuf:"bytes,15,opt,name=location_value,json=locationValue,oneof"`
}

func (*Task_LocationNull) isTask_Location()  {}
func (*Task_LocationValue) isTask_Location() {}

func (m *Task) GetLocation() isTask_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Task) GetStress() int64 {
	if m != nil {
		return m.Stress
	}
	return 0
}

func (m *Task) GetLocationNull() bool {
	if x, ok := m.GetLocation().(*Task_LocationNull); ok {
		return x.LocationNull
	}
	return false
}

func (m *Task) GetLocationValue() *routes.Location {
	if x, ok := m.GetLocation().(*Task_LocationValue); ok {
		return x.LocationValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Task) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Task_OneofMarshaler, _Task_OneofUnmarshaler, _Task_OneofSizer, []interface{}{
		(*Task_LocationNull)(nil),
		(*Task_LocationValue)(nil),
	}
}

func _Task_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Task)
	// location
	switch x := m.Location.(type) {
	case *Task_LocationNull:
		t := uint64(0)
		if x.LocationNull {
			t = 1
		}
		_ = b.EncodeVarint(14<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *Task_LocationValue:
		_ = b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LocationValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Task.Location has unexpected type %T", x)
	}
	return nil
}

func _Task_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Task)
	switch tag {
	case 14: // location.location_null
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Location = &Task_LocationNull{x != 0}
		return true, err
	case 15: // location.location_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(routes.Location)
		err := b.DecodeMessage(msg)
		m.Location = &Task_LocationValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Task_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Task)
	// location
	switch x := m.Location.(type) {
	case *Task_LocationNull:
		n += 1 // tag and wire
		n += 1
	case *Task_LocationValue:
		s := proto.Size(x.LocationValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Day)(nil), "shared.model.schedule.Day")
	proto.RegisterType((*Task)(nil), "shared.model.schedule.Task")
}
func (this *Day) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Day)
	if !ok {
		that2, ok := that.(Day)
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
	if this.Datetime != that1.Datetime {
		return false
	}
	if len(this.Tasks) != len(that1.Tasks) {
		return false
	}
	for i := range this.Tasks {
		if !this.Tasks[i].Equal(that1.Tasks[i]) {
			return false
		}
	}
	if that1.Weather == nil {
		if this.Weather != nil {
			return false
		}
	} else if this.Weather == nil {
		return false
	} else if !this.Weather.Equal(that1.Weather) {
		return false
	}
	if that1.Traffic == nil {
		if this.Traffic != nil {
			return false
		}
	} else if this.Traffic == nil {
		return false
	} else if !this.Traffic.Equal(that1.Traffic) {
		return false
	}
	return true
}
func (this *Day_WeatherNull) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Day_WeatherNull)
	if !ok {
		that2, ok := that.(Day_WeatherNull)
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
	if this.WeatherNull != that1.WeatherNull {
		return false
	}
	return true
}
func (this *Day_TrafficNull) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Day_TrafficNull)
	if !ok {
		that2, ok := that.(Day_TrafficNull)
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
	if this.TrafficNull != that1.TrafficNull {
		return false
	}
	return true
}
func (this *Task) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Task)
	if !ok {
		that2, ok := that.(Task)
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
	if this.Duration != that1.Duration {
		return false
	}
	if this.Stress != that1.Stress {
		return false
	}
	if that1.Location == nil {
		if this.Location != nil {
			return false
		}
	} else if this.Location == nil {
		return false
	} else if !this.Location.Equal(that1.Location) {
		return false
	}
	return true
}
func (this *Task_LocationNull) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Task_LocationNull)
	if !ok {
		that2, ok := that.(Task_LocationNull)
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
	if this.LocationNull != that1.LocationNull {
		return false
	}
	return true
}
func (this *Task_LocationValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Task_LocationValue)
	if !ok {
		that2, ok := that.(Task_LocationValue)
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
	if !this.LocationValue.Equal(that1.LocationValue) {
		return false
	}
	return true
}
func (this *Day) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&schedule.Day{")
	s = append(s, "Datetime: "+fmt.Sprintf("%#v", this.Datetime)+",\n")
	if this.Tasks != nil {
		s = append(s, "Tasks: "+fmt.Sprintf("%#v", this.Tasks)+",\n")
	}
	if this.Weather != nil {
		s = append(s, "Weather: "+fmt.Sprintf("%#v", this.Weather)+",\n")
	}
	if this.Traffic != nil {
		s = append(s, "Traffic: "+fmt.Sprintf("%#v", this.Traffic)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Day_WeatherNull) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&schedule.Day_WeatherNull{` +
		`WeatherNull:` + fmt.Sprintf("%#v", this.WeatherNull) + `}`}, ", ")
	return s
}
func (this *Day_TrafficNull) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&schedule.Day_TrafficNull{` +
		`TrafficNull:` + fmt.Sprintf("%#v", this.TrafficNull) + `}`}, ", ")
	return s
}
func (this *Task) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&schedule.Task{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Duration: "+fmt.Sprintf("%#v", this.Duration)+",\n")
	s = append(s, "Stress: "+fmt.Sprintf("%#v", this.Stress)+",\n")
	if this.Location != nil {
		s = append(s, "Location: "+fmt.Sprintf("%#v", this.Location)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Task_LocationNull) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&schedule.Task_LocationNull{` +
		`LocationNull:` + fmt.Sprintf("%#v", this.LocationNull) + `}`}, ", ")
	return s
}
func (this *Task_LocationValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&schedule.Task_LocationValue{` +
		`LocationValue:` + fmt.Sprintf("%#v", this.LocationValue) + `}`}, ", ")
	return s
}
func valueToGoStringSchedule(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Day) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Day) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Datetime != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Datetime))
	}
	if len(m.Tasks) > 0 {
		for _, msg := range m.Tasks {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSchedule(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Weather != nil {
		nn1, err := m.Weather.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Traffic != nil {
		nn2, err := m.Traffic.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	return i, nil
}

func (m *Day_WeatherNull) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x60
	i++
	if m.WeatherNull {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *Day_TrafficNull) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x70
	i++
	if m.TrafficNull {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Duration != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Duration))
	}
	if m.Stress != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Stress))
	}
	if m.Location != nil {
		nn3, err := m.Location.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *Task_LocationNull) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x70
	i++
	if m.LocationNull {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *Task_LocationValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.LocationValue != nil {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.LocationValue.Size()))
		n4, err := m.LocationValue.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func encodeVarintSchedule(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Day) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Datetime != 0 {
		n += 1 + sovSchedule(uint64(m.Datetime))
	}
	if len(m.Tasks) > 0 {
		for _, e := range m.Tasks {
			l = e.Size()
			n += 1 + l + sovSchedule(uint64(l))
		}
	}
	if m.Weather != nil {
		n += m.Weather.Size()
	}
	if m.Traffic != nil {
		n += m.Traffic.Size()
	}
	return n
}

func (m *Day_WeatherNull) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *Day_TrafficNull) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *Task) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovSchedule(uint64(m.Duration))
	}
	if m.Stress != 0 {
		n += 1 + sovSchedule(uint64(m.Stress))
	}
	if m.Location != nil {
		n += m.Location.Size()
	}
	return n
}

func (m *Task_LocationNull) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *Task_LocationValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LocationValue != nil {
		l = m.LocationValue.Size()
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func sovSchedule(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchedule(x uint64) (n int) {
	return sovSchedule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Day) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Day{`,
		`Datetime:` + fmt.Sprintf("%v", this.Datetime) + `,`,
		`Tasks:` + strings.Replace(fmt.Sprintf("%v", this.Tasks), "Task", "Task", 1) + `,`,
		`Weather:` + fmt.Sprintf("%v", this.Weather) + `,`,
		`Traffic:` + fmt.Sprintf("%v", this.Traffic) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Day_WeatherNull) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Day_WeatherNull{`,
		`WeatherNull:` + fmt.Sprintf("%v", this.WeatherNull) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Day_TrafficNull) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Day_TrafficNull{`,
		`TrafficNull:` + fmt.Sprintf("%v", this.TrafficNull) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Task) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Task{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Duration:` + fmt.Sprintf("%v", this.Duration) + `,`,
		`Stress:` + fmt.Sprintf("%v", this.Stress) + `,`,
		`Location:` + fmt.Sprintf("%v", this.Location) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Task_LocationNull) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Task_LocationNull{`,
		`LocationNull:` + fmt.Sprintf("%v", this.LocationNull) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Task_LocationValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Task_LocationValue{`,
		`LocationValue:` + strings.Replace(fmt.Sprintf("%v", this.LocationValue), "Location", "routes.Location", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSchedule(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Day) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: Day: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Day: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Datetime", wireType)
			}
			m.Datetime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Datetime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tasks = append(m.Tasks, &Task{})
			if err := m.Tasks[len(m.Tasks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeatherNull", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Weather = &Day_WeatherNull{b}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrafficNull", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Traffic = &Day_TrafficNull{b}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchedule
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
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stress", wireType)
			}
			m.Stress = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stress |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocationNull", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Location = &Task_LocationNull{b}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocationValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &routes.Location{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Location = &Task_LocationValue{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchedule
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
func skipSchedule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
				return 0, ErrInvalidLengthSchedule
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchedule
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
				next, err := skipSchedule(dAtA[start:])
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
	ErrInvalidLengthSchedule = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedule   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("shared/model/schedule/schedule.proto", fileDescriptor_schedule_a5fb929fc46108ca)
}

var fileDescriptor_schedule_a5fb929fc46108ca = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0xed, 0x9b, 0xcb, 0xa5, 0xd7, 0x29, 0x45, 0xb2, 0x04, 0x8a, 0x8a, 0xb0, 0xa2, 0x02,
	0x52, 0xa6, 0x44, 0x2d, 0x0b, 0x5d, 0x2b, 0x84, 0x3a, 0x20, 0x86, 0x08, 0x31, 0xb0, 0x20, 0x27,
	0x71, 0x9b, 0xa8, 0x4e, 0x5c, 0xd9, 0x0e, 0x88, 0x8d, 0x47, 0xe0, 0x31, 0x78, 0x00, 0x1e, 0x02,
	0xb6, 0x8e, 0x1d, 0x69, 0xba, 0x30, 0xf6, 0x11, 0x50, 0x1c, 0x37, 0xa2, 0x88, 0x29, 0xe7, 0x7c,
	0xe7, 0xf7, 0xe5, 0xfc, 0x49, 0xd0, 0x53, 0x95, 0x53, 0xc9, 0xb2, 0xa8, 0x14, 0x19, 0xe3, 0x91,
	0x4a, 0x73, 0x96, 0xd5, 0x9c, 0xf5, 0x41, 0xb8, 0x95, 0x42, 0x0b, 0xfc, 0xa0, 0xa3, 0x42, 0x43,
	0x85, 0xe7, 0xe2, 0xd8, 0xbf, 0x30, 0x4b, 0x51, 0x6b, 0xa6, 0xec, 0xa3, 0x33, 0x4e, 0xbe, 0x43,
	0xe4, 0xbc, 0xa4, 0x9f, 0xf1, 0x18, 0x0d, 0x32, 0xaa, 0x99, 0x2e, 0x4a, 0xe6, 0x41, 0x1f, 0x06,
	0x4e, 0xdc, 0xe7, 0x78, 0x8a, 0xee, 0x68, 0xaa, 0x36, 0xca, 0xbb, 0xf2, 0x9d, 0xc0, 0x9d, 0x3d,
	0x0a, 0xff, 0xdb, 0x2c, 0x7c, 0x4b, 0xd5, 0x26, 0xee, 0x48, 0xfc, 0x04, 0x0d, 0x3f, 0x31, 0xaa,
	0x73, 0x26, 0x3f, 0x54, 0x35, 0xe7, 0xde, 0xd0, 0x87, 0xc1, 0x60, 0x09, 0x62, 0xd7, 0xaa, 0x6f,
	0x6a, 0xce, 0x5b, 0x48, 0x4b, 0xba, 0x5a, 0x15, 0x69, 0x07, 0x8d, 0x0c, 0x04, 0x63, 0xd7, 0xaa,
	0x2d, 0xb4, 0xb8, 0x45, 0x77, 0xad, 0xa7, 0x0d, 0x6d, 0x65, 0xf2, 0x13, 0xa2, 0xeb, 0xb6, 0x1f,
	0xc6, 0xe8, 0xba, 0xa2, 0x76, 0xe6, 0xdb, 0xd8, 0xc4, 0x66, 0x97, 0x5a, 0x52, 0x5d, 0x88, 0xca,
	0xbb, 0xb2, 0xbb, 0xd8, 0x1c, 0x3f, 0x44, 0x37, 0x4a, 0x4b, 0xa6, 0x94, 0xe7, 0x98, 0x8a, 0xcd,
	0xf0, 0x33, 0x74, 0x8f, 0x8b, 0xd4, 0x30, 0x7f, 0x0f, 0x03, 0xe2, 0xe1, 0x59, 0x36, 0x23, 0xbf,
	0x42, 0xa3, 0x1e, 0xfb, 0x48, 0x79, 0xcd, 0xbc, 0xfb, 0x3e, 0x0c, 0xdc, 0xd9, 0xe3, 0xcb, 0x9b,
	0xd8, 0x13, 0xbf, 0xb6, 0xe8, 0x12, 0xc4, 0xfd, 0xdb, 0xdf, 0xb5, 0xae, 0x05, 0x42, 0x83, 0xb3,
	0xb0, 0x10, 0xbb, 0x03, 0x01, 0xfb, 0x03, 0x01, 0xa7, 0x03, 0x81, 0x5f, 0x1a, 0x02, 0xbf, 0x35,
	0x04, 0xfe, 0x68, 0x08, 0xdc, 0x35, 0x04, 0xfe, 0x6a, 0x08, 0xfc, 0xdd, 0x10, 0x70, 0x6a, 0x08,
	0xfc, 0x7a, 0x24, 0x60, 0x77, 0x24, 0x60, 0x7f, 0x24, 0xe0, 0xfd, 0x7c, 0x5d, 0xe8, 0xbc, 0x4e,
	0xc2, 0x54, 0x94, 0x91, 0xcc, 0x8b, 0x4a, 0x4c, 0xe7, 0xf3, 0x17, 0xd1, 0x56, 0x8a, 0xb5, 0xa4,
	0x65, 0xc9, 0xa2, 0x84, 0xa6, 0x1b, 0x56, 0x65, 0x11, 0x2f, 0x92, 0x7f, 0xfe, 0x9d, 0xe4, 0xc6,
	0x7c, 0xfa, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xfa, 0x50, 0x43, 0x5b, 0x02, 0x00,
	0x00,
}
