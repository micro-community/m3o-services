// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/home/home.proto

package go_micro_api_home

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

type App struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Icon                 string   `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_33e371a295d94392, []int{0}
}

func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *App) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

type User struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	ProfilePictureUrl    string   `protobuf:"bytes,3,opt,name=profile_picture_url,json=profilePictureUrl,proto3" json:"profile_picture_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_33e371a295d94392, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetProfilePictureUrl() string {
	if m != nil {
		return m.ProfilePictureUrl
	}
	return ""
}

type ReadUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserRequest) Reset()         { *m = ReadUserRequest{} }
func (m *ReadUserRequest) String() string { return proto.CompactTextString(m) }
func (*ReadUserRequest) ProtoMessage()    {}
func (*ReadUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33e371a295d94392, []int{2}
}

func (m *ReadUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserRequest.Unmarshal(m, b)
}
func (m *ReadUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserRequest.Marshal(b, m, deterministic)
}
func (m *ReadUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserRequest.Merge(m, src)
}
func (m *ReadUserRequest) XXX_Size() int {
	return xxx_messageInfo_ReadUserRequest.Size(m)
}
func (m *ReadUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserRequest proto.InternalMessageInfo

type ReadUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserResponse) Reset()         { *m = ReadUserResponse{} }
func (m *ReadUserResponse) String() string { return proto.CompactTextString(m) }
func (*ReadUserResponse) ProtoMessage()    {}
func (*ReadUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33e371a295d94392, []int{3}
}

func (m *ReadUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserResponse.Unmarshal(m, b)
}
func (m *ReadUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserResponse.Marshal(b, m, deterministic)
}
func (m *ReadUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserResponse.Merge(m, src)
}
func (m *ReadUserResponse) XXX_Size() int {
	return xxx_messageInfo_ReadUserResponse.Size(m)
}
func (m *ReadUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserResponse proto.InternalMessageInfo

func (m *ReadUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListAppsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAppsRequest) Reset()         { *m = ListAppsRequest{} }
func (m *ListAppsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAppsRequest) ProtoMessage()    {}
func (*ListAppsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33e371a295d94392, []int{4}
}

func (m *ListAppsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppsRequest.Unmarshal(m, b)
}
func (m *ListAppsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppsRequest.Marshal(b, m, deterministic)
}
func (m *ListAppsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppsRequest.Merge(m, src)
}
func (m *ListAppsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAppsRequest.Size(m)
}
func (m *ListAppsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppsRequest proto.InternalMessageInfo

type ListAppsResponse struct {
	Apps                 []*App   `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAppsResponse) Reset()         { *m = ListAppsResponse{} }
func (m *ListAppsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAppsResponse) ProtoMessage()    {}
func (*ListAppsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33e371a295d94392, []int{5}
}

func (m *ListAppsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppsResponse.Unmarshal(m, b)
}
func (m *ListAppsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppsResponse.Marshal(b, m, deterministic)
}
func (m *ListAppsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppsResponse.Merge(m, src)
}
func (m *ListAppsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAppsResponse.Size(m)
}
func (m *ListAppsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppsResponse proto.InternalMessageInfo

func (m *ListAppsResponse) GetApps() []*App {
	if m != nil {
		return m.Apps
	}
	return nil
}

func init() {
	proto.RegisterType((*App)(nil), "go.micro.api.home.App")
	proto.RegisterType((*User)(nil), "go.micro.api.home.User")
	proto.RegisterType((*ReadUserRequest)(nil), "go.micro.api.home.ReadUserRequest")
	proto.RegisterType((*ReadUserResponse)(nil), "go.micro.api.home.ReadUserResponse")
	proto.RegisterType((*ListAppsRequest)(nil), "go.micro.api.home.ListAppsRequest")
	proto.RegisterType((*ListAppsResponse)(nil), "go.micro.api.home.ListAppsResponse")
}

func init() { proto.RegisterFile("proto/home/home.proto", fileDescriptor_33e371a295d94392) }

var fileDescriptor_33e371a295d94392 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x1b, 0x24, 0x9d, 0x82, 0xb6, 0x2b, 0x6a, 0xa8, 0x08, 0x25, 0x5e, 0x8a, 0xc2,
	0x0a, 0xf5, 0xae, 0xf4, 0xe6, 0x41, 0x44, 0x52, 0x7a, 0xf0, 0x54, 0x62, 0x3a, 0xad, 0x0b, 0x49,
	0x77, 0xdc, 0x4d, 0x0e, 0xfe, 0x2b, 0x7f, 0xa2, 0xec, 0x24, 0xb1, 0x10, 0x83, 0x97, 0xb0, 0xfb,
	0xde, 0xcb, 0x37, 0x6f, 0x42, 0xe0, 0x8c, 0x8c, 0x2e, 0xf4, 0xdd, 0x87, 0xce, 0x91, 0x1f, 0x92,
	0xef, 0x62, 0xbc, 0xd3, 0x32, 0x57, 0xa9, 0xd1, 0x32, 0x21, 0x25, 0x9d, 0x11, 0xbd, 0x41, 0x7f,
	0x41, 0x24, 0x8e, 0xa1, 0xa7, 0x36, 0xa1, 0x37, 0xf5, 0x66, 0x83, 0xb8, 0xa7, 0x36, 0x42, 0x80,
	0xbf, 0x4f, 0x72, 0x0c, 0x7b, 0xac, 0xf0, 0x59, 0x4c, 0x20, 0x48, 0x93, 0x02, 0x77, 0xda, 0x7c,
	0x85, 0x7d, 0xd6, 0x7f, 0xef, 0x2e, 0xaf, 0x52, 0xbd, 0x0f, 0xfd, 0x2a, 0xef, 0xce, 0x91, 0x01,
	0x7f, 0x65, 0xd1, 0x88, 0x2b, 0x80, 0xad, 0x32, 0xb6, 0x58, 0x33, 0xb1, 0x9a, 0x31, 0x60, 0xe5,
	0xc5, 0x61, 0x2f, 0x61, 0x90, 0x25, 0x8d, 0x5b, 0xcd, 0x0b, 0x9c, 0xc0, 0xa6, 0x84, 0x53, 0x32,
	0x7a, 0xab, 0x32, 0x5c, 0x93, 0x4a, 0x8b, 0xd2, 0xe0, 0xba, 0x34, 0x59, 0x3d, 0x7e, 0x5c, 0x5b,
	0xaf, 0x95, 0xb3, 0x32, 0x59, 0x34, 0x86, 0x93, 0x18, 0x93, 0x8d, 0x9b, 0x1b, 0xe3, 0x67, 0x89,
	0xb6, 0x88, 0x1e, 0x61, 0x74, 0x90, 0x2c, 0xe9, 0xbd, 0x45, 0x71, 0x0b, 0x7e, 0x69, 0xd1, 0x70,
	0x99, 0xe1, 0xfc, 0x42, 0xfe, 0xf9, 0x2e, 0x92, 0xe3, 0x1c, 0x72, 0xcc, 0x67, 0x65, 0x8b, 0x05,
	0x91, 0x6d, 0x98, 0x0f, 0x30, 0x3a, 0x48, 0x35, 0xf3, 0x06, 0xfc, 0x84, 0xc8, 0x86, 0xde, 0xb4,
	0x3f, 0x1b, 0xce, 0xcf, 0x3b, 0x98, 0x0b, 0xa2, 0x98, 0x33, 0xf3, 0x6f, 0x0f, 0xfc, 0x27, 0x9d,
	0xa3, 0x58, 0x42, 0xd0, 0x94, 0x13, 0x51, 0xc7, 0x2b, 0xad, 0x65, 0x26, 0xd7, 0xff, 0x66, 0xea,
	0x26, 0x4b, 0x08, 0x9a, 0x76, 0x9d, 0xd0, 0xd6, 0x36, 0x9d, 0xd0, 0xf6, 0x7a, 0xef, 0x47, 0xfc,
	0x0b, 0xdd, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x89, 0xcb, 0x1f, 0x8f, 0x5b, 0x02, 0x00, 0x00,
}