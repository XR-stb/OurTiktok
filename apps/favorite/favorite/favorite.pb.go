// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: favorite.proto

package favorite

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
	return file_favorite_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_favorite_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{0}
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
	TotalFavorited  string `protobuf:"bytes,9,opt,name=TotalFavorited,proto3" json:"TotalFavorited,omitempty"`
	WorkCount       int64  `protobuf:"varint,10,opt,name=WorkCount,proto3" json:"WorkCount,omitempty"`
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[0]
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
	return file_favorite_proto_rawDescGZIP(), []int{0}
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

func (x *UserInfo) GetTotalFavorited() string {
	if x != nil {
		return x.TotalFavorited
	}
	return ""
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
	Author        *UserInfo `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	PlayUrl       string    `protobuf:"bytes,3,opt,name=PlayUrl,proto3" json:"PlayUrl,omitempty"`
	CoverUrl      string    `protobuf:"bytes,4,opt,name=CoverUrl,proto3" json:"CoverUrl,omitempty"`
	FavoriteCount int64     `protobuf:"varint,5,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`
	CommentCount  int64     `protobuf:"varint,6,opt,name=CommentCount,proto3" json:"CommentCount,omitempty"`
	IsFavorite    bool      `protobuf:"varint,7,opt,name=IsFavorite,proto3" json:"IsFavorite,omitempty"`
	Title         string    `protobuf:"bytes,8,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[1]
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
	return file_favorite_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
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

type VideoFavorite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FavoriteCount int64 `protobuf:"varint,1,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`
	IsFavorite    bool  `protobuf:"varint,2,opt,name=IsFavorite,proto3" json:"IsFavorite,omitempty"`
}

func (x *VideoFavorite) Reset() {
	*x = VideoFavorite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFavorite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFavorite) ProtoMessage() {}

func (x *VideoFavorite) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFavorite.ProtoReflect.Descriptor instead.
func (*VideoFavorite) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{2}
}

func (x *VideoFavorite) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *VideoFavorite) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type UserFavorite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FavoriteCount  int64 `protobuf:"varint,1,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`
	TotalFavorited int64 `protobuf:"varint,2,opt,name=TotalFavorited,proto3" json:"TotalFavorited,omitempty"`
}

func (x *UserFavorite) Reset() {
	*x = UserFavorite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFavorite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFavorite) ProtoMessage() {}

func (x *UserFavorite) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFavorite.ProtoReflect.Descriptor instead.
func (*UserFavorite) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{3}
}

func (x *UserFavorite) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *UserFavorite) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

type ActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	VideoId    int64 `protobuf:"varint,2,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	ActionType int32 `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (x *ActionReq) Reset() {
	*x = ActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionReq) ProtoMessage() {}

func (x *ActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionReq.ProtoReflect.Descriptor instead.
func (*ActionReq) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{4}
}

func (x *ActionReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ActionReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *ActionReq) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type ActionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=Status,proto3,enum=favorite.Status" json:"Status,omitempty"`
}

func (x *ActionRes) Reset() {
	*x = ActionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRes) ProtoMessage() {}

func (x *ActionRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRes.ProtoReflect.Descriptor instead.
func (*ActionRes) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{5}
}

func (x *ActionRes) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ThisId int64 `protobuf:"varint,2,opt,name=ThisId,proto3" json:"ThisId,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{6}
}

func (x *ListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListReq) GetThisId() int64 {
	if x != nil {
		return x.ThisId
	}
	return 0
}

type ListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Status   `protobuf:"varint,1,opt,name=Status,proto3,enum=favorite.Status" json:"Status,omitempty"`
	VideoList []*Video `protobuf:"bytes,2,rep,name=VideoList,proto3" json:"VideoList,omitempty"`
}

func (x *ListRes) Reset() {
	*x = ListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRes) ProtoMessage() {}

func (x *ListRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRes.ProtoReflect.Descriptor instead.
func (*ListRes) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{7}
}

func (x *ListRes) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *ListRes) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type GetVideoFavoriteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64   `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	VideoIds    []int64 `protobuf:"varint,2,rep,packed,name=VideoIds,proto3" json:"VideoIds,omitempty"`
	AllFavorite bool    `protobuf:"varint,3,opt,name=AllFavorite,proto3" json:"AllFavorite,omitempty"`
}

func (x *GetVideoFavoriteReq) Reset() {
	*x = GetVideoFavoriteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoFavoriteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoFavoriteReq) ProtoMessage() {}

func (x *GetVideoFavoriteReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoFavoriteReq.ProtoReflect.Descriptor instead.
func (*GetVideoFavoriteReq) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{8}
}

func (x *GetVideoFavoriteReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetVideoFavoriteReq) GetVideoIds() []int64 {
	if x != nil {
		return x.VideoIds
	}
	return nil
}

func (x *GetVideoFavoriteReq) GetAllFavorite() bool {
	if x != nil {
		return x.AllFavorite
	}
	return false
}

type GetVideoFavoriteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Favorites []*VideoFavorite `protobuf:"bytes,1,rep,name=Favorites,proto3" json:"Favorites,omitempty"`
}

func (x *GetVideoFavoriteRes) Reset() {
	*x = GetVideoFavoriteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoFavoriteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoFavoriteRes) ProtoMessage() {}

func (x *GetVideoFavoriteRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoFavoriteRes.ProtoReflect.Descriptor instead.
func (*GetVideoFavoriteRes) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{9}
}

func (x *GetVideoFavoriteRes) GetFavorites() []*VideoFavorite {
	if x != nil {
		return x.Favorites
	}
	return nil
}

type GetUserFavoriteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []int64 `protobuf:"varint,1,rep,packed,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUserFavoriteReq) Reset() {
	*x = GetUserFavoriteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserFavoriteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserFavoriteReq) ProtoMessage() {}

func (x *GetUserFavoriteReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserFavoriteReq.ProtoReflect.Descriptor instead.
func (*GetUserFavoriteReq) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserFavoriteReq) GetUsers() []int64 {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserFavoriteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Favorites []*UserFavorite `protobuf:"bytes,2,rep,name=Favorites,proto3" json:"Favorites,omitempty"`
}

func (x *GetUserFavoriteRes) Reset() {
	*x = GetUserFavoriteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserFavoriteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserFavoriteRes) ProtoMessage() {}

func (x *GetUserFavoriteRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserFavoriteRes.ProtoReflect.Descriptor instead.
func (*GetUserFavoriteRes) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserFavoriteRes) GetFavorites() []*UserFavorite {
	if x != nil {
		return x.Favorites
	}
	return nil
}

var File_favorite_proto protoreflect.FileDescriptor

var file_favorite_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22, 0xe6, 0x02, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x28, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xf9, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x79,
	0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12,
	0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49,
	0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0x55, 0x0a, 0x0d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22, 0x5c, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x35, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x28, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x39, 0x0a, 0x07, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x54, 0x68, 0x69, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54,
	0x68, 0x69, 0x73, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x28, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x6c, 0x6c, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22, 0x4c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a,
	0x09, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x2a, 0x1b, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x4f, 0x4b, 0x10, 0x01, 0x32, 0x8d, 0x02, 0x0a, 0x08, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x13, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x11, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_favorite_proto_rawDescOnce sync.Once
	file_favorite_proto_rawDescData = file_favorite_proto_rawDesc
)

func file_favorite_proto_rawDescGZIP() []byte {
	file_favorite_proto_rawDescOnce.Do(func() {
		file_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorite_proto_rawDescData)
	})
	return file_favorite_proto_rawDescData
}

var file_favorite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_favorite_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: favorite.Status
	(*UserInfo)(nil),            // 1: favorite.UserInfo
	(*Video)(nil),               // 2: favorite.Video
	(*VideoFavorite)(nil),       // 3: favorite.VideoFavorite
	(*UserFavorite)(nil),        // 4: favorite.UserFavorite
	(*ActionReq)(nil),           // 5: favorite.ActionReq
	(*ActionRes)(nil),           // 6: favorite.ActionRes
	(*ListReq)(nil),             // 7: favorite.ListReq
	(*ListRes)(nil),             // 8: favorite.ListRes
	(*GetVideoFavoriteReq)(nil), // 9: favorite.GetVideoFavoriteReq
	(*GetVideoFavoriteRes)(nil), // 10: favorite.GetVideoFavoriteRes
	(*GetUserFavoriteReq)(nil),  // 11: favorite.GetUserFavoriteReq
	(*GetUserFavoriteRes)(nil),  // 12: favorite.GetUserFavoriteRes
}
var file_favorite_proto_depIdxs = []int32{
	1,  // 0: favorite.Video.Author:type_name -> favorite.UserInfo
	0,  // 1: favorite.ActionRes.Status:type_name -> favorite.Status
	0,  // 2: favorite.ListRes.Status:type_name -> favorite.Status
	2,  // 3: favorite.ListRes.VideoList:type_name -> favorite.Video
	3,  // 4: favorite.GetVideoFavoriteRes.Favorites:type_name -> favorite.VideoFavorite
	4,  // 5: favorite.GetUserFavoriteRes.Favorites:type_name -> favorite.UserFavorite
	5,  // 6: favorite.Favorite.Action:input_type -> favorite.ActionReq
	7,  // 7: favorite.Favorite.List:input_type -> favorite.ListReq
	9,  // 8: favorite.Favorite.GetVideoFavorite:input_type -> favorite.GetVideoFavoriteReq
	11, // 9: favorite.Favorite.GetUserFavorite:input_type -> favorite.GetUserFavoriteReq
	6,  // 10: favorite.Favorite.Action:output_type -> favorite.ActionRes
	8,  // 11: favorite.Favorite.List:output_type -> favorite.ListRes
	10, // 12: favorite.Favorite.GetVideoFavorite:output_type -> favorite.GetVideoFavoriteRes
	12, // 13: favorite.Favorite.GetUserFavorite:output_type -> favorite.GetUserFavoriteRes
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_favorite_proto_init() }
func file_favorite_proto_init() {
	if File_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_favorite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFavorite); i {
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
		file_favorite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFavorite); i {
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
		file_favorite_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionReq); i {
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
		file_favorite_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRes); i {
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
		file_favorite_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_favorite_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRes); i {
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
		file_favorite_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoFavoriteReq); i {
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
		file_favorite_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoFavoriteRes); i {
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
		file_favorite_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserFavoriteReq); i {
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
		file_favorite_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserFavoriteRes); i {
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
			RawDescriptor: file_favorite_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorite_proto_goTypes,
		DependencyIndexes: file_favorite_proto_depIdxs,
		EnumInfos:         file_favorite_proto_enumTypes,
		MessageInfos:      file_favorite_proto_msgTypes,
	}.Build()
	File_favorite_proto = out.File
	file_favorite_proto_rawDesc = nil
	file_favorite_proto_goTypes = nil
	file_favorite_proto_depIdxs = nil
}
