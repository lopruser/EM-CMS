// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: category.proto

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

type CategoryCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentId         uint32                 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Kind             string                 `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	UrlPath          string                 `protobuf:"bytes,4,opt,name=url_path,json=urlPath,proto3" json:"url_path,omitempty"`
	Sort             uint32                 `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Summary          string                 `protobuf:"bytes,6,opt,name=summary,proto3" json:"summary,omitempty"`
	TemplatePath     string                 `protobuf:"bytes,7,opt,name=template_path,json=templatePath,proto3" json:"template_path,omitempty"`
	PostTemplatePath string                 `protobuf:"bytes,8,opt,name=post_template_path,json=postTemplatePath,proto3" json:"post_template_path,omitempty"`
	Status           int32                  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	IsHidden         int32                  `protobuf:"varint,10,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	IsMain           int32                  `protobuf:"varint,11,opt,name=is_main,json=isMain,proto3" json:"is_main,omitempty"`
	Thumbnail        []*CategoryCreate_Path `protobuf:"bytes,12,rep,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *CategoryCreate) Reset() {
	*x = CategoryCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryCreate) ProtoMessage() {}

func (x *CategoryCreate) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryCreate.ProtoReflect.Descriptor instead.
func (*CategoryCreate) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryCreate) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CategoryCreate) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CategoryCreate) GetUrlPath() string {
	if x != nil {
		return x.UrlPath
	}
	return ""
}

func (x *CategoryCreate) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *CategoryCreate) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CategoryCreate) GetTemplatePath() string {
	if x != nil {
		return x.TemplatePath
	}
	return ""
}

func (x *CategoryCreate) GetPostTemplatePath() string {
	if x != nil {
		return x.PostTemplatePath
	}
	return ""
}

func (x *CategoryCreate) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CategoryCreate) GetIsHidden() int32 {
	if x != nil {
		return x.IsHidden
	}
	return 0
}

func (x *CategoryCreate) GetIsMain() int32 {
	if x != nil {
		return x.IsMain
	}
	return 0
}

func (x *CategoryCreate) GetThumbnail() []*CategoryCreate_Path {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

type CategoryEdit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentId         uint32               `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Kind             string               `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	UrlPath          string               `protobuf:"bytes,4,opt,name=url_path,json=urlPath,proto3" json:"url_path,omitempty"`
	Sort             uint32               `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Summary          string               `protobuf:"bytes,6,opt,name=summary,proto3" json:"summary,omitempty"`
	TemplatePath     string               `protobuf:"bytes,7,opt,name=template_path,json=templatePath,proto3" json:"template_path,omitempty"`
	PostTemplatePath string               `protobuf:"bytes,8,opt,name=post_template_path,json=postTemplatePath,proto3" json:"post_template_path,omitempty"`
	Status           int32                `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	IsHidden         int32                `protobuf:"varint,10,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	IsMain           int32                `protobuf:"varint,11,opt,name=is_main,json=isMain,proto3" json:"is_main,omitempty"`
	Thumbnail        []*CategoryEdit_Path `protobuf:"bytes,12,rep,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Id               uint32               `protobuf:"varint,13,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CategoryEdit) Reset() {
	*x = CategoryEdit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryEdit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryEdit) ProtoMessage() {}

func (x *CategoryEdit) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryEdit.ProtoReflect.Descriptor instead.
func (*CategoryEdit) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryEdit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryEdit) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CategoryEdit) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CategoryEdit) GetUrlPath() string {
	if x != nil {
		return x.UrlPath
	}
	return ""
}

func (x *CategoryEdit) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *CategoryEdit) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CategoryEdit) GetTemplatePath() string {
	if x != nil {
		return x.TemplatePath
	}
	return ""
}

func (x *CategoryEdit) GetPostTemplatePath() string {
	if x != nil {
		return x.PostTemplatePath
	}
	return ""
}

func (x *CategoryEdit) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CategoryEdit) GetIsHidden() int32 {
	if x != nil {
		return x.IsHidden
	}
	return 0
}

func (x *CategoryEdit) GetIsMain() int32 {
	if x != nil {
		return x.IsMain
	}
	return 0
}

func (x *CategoryEdit) GetThumbnail() []*CategoryEdit_Path {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *CategoryEdit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CategoryDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*protobuf.Id `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CategoryDelete) Reset() {
	*x = CategoryDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDelete) ProtoMessage() {}

func (x *CategoryDelete) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDelete.ProtoReflect.Descriptor instead.
func (*CategoryDelete) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryDelete) GetCategories() []*protobuf.Id {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CategoryCreate_Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CategoryCreate_Path) Reset() {
	*x = CategoryCreate_Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryCreate_Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryCreate_Path) ProtoMessage() {}

func (x *CategoryCreate_Path) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryCreate_Path.ProtoReflect.Descriptor instead.
func (*CategoryCreate_Path) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CategoryCreate_Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type CategoryEdit_Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CategoryEdit_Path) Reset() {
	*x = CategoryEdit_Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryEdit_Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryEdit_Path) ProtoMessage() {}

func (x *CategoryEdit_Path) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryEdit_Path.ProtoReflect.Descriptor instead.
func (*CategoryEdit_Path) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CategoryEdit_Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_category_proto protoreflect.FileDescriptor

var file_category_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x20, 0x45, 0x74, 0x70, 0x6d,
	0x6c, 0x73, 0x2f, 0x45, 0x74, 0x70, 0x6d, 0x6c, 0x73, 0x2d, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x03, 0x0a, 0x0e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f,
	0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x3b, 0x0a, 0x09,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x1a, 0x1a, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xa4, 0x03, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x45, 0x64, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x72, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x72, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69,
	0x73, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x39, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x64, 0x69, 0x74,
	0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x1a, 0x1a, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x41, 0x0a, 0x0e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2f,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x64, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32,
	0x89, 0x03, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x61, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22,
	0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x5b, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x64, 0x69, 0x74, 0x1a,
	0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x1a, 0x19,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a,
	0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x5a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x12, 0x2e, 0x65, 0x6d, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15,
	0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_category_proto_rawDescOnce sync.Once
	file_category_proto_rawDescData = file_category_proto_rawDesc
)

func file_category_proto_rawDescGZIP() []byte {
	file_category_proto_rawDescOnce.Do(func() {
		file_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_category_proto_rawDescData)
	})
	return file_category_proto_rawDescData
}

var file_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_category_proto_goTypes = []interface{}{
	(*CategoryCreate)(nil),      // 0: protobuf.CategoryCreate
	(*CategoryEdit)(nil),        // 1: protobuf.CategoryEdit
	(*CategoryDelete)(nil),      // 2: protobuf.CategoryDelete
	(*CategoryCreate_Path)(nil), // 3: protobuf.CategoryCreate.Path
	(*CategoryEdit_Path)(nil),   // 4: protobuf.CategoryEdit.Path
	(*protobuf.Id)(nil),         // 5: em_protobuf.Id
	(*protobuf.Empty)(nil),      // 6: em_protobuf.Empty
	(*protobuf.Response)(nil),   // 7: em_protobuf.Response
}
var file_category_proto_depIdxs = []int32{
	3, // 0: protobuf.CategoryCreate.thumbnail:type_name -> protobuf.CategoryCreate.Path
	4, // 1: protobuf.CategoryEdit.thumbnail:type_name -> protobuf.CategoryEdit.Path
	5, // 2: protobuf.CategoryDelete.categories:type_name -> em_protobuf.Id
	0, // 3: protobuf.Category.Create:input_type -> protobuf.CategoryCreate
	1, // 4: protobuf.Category.Edit:input_type -> protobuf.CategoryEdit
	2, // 5: protobuf.Category.Delete:input_type -> protobuf.CategoryDelete
	6, // 6: protobuf.Category.GetTree:input_type -> em_protobuf.Empty
	7, // 7: protobuf.Category.Create:output_type -> em_protobuf.Response
	7, // 8: protobuf.Category.Edit:output_type -> em_protobuf.Response
	7, // 9: protobuf.Category.Delete:output_type -> em_protobuf.Response
	7, // 10: protobuf.Category.GetTree:output_type -> em_protobuf.Response
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_category_proto_init() }
func file_category_proto_init() {
	if File_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryCreate); i {
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
		file_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryEdit); i {
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
		file_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDelete); i {
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
		file_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryCreate_Path); i {
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
		file_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryEdit_Path); i {
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
			RawDescriptor: file_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_category_proto_goTypes,
		DependencyIndexes: file_category_proto_depIdxs,
		MessageInfos:      file_category_proto_msgTypes,
	}.Build()
	File_category_proto = out.File
	file_category_proto_rawDesc = nil
	file_category_proto_goTypes = nil
	file_category_proto_depIdxs = nil
}
