// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notes/proto/notes.proto

package go_micro_srv_notes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Note struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{0}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateNoteRequest struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNoteRequest) Reset()         { *m = CreateNoteRequest{} }
func (m *CreateNoteRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNoteRequest) ProtoMessage()    {}
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{1}
}

func (m *CreateNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNoteRequest.Unmarshal(m, b)
}
func (m *CreateNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNoteRequest.Marshal(b, m, deterministic)
}
func (m *CreateNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNoteRequest.Merge(m, src)
}
func (m *CreateNoteRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNoteRequest.Size(m)
}
func (m *CreateNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNoteRequest proto.InternalMessageInfo

func (m *CreateNoteRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type CreateNoteResponse struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNoteResponse) Reset()         { *m = CreateNoteResponse{} }
func (m *CreateNoteResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNoteResponse) ProtoMessage()    {}
func (*CreateNoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{2}
}

func (m *CreateNoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNoteResponse.Unmarshal(m, b)
}
func (m *CreateNoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNoteResponse.Marshal(b, m, deterministic)
}
func (m *CreateNoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNoteResponse.Merge(m, src)
}
func (m *CreateNoteResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNoteResponse.Size(m)
}
func (m *CreateNoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNoteResponse proto.InternalMessageInfo

func (m *CreateNoteResponse) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type UpdateNoteRequest struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNoteRequest) Reset()         { *m = UpdateNoteRequest{} }
func (m *UpdateNoteRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNoteRequest) ProtoMessage()    {}
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{3}
}

func (m *UpdateNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNoteRequest.Unmarshal(m, b)
}
func (m *UpdateNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNoteRequest.Marshal(b, m, deterministic)
}
func (m *UpdateNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNoteRequest.Merge(m, src)
}
func (m *UpdateNoteRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNoteRequest.Size(m)
}
func (m *UpdateNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNoteRequest proto.InternalMessageInfo

func (m *UpdateNoteRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type UpdateNoteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNoteResponse) Reset()         { *m = UpdateNoteResponse{} }
func (m *UpdateNoteResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateNoteResponse) ProtoMessage()    {}
func (*UpdateNoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{4}
}

func (m *UpdateNoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNoteResponse.Unmarshal(m, b)
}
func (m *UpdateNoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNoteResponse.Marshal(b, m, deterministic)
}
func (m *UpdateNoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNoteResponse.Merge(m, src)
}
func (m *UpdateNoteResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateNoteResponse.Size(m)
}
func (m *UpdateNoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNoteResponse proto.InternalMessageInfo

type DeleteNoteRequest struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNoteRequest) Reset()         { *m = DeleteNoteRequest{} }
func (m *DeleteNoteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNoteRequest) ProtoMessage()    {}
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{5}
}

func (m *DeleteNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNoteRequest.Unmarshal(m, b)
}
func (m *DeleteNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNoteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNoteRequest.Merge(m, src)
}
func (m *DeleteNoteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNoteRequest.Size(m)
}
func (m *DeleteNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNoteRequest proto.InternalMessageInfo

func (m *DeleteNoteRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type DeleteNoteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNoteResponse) Reset()         { *m = DeleteNoteResponse{} }
func (m *DeleteNoteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteNoteResponse) ProtoMessage()    {}
func (*DeleteNoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{6}
}

func (m *DeleteNoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNoteResponse.Unmarshal(m, b)
}
func (m *DeleteNoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNoteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteNoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNoteResponse.Merge(m, src)
}
func (m *DeleteNoteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteNoteResponse.Size(m)
}
func (m *DeleteNoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNoteResponse proto.InternalMessageInfo

type ListNotesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNotesRequest) Reset()         { *m = ListNotesRequest{} }
func (m *ListNotesRequest) String() string { return proto.CompactTextString(m) }
func (*ListNotesRequest) ProtoMessage()    {}
func (*ListNotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{7}
}

func (m *ListNotesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotesRequest.Unmarshal(m, b)
}
func (m *ListNotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotesRequest.Marshal(b, m, deterministic)
}
func (m *ListNotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotesRequest.Merge(m, src)
}
func (m *ListNotesRequest) XXX_Size() int {
	return xxx_messageInfo_ListNotesRequest.Size(m)
}
func (m *ListNotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotesRequest proto.InternalMessageInfo

type ListNotesResponse struct {
	Notes                []*Note  `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNotesResponse) Reset()         { *m = ListNotesResponse{} }
func (m *ListNotesResponse) String() string { return proto.CompactTextString(m) }
func (*ListNotesResponse) ProtoMessage()    {}
func (*ListNotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f19915b807268a5, []int{8}
}

func (m *ListNotesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotesResponse.Unmarshal(m, b)
}
func (m *ListNotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotesResponse.Marshal(b, m, deterministic)
}
func (m *ListNotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotesResponse.Merge(m, src)
}
func (m *ListNotesResponse) XXX_Size() int {
	return xxx_messageInfo_ListNotesResponse.Size(m)
}
func (m *ListNotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotesResponse proto.InternalMessageInfo

func (m *ListNotesResponse) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*Note)(nil), "go.micro.srv.notes.Note")
	proto.RegisterType((*CreateNoteRequest)(nil), "go.micro.srv.notes.CreateNoteRequest")
	proto.RegisterType((*CreateNoteResponse)(nil), "go.micro.srv.notes.CreateNoteResponse")
	proto.RegisterType((*UpdateNoteRequest)(nil), "go.micro.srv.notes.UpdateNoteRequest")
	proto.RegisterType((*UpdateNoteResponse)(nil), "go.micro.srv.notes.UpdateNoteResponse")
	proto.RegisterType((*DeleteNoteRequest)(nil), "go.micro.srv.notes.DeleteNoteRequest")
	proto.RegisterType((*DeleteNoteResponse)(nil), "go.micro.srv.notes.DeleteNoteResponse")
	proto.RegisterType((*ListNotesRequest)(nil), "go.micro.srv.notes.ListNotesRequest")
	proto.RegisterType((*ListNotesResponse)(nil), "go.micro.srv.notes.ListNotesResponse")
}

func init() { proto.RegisterFile("notes/proto/notes.proto", fileDescriptor_7f19915b807268a5) }

var fileDescriptor_7f19915b807268a5 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x93, 0x56, 0x1c, 0x45, 0xec, 0x50, 0x70, 0xe9, 0xa9, 0x04, 0x95, 0x1c, 0x64,
	0x0b, 0xf5, 0x09, 0x6a, 0x3d, 0x8a, 0x87, 0x14, 0x11, 0xbc, 0x48, 0x6d, 0x06, 0x09, 0xb4, 0xdd,
	0x98, 0x5d, 0xc5, 0x47, 0xf4, 0xb1, 0x64, 0x67, 0x23, 0x06, 0x77, 0xb1, 0x48, 0xbd, 0xcd, 0x4e,
	0xfe, 0xf9, 0xfe, 0x61, 0x66, 0x02, 0x27, 0x1b, 0x65, 0x48, 0x8f, 0xab, 0x5a, 0x19, 0x35, 0xe6,
	0x58, 0x72, 0x8c, 0xf8, 0xac, 0xe4, 0xba, 0x5c, 0xd6, 0x4a, 0xea, 0xfa, 0x4d, 0xf2, 0x97, 0xf4,
	0x01, 0x92, 0x5b, 0x65, 0x08, 0x8f, 0xa0, 0x53, 0x16, 0x22, 0x1a, 0x45, 0xd9, 0x7e, 0xde, 0x29,
	0x0b, 0x14, 0xb0, 0xb7, 0xac, 0x69, 0x61, 0xa8, 0x10, 0x9d, 0x51, 0x94, 0xc5, 0xf9, 0xd7, 0x13,
	0x07, 0xd0, 0x35, 0xa5, 0x59, 0x91, 0x88, 0x59, 0xec, 0x1e, 0x88, 0x90, 0x18, 0x7a, 0x37, 0x22,
	0xe1, 0x24, 0xc7, 0xe9, 0x14, 0xfa, 0x33, 0x2e, 0xb2, 0x0e, 0x39, 0xbd, 0xbc, 0x92, 0x36, 0x78,
	0x01, 0x89, 0x75, 0x66, 0xab, 0x83, 0x89, 0x90, 0x7e, 0x4f, 0x92, 0xe5, 0xac, 0x4a, 0xaf, 0x00,
	0xdb, 0x08, 0x5d, 0xa9, 0x8d, 0xa6, 0x3f, 0x32, 0xa6, 0xd0, 0xbf, 0xab, 0x8a, 0x9d, 0xda, 0x18,
	0x00, 0xb6, 0x11, 0xae, 0x0d, 0x0b, 0xbe, 0xa6, 0x15, 0xed, 0x08, 0x6e, 0x23, 0x1a, 0x30, 0xc2,
	0xf1, 0x4d, 0xa9, 0x8d, 0xcd, 0xe9, 0x86, 0x9b, 0xce, 0xa0, 0xdf, 0xca, 0x35, 0x83, 0x90, 0xd0,
	0x65, 0xa4, 0x88, 0x46, 0xf1, 0xaf, 0x6e, 0x4e, 0x36, 0xf9, 0x88, 0xa1, 0xcb, 0x04, 0x9c, 0x43,
	0x62, 0x71, 0x78, 0x1a, 0x2a, 0xf9, 0x69, 0x3e, 0x3c, 0xdb, 0xa2, 0x6a, 0xda, 0xb9, 0x87, 0x9e,
	0xdb, 0x16, 0x06, 0x0b, 0xbc, 0x63, 0x18, 0x9e, 0x6f, 0x93, 0x7d, 0x83, 0xdd, 0x98, 0xc2, 0x60,
	0x6f, 0x0b, 0x61, 0xb0, 0x3f, 0x69, 0x0b, 0x76, 0x8b, 0x0d, 0x83, 0xbd, 0xbb, 0x09, 0x83, 0xfd,
	0xdb, 0xc0, 0x47, 0x38, 0x74, 0xd9, 0xb9, 0xa9, 0x69, 0xb1, 0xfe, 0x67, 0x7c, 0x16, 0x3d, 0xf5,
	0xf8, 0x9f, 0xbe, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x94, 0x95, 0xe9, 0xbd, 0xee, 0x03, 0x00,
	0x00,
}
