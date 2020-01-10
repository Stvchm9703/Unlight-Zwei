// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SkillList struct {
	_tab flatbuffers.Table
}

func GetRootAsSkillList(buf []byte, offset flatbuffers.UOffsetT) *SkillList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SkillList{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SkillList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SkillList) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SkillList) SkillList(obj *Skill, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SkillList) SkillListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SkillList) LastUpdated() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func SkillListStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SkillListAddSkillList(builder *flatbuffers.Builder, skillList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(skillList), 0)
}
func SkillListStartSkillListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SkillListAddLastUpdated(builder *flatbuffers.Builder, lastUpdated flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lastUpdated), 0)
}
func SkillListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}