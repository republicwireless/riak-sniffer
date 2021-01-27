// Code generated by protoc-gen-go.
// source: riak_search.proto
// DO NOT EDIT!

package riak

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type RpbSearchDoc struct {
	Fields           []*RpbPair `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *RpbSearchDoc) Reset()         { *this = RpbSearchDoc{} }
func (this *RpbSearchDoc) String() string { return proto.CompactTextString(this) }
func (*RpbSearchDoc) ProtoMessage()       {}

type RpbSearchQueryReq struct {
	Q                []byte   `protobuf:"bytes,1,req,name=q" json:"q,omitempty"`
	Index            []byte   `protobuf:"bytes,2,req,name=index" json:"index,omitempty"`
	Rows             *uint32  `protobuf:"varint,3,opt,name=rows" json:"rows,omitempty"`
	Start            *uint32  `protobuf:"varint,4,opt,name=start" json:"start,omitempty"`
	Sort             []byte   `protobuf:"bytes,5,opt,name=sort" json:"sort,omitempty"`
	Filter           []byte   `protobuf:"bytes,6,opt,name=filter" json:"filter,omitempty"`
	Df               []byte   `protobuf:"bytes,7,opt,name=df" json:"df,omitempty"`
	Op               []byte   `protobuf:"bytes,8,opt,name=op" json:"op,omitempty"`
	Fl               [][]byte `protobuf:"bytes,9,rep,name=fl" json:"fl,omitempty"`
	Presort          []byte   `protobuf:"bytes,10,opt,name=presort" json:"presort,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *RpbSearchQueryReq) Reset()         { *this = RpbSearchQueryReq{} }
func (this *RpbSearchQueryReq) String() string { return proto.CompactTextString(this) }
func (*RpbSearchQueryReq) ProtoMessage()       {}

func (this *RpbSearchQueryReq) GetQ() []byte {
	if this != nil {
		return this.Q
	}
	return nil
}

func (this *RpbSearchQueryReq) GetIndex() []byte {
	if this != nil {
		return this.Index
	}
	return nil
}

func (this *RpbSearchQueryReq) GetRows() uint32 {
	if this != nil && this.Rows != nil {
		return *this.Rows
	}
	return 0
}

func (this *RpbSearchQueryReq) GetStart() uint32 {
	if this != nil && this.Start != nil {
		return *this.Start
	}
	return 0
}

func (this *RpbSearchQueryReq) GetSort() []byte {
	if this != nil {
		return this.Sort
	}
	return nil
}

func (this *RpbSearchQueryReq) GetFilter() []byte {
	if this != nil {
		return this.Filter
	}
	return nil
}

func (this *RpbSearchQueryReq) GetDf() []byte {
	if this != nil {
		return this.Df
	}
	return nil
}

func (this *RpbSearchQueryReq) GetOp() []byte {
	if this != nil {
		return this.Op
	}
	return nil
}

func (this *RpbSearchQueryReq) GetPresort() []byte {
	if this != nil {
		return this.Presort
	}
	return nil
}

type RpbSearchQueryResp struct {
	Docs             []*RpbSearchDoc `protobuf:"bytes,1,rep,name=docs" json:"docs,omitempty"`
	MaxScore         *float32        `protobuf:"fixed32,2,opt,name=max_score" json:"max_score,omitempty"`
	NumFound         *uint32         `protobuf:"varint,3,opt,name=num_found" json:"num_found,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *RpbSearchQueryResp) Reset()         { *this = RpbSearchQueryResp{} }
func (this *RpbSearchQueryResp) String() string { return proto.CompactTextString(this) }
func (*RpbSearchQueryResp) ProtoMessage()       {}

func (this *RpbSearchQueryResp) GetMaxScore() float32 {
	if this != nil && this.MaxScore != nil {
		return *this.MaxScore
	}
	return 0
}

func (this *RpbSearchQueryResp) GetNumFound() uint32 {
	if this != nil && this.NumFound != nil {
		return *this.NumFound
	}
	return 0
}

func init() {
}
