// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BidirectionalSequenceRNNOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsBidirectionalSequenceRNNOptions(buf []byte, offset flatbuffers.UOffsetT) *BidirectionalSequenceRNNOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BidirectionalSequenceRNNOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBidirectionalSequenceRNNOptions(buf []byte, offset flatbuffers.UOffsetT) *BidirectionalSequenceRNNOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BidirectionalSequenceRNNOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BidirectionalSequenceRNNOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BidirectionalSequenceRNNOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BidirectionalSequenceRNNOptions) TimeMajor() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BidirectionalSequenceRNNOptions) MutateTimeMajor(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *BidirectionalSequenceRNNOptions) FusedActivationFunction() ActivationFunctionType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return ActivationFunctionType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *BidirectionalSequenceRNNOptions) MutateFusedActivationFunction(n ActivationFunctionType) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func (rcv *BidirectionalSequenceRNNOptions) MergeOutputs() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BidirectionalSequenceRNNOptions) MutateMergeOutputs(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *BidirectionalSequenceRNNOptions) AsymmetricQuantizeInputs() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BidirectionalSequenceRNNOptions) MutateAsymmetricQuantizeInputs(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func BidirectionalSequenceRNNOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func BidirectionalSequenceRNNOptionsAddTimeMajor(builder *flatbuffers.Builder, timeMajor bool) {
	builder.PrependBoolSlot(0, timeMajor, false)
}
func BidirectionalSequenceRNNOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction ActivationFunctionType) {
	builder.PrependInt8Slot(1, int8(fusedActivationFunction), 0)
}
func BidirectionalSequenceRNNOptionsAddMergeOutputs(builder *flatbuffers.Builder, mergeOutputs bool) {
	builder.PrependBoolSlot(2, mergeOutputs, false)
}
func BidirectionalSequenceRNNOptionsAddAsymmetricQuantizeInputs(builder *flatbuffers.Builder, asymmetricQuantizeInputs bool) {
	builder.PrependBoolSlot(3, asymmetricQuantizeInputs, false)
}
func BidirectionalSequenceRNNOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
