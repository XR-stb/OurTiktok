// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: feed.proto

package feed

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

type Status int32

const (
	Status_OK    Status = 0
	Status_notOK Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "notOK",
	}
	Status_value = map[string]int32{
		"OK":    0,
		"notOK": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_feed_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_feed_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{0}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username        string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	FollowCount     int64  `protobuf:"varint,3,opt,name=FollowCount,proto3" json:"FollowCount,omitempty"`
	FollowerCount   int64  `protobuf:"varint,4,opt,name=FollowerCount,proto3" json:"FollowerCount,omitempty"`
	IsFollow        bool   `protobuf:"varint,5,opt,name=IsFollow,proto3" json:"IsFollow,omitempty"`
	Avatar          string `protobuf:"bytes,6,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	BackgroundImage string `protobuf:"bytes,7,opt,name=BackgroundImage,proto3" json:"BackgroundImage,omitempty"`
	Signature       string `protobuf:"bytes,8,opt,name=Signature,proto3" json:"Signature,omitempty"`
	TotalFavorited  int64  `protobuf:"varint,9,opt,name=TotalFavorited,proto3" json:"TotalFavorited,omitempty"`
	WorkCount       int64  `protobuf:"varint,10,opt,name=WorkCount,proto3" json:"WorkCount,omitempty"`
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *UserInfo) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *UserInfo) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfo) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *UserInfo) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UserInfo) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *UserInfo) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *UserInfo) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AuthorId      int64     `protobuf:"varint,2,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	Author        *UserInfo `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	PlayUrl       string    `protobuf:"bytes,4,opt,name=PlayUrl,proto3" json:"PlayUrl,omitempty"`
	CoverUrl      string    `protobuf:"bytes,5,opt,name=CoverUrl,proto3" json:"CoverUrl,omitempty"`
	FavoriteCount int64     `protobuf:"varint,6,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`
	CommentCount  int64     `protobuf:"varint,7,opt,name=CommentCount,proto3" json:"CommentCount,omitempty"`
	IsFavorite    bool      `protobuf:"varint,8,opt,name=IsFavorite,proto3" json:"IsFavorite,omitempty"`
	Title         string    `protobuf:"bytes,9,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Video) GetAuthor() *UserInfo {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type FeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	LatestTime int64 `protobuf:"varint,2,opt,name=LatestTime,proto3" json:"LatestTime,omitempty"`
}

func (x *FeedReq) Reset() {
	*x = FeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedReq) ProtoMessage() {}

func (x *FeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedReq.ProtoReflect.Descriptor instead.
func (*FeedReq) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{2}
}

func (x *FeedReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FeedReq) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

type FeedRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Status   `protobuf:"varint,1,opt,name=Status,proto3,enum=feed.Status" json:"Status,omitempty"`
	NextTime  int64    `protobuf:"varint,2,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
	VideoList []*Video `protobuf:"bytes,3,rep,name=VideoList,proto3" json:"VideoList,omitempty"`
}

func (x *FeedRes) Reset() {
	*x = FeedRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRes) ProtoMessage() {}

func (x *FeedRes) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRes.ProtoReflect.Descriptor instead.
func (*FeedRes) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{3}
}

func (x *FeedRes) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *FeedRes) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

func (x *FeedRes) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_feed_proto protoreflect.FileDescriptor

var file_feed_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x65,
	0x65, 0x64, 0x22, 0xe6, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x57, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x05,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x79,
	0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12,
	0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49,
	0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0x41, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x76, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x1b, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x4f, 0x4b, 0x10, 0x01, 0x32, 0x2c, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0d, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feed_proto_rawDescOnce sync.Once
	file_feed_proto_rawDescData = file_feed_proto_rawDesc
)

func file_feed_proto_rawDescGZIP() []byte {
	file_feed_proto_rawDescOnce.Do(func() {
		file_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_proto_rawDescData)
	})
	return file_feed_proto_rawDescData
}

var file_feed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_feed_proto_goTypes = []interface{}{
	(Status)(0),      // 0: feed.Status
	(*UserInfo)(nil), // 1: feed.UserInfo
	(*Video)(nil),    // 2: feed.Video
	(*FeedReq)(nil),  // 3: feed.FeedReq
	(*FeedRes)(nil),  // 4: feed.FeedRes
}
var file_feed_proto_depIdxs = []int32{
	1, // 0: feed.Video.Author:type_name -> feed.UserInfo
	0, // 1: feed.FeedRes.Status:type_name -> feed.Status
	2, // 2: feed.FeedRes.VideoList:type_name -> feed.Video
	3, // 3: feed.Feed.Feed:input_type -> feed.FeedReq
	4, // 4: feed.Feed.Feed:output_type -> feed.FeedRes
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_feed_proto_init() }
func file_feed_proto_init() {
	if File_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedReq); i {
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
		file_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedRes); i {
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
			RawDescriptor: file_feed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_proto_goTypes,
		DependencyIndexes: file_feed_proto_depIdxs,
		EnumInfos:         file_feed_proto_enumTypes,
		MessageInfos:      file_feed_proto_msgTypes,
	}.Build()
	File_feed_proto = out.File
	file_feed_proto_rawDesc = nil
	file_feed_proto_goTypes = nil
	file_feed_proto_depIdxs = nil
}
