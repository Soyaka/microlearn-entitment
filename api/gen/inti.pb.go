// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: inti.proto

package entitment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookmarkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ContentId string `protobuf:"bytes,3,opt,name=contentId,proto3" json:"contentId,omitempty"`
}

func (x *BookmarkReq) Reset() {
	*x = BookmarkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookmarkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookmarkReq) ProtoMessage() {}

func (x *BookmarkReq) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookmarkReq.ProtoReflect.Descriptor instead.
func (*BookmarkReq) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{0}
}

func (x *BookmarkReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookmarkReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookmarkReq) GetContentId() string {
	if x != nil {
		return x.ContentId
	}
	return ""
}

type BookmarkRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookmarks []string `protobuf:"bytes,1,rep,name=bookmarks,proto3" json:"bookmarks,omitempty"`
}

func (x *BookmarkRes) Reset() {
	*x = BookmarkRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookmarkRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookmarkRes) ProtoMessage() {}

func (x *BookmarkRes) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookmarkRes.ProtoReflect.Descriptor instead.
func (*BookmarkRes) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{1}
}

func (x *BookmarkRes) GetBookmarks() []string {
	if x != nil {
		return x.Bookmarks
	}
	return nil
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{2}
}

func (x *UserID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InterestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Tech   string `protobuf:"bytes,3,opt,name=tech,proto3" json:"tech,omitempty"`
	Level  string `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	Cert   string `protobuf:"bytes,5,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *InterestReq) Reset() {
	*x = InterestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterestReq) ProtoMessage() {}

func (x *InterestReq) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterestReq.ProtoReflect.Descriptor instead.
func (*InterestReq) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{3}
}

func (x *InterestReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InterestReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InterestReq) GetTech() string {
	if x != nil {
		return x.Tech
	}
	return ""
}

func (x *InterestReq) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *InterestReq) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

type InterestRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tech  string `protobuf:"bytes,1,opt,name=tech,proto3" json:"tech,omitempty"`
	Level string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Cert  string `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *InterestRes) Reset() {
	*x = InterestRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterestRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterestRes) ProtoMessage() {}

func (x *InterestRes) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterestRes.ProtoReflect.Descriptor instead.
func (*InterestRes) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{4}
}

func (x *InterestRes) GetTech() string {
	if x != nil {
		return x.Tech
	}
	return ""
}

func (x *InterestRes) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *InterestRes) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

type ProgressReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ContentId string `protobuf:"bytes,3,opt,name=contentId,proto3" json:"contentId,omitempty"`
	Progress  int32  `protobuf:"varint,4,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *ProgressReq) Reset() {
	*x = ProgressReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressReq) ProtoMessage() {}

func (x *ProgressReq) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressReq.ProtoReflect.Descriptor instead.
func (*ProgressReq) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{5}
}

func (x *ProgressReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProgressReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProgressReq) GetContentId() string {
	if x != nil {
		return x.ContentId
	}
	return ""
}

func (x *ProgressReq) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

type ProgressRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentId string `protobuf:"bytes,3,opt,name=contentId,proto3" json:"contentId,omitempty"`
	Progress  int32  `protobuf:"varint,4,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *ProgressRes) Reset() {
	*x = ProgressRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressRes) ProtoMessage() {}

func (x *ProgressRes) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressRes.ProtoReflect.Descriptor instead.
func (*ProgressRes) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{6}
}

func (x *ProgressRes) GetContentId() string {
	if x != nil {
		return x.ContentId
	}
	return ""
}

func (x *ProgressRes) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

type ProgressListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress []*ProgressRes `protobuf:"bytes,1,rep,name=progress,proto3" json:"progress,omitempty"`
}

func (x *ProgressListRes) Reset() {
	*x = ProgressListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressListRes) ProtoMessage() {}

func (x *ProgressListRes) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressListRes.ProtoReflect.Descriptor instead.
func (*ProgressListRes) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{7}
}

func (x *ProgressListRes) GetProgress() []*ProgressRes {
	if x != nil {
		return x.Progress
	}
	return nil
}

type SubscriptionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"id"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty" gorm:"userId"`
	ContentId string `protobuf:"bytes,3,opt,name=contentId,proto3" json:"contentId,omitempty" gorm:"contentId" sqlite:"contentId"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"updated_at"`
}

func (x *SubscriptionReq) Reset() {
	*x = SubscriptionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionReq) ProtoMessage() {}

func (x *SubscriptionReq) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionReq.ProtoReflect.Descriptor instead.
func (*SubscriptionReq) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{8}
}

func (x *SubscriptionReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubscriptionReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SubscriptionReq) GetContentId() string {
	if x != nil {
		return x.ContentId
	}
	return ""
}

func (x *SubscriptionReq) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SubscriptionReq) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SubscriptionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subs []string `protobuf:"bytes,1,rep,name=subs,proto3" json:"subs,omitempty"`
}

func (x *SubscriptionRes) Reset() {
	*x = SubscriptionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inti_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRes) ProtoMessage() {}

func (x *SubscriptionRes) ProtoReflect() protoreflect.Message {
	mi := &file_inti_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRes.ProtoReflect.Descriptor instead.
func (*SubscriptionRes) Descriptor() ([]byte, []int) {
	return file_inti_proto_rawDescGZIP(), []int{9}
}

func (x *SubscriptionRes) GetSubs() []string {
	if x != nil {
		return x.Subs
	}
	return nil
}

var File_inti_proto protoreflect.FileDescriptor

var file_inti_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0b,
	0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x2b, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0x18,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x22, 0x4b, 0x0a,
	0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x63, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x22, 0x6f, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x47, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x75, 0x62, 0x73,
	0x32, 0x37, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x07, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x32, 0x38, 0x0a, 0x0f, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x12, 0x07, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x32, 0x3f, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x07, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x32, 0x43, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6f, 0x79, 0x61, 0x6b, 0x61, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inti_proto_rawDescOnce sync.Once
	file_inti_proto_rawDescData = file_inti_proto_rawDesc
)

func file_inti_proto_rawDescGZIP() []byte {
	file_inti_proto_rawDescOnce.Do(func() {
		file_inti_proto_rawDescData = protoimpl.X.CompressGZIP(file_inti_proto_rawDescData)
	})
	return file_inti_proto_rawDescData
}

var file_inti_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_inti_proto_goTypes = []interface{}{
	(*BookmarkReq)(nil),     // 0: BookmarkReq
	(*BookmarkRes)(nil),     // 1: BookmarkRes
	(*UserID)(nil),          // 2: UserID
	(*InterestReq)(nil),     // 3: InterestReq
	(*InterestRes)(nil),     // 4: InterestRes
	(*ProgressReq)(nil),     // 5: ProgressReq
	(*ProgressRes)(nil),     // 6: ProgressRes
	(*ProgressListRes)(nil), // 7: ProgressListRes
	(*SubscriptionReq)(nil), // 8: SubscriptionReq
	(*SubscriptionRes)(nil), // 9: SubscriptionRes
}
var file_inti_proto_depIdxs = []int32{
	6, // 0: ProgressListRes.progress:type_name -> ProgressRes
	2, // 1: BookmarkService.GetBookmark:input_type -> UserID
	2, // 2: InterestService.GetInterests:input_type -> UserID
	2, // 3: ProgressService.GetProgressList:input_type -> UserID
	2, // 4: SubscriptionService.GetSubscription:input_type -> UserID
	1, // 5: BookmarkService.GetBookmark:output_type -> BookmarkRes
	4, // 6: InterestService.GetInterests:output_type -> InterestRes
	7, // 7: ProgressService.GetProgressList:output_type -> ProgressListRes
	9, // 8: SubscriptionService.GetSubscription:output_type -> SubscriptionRes
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inti_proto_init() }
func file_inti_proto_init() {
	if File_inti_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inti_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookmarkReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookmarkRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterestReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterestRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressListRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inti_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inti_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_inti_proto_goTypes,
		DependencyIndexes: file_inti_proto_depIdxs,
		MessageInfos:      file_inti_proto_msgTypes,
	}.Build()
	File_inti_proto = out.File
	file_inti_proto_rawDesc = nil
	file_inti_proto_goTypes = nil
	file_inti_proto_depIdxs = nil
}
