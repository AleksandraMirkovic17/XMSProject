// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: post_service.proto

package dislinked

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reaction_ReactionType int32

const (
	Reaction_LIKE    Reaction_ReactionType = 0
	Reaction_DISLIKE Reaction_ReactionType = 1
)

// Enum value maps for Reaction_ReactionType.
var (
	Reaction_ReactionType_name = map[int32]string{
		0: "LIKE",
		1: "DISLIKE",
	}
	Reaction_ReactionType_value = map[string]int32{
		"LIKE":    0,
		"DISLIKE": 1,
	}
)

func (x Reaction_ReactionType) Enum() *Reaction_ReactionType {
	p := new(Reaction_ReactionType)
	*p = x
	return p
}

func (x Reaction_ReactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reaction_ReactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_post_service_proto_enumTypes[0].Descriptor()
}

func (Reaction_ReactionType) Type() protoreflect.EnumType {
	return &file_post_service_proto_enumTypes[0]
}

func (x Reaction_ReactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reaction_ReactionType.Descriptor instead.
func (Reaction_ReactionType) EnumDescriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{12, 0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{0}
}

type GetMultipleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetMultipleResponse) Reset() {
	*x = GetMultipleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultipleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultipleResponse) ProtoMessage() {}

func (x *GetMultipleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultipleResponse.ProtoReflect.Descriptor instead.
func (*GetMultipleResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetMultipleResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type NewPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *NewPost) Reset() {
	*x = NewPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPost) ProtoMessage() {}

func (x *NewPost) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPost.ProtoReflect.Descriptor instead.
func (*NewPost) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *NewPost) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type CommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId  string   `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
	Comment *Comment `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentRequest) Reset() {
	*x = CommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentRequest) ProtoMessage() {}

func (x *CommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentRequest.ProtoReflect.Descriptor instead.
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *CommentRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CommentRequest) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=commentId,proto3" json:"commentId,omitempty"`
}

func (x *CommentResponse) Reset() {
	*x = CommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentResponse) ProtoMessage() {}

func (x *CommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentResponse.ProtoReflect.Descriptor instead.
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{4}
}

func (x *CommentResponse) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{7}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetByUserRequest) Reset() {
	*x = GetByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserRequest) ProtoMessage() {}

func (x *GetByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserRequest.ProtoReflect.Descriptor instead.
func (*GetByUserRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetByUserRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ReactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reaction *Reaction `protobuf:"bytes,1,opt,name=reaction,proto3" json:"reaction,omitempty"`
}

func (x *ReactionRequest) Reset() {
	*x = ReactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionRequest) ProtoMessage() {}

func (x *ReactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionRequest.ProtoReflect.Descriptor instead.
func (*ReactionRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{10}
}

func (x *ReactionRequest) GetReaction() *Reaction {
	if x != nil {
		return x.Reaction
	}
	return nil
}

type ReactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *ReactionResponse) Reset() {
	*x = ReactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionResponse) ProtoMessage() {}

func (x *ReactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionResponse.ProtoReflect.Descriptor instead.
func (*ReactionResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{11}
}

func (x *ReactionResponse) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string                `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PostId       string                `protobuf:"bytes,2,opt,name=postId,proto3" json:"postId,omitempty"`
	ReactionType Reaction_ReactionType `protobuf:"varint,3,opt,name=reactionType,proto3,enum=Reaction_ReactionType" json:"reactionType,omitempty"`
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{12}
}

func (x *Reaction) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Reaction) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Reaction) GetReactionType() Reaction_ReactionType {
	if x != nil {
		return x.ReactionType
	}
	return Reaction_LIKE
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User      string                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Posttext  string                 `protobuf:"bytes,3,opt,name=posttext,proto3" json:"posttext,omitempty"`
	Image     string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Links     []string               `protobuf:"bytes,5,rep,name=links,proto3" json:"links,omitempty"`
	Date      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Reactions []*Reaction            `protobuf:"bytes,7,rep,name=reactions,proto3" json:"reactions,omitempty"`
	Comments  []*Comment             `protobuf:"bytes,8,rep,name=comments,proto3" json:"comments,omitempty"`
	Isdeleted bool                   `protobuf:"varint,9,opt,name=isdeleted,proto3" json:"isdeleted,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{13}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Post) GetPosttext() string {
	if x != nil {
		return x.Posttext
	}
	return ""
}

func (x *Post) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Post) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Post) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Post) GetReactions() []*Reaction {
	if x != nil {
		return x.Reactions
	}
	return nil
}

func (x *Post) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Post) GetIsdeleted() bool {
	if x != nil {
		return x.Isdeleted
	}
	return false
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content  string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Date     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Username string                 `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{14}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_post_service_proto protoreflect.FileDescriptor

var file_post_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x32, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x22, 0x24, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x0f, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x26, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x2a, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0c,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x22,
	0x8f, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x22, 0x7f, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0xb3, 0x03, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x35, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x12,
	0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x35, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x4e, 0x65, 0x77, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x05, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x12, 0x51, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x63, 0x74, 0x54, 0x6f, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x0b,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x3a, 0x08, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x16, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x7b,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_service_proto_rawDescOnce sync.Once
	file_post_service_proto_rawDescData = file_post_service_proto_rawDesc
)

func file_post_service_proto_rawDescGZIP() []byte {
	file_post_service_proto_rawDescOnce.Do(func() {
		file_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_service_proto_rawDescData)
	})
	return file_post_service_proto_rawDescData
}

var file_post_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_post_service_proto_goTypes = []interface{}{
	(Reaction_ReactionType)(0),    // 0: Reaction.ReactionType
	(*Empty)(nil),                 // 1: Empty
	(*GetMultipleResponse)(nil),   // 2: GetMultipleResponse
	(*NewPost)(nil),               // 3: NewPost
	(*CommentRequest)(nil),        // 4: CommentRequest
	(*CommentResponse)(nil),       // 5: CommentResponse
	(*GetRequest)(nil),            // 6: GetRequest
	(*GetResponse)(nil),           // 7: GetResponse
	(*GetAllRequest)(nil),         // 8: GetAllRequest
	(*GetAllResponse)(nil),        // 9: GetAllResponse
	(*GetByUserRequest)(nil),      // 10: GetByUserRequest
	(*ReactionRequest)(nil),       // 11: ReactionRequest
	(*ReactionResponse)(nil),      // 12: ReactionResponse
	(*Reaction)(nil),              // 13: Reaction
	(*Post)(nil),                  // 14: Post
	(*Comment)(nil),               // 15: Comment
	(*timestamppb.Timestamp)(nil), // 16: google.protobuf.Timestamp
}
var file_post_service_proto_depIdxs = []int32{
	14, // 0: GetMultipleResponse.posts:type_name -> Post
	14, // 1: NewPost.post:type_name -> Post
	15, // 2: CommentRequest.comment:type_name -> Comment
	14, // 3: GetResponse.post:type_name -> Post
	14, // 4: GetAllResponse.posts:type_name -> Post
	13, // 5: ReactionRequest.reaction:type_name -> Reaction
	0,  // 6: Reaction.reactionType:type_name -> Reaction.ReactionType
	16, // 7: Post.date:type_name -> google.protobuf.Timestamp
	13, // 8: Post.reactions:type_name -> Reaction
	15, // 9: Post.comments:type_name -> Comment
	16, // 10: Comment.date:type_name -> google.protobuf.Timestamp
	6,  // 11: PostService.Get:input_type -> GetRequest
	1,  // 12: PostService.GetAll:input_type -> Empty
	10, // 13: PostService.GetByUser:input_type -> GetByUserRequest
	3,  // 14: PostService.CreatePost:input_type -> NewPost
	11, // 15: PostService.ReactToPost:input_type -> ReactionRequest
	4,  // 16: PostService.CreateCommentOnPost:input_type -> CommentRequest
	7,  // 17: PostService.Get:output_type -> GetResponse
	2,  // 18: PostService.GetAll:output_type -> GetMultipleResponse
	9,  // 19: PostService.GetByUser:output_type -> GetAllResponse
	3,  // 20: PostService.CreatePost:output_type -> NewPost
	12, // 21: PostService.ReactToPost:output_type -> ReactionResponse
	5,  // 22: PostService.CreateCommentOnPost:output_type -> CommentResponse
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_post_service_proto_init() }
func file_post_service_proto_init() {
	if File_post_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_post_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultipleResponse); i {
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
		file_post_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPost); i {
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
		file_post_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentRequest); i {
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
		file_post_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentResponse); i {
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
		file_post_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_post_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_post_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_post_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_post_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUserRequest); i {
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
		file_post_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionRequest); i {
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
		file_post_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionResponse); i {
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
		file_post_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
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
		file_post_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_post_service_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
			RawDescriptor: file_post_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_service_proto_goTypes,
		DependencyIndexes: file_post_service_proto_depIdxs,
		EnumInfos:         file_post_service_proto_enumTypes,
		MessageInfos:      file_post_service_proto_msgTypes,
	}.Build()
	File_post_service_proto = out.File
	file_post_service_proto_rawDesc = nil
	file_post_service_proto_goTypes = nil
	file_post_service_proto_depIdxs = nil
}
