// Code generated by protoc-gen-go.
// source: history/action.proto
// DO NOT EDIT!

/*
Package ProtobufHistory is a generated protocol buffer package.

It is generated from these files:
	history/action.proto
	history/frame.proto
	history/history.proto

It has these top-level messages:
	DocAction
	InitAction
	NudgeAction
	ScaleAction
	RotateAction
	NudgeHandleAction
	AddHandleAction
	RemoveHandleAction
	InsertAction
	HistoryFrame
	DocHistory
*/
package ProtobufHistory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import index "geometry"
import posn "geometry"
import ProtobufHistory1 "geometry"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DocAction struct {
	// Types that are valid to be assigned to Action:
	//	*DocAction_InitAction
	//	*DocAction_NudgeAction
	//	*DocAction_ScaleAction
	//	*DocAction_RotateAction
	Action isDocAction_Action `protobuf_oneof:"action"`
}

func (m *DocAction) Reset()                    { *m = DocAction{} }
func (m *DocAction) String() string            { return proto.CompactTextString(m) }
func (*DocAction) ProtoMessage()               {}
func (*DocAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isDocAction_Action interface {
	isDocAction_Action()
}

type DocAction_InitAction struct {
	InitAction *InitAction `protobuf:"bytes,1,opt,name=initAction,oneof"`
}
type DocAction_NudgeAction struct {
	NudgeAction *NudgeAction `protobuf:"bytes,2,opt,name=nudgeAction,oneof"`
}
type DocAction_ScaleAction struct {
	ScaleAction *ScaleAction `protobuf:"bytes,3,opt,name=scaleAction,oneof"`
}
type DocAction_RotateAction struct {
	RotateAction *RotateAction `protobuf:"bytes,4,opt,name=rotateAction,oneof"`
}

func (*DocAction_InitAction) isDocAction_Action()   {}
func (*DocAction_NudgeAction) isDocAction_Action()  {}
func (*DocAction_ScaleAction) isDocAction_Action()  {}
func (*DocAction_RotateAction) isDocAction_Action() {}

func (m *DocAction) GetAction() isDocAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *DocAction) GetInitAction() *InitAction {
	if x, ok := m.GetAction().(*DocAction_InitAction); ok {
		return x.InitAction
	}
	return nil
}

func (m *DocAction) GetNudgeAction() *NudgeAction {
	if x, ok := m.GetAction().(*DocAction_NudgeAction); ok {
		return x.NudgeAction
	}
	return nil
}

func (m *DocAction) GetScaleAction() *ScaleAction {
	if x, ok := m.GetAction().(*DocAction_ScaleAction); ok {
		return x.ScaleAction
	}
	return nil
}

func (m *DocAction) GetRotateAction() *RotateAction {
	if x, ok := m.GetAction().(*DocAction_RotateAction); ok {
		return x.RotateAction
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DocAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DocAction_OneofMarshaler, _DocAction_OneofUnmarshaler, _DocAction_OneofSizer, []interface{}{
		(*DocAction_InitAction)(nil),
		(*DocAction_NudgeAction)(nil),
		(*DocAction_ScaleAction)(nil),
		(*DocAction_RotateAction)(nil),
	}
}

func _DocAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DocAction)
	// action
	switch x := m.Action.(type) {
	case *DocAction_InitAction:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitAction); err != nil {
			return err
		}
	case *DocAction_NudgeAction:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NudgeAction); err != nil {
			return err
		}
	case *DocAction_ScaleAction:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ScaleAction); err != nil {
			return err
		}
	case *DocAction_RotateAction:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RotateAction); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DocAction.Action has unexpected type %T", x)
	}
	return nil
}

func _DocAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DocAction)
	switch tag {
	case 1: // action.initAction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InitAction)
		err := b.DecodeMessage(msg)
		m.Action = &DocAction_InitAction{msg}
		return true, err
	case 2: // action.nudgeAction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NudgeAction)
		err := b.DecodeMessage(msg)
		m.Action = &DocAction_NudgeAction{msg}
		return true, err
	case 3: // action.scaleAction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ScaleAction)
		err := b.DecodeMessage(msg)
		m.Action = &DocAction_ScaleAction{msg}
		return true, err
	case 4: // action.rotateAction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RotateAction)
		err := b.DecodeMessage(msg)
		m.Action = &DocAction_RotateAction{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DocAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DocAction)
	// action
	switch x := m.Action.(type) {
	case *DocAction_InitAction:
		s := proto.Size(x.InitAction)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DocAction_NudgeAction:
		s := proto.Size(x.NudgeAction)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DocAction_ScaleAction:
		s := proto.Size(x.ScaleAction)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DocAction_RotateAction:
		s := proto.Size(x.RotateAction)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InitAction struct {
}

func (m *InitAction) Reset()                    { *m = InitAction{} }
func (m *InitAction) String() string            { return proto.CompactTextString(m) }
func (*InitAction) ProtoMessage()               {}
func (*InitAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NudgeAction struct {
	Indexes []*index.Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	Xd      float64        `protobuf:"fixed64,2,opt,name=xd" json:"xd,omitempty"`
	Yd      float64        `protobuf:"fixed64,3,opt,name=yd" json:"yd,omitempty"`
}

func (m *NudgeAction) Reset()                    { *m = NudgeAction{} }
func (m *NudgeAction) String() string            { return proto.CompactTextString(m) }
func (*NudgeAction) ProtoMessage()               {}
func (*NudgeAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NudgeAction) GetIndexes() []*index.Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *NudgeAction) GetXd() float64 {
	if m != nil {
		return m.Xd
	}
	return 0
}

func (m *NudgeAction) GetYd() float64 {
	if m != nil {
		return m.Yd
	}
	return 0
}

type ScaleAction struct {
	Indexes []*index.Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	Xs      float64        `protobuf:"fixed64,2,opt,name=xs" json:"xs,omitempty"`
	Ys      float64        `protobuf:"fixed64,3,opt,name=ys" json:"ys,omitempty"`
	Origin  *posn.Posn     `protobuf:"bytes,4,opt,name=origin" json:"origin,omitempty"`
}

func (m *ScaleAction) Reset()                    { *m = ScaleAction{} }
func (m *ScaleAction) String() string            { return proto.CompactTextString(m) }
func (*ScaleAction) ProtoMessage()               {}
func (*ScaleAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ScaleAction) GetIndexes() []*index.Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *ScaleAction) GetXs() float64 {
	if m != nil {
		return m.Xs
	}
	return 0
}

func (m *ScaleAction) GetYs() float64 {
	if m != nil {
		return m.Ys
	}
	return 0
}

func (m *ScaleAction) GetOrigin() *posn.Posn {
	if m != nil {
		return m.Origin
	}
	return nil
}

type RotateAction struct {
	Indexes []*index.Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	Angle   float64        `protobuf:"fixed64,2,opt,name=angle" json:"angle,omitempty"`
	Origin  *posn.Posn     `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
}

func (m *RotateAction) Reset()                    { *m = RotateAction{} }
func (m *RotateAction) String() string            { return proto.CompactTextString(m) }
func (*RotateAction) ProtoMessage()               {}
func (*RotateAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RotateAction) GetIndexes() []*index.Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *RotateAction) GetAngle() float64 {
	if m != nil {
		return m.Angle
	}
	return 0
}

func (m *RotateAction) GetOrigin() *posn.Posn {
	if m != nil {
		return m.Origin
	}
	return nil
}

type NudgeHandleAction struct {
	Indexes []*index.Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	Xd      float64        `protobuf:"fixed64,2,opt,name=xd" json:"xd,omitempty"`
	Yd      float64        `protobuf:"fixed64,3,opt,name=yd" json:"yd,omitempty"`
	Handle  posn.Handle    `protobuf:"varint,4,opt,name=handle,enum=Handle" json:"handle,omitempty"`
}

func (m *NudgeHandleAction) Reset()                    { *m = NudgeHandleAction{} }
func (m *NudgeHandleAction) String() string            { return proto.CompactTextString(m) }
func (*NudgeHandleAction) ProtoMessage()               {}
func (*NudgeHandleAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NudgeHandleAction) GetIndexes() []*index.Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *NudgeHandleAction) GetXd() float64 {
	if m != nil {
		return m.Xd
	}
	return 0
}

func (m *NudgeHandleAction) GetYd() float64 {
	if m != nil {
		return m.Yd
	}
	return 0
}

func (m *NudgeHandleAction) GetHandle() posn.Handle {
	if m != nil {
		return m.Handle
	}
	return posn.Handle_UNKNOWN
}

type AddHandleAction struct {
	Indexes []*index.Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	Handle  posn.Handle    `protobuf:"varint,2,opt,name=handle,enum=Handle" json:"handle,omitempty"`
	Posn    *posn.Posn     `protobuf:"bytes,3,opt,name=posn" json:"posn,omitempty"`
	Reflect bool           `protobuf:"varint,4,opt,name=reflect" json:"reflect,omitempty"`
}

func (m *AddHandleAction) Reset()                    { *m = AddHandleAction{} }
func (m *AddHandleAction) String() string            { return proto.CompactTextString(m) }
func (*AddHandleAction) ProtoMessage()               {}
func (*AddHandleAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddHandleAction) GetIndexes() []*index.Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *AddHandleAction) GetHandle() posn.Handle {
	if m != nil {
		return m.Handle
	}
	return posn.Handle_UNKNOWN
}

func (m *AddHandleAction) GetPosn() *posn.Posn {
	if m != nil {
		return m.Posn
	}
	return nil
}

func (m *AddHandleAction) GetReflect() bool {
	if m != nil {
		return m.Reflect
	}
	return false
}

type RemoveHandleAction struct {
	Indexes []*index.Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	Handle  posn.Handle    `protobuf:"varint,2,opt,name=handle,enum=Handle" json:"handle,omitempty"`
	Reflect bool           `protobuf:"varint,3,opt,name=reflect" json:"reflect,omitempty"`
}

func (m *RemoveHandleAction) Reset()                    { *m = RemoveHandleAction{} }
func (m *RemoveHandleAction) String() string            { return proto.CompactTextString(m) }
func (*RemoveHandleAction) ProtoMessage()               {}
func (*RemoveHandleAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RemoveHandleAction) GetIndexes() []*index.Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *RemoveHandleAction) GetHandle() posn.Handle {
	if m != nil {
		return m.Handle
	}
	return posn.Handle_UNKNOWN
}

func (m *RemoveHandleAction) GetReflect() bool {
	if m != nil {
		return m.Reflect
	}
	return false
}

type InsertAction struct {
	Items []*InsertAction_ItemIndex `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *InsertAction) Reset()                    { *m = InsertAction{} }
func (m *InsertAction) String() string            { return proto.CompactTextString(m) }
func (*InsertAction) ProtoMessage()               {}
func (*InsertAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *InsertAction) GetItems() []*InsertAction_ItemIndex {
	if m != nil {
		return m.Items
	}
	return nil
}

type InsertAction_ItemIndex struct {
	Item  *ProtobufHistory1.Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
	Index *index.Index           `protobuf:"bytes,2,opt,name=index" json:"index,omitempty"`
}

func (m *InsertAction_ItemIndex) Reset()                    { *m = InsertAction_ItemIndex{} }
func (m *InsertAction_ItemIndex) String() string            { return proto.CompactTextString(m) }
func (*InsertAction_ItemIndex) ProtoMessage()               {}
func (*InsertAction_ItemIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

func (m *InsertAction_ItemIndex) GetItem() *ProtobufHistory1.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *InsertAction_ItemIndex) GetIndex() *index.Index {
	if m != nil {
		return m.Index
	}
	return nil
}

func init() {
	proto.RegisterType((*DocAction)(nil), "ProtobufHistory.DocAction")
	proto.RegisterType((*InitAction)(nil), "ProtobufHistory.InitAction")
	proto.RegisterType((*NudgeAction)(nil), "ProtobufHistory.NudgeAction")
	proto.RegisterType((*ScaleAction)(nil), "ProtobufHistory.ScaleAction")
	proto.RegisterType((*RotateAction)(nil), "ProtobufHistory.RotateAction")
	proto.RegisterType((*NudgeHandleAction)(nil), "ProtobufHistory.NudgeHandleAction")
	proto.RegisterType((*AddHandleAction)(nil), "ProtobufHistory.AddHandleAction")
	proto.RegisterType((*RemoveHandleAction)(nil), "ProtobufHistory.RemoveHandleAction")
	proto.RegisterType((*InsertAction)(nil), "ProtobufHistory.InsertAction")
	proto.RegisterType((*InsertAction_ItemIndex)(nil), "ProtobufHistory.InsertAction.ItemIndex")
}

func init() { proto.RegisterFile("history/action.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xd2, 0x36, 0xdd, 0x6e, 0xaa, 0x4d, 0x98, 0x22, 0x95, 0xb2, 0x89, 0x2a, 0x2f, 0x8c,
	0x97, 0x4c, 0x2a, 0xcf, 0x48, 0x0c, 0x78, 0x48, 0x5f, 0x60, 0x32, 0xfc, 0x81, 0xac, 0xf6, 0x32,
	0x4b, 0x89, 0x3d, 0x62, 0x6f, 0x6a, 0x7e, 0x02, 0x3f, 0x82, 0x27, 0xfe, 0x28, 0xb2, 0x9d, 0x0f,
	0x97, 0x14, 0x69, 0x48, 0x7d, 0x8a, 0xee, 0xc9, 0x3d, 0xe7, 0xdc, 0x73, 0xa5, 0x6b, 0x98, 0xdd,
	0x31, 0xa9, 0x44, 0x59, 0x5d, 0xa6, 0x1b, 0xc5, 0x04, 0x8f, 0xef, 0x4b, 0xa1, 0x04, 0x3a, 0xbd,
	0xd6, 0x9f, 0x9b, 0x87, 0xdb, 0xc4, 0xfe, 0x5d, 0xcc, 0x32, 0x2a, 0x0a, 0xaa, 0xca, 0xea, 0x92,
	0x71, 0x42, 0xb7, 0xb6, 0x6d, 0xf1, 0xbc, 0x45, 0xef, 0x85, 0xe4, 0x3d, 0x90, 0x29, 0x5a, 0x58,
	0x30, 0xfa, 0xe5, 0xc3, 0xf1, 0x67, 0xb1, 0xb9, 0x32, 0x26, 0xe8, 0x3d, 0x00, 0xe3, 0x4c, 0xd9,
	0x6a, 0xee, 0x2d, 0xbd, 0x8b, 0x70, 0xf5, 0x2a, 0xfe, 0xcb, 0x33, 0x5e, 0xb7, 0x2d, 0xc9, 0x00,
	0x3b, 0x04, 0xf4, 0x01, 0x42, 0xfe, 0x40, 0x32, 0x5a, 0xf3, 0x7d, 0xc3, 0x3f, 0xeb, 0xf1, 0xbf,
	0x74, 0x3d, 0xc9, 0x00, 0xbb, 0x14, 0xad, 0x20, 0x37, 0x69, 0xde, 0x28, 0x0c, 0xff, 0xa1, 0xf0,
	0xad, 0xeb, 0xd1, 0x0a, 0x0e, 0x05, 0x7d, 0x82, 0x69, 0x29, 0x54, 0xaa, 0x1a, 0x89, 0x91, 0x91,
	0x38, 0xef, 0x49, 0x60, 0xa7, 0x29, 0x19, 0xe0, 0x1d, 0xd2, 0xc7, 0x23, 0x08, 0xec, 0xda, 0xa3,
	0x29, 0x40, 0x17, 0x37, 0xfa, 0x0a, 0xa1, 0x33, 0x3c, 0x5a, 0xc2, 0xc4, 0x6c, 0x9d, 0xca, 0xb9,
	0xb7, 0x1c, 0x5e, 0x84, 0xab, 0x20, 0x5e, 0xeb, 0x1a, 0x37, 0x30, 0x3a, 0x01, 0x7f, 0x4b, 0xcc,
	0x22, 0x3c, 0xec, 0x6f, 0x89, 0xae, 0x2b, 0x62, 0x62, 0x79, 0xd8, 0xaf, 0x48, 0xc4, 0x21, 0x74,
	0xb2, 0x3c, 0x51, 0x50, 0xb6, 0x82, 0xa6, 0xae, 0x64, 0x2b, 0x28, 0xd1, 0x39, 0x04, 0xa2, 0x64,
	0x19, 0x6b, 0x82, 0x8f, 0xe3, 0x6b, 0x21, 0x39, 0xae, 0xc1, 0x88, 0xc2, 0xd4, 0x0d, 0xfe, 0x04,
	0xc3, 0x19, 0x8c, 0x53, 0x9e, 0xe5, 0xb4, 0xf6, 0xb4, 0x85, 0x63, 0x33, 0xdc, 0x67, 0xf3, 0x08,
	0xcf, 0xcc, 0x9e, 0x92, 0x94, 0x93, 0xfc, 0x60, 0xdb, 0x42, 0xaf, 0x21, 0xb8, 0x33, 0x8a, 0x26,
	0xdc, 0xc9, 0x6a, 0x12, 0x5b, 0x03, 0x5c, 0xc3, 0xd1, 0x4f, 0x0f, 0x4e, 0xaf, 0x08, 0xf9, 0x4f,
	0xdb, 0x4e, 0xd6, 0xdf, 0x2b, 0x8b, 0x5e, 0xc2, 0x48, 0xdf, 0xd1, 0x6e, 0x56, 0x03, 0xa1, 0x39,
	0x4c, 0x4a, 0x7a, 0x9b, 0xd3, 0x8d, 0x32, 0x33, 0x1d, 0xe1, 0xa6, 0x8c, 0x7e, 0x00, 0xc2, 0xb4,
	0x10, 0x8f, 0xf4, 0xd0, 0xd3, 0x38, 0x96, 0xc3, 0x5d, 0xcb, 0xdf, 0x1e, 0x4c, 0xd7, 0x5c, 0xd2,
	0x52, 0xb5, 0xf7, 0x3c, 0xd6, 0xb7, 0xde, 0x78, 0xbd, 0xd9, 0x73, 0xca, 0x5d, 0x77, 0xbc, 0x56,
	0xb4, 0xb0, 0xc3, 0x58, 0xd6, 0xe2, 0x3b, 0x1c, 0xb7, 0x18, 0x7a, 0x0b, 0x23, 0x8d, 0xd6, 0xaf,
	0xc2, 0x8b, 0xbe, 0x94, 0xa2, 0x05, 0x36, 0x2d, 0xe8, 0x0c, 0xc6, 0x26, 0x4d, 0xfd, 0x02, 0x34,
	0x11, 0x2d, 0x78, 0x13, 0x98, 0x97, 0xe7, 0xdd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x2c,
	0x88, 0x36, 0xe2, 0x04, 0x00, 0x00,
}
