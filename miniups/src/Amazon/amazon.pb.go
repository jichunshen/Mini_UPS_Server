// Code generated by protoc-gen-go.
// source: amazon.proto
// DO NOT EDIT!

/*
Package Amazon is a generated protocol buffer package.

It is generated from these files:
	amazon.proto

It has these top-level messages:
	AProduct
	AInitWarehouse
	AConnect
	AConnected
	APack
	APutOnTruck
	APurchaseMore
	ACommands
	AResponses
*/
package Amazon

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type AProduct struct {
	Id               *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Description      *string `protobuf:"bytes,2,req,name=description" json:"description,omitempty"`
	Count            *int32  `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AProduct) Reset()         { *m = AProduct{} }
func (m *AProduct) String() string { return proto.CompactTextString(m) }
func (*AProduct) ProtoMessage()    {}

func (m *AProduct) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *AProduct) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *AProduct) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type AInitWarehouse struct {
	X                *int32 `protobuf:"varint,1,req,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,req,name=y" json:"y,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AInitWarehouse) Reset()         { *m = AInitWarehouse{} }
func (m *AInitWarehouse) String() string { return proto.CompactTextString(m) }
func (*AInitWarehouse) ProtoMessage()    {}

func (m *AInitWarehouse) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *AInitWarehouse) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type AConnect struct {
	Worldid          *int64            `protobuf:"varint,1,req,name=worldid" json:"worldid,omitempty"`
	Initwh           []*AInitWarehouse `protobuf:"bytes,2,rep,name=initwh" json:"initwh,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *AConnect) Reset()         { *m = AConnect{} }
func (m *AConnect) String() string { return proto.CompactTextString(m) }
func (*AConnect) ProtoMessage()    {}

func (m *AConnect) GetWorldid() int64 {
	if m != nil && m.Worldid != nil {
		return *m.Worldid
	}
	return 0
}

func (m *AConnect) GetInitwh() []*AInitWarehouse {
	if m != nil {
		return m.Initwh
	}
	return nil
}

type AConnected struct {
	Error            *string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AConnected) Reset()         { *m = AConnected{} }
func (m *AConnected) String() string { return proto.CompactTextString(m) }
func (*AConnected) ProtoMessage()    {}

func (m *AConnected) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type APack struct {
	Whnum            *int32      `protobuf:"varint,1,req,name=whnum" json:"whnum,omitempty"`
	Things           []*AProduct `protobuf:"bytes,2,rep,name=things" json:"things,omitempty"`
	Shipid           *int64      `protobuf:"varint,3,req,name=shipid" json:"shipid,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *APack) Reset()         { *m = APack{} }
func (m *APack) String() string { return proto.CompactTextString(m) }
func (*APack) ProtoMessage()    {}

func (m *APack) GetWhnum() int32 {
	if m != nil && m.Whnum != nil {
		return *m.Whnum
	}
	return 0
}

func (m *APack) GetThings() []*AProduct {
	if m != nil {
		return m.Things
	}
	return nil
}

func (m *APack) GetShipid() int64 {
	if m != nil && m.Shipid != nil {
		return *m.Shipid
	}
	return 0
}

type APutOnTruck struct {
	Whnum            *int32 `protobuf:"varint,1,req,name=whnum" json:"whnum,omitempty"`
	Truckid          *int32 `protobuf:"varint,2,req,name=truckid" json:"truckid,omitempty"`
	Shipid           *int64 `protobuf:"varint,3,req,name=shipid" json:"shipid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *APutOnTruck) Reset()         { *m = APutOnTruck{} }
func (m *APutOnTruck) String() string { return proto.CompactTextString(m) }
func (*APutOnTruck) ProtoMessage()    {}

func (m *APutOnTruck) GetWhnum() int32 {
	if m != nil && m.Whnum != nil {
		return *m.Whnum
	}
	return 0
}

func (m *APutOnTruck) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *APutOnTruck) GetShipid() int64 {
	if m != nil && m.Shipid != nil {
		return *m.Shipid
	}
	return 0
}

type APurchaseMore struct {
	Whnum            *int32      `protobuf:"varint,1,req,name=whnum" json:"whnum,omitempty"`
	Things           []*AProduct `protobuf:"bytes,2,rep,name=things" json:"things,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *APurchaseMore) Reset()         { *m = APurchaseMore{} }
func (m *APurchaseMore) String() string { return proto.CompactTextString(m) }
func (*APurchaseMore) ProtoMessage()    {}

func (m *APurchaseMore) GetWhnum() int32 {
	if m != nil && m.Whnum != nil {
		return *m.Whnum
	}
	return 0
}

func (m *APurchaseMore) GetThings() []*AProduct {
	if m != nil {
		return m.Things
	}
	return nil
}

type ACommands struct {
	Buy              []*APurchaseMore `protobuf:"bytes,1,rep,name=buy" json:"buy,omitempty"`
	Load             []*APutOnTruck   `protobuf:"bytes,2,rep,name=load" json:"load,omitempty"`
	Topack           []*APack         `protobuf:"bytes,3,rep,name=topack" json:"topack,omitempty"`
	Simspeed         *uint32          `protobuf:"varint,4,opt,name=simspeed" json:"simspeed,omitempty"`
	Disconnect       *bool            `protobuf:"varint,5,opt,name=disconnect" json:"disconnect,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ACommands) Reset()         { *m = ACommands{} }
func (m *ACommands) String() string { return proto.CompactTextString(m) }
func (*ACommands) ProtoMessage()    {}

func (m *ACommands) GetBuy() []*APurchaseMore {
	if m != nil {
		return m.Buy
	}
	return nil
}

func (m *ACommands) GetLoad() []*APutOnTruck {
	if m != nil {
		return m.Load
	}
	return nil
}

func (m *ACommands) GetTopack() []*APack {
	if m != nil {
		return m.Topack
	}
	return nil
}

func (m *ACommands) GetSimspeed() uint32 {
	if m != nil && m.Simspeed != nil {
		return *m.Simspeed
	}
	return 0
}

func (m *ACommands) GetDisconnect() bool {
	if m != nil && m.Disconnect != nil {
		return *m.Disconnect
	}
	return false
}

type AResponses struct {
	Arrived          []*APurchaseMore `protobuf:"bytes,1,rep,name=arrived" json:"arrived,omitempty"`
	Ready            []int64          `protobuf:"varint,2,rep,name=ready" json:"ready,omitempty"`
	Loaded           []int64          `protobuf:"varint,3,rep,name=loaded" json:"loaded,omitempty"`
	Error            *string          `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	Finished         *bool            `protobuf:"varint,5,opt,name=finished" json:"finished,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AResponses) Reset()         { *m = AResponses{} }
func (m *AResponses) String() string { return proto.CompactTextString(m) }
func (*AResponses) ProtoMessage()    {}

func (m *AResponses) GetArrived() []*APurchaseMore {
	if m != nil {
		return m.Arrived
	}
	return nil
}

func (m *AResponses) GetReady() []int64 {
	if m != nil {
		return m.Ready
	}
	return nil
}

func (m *AResponses) GetLoaded() []int64 {
	if m != nil {
		return m.Loaded
	}
	return nil
}

func (m *AResponses) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *AResponses) GetFinished() bool {
	if m != nil && m.Finished != nil {
		return *m.Finished
	}
	return false
}

func init() {
}