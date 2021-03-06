// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pipeline.proto

package spinnakerpb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Parameter struct {
	Default              string              `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
	Description          string              `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	HasOptions           bool                `protobuf:"varint,3,opt,name=hasOptions,proto3" json:"hasOptions,omitempty"`
	Label                string              `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Name                 string              `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool                `protobuf:"varint,6,opt,name=required,proto3" json:"required,omitempty"`
	Options              []*Parameter_Option `protobuf:"bytes,7,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{0}
}
func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return m.Size()
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *Parameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Parameter) GetHasOptions() bool {
	if m != nil {
		return m.HasOptions
	}
	return false
}

func (m *Parameter) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *Parameter) GetOptions() []*Parameter_Option {
	if m != nil {
		return m.Options
	}
	return nil
}

type Parameter_Option struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter_Option) Reset()         { *m = Parameter_Option{} }
func (m *Parameter_Option) String() string { return proto.CompactTextString(m) }
func (*Parameter_Option) ProtoMessage()    {}
func (*Parameter_Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{0, 0}
}
func (m *Parameter_Option) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Parameter_Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Parameter_Option.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Parameter_Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter_Option.Merge(m, src)
}
func (m *Parameter_Option) XXX_Size() int {
	return m.Size()
}
func (m *Parameter_Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter_Option proto.InternalMessageInfo

func (m *Parameter_Option) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Pipeline struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Application          string              `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	Name                 string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ExpectedArtifacts    []*ExpectedArtifact `protobuf:"bytes,4,rep,name=expectedArtifacts,proto3" json:"expectedArtifacts,omitempty"`
	KeepWaitingPipelines bool                `protobuf:"varint,5,opt,name=keepWaitingPipelines,proto3" json:"keepWaitingPipelines,omitempty"`
	LastModifiedBy       string              `protobuf:"bytes,6,opt,name=lastModifiedBy,proto3" json:"lastModifiedBy,omitempty"`
	LimitConcurrent      bool                `protobuf:"varint,7,opt,name=limitConcurrent,proto3" json:"limitConcurrent,omitempty"`
	Notifications        []*Notification     `protobuf:"bytes,8,rep,name=notifications,proto3" json:"notifications,omitempty"`
	ParameterConfig      []*Parameter        `protobuf:"bytes,9,rep,name=parameterConfig,proto3" json:"parameterConfig,omitempty"`
	Stages               []*Stage            `protobuf:"bytes,10,rep,name=stages,proto3" json:"stages,omitempty"`
	Triggers             []*Trigger          `protobuf:"bytes,11,rep,name=triggers,proto3" json:"triggers,omitempty"`
	UpdateTs             uint64              `protobuf:"varint,12,opt,name=updateTs,proto3" json:"updateTs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}
func (*Pipeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{1}
}
func (m *Pipeline) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pipeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pipeline.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipeline.Merge(m, src)
}
func (m *Pipeline) XXX_Size() int {
	return m.Size()
}
func (m *Pipeline) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipeline.DiscardUnknown(m)
}

var xxx_messageInfo_Pipeline proto.InternalMessageInfo

func (m *Pipeline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pipeline) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *Pipeline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pipeline) GetExpectedArtifacts() []*ExpectedArtifact {
	if m != nil {
		return m.ExpectedArtifacts
	}
	return nil
}

func (m *Pipeline) GetKeepWaitingPipelines() bool {
	if m != nil {
		return m.KeepWaitingPipelines
	}
	return false
}

func (m *Pipeline) GetLastModifiedBy() string {
	if m != nil {
		return m.LastModifiedBy
	}
	return ""
}

func (m *Pipeline) GetLimitConcurrent() bool {
	if m != nil {
		return m.LimitConcurrent
	}
	return false
}

func (m *Pipeline) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *Pipeline) GetParameterConfig() []*Parameter {
	if m != nil {
		return m.ParameterConfig
	}
	return nil
}

func (m *Pipeline) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

func (m *Pipeline) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func (m *Pipeline) GetUpdateTs() uint64 {
	if m != nil {
		return m.UpdateTs
	}
	return 0
}

func init() {
	proto.RegisterType((*Parameter)(nil), "spinnakerpb.Parameter")
	proto.RegisterType((*Parameter_Option)(nil), "spinnakerpb.Parameter.Option")
	proto.RegisterType((*Pipeline)(nil), "spinnakerpb.Pipeline")
}

func init() { proto.RegisterFile("pipeline.proto", fileDescriptor_7ac67a7adf3df9c7) }

var fileDescriptor_7ac67a7adf3df9c7 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcf, 0x8e, 0xd3, 0x3c,
	0x10, 0x57, 0xda, 0x6e, 0x9b, 0x4e, 0x77, 0xbb, 0xfa, 0xac, 0xea, 0x93, 0xa9, 0x44, 0x14, 0xed,
	0x01, 0x45, 0x1c, 0x2a, 0xb4, 0x1c, 0x38, 0x02, 0xbb, 0xe2, 0x84, 0x80, 0x95, 0x59, 0x89, 0xb3,
	0x9b, 0x4c, 0x82, 0xb5, 0xa9, 0x63, 0x6c, 0x07, 0xc1, 0x0b, 0xf0, 0x1c, 0x3c, 0x0e, 0x47, 0x1e,
	0x01, 0xf5, 0x49, 0x50, 0x1c, 0x27, 0x4a, 0x4b, 0x6f, 0xf9, 0xfd, 0x99, 0xb1, 0x7f, 0xe3, 0x09,
	0x2c, 0x95, 0x50, 0x58, 0x0a, 0x89, 0x1b, 0xa5, 0x2b, 0x5b, 0x91, 0x85, 0x51, 0x42, 0x4a, 0xfe,
	0x80, 0x5a, 0x6d, 0xd7, 0x4b, 0xae, 0xad, 0xc8, 0x79, 0x6a, 0x5b, 0x71, 0x4d, 0x64, 0x65, 0x45,
	0x2e, 0x52, 0x6e, 0x45, 0x25, 0x3d, 0xb7, 0x30, 0x96, 0x17, 0xbe, 0x7a, 0x7d, 0x61, 0xb5, 0x28,
	0x0a, 0xd4, 0x2d, 0xbc, 0xfa, 0x31, 0x82, 0xf9, 0x1d, 0xd7, 0x7c, 0x87, 0x16, 0x35, 0xa1, 0x30,
	0xcb, 0x30, 0xe7, 0x75, 0x69, 0x69, 0x10, 0x07, 0xc9, 0x9c, 0x75, 0x90, 0xc4, 0xb0, 0xc8, 0xd0,
	0xa4, 0x5a, 0xa8, 0xa6, 0x31, 0x1d, 0x39, 0x75, 0x48, 0x91, 0x08, 0xe0, 0x33, 0x37, 0x1f, 0x1c,
	0x30, 0x74, 0x1c, 0x07, 0x49, 0xc8, 0x06, 0x0c, 0x59, 0xc1, 0x59, 0xc9, 0xb7, 0x58, 0xd2, 0x89,
	0xab, 0x6d, 0x01, 0x21, 0x30, 0x91, 0x7c, 0x87, 0xf4, 0xcc, 0x91, 0xee, 0x9b, 0xac, 0x21, 0xd4,
	0xf8, 0xa5, 0x16, 0x1a, 0x33, 0x3a, 0x75, 0x7d, 0x7a, 0x4c, 0x5e, 0xc0, 0xac, 0xf2, 0x47, 0xcc,
	0xe2, 0x71, 0xb2, 0xb8, 0x7e, 0xbc, 0x19, 0x8c, 0x63, 0xd3, 0x47, 0xd9, 0xb4, 0xc7, 0xb2, 0xce,
	0xbd, 0x8e, 0x60, 0xda, 0x52, 0xcd, 0x45, 0xbe, 0xf2, 0xb2, 0x46, 0x1f, 0xb1, 0x05, 0x57, 0x3f,
	0x27, 0x10, 0xde, 0xf9, 0x41, 0x93, 0x25, 0x8c, 0x44, 0xe6, 0xf5, 0x91, 0xc8, 0x9a, 0xf4, 0x5c,
	0xa9, 0xd2, 0x8f, 0xb5, 0x4b, 0x3f, 0xa0, 0xfa, 0x1c, 0xe3, 0x41, 0x8e, 0xb7, 0xf0, 0x1f, 0x7e,
	0x53, 0x98, 0x5a, 0xcc, 0x5e, 0xfb, 0x57, 0x32, 0x74, 0x72, 0xe2, 0xd6, 0x6f, 0x8e, 0x5c, 0xec,
	0xdf, 0x3a, 0x72, 0x0d, 0xab, 0x07, 0x44, 0xf5, 0x89, 0x0b, 0x2b, 0x64, 0xd1, 0xdd, 0xd4, 0xb8,
	0xc1, 0x85, 0xec, 0xa4, 0x46, 0x9e, 0xc0, 0xb2, 0xe4, 0xc6, 0xbe, 0xab, 0x32, 0x91, 0x0b, 0xcc,
	0x6e, 0xbe, 0xbb, 0x71, 0xce, 0xd9, 0x11, 0x4b, 0x12, 0xb8, 0x2c, 0xc5, 0x4e, 0xd8, 0xdb, 0x4a,
	0xa6, 0xb5, 0xd6, 0x28, 0x2d, 0x9d, 0xb9, 0xb6, 0xc7, 0x34, 0x79, 0x09, 0x17, 0xc3, 0x05, 0x33,
	0x34, 0x74, 0x71, 0x1e, 0x1d, 0xc4, 0x79, 0x3f, 0x70, 0xb0, 0x43, 0x3f, 0x79, 0x05, 0x97, 0xaa,
	0x7b, 0xa3, 0xdb, 0x4a, 0xe6, 0xa2, 0xa0, 0x73, 0xd7, 0xe2, 0xff, 0xd3, 0xef, 0xc8, 0x8e, 0xed,
	0xe4, 0x29, 0x4c, 0xdd, 0x3e, 0x1b, 0x0a, 0xae, 0x90, 0x1c, 0x14, 0x7e, 0x6c, 0x24, 0xe6, 0x1d,
	0xe4, 0x19, 0x84, 0x7e, 0xdd, 0x0d, 0x5d, 0x38, 0xf7, 0xea, 0xc0, 0x7d, 0xdf, 0x8a, 0xac, 0x77,
	0x35, 0xbb, 0x57, 0xab, 0x8c, 0x5b, 0xbc, 0x37, 0xf4, 0x3c, 0x0e, 0x92, 0x09, 0xeb, 0xf1, 0xcd,
	0xf9, 0xaf, 0x7d, 0x14, 0xfc, 0xde, 0x47, 0xc1, 0x9f, 0x7d, 0x14, 0x6c, 0xa7, 0xee, 0x07, 0x7a,
	0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x05, 0x3f, 0xc9, 0x86, 0x9f, 0x03, 0x00, 0x00,
}

func (m *Parameter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Parameter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Parameter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Options[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPipeline(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Required {
		i--
		if m.Required {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x22
	}
	if m.HasOptions {
		i--
		if m.HasOptions {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Default) > 0 {
		i -= len(m.Default)
		copy(dAtA[i:], m.Default)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Default)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Parameter_Option) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Parameter_Option) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Parameter_Option) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pipeline) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pipeline) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pipeline) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdateTs != 0 {
		i = encodeVarintPipeline(dAtA, i, uint64(m.UpdateTs))
		i--
		dAtA[i] = 0x60
	}
	if len(m.Triggers) > 0 {
		for iNdEx := len(m.Triggers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Triggers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPipeline(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Stages) > 0 {
		for iNdEx := len(m.Stages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Stages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPipeline(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.ParameterConfig) > 0 {
		for iNdEx := len(m.ParameterConfig) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParameterConfig[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPipeline(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Notifications) > 0 {
		for iNdEx := len(m.Notifications) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Notifications[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPipeline(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.LimitConcurrent {
		i--
		if m.LimitConcurrent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.LastModifiedBy) > 0 {
		i -= len(m.LastModifiedBy)
		copy(dAtA[i:], m.LastModifiedBy)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.LastModifiedBy)))
		i--
		dAtA[i] = 0x32
	}
	if m.KeepWaitingPipelines {
		i--
		if m.KeepWaitingPipelines {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.ExpectedArtifacts) > 0 {
		for iNdEx := len(m.ExpectedArtifacts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExpectedArtifacts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPipeline(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Application) > 0 {
		i -= len(m.Application)
		copy(dAtA[i:], m.Application)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Application)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPipeline(dAtA []byte, offset int, v uint64) int {
	offset -= sovPipeline(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Parameter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Default)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.HasOptions {
		n += 2
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.Required {
		n += 2
	}
	if len(m.Options) > 0 {
		for _, e := range m.Options {
			l = e.Size()
			n += 1 + l + sovPipeline(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Parameter_Option) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Pipeline) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	l = len(m.Application)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	if len(m.ExpectedArtifacts) > 0 {
		for _, e := range m.ExpectedArtifacts {
			l = e.Size()
			n += 1 + l + sovPipeline(uint64(l))
		}
	}
	if m.KeepWaitingPipelines {
		n += 2
	}
	l = len(m.LastModifiedBy)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.LimitConcurrent {
		n += 2
	}
	if len(m.Notifications) > 0 {
		for _, e := range m.Notifications {
			l = e.Size()
			n += 1 + l + sovPipeline(uint64(l))
		}
	}
	if len(m.ParameterConfig) > 0 {
		for _, e := range m.ParameterConfig {
			l = e.Size()
			n += 1 + l + sovPipeline(uint64(l))
		}
	}
	if len(m.Stages) > 0 {
		for _, e := range m.Stages {
			l = e.Size()
			n += 1 + l + sovPipeline(uint64(l))
		}
	}
	if len(m.Triggers) > 0 {
		for _, e := range m.Triggers {
			l = e.Size()
			n += 1 + l + sovPipeline(uint64(l))
		}
	}
	if m.UpdateTs != 0 {
		n += 1 + sovPipeline(uint64(m.UpdateTs))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPipeline(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPipeline(x uint64) (n int) {
	return sovPipeline(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Parameter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPipeline
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
			return fmt.Errorf("proto: Parameter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Parameter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Default", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Default = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasOptions", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasOptions = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Required", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Required = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, &Parameter_Option{})
			if err := m.Options[len(m.Options)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPipeline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPipeline
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPipeline
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
func (m *Parameter_Option) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPipeline
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
			return fmt.Errorf("proto: Option: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Option: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPipeline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPipeline
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPipeline
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
func (m *Pipeline) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPipeline
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
			return fmt.Errorf("proto: Pipeline: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pipeline: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Application = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedArtifacts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpectedArtifacts = append(m.ExpectedArtifacts, &ExpectedArtifact{})
			if err := m.ExpectedArtifacts[len(m.ExpectedArtifacts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeepWaitingPipelines", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.KeepWaitingPipelines = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModifiedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastModifiedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitConcurrent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LimitConcurrent = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notifications", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notifications = append(m.Notifications, &Notification{})
			if err := m.Notifications[len(m.Notifications)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParameterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParameterConfig = append(m.ParameterConfig, &Parameter{})
			if err := m.ParameterConfig[len(m.ParameterConfig)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stages = append(m.Stages, &Stage{})
			if err := m.Stages[len(m.Stages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Triggers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
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
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPipeline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Triggers = append(m.Triggers, &Trigger{})
			if err := m.Triggers[len(m.Triggers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTs", wireType)
			}
			m.UpdateTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateTs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPipeline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPipeline
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPipeline
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
func skipPipeline(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPipeline
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
					return 0, ErrIntOverflowPipeline
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
					return 0, ErrIntOverflowPipeline
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
				return 0, ErrInvalidLengthPipeline
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPipeline
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPipeline
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
				next, err := skipPipeline(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPipeline
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
	ErrInvalidLengthPipeline = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPipeline   = fmt.Errorf("proto: integer overflow")
)
