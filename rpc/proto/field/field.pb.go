// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: field.proto

package field

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
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

type MetricList struct {
	Metrics              []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MetricList) Reset()         { *m = MetricList{} }
func (m *MetricList) String() string { return proto.CompactTextString(m) }
func (*MetricList) ProtoMessage()    {}
func (*MetricList) Descriptor() ([]byte, []int) {
	return fileDescriptor_04234ff7fdd53e6e, []int{0}
}
func (m *MetricList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetricList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetricList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetricList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricList.Merge(m, src)
}
func (m *MetricList) XXX_Size() int {
	return m.Size()
}
func (m *MetricList) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricList.DiscardUnknown(m)
}

var xxx_messageInfo_MetricList proto.InternalMessageInfo

func (m *MetricList) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Metric struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tags                 string   `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Fields               []*Field `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_04234ff7fdd53e6e, []int{1}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return m.Size()
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Metric) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Metric) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Field struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Field:
	//	*Field_Sum
	//	*Field_Min
	Field                isField_Field `protobuf_oneof:"field"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_04234ff7fdd53e6e, []int{2}
}
func (m *Field) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Field.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return m.Size()
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

type isField_Field interface {
	isField_Field()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Field_Sum struct {
	Sum float64 `protobuf:"fixed64,100,opt,name=sum,proto3,oneof"`
}
type Field_Min struct {
	Min float64 `protobuf:"fixed64,101,opt,name=min,proto3,oneof"`
}

func (*Field_Sum) isField_Field() {}
func (*Field_Min) isField_Field() {}

func (m *Field) GetField() isField_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetSum() float64 {
	if x, ok := m.GetField().(*Field_Sum); ok {
		return x.Sum
	}
	return 0
}

func (m *Field) GetMin() float64 {
	if x, ok := m.GetField().(*Field_Min); ok {
		return x.Min
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Field) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Field_OneofMarshaler, _Field_OneofUnmarshaler, _Field_OneofSizer, []interface{}{
		(*Field_Sum)(nil),
		(*Field_Min)(nil),
	}
}

func _Field_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Field)
	// field
	switch x := m.Field.(type) {
	case *Field_Sum:
		_ = b.EncodeVarint(100<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.Sum))
	case *Field_Min:
		_ = b.EncodeVarint(101<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.Min))
	case nil:
	default:
		return fmt.Errorf("Field.Field has unexpected type %T", x)
	}
	return nil
}

func _Field_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Field)
	switch tag {
	case 100: // field.sum
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Field = &Field_Sum{math.Float64frombits(x)}
		return true, err
	case 101: // field.min
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Field = &Field_Min{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Field_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Field)
	// field
	switch x := m.Field.(type) {
	case *Field_Sum:
		n += 2 // tag and wire
		n += 8
	case *Field_Min:
		n += 2 // tag and wire
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*MetricList)(nil), "field.MetricList")
	proto.RegisterType((*Metric)(nil), "field.Metric")
	proto.RegisterType((*Field)(nil), "field.Field")
}

func init() { proto.RegisterFile("field.proto", fileDescriptor_04234ff7fdd53e6e) }

var fileDescriptor_04234ff7fdd53e6e = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcb, 0x4c, 0xcd,
	0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x4c, 0xb9, 0xb8, 0x7c,
	0x53, 0x4b, 0x8a, 0x32, 0x93, 0x7d, 0x32, 0x8b, 0x4b, 0x84, 0xd4, 0xb9, 0xd8, 0x73, 0xc1, 0xbc,
	0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x5e, 0x3d, 0x88, 0x1e, 0x88, 0x9a, 0x20, 0x98,
	0xac, 0x52, 0x09, 0x17, 0x1b, 0x44, 0x48, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe1, 0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x2d,
	0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x08, 0x80, 0x74, 0x94,
	0x24, 0xa6, 0x17, 0x4b, 0x30, 0x43, 0x74, 0x80, 0xd8, 0x42, 0x2a, 0x5c, 0x6c, 0x60, 0x8b, 0x8a,
	0x25, 0x58, 0xc0, 0xf6, 0xf2, 0x40, 0xed, 0x75, 0x03, 0x91, 0x41, 0x50, 0x39, 0x25, 0x1f, 0x2e,
	0x56, 0xb0, 0x00, 0x56, 0x4b, 0x85, 0xb8, 0x98, 0x8b, 0x4b, 0x73, 0x25, 0x52, 0x14, 0x18, 0x35,
	0x18, 0x3d, 0x18, 0x82, 0x40, 0x1c, 0x90, 0x58, 0x6e, 0x66, 0x9e, 0x44, 0x2a, 0x4c, 0x2c, 0x37,
	0x33, 0xcf, 0x89, 0x9d, 0x0b, 0xe2, 0x75, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0x70, 0xd0, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x57, 0xa7, 0xe8, 0x29, 0x01, 0x00, 0x00,
}

func (m *MetricList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, msg := range m.Metrics {
			dAtA[i] = 0xa
			i++
			i = encodeVarintField(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Metric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metric) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintField(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintField(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.Tags) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintField(dAtA, i, uint64(len(m.Tags)))
		i += copy(dAtA[i:], m.Tags)
	}
	if len(m.Fields) > 0 {
		for _, msg := range m.Fields {
			dAtA[i] = 0x22
			i++
			i = encodeVarintField(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Field) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Field) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintField(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Field != nil {
		nn1, err := m.Field.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Field_Sum) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa1
	i++
	dAtA[i] = 0x6
	i++
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Sum))))
	i += 8
	return i, nil
}
func (m *Field_Min) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa9
	i++
	dAtA[i] = 0x6
	i++
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Min))))
	i += 8
	return i, nil
}
func encodeVarintField(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MetricList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, e := range m.Metrics {
			l = e.Size()
			n += 1 + l + sovField(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Metric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovField(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovField(uint64(m.Timestamp))
	}
	l = len(m.Tags)
	if l > 0 {
		n += 1 + l + sovField(uint64(l))
	}
	if len(m.Fields) > 0 {
		for _, e := range m.Fields {
			l = e.Size()
			n += 1 + l + sovField(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Field) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovField(uint64(l))
	}
	if m.Field != nil {
		n += m.Field.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Field_Sum) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 10
	return n
}
func (m *Field_Min) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 10
	return n
}

func sovField(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozField(x uint64) (n int) {
	return sovField(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetricList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowField
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowField
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthField
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthField
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, &Metric{})
			if err := m.Metrics[len(m.Metrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipField(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthField
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthField
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Metric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowField
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowField
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthField
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthField
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowField
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowField
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthField
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthField
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowField
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthField
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthField
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fields = append(m.Fields, &Field{})
			if err := m.Fields[len(m.Fields)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipField(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthField
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthField
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Field) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowField
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Field: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Field: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowField
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthField
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthField
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 100:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Field = &Field_Sum{float64(math.Float64frombits(v))}
		case 101:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Field = &Field_Min{float64(math.Float64frombits(v))}
		default:
			iNdEx = preIndex
			skippy, err := skipField(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthField
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthField
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipField(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowField
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
					return 0, ErrIntOverflowField
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
					return 0, ErrIntOverflowField
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
			if length < 0 {
				return 0, ErrInvalidLengthField
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthField
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowField
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
				next, err := skipField(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthField
				}
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
	ErrInvalidLengthField = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowField   = fmt.Errorf("proto: integer overflow")
)
