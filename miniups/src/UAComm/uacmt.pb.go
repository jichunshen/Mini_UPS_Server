// Code generated by protoc-gen-go.
// source: uacmt.proto
// DO NOT EDIT!

/*
Package UAComm is a generated protocol buffer package.

It is generated from these files:
	uacmt.proto

It has these top-level messages:
	Loaded
	Product
	Warehouse
	Purchase
	ChangeDes
	AUCommands
	PackageStatus
	StartLoad
	DesChangeResult
	PackageID
	UACommands
*/
package UAComm

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// A -> U
type Loaded struct {
	Packageid        *int64 `protobuf:"varint,1,req,name=packageid" json:"packageid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Loaded) Reset()         { *m = Loaded{} }
func (m *Loaded) String() string { return proto.CompactTextString(m) }
func (*Loaded) ProtoMessage()    {}

func (m *Loaded) GetPackageid() int64 {
	if m != nil && m.Packageid != nil {
		return *m.Packageid
	}
	return 0
}

type Product struct {
	Id               *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Description      *string `protobuf:"bytes,2,req,name=description" json:"description,omitempty"`
	Count            *int32  `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}

func (m *Product) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Product) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Product) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type Warehouse struct {
	Whid             *int32 `protobuf:"varint,1,req,name=whid" json:"whid,omitempty"`
	Whx              *int32 `protobuf:"varint,2,req,name=whx" json:"whx,omitempty"`
	Why              *int32 `protobuf:"varint,3,req,name=why" json:"why,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Warehouse) Reset()         { *m = Warehouse{} }
func (m *Warehouse) String() string { return proto.CompactTextString(m) }
func (*Warehouse) ProtoMessage()    {}

func (m *Warehouse) GetWhid() int32 {
	if m != nil && m.Whid != nil {
		return *m.Whid
	}
	return 0
}

func (m *Warehouse) GetWhx() int32 {
	if m != nil && m.Whx != nil {
		return *m.Whx
	}
	return 0
}

func (m *Warehouse) GetWhy() int32 {
	if m != nil && m.Why != nil {
		return *m.Why
	}
	return 0
}

type Purchase struct {
	Whid             *int32     `protobuf:"varint,1,req,name=whid" json:"whid,omitempty"`
	X                *int32     `protobuf:"varint,2,req,name=x" json:"x,omitempty"`
	Y                *int32     `protobuf:"varint,3,req,name=y" json:"y,omitempty"`
	Things           []*Product `protobuf:"bytes,4,rep,name=things" json:"things,omitempty"`
	Orderid          *int64     `protobuf:"varint,5,req,name=orderid" json:"orderid,omitempty"`
	Upsuserid        *string    `protobuf:"bytes,6,opt,name=upsuserid" json:"upsuserid,omitempty"`
	Isprime          *bool      `protobuf:"varint,7,req,name=isprime" json:"isprime,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Purchase) Reset()         { *m = Purchase{} }
func (m *Purchase) String() string { return proto.CompactTextString(m) }
func (*Purchase) ProtoMessage()    {}

func (m *Purchase) GetWhid() int32 {
	if m != nil && m.Whid != nil {
		return *m.Whid
	}
	return 0
}

func (m *Purchase) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Purchase) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *Purchase) GetThings() []*Product {
	if m != nil {
		return m.Things
	}
	return nil
}

func (m *Purchase) GetOrderid() int64 {
	if m != nil && m.Orderid != nil {
		return *m.Orderid
	}
	return 0
}

func (m *Purchase) GetUpsuserid() string {
	if m != nil && m.Upsuserid != nil {
		return *m.Upsuserid
	}
	return ""
}

func (m *Purchase) GetIsprime() bool {
	if m != nil && m.Isprime != nil {
		return *m.Isprime
	}
	return false
}

type ChangeDes struct {
	Pkgid            *int64 `protobuf:"varint,1,req,name=pkgid" json:"pkgid,omitempty"`
	X                *int32 `protobuf:"varint,2,req,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,3,req,name=y" json:"y,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChangeDes) Reset()         { *m = ChangeDes{} }
func (m *ChangeDes) String() string { return proto.CompactTextString(m) }
func (*ChangeDes) ProtoMessage()    {}

func (m *ChangeDes) GetPkgid() int64 {
	if m != nil && m.Pkgid != nil {
		return *m.Pkgid
	}
	return 0
}

func (m *ChangeDes) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *ChangeDes) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

// A to U Command
type AUCommands struct {
	Pc               *Purchase    `protobuf:"bytes,1,opt,name=pc" json:"pc,omitempty"`
	Ldd              *Loaded      `protobuf:"bytes,2,opt,name=ldd" json:"ldd,omitempty"`
	Chdes            *ChangeDes   `protobuf:"bytes,3,opt,name=chdes" json:"chdes,omitempty"`
	Disconnect       *bool        `protobuf:"varint,4,opt,name=disconnect" json:"disconnect,omitempty"`
	Whinfo           []*Warehouse `protobuf:"bytes,5,rep,name=whinfo" json:"whinfo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AUCommands) Reset()         { *m = AUCommands{} }
func (m *AUCommands) String() string { return proto.CompactTextString(m) }
func (*AUCommands) ProtoMessage()    {}

func (m *AUCommands) GetPc() *Purchase {
	if m != nil {
		return m.Pc
	}
	return nil
}

func (m *AUCommands) GetLdd() *Loaded {
	if m != nil {
		return m.Ldd
	}
	return nil
}

func (m *AUCommands) GetChdes() *ChangeDes {
	if m != nil {
		return m.Chdes
	}
	return nil
}

func (m *AUCommands) GetDisconnect() bool {
	if m != nil && m.Disconnect != nil {
		return *m.Disconnect
	}
	return false
}

func (m *AUCommands) GetWhinfo() []*Warehouse {
	if m != nil {
		return m.Whinfo
	}
	return nil
}

// U -> A
type PackageStatus struct {
	Pkgid            *int64  `protobuf:"varint,1,req,name=pkgid" json:"pkgid,omitempty"`
	Status           *string `protobuf:"bytes,2,req,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PackageStatus) Reset()         { *m = PackageStatus{} }
func (m *PackageStatus) String() string { return proto.CompactTextString(m) }
func (*PackageStatus) ProtoMessage()    {}

func (m *PackageStatus) GetPkgid() int64 {
	if m != nil && m.Pkgid != nil {
		return *m.Pkgid
	}
	return 0
}

func (m *PackageStatus) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

type StartLoad struct {
	Whid *int32 `protobuf:"varint,1,req,name=whid" json:"whid,omitempty"`
	// required int64 packageid = 2;
	Truckid          *int32  `protobuf:"varint,2,req,name=truckid" json:"truckid,omitempty"`
	Packageid        []int64 `protobuf:"varint,3,rep,name=packageid" json:"packageid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StartLoad) Reset()         { *m = StartLoad{} }
func (m *StartLoad) String() string { return proto.CompactTextString(m) }
func (*StartLoad) ProtoMessage()    {}

func (m *StartLoad) GetWhid() int32 {
	if m != nil && m.Whid != nil {
		return *m.Whid
	}
	return 0
}

func (m *StartLoad) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *StartLoad) GetPackageid() []int64 {
	if m != nil {
		return m.Packageid
	}
	return nil
}

type DesChangeResult struct {
	Pkgid            *int64 `protobuf:"varint,1,req,name=pkgid" json:"pkgid,omitempty"`
	Res              *bool  `protobuf:"varint,2,req,name=res" json:"res,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DesChangeResult) Reset()         { *m = DesChangeResult{} }
func (m *DesChangeResult) String() string { return proto.CompactTextString(m) }
func (*DesChangeResult) ProtoMessage()    {}

func (m *DesChangeResult) GetPkgid() int64 {
	if m != nil && m.Pkgid != nil {
		return *m.Pkgid
	}
	return 0
}

func (m *DesChangeResult) GetRes() bool {
	if m != nil && m.Res != nil {
		return *m.Res
	}
	return false
}

type PackageID struct {
	Pkgid            *int64 `protobuf:"varint,1,req,name=pkgid" json:"pkgid,omitempty"`
	Orderid          *int64 `protobuf:"varint,2,req,name=orderid" json:"orderid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PackageID) Reset()         { *m = PackageID{} }
func (m *PackageID) String() string { return proto.CompactTextString(m) }
func (*PackageID) ProtoMessage()    {}

func (m *PackageID) GetPkgid() int64 {
	if m != nil && m.Pkgid != nil {
		return *m.Pkgid
	}
	return 0
}

func (m *PackageID) GetOrderid() int64 {
	if m != nil && m.Orderid != nil {
		return *m.Orderid
	}
	return 0
}

// U to A Command
type UACommands struct {
	Sl               []*StartLoad     `protobuf:"bytes,1,rep,name=sl" json:"sl,omitempty"`
	Pid              *PackageID       `protobuf:"bytes,2,opt,name=pid" json:"pid,omitempty"`
	Pks              []*PackageStatus `protobuf:"bytes,3,rep,name=pks" json:"pks,omitempty"`
	Dcr              *DesChangeResult `protobuf:"bytes,4,opt,name=dcr" json:"dcr,omitempty"`
	Worldid          *int64           `protobuf:"varint,5,opt,name=worldid" json:"worldid,omitempty"`
	Finished         *bool            `protobuf:"varint,6,opt,name=finished" json:"finished,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *UACommands) Reset()         { *m = UACommands{} }
func (m *UACommands) String() string { return proto.CompactTextString(m) }
func (*UACommands) ProtoMessage()    {}

func (m *UACommands) GetSl() []*StartLoad {
	if m != nil {
		return m.Sl
	}
	return nil
}

func (m *UACommands) GetPid() *PackageID {
	if m != nil {
		return m.Pid
	}
	return nil
}

func (m *UACommands) GetPks() []*PackageStatus {
	if m != nil {
		return m.Pks
	}
	return nil
}

func (m *UACommands) GetDcr() *DesChangeResult {
	if m != nil {
		return m.Dcr
	}
	return nil
}

func (m *UACommands) GetWorldid() int64 {
	if m != nil && m.Worldid != nil {
		return *m.Worldid
	}
	return 0
}

func (m *UACommands) GetFinished() bool {
	if m != nil && m.Finished != nil {
		return *m.Finished
	}
	return false
}

func init() {
}