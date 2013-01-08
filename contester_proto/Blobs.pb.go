// Code generated by protoc-gen-go.
// source: Blobs.proto
// DO NOT EDIT!

package contester_proto

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Blob_CompressionInfo_CompressionType int32

const (
	Blob_CompressionInfo_METHOD_NONE Blob_CompressionInfo_CompressionType = 0
	Blob_CompressionInfo_METHOD_ZLIB Blob_CompressionInfo_CompressionType = 1
)

var Blob_CompressionInfo_CompressionType_name = map[int32]string{
	0: "METHOD_NONE",
	1: "METHOD_ZLIB",
}
var Blob_CompressionInfo_CompressionType_value = map[string]int32{
	"METHOD_NONE": 0,
	"METHOD_ZLIB": 1,
}

func (x Blob_CompressionInfo_CompressionType) Enum() *Blob_CompressionInfo_CompressionType {
	p := new(Blob_CompressionInfo_CompressionType)
	*p = x
	return p
}
func (x Blob_CompressionInfo_CompressionType) String() string {
	return proto.EnumName(Blob_CompressionInfo_CompressionType_name, int32(x))
}
func (x Blob_CompressionInfo_CompressionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Blob_CompressionInfo_CompressionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Blob_CompressionInfo_CompressionType_value, data, "Blob_CompressionInfo_CompressionType")
	if err != nil {
		return err
	}
	*x = Blob_CompressionInfo_CompressionType(value)
	return nil
}

type Blob struct {
	Data             []byte                `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Compression      *Blob_CompressionInfo `protobuf:"bytes,2,opt,name=compression" json:"compression,omitempty"`
	Sha1             []byte                `protobuf:"bytes,3,opt,name=sha1" json:"sha1,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (this *Blob) Reset()         { *this = Blob{} }
func (this *Blob) String() string { return proto.CompactTextString(this) }
func (*Blob) ProtoMessage()       {}

func (this *Blob) GetData() []byte {
	if this != nil {
		return this.Data
	}
	return nil
}

func (this *Blob) GetCompression() *Blob_CompressionInfo {
	if this != nil {
		return this.Compression
	}
	return nil
}

func (this *Blob) GetSha1() []byte {
	if this != nil {
		return this.Sha1
	}
	return nil
}

type Blob_CompressionInfo struct {
	Method           *Blob_CompressionInfo_CompressionType `protobuf:"varint,1,opt,name=method,enum=contester.proto.Blob_CompressionInfo_CompressionType" json:"method,omitempty"`
	OriginalSize     *uint32                               `protobuf:"varint,2,opt,name=original_size" json:"original_size,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (this *Blob_CompressionInfo) Reset()         { *this = Blob_CompressionInfo{} }
func (this *Blob_CompressionInfo) String() string { return proto.CompactTextString(this) }
func (*Blob_CompressionInfo) ProtoMessage()       {}

func (this *Blob_CompressionInfo) GetMethod() Blob_CompressionInfo_CompressionType {
	if this != nil && this.Method != nil {
		return *this.Method
	}
	return 0
}

func (this *Blob_CompressionInfo) GetOriginalSize() uint32 {
	if this != nil && this.OriginalSize != nil {
		return *this.OriginalSize
	}
	return 0
}

type Module struct {
	Name             *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Data             *Blob   `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	Type             *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Module) Reset()         { *this = Module{} }
func (this *Module) String() string { return proto.CompactTextString(this) }
func (*Module) ProtoMessage()       {}

func (this *Module) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Module) GetData() *Blob {
	if this != nil {
		return this.Data
	}
	return nil
}

func (this *Module) GetType() string {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return ""
}

type FileBlob struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Data             *Blob   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *FileBlob) Reset()         { *this = FileBlob{} }
func (this *FileBlob) String() string { return proto.CompactTextString(this) }
func (*FileBlob) ProtoMessage()       {}

func (this *FileBlob) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *FileBlob) GetData() *Blob {
	if this != nil {
		return this.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("contester.proto.Blob_CompressionInfo_CompressionType", Blob_CompressionInfo_CompressionType_name, Blob_CompressionInfo_CompressionType_value)
}
