// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: post.proto

package protobuf

import (
	protobuf "github.com/Etpmls/Etpmls-Micro/v2/protobuf"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PostCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId   uint32             `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Content      string             `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	TemplatePath string             `protobuf:"bytes,4,opt,name=template_path,json=templatePath,proto3" json:"template_path,omitempty"`
	UrlPath      string             `protobuf:"bytes,5,opt,name=url_path,json=urlPath,proto3" json:"url_path,omitempty"`
	Language     string             `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	Summary      string             `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	Sort         uint32             `protobuf:"varint,8,opt,name=sort,proto3" json:"sort,omitempty"`
	Parameter    string             `protobuf:"bytes,9,opt,name=parameter,proto3" json:"parameter,omitempty"`
	Status       uint32             `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	Thumbnail    []*PostCreate_Path `protobuf:"bytes,11,rep,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	ContentImage []*PostCreate_Path `protobuf:"bytes,12,rep,name=content_image,json=contentImage,proto3" json:"content_image,omitempty"`
}

func (x *PostCreate) Reset() {
	*x = PostCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCreate) ProtoMessage() {}

func (x *PostCreate) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCreate.ProtoReflect.Descriptor instead.
func (*PostCreate) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *PostCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostCreate) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *PostCreate) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostCreate) GetTemplatePath() string {
	if x != nil {
		return x.TemplatePath
	}
	return ""
}

func (x *PostCreate) GetUrlPath() string {
	if x != nil {
		return x.UrlPath
	}
	return ""
}

func (x *PostCreate) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *PostCreate) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *PostCreate) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *PostCreate) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

func (x *PostCreate) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PostCreate) GetThumbnail() []*PostCreate_Path {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *PostCreate) GetContentImage() []*PostCreate_Path {
	if x != nil {
		return x.ContentImage
	}
	return nil
}

type PostEdit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId   uint32           `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Content      string           `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	TemplatePath string           `protobuf:"bytes,4,opt,name=template_path,json=templatePath,proto3" json:"template_path,omitempty"`
	UrlPath      string           `protobuf:"bytes,5,opt,name=url_path,json=urlPath,proto3" json:"url_path,omitempty"`
	Language     string           `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	Summary      string           `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	Sort         uint32           `protobuf:"varint,8,opt,name=sort,proto3" json:"sort,omitempty"`
	Parameter    string           `protobuf:"bytes,9,opt,name=parameter,proto3" json:"parameter,omitempty"`
	Status       uint32           `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	Thumbnail    []*PostEdit_Path `protobuf:"bytes,11,rep,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	ContentImage []*PostEdit_Path `protobuf:"bytes,12,rep,name=content_image,json=contentImage,proto3" json:"content_image,omitempty"`
	Id           uint32           `protobuf:"varint,13,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PostEdit) Reset() {
	*x = PostEdit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEdit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEdit) ProtoMessage() {}

func (x *PostEdit) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEdit.ProtoReflect.Descriptor instead.
func (*PostEdit) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *PostEdit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostEdit) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *PostEdit) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostEdit) GetTemplatePath() string {
	if x != nil {
		return x.TemplatePath
	}
	return ""
}

func (x *PostEdit) GetUrlPath() string {
	if x != nil {
		return x.UrlPath
	}
	return ""
}

func (x *PostEdit) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *PostEdit) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *PostEdit) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *PostEdit) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

func (x *PostEdit) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PostEdit) GetThumbnail() []*PostEdit_Path {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *PostEdit) GetContentImage() []*PostEdit_Path {
	if x != nil {
		return x.ContentImage
	}
	return nil
}

func (x *PostEdit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PostDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*protobuf.Id `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PostDelete) Reset() {
	*x = PostDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDelete) ProtoMessage() {}

func (x *PostDelete) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDelete.ProtoReflect.Descriptor instead.
func (*PostDelete) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *PostDelete) GetPosts() []*protobuf.Id {
	if x != nil {
		return x.Posts
	}
	return nil
}

type PostCreate_Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PostCreate_Path) Reset() {
	*x = PostCreate_Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCreate_Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCreate_Path) ProtoMessage() {}

func (x *PostCreate_Path) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCreate_Path.ProtoReflect.Descriptor instead.
func (*PostCreate_Path) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PostCreate_Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type PostEdit_Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PostEdit_Path) Reset() {
	*x = PostEdit_Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEdit_Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEdit_Path) ProtoMessage() {}

func (x *PostEdit_Path) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEdit_Path.ProtoReflect.Descriptor instead.
func (*PostEdit_Path) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PostEdit_Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x20, 0x45, 0x74, 0x70, 0x6d, 0x6c, 0x73, 0x2f, 0x45,
	0x74, 0x70, 0x6d, 0x6c, 0x73, 0x2d, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x03, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72, 0x6c,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x3e, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x0a,
	0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xba, 0x03, 0x0a, 0x08, 0x50, 0x6f,
	0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72,
	0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72,
	0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x1a, 0x0a, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x33, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x64, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x32, 0xec, 0x02, 0x0a, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x53, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x6d,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x15, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x65, 0x64, 0x69,
	0x74, 0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x2a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x59, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x65, 0x6d, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData = file_post_proto_rawDesc
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_proto_rawDescData)
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_post_proto_goTypes = []interface{}{
	(*PostCreate)(nil),          // 0: protobuf.PostCreate
	(*PostEdit)(nil),            // 1: protobuf.PostEdit
	(*PostDelete)(nil),          // 2: protobuf.PostDelete
	(*PostCreate_Path)(nil),     // 3: protobuf.PostCreate.Path
	(*PostEdit_Path)(nil),       // 4: protobuf.PostEdit.Path
	(*protobuf.Id)(nil),         // 5: em_protobuf.Id
	(*protobuf.Pagination)(nil), // 6: em_protobuf.Pagination
	(*protobuf.Response)(nil),   // 7: em_protobuf.Response
}
var file_post_proto_depIdxs = []int32{
	3, // 0: protobuf.PostCreate.thumbnail:type_name -> protobuf.PostCreate.Path
	3, // 1: protobuf.PostCreate.content_image:type_name -> protobuf.PostCreate.Path
	4, // 2: protobuf.PostEdit.thumbnail:type_name -> protobuf.PostEdit.Path
	4, // 3: protobuf.PostEdit.content_image:type_name -> protobuf.PostEdit.Path
	5, // 4: protobuf.PostDelete.posts:type_name -> em_protobuf.Id
	0, // 5: protobuf.Post.Create:input_type -> protobuf.PostCreate
	1, // 6: protobuf.Post.Edit:input_type -> protobuf.PostEdit
	2, // 7: protobuf.Post.Delete:input_type -> protobuf.PostDelete
	6, // 8: protobuf.Post.GetAll:input_type -> em_protobuf.Pagination
	7, // 9: protobuf.Post.Create:output_type -> em_protobuf.Response
	7, // 10: protobuf.Post.Edit:output_type -> em_protobuf.Response
	7, // 11: protobuf.Post.Delete:output_type -> em_protobuf.Response
	7, // 12: protobuf.Post.GetAll:output_type -> em_protobuf.Response
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCreate); i {
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
		file_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEdit); i {
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
		file_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostDelete); i {
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
		file_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCreate_Path); i {
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
		file_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEdit_Path); i {
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
			RawDescriptor: file_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_rawDesc = nil
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
