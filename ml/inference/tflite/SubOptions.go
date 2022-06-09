// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SubOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsSubOptions(buf []byte, offset flatbuffers.UOffsetT) *SubOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SubOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSubOptions(buf []byte, offset flatbuffers.UOffsetT) *SubOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SubOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SubOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SubOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SubOptions) FusedActivationFunction() ActivationFunctionType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return ActivationFunctionType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *SubOptions) MutateFusedActivationFunction(n ActivationFunctionType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *SubOptions) PotScaleInt16() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *SubOptions) MutatePotScaleInt16(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func SubOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SubOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction ActivationFunctionType) {
	builder.PrependInt8Slot(0, int8(fusedActivationFunction), 0)
}
func SubOptionsAddPotScaleInt16(builder *flatbuffers.Builder, potScaleInt16 bool) {
	builder.PrependBoolSlot(1, potScaleInt16, true)
}
func SubOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
