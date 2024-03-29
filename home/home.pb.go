// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: home.proto

package home

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=Ping,proto3" json:"Ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=Pong,proto3" json:"Pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// top
type TopInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,10,opt,name=Title,proto3" json:"Title,omitempty"`
	Url   string `protobuf:"bytes,20,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *TopInfo) Reset() {
	*x = TopInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopInfo) ProtoMessage() {}

func (x *TopInfo) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopInfo.ProtoReflect.Descriptor instead.
func (*TopInfo) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{2}
}

func (x *TopInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TopInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type SearchTopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Where int64 `protobuf:"varint,10,opt,name=Where,proto3" json:"Where,omitempty"`
}

func (x *SearchTopRequest) Reset() {
	*x = SearchTopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTopRequest) ProtoMessage() {}

func (x *SearchTopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTopRequest.ProtoReflect.Descriptor instead.
func (*SearchTopRequest) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{3}
}

func (x *SearchTopRequest) GetWhere() int64 {
	if x != nil {
		return x.Where
	}
	return 0
}

type SearchTopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopInfos []*TopInfo `protobuf:"bytes,10,rep,name=TopInfos,proto3" json:"TopInfos,omitempty"`
}

func (x *SearchTopResponse) Reset() {
	*x = SearchTopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTopResponse) ProtoMessage() {}

func (x *SearchTopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTopResponse.ProtoReflect.Descriptor instead.
func (*SearchTopResponse) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{4}
}

func (x *SearchTopResponse) GetTopInfos() []*TopInfo {
	if x != nil {
		return x.TopInfos
	}
	return nil
}

// Slideshow
type Slideshow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,10,opt,name=Title,proto3" json:"Title,omitempty"`
	Url   string `protobuf:"bytes,20,opt,name=Url,proto3" json:"Url,omitempty"`
	Image string `protobuf:"bytes,30,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *Slideshow) Reset() {
	*x = Slideshow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slideshow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slideshow) ProtoMessage() {}

func (x *Slideshow) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slideshow.ProtoReflect.Descriptor instead.
func (*Slideshow) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{5}
}

func (x *Slideshow) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Slideshow) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Slideshow) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type SearchSlideshowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Where int64 `protobuf:"varint,10,opt,name=Where,proto3" json:"Where,omitempty"`
}

func (x *SearchSlideshowRequest) Reset() {
	*x = SearchSlideshowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSlideshowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSlideshowRequest) ProtoMessage() {}

func (x *SearchSlideshowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSlideshowRequest.ProtoReflect.Descriptor instead.
func (*SearchSlideshowRequest) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{6}
}

func (x *SearchSlideshowRequest) GetWhere() int64 {
	if x != nil {
		return x.Where
	}
	return 0
}

type SearchSlideshowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slideshows []*Slideshow `protobuf:"bytes,10,rep,name=Slideshows,proto3" json:"Slideshows,omitempty"`
}

func (x *SearchSlideshowResponse) Reset() {
	*x = SearchSlideshowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSlideshowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSlideshowResponse) ProtoMessage() {}

func (x *SearchSlideshowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSlideshowResponse.ProtoReflect.Descriptor instead.
func (*SearchSlideshowResponse) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{7}
}

func (x *SearchSlideshowResponse) GetSlideshows() []*Slideshow {
	if x != nil {
		return x.Slideshows
	}
	return nil
}

// MainBody
type MainBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,10,opt,name=Title,proto3" json:"Title,omitempty"`
	Url   string `protobuf:"bytes,20,opt,name=Url,proto3" json:"Url,omitempty"`
	Image string `protobuf:"bytes,30,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *MainBody) Reset() {
	*x = MainBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainBody) ProtoMessage() {}

func (x *MainBody) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainBody.ProtoReflect.Descriptor instead.
func (*MainBody) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{8}
}

func (x *MainBody) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MainBody) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MainBody) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type SearchMainBodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Where int64 `protobuf:"varint,10,opt,name=Where,proto3" json:"Where,omitempty"`
}

func (x *SearchMainBodyRequest) Reset() {
	*x = SearchMainBodyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMainBodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMainBodyRequest) ProtoMessage() {}

func (x *SearchMainBodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMainBodyRequest.ProtoReflect.Descriptor instead.
func (*SearchMainBodyRequest) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{9}
}

func (x *SearchMainBodyRequest) GetWhere() int64 {
	if x != nil {
		return x.Where
	}
	return 0
}

type SearchMainBodyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainBodys []*MainBody `protobuf:"bytes,10,rep,name=MainBodys,proto3" json:"MainBodys,omitempty"`
}

func (x *SearchMainBodyResponse) Reset() {
	*x = SearchMainBodyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMainBodyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMainBodyResponse) ProtoMessage() {}

func (x *SearchMainBodyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMainBodyResponse.ProtoReflect.Descriptor instead.
func (*SearchMainBodyResponse) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{10}
}

func (x *SearchMainBodyResponse) GetMainBodys() []*MainBody {
	if x != nil {
		return x.MainBodys
	}
	return nil
}

// Bottom
type Bottom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,10,opt,name=Title,proto3" json:"Title,omitempty"`
	Url   string `protobuf:"bytes,20,opt,name=Url,proto3" json:"Url,omitempty"`
	Image string `protobuf:"bytes,30,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *Bottom) Reset() {
	*x = Bottom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bottom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bottom) ProtoMessage() {}

func (x *Bottom) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bottom.ProtoReflect.Descriptor instead.
func (*Bottom) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{11}
}

func (x *Bottom) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Bottom) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Bottom) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type SearchBottomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Where int64 `protobuf:"varint,10,opt,name=Where,proto3" json:"Where,omitempty"`
}

func (x *SearchBottomRequest) Reset() {
	*x = SearchBottomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBottomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBottomRequest) ProtoMessage() {}

func (x *SearchBottomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBottomRequest.ProtoReflect.Descriptor instead.
func (*SearchBottomRequest) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{12}
}

func (x *SearchBottomRequest) GetWhere() int64 {
	if x != nil {
		return x.Where
	}
	return 0
}

type SearchBottomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bottoms []*Bottom `protobuf:"bytes,10,rep,name=Bottoms,proto3" json:"Bottoms,omitempty"`
}

func (x *SearchBottomResponse) Reset() {
	*x = SearchBottomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBottomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBottomResponse) ProtoMessage() {}

func (x *SearchBottomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBottomResponse.ProtoReflect.Descriptor instead.
func (*SearchBottomResponse) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{13}
}

func (x *SearchBottomResponse) GetBottoms() []*Bottom {
	if x != nil {
		return x.Bottoms
	}
	return nil
}

var File_home_proto protoreflect.FileDescriptor

var file_home_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x6f,
	0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x6f, 0x6e,
	0x67, 0x22, 0x31, 0x0a, 0x07, 0x54, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x22, 0x3e,
	0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x54, 0x6f, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x54, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x49,
	0x0a, 0x09, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x16, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x22, 0x4a, 0x0a, 0x17, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f,
	0x77, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e,
	0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x52, 0x0a, 0x53, 0x6c, 0x69, 0x64, 0x65,
	0x73, 0x68, 0x6f, 0x77, 0x73, 0x22, 0x48, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x2d, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x22, 0x46,
	0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x6f, 0x64, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x4d, 0x61, 0x69, 0x6e,
	0x42, 0x6f, 0x64, 0x79, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x09, 0x4d, 0x61, 0x69,
	0x6e, 0x42, 0x6f, 0x64, 0x79, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2b,
	0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x74,
	0x6f, 0x6d, 0x52, 0x07, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x73, 0x32, 0xcf, 0x02, 0x0a, 0x04,
	0x48, 0x6f, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x12, 0x16, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x12, 0x1c, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73,
	0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x12, 0x19, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_home_proto_rawDescOnce sync.Once
	file_home_proto_rawDescData = file_home_proto_rawDesc
)

func file_home_proto_rawDescGZIP() []byte {
	file_home_proto_rawDescOnce.Do(func() {
		file_home_proto_rawDescData = protoimpl.X.CompressGZIP(file_home_proto_rawDescData)
	})
	return file_home_proto_rawDescData
}

var file_home_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_home_proto_goTypes = []interface{}{
	(*Request)(nil),                 // 0: home.Request
	(*Response)(nil),                // 1: home.Response
	(*TopInfo)(nil),                 // 2: home.TopInfo
	(*SearchTopRequest)(nil),        // 3: home.SearchTopRequest
	(*SearchTopResponse)(nil),       // 4: home.SearchTopResponse
	(*Slideshow)(nil),               // 5: home.Slideshow
	(*SearchSlideshowRequest)(nil),  // 6: home.SearchSlideshowRequest
	(*SearchSlideshowResponse)(nil), // 7: home.SearchSlideshowResponse
	(*MainBody)(nil),                // 8: home.MainBody
	(*SearchMainBodyRequest)(nil),   // 9: home.SearchMainBodyRequest
	(*SearchMainBodyResponse)(nil),  // 10: home.SearchMainBodyResponse
	(*Bottom)(nil),                  // 11: home.Bottom
	(*SearchBottomRequest)(nil),     // 12: home.SearchBottomRequest
	(*SearchBottomResponse)(nil),    // 13: home.SearchBottomResponse
}
var file_home_proto_depIdxs = []int32{
	2,  // 0: home.SearchTopResponse.TopInfos:type_name -> home.TopInfo
	5,  // 1: home.SearchSlideshowResponse.Slideshows:type_name -> home.Slideshow
	8,  // 2: home.SearchMainBodyResponse.MainBodys:type_name -> home.MainBody
	11, // 3: home.SearchBottomResponse.Bottoms:type_name -> home.Bottom
	0,  // 4: home.Home.Ping:input_type -> home.Request
	3,  // 5: home.Home.SearchTop:input_type -> home.SearchTopRequest
	6,  // 6: home.Home.SearchSlideshow:input_type -> home.SearchSlideshowRequest
	9,  // 7: home.Home.SearchMainBody:input_type -> home.SearchMainBodyRequest
	12, // 8: home.Home.SearchBottom:input_type -> home.SearchBottomRequest
	1,  // 9: home.Home.Ping:output_type -> home.Response
	4,  // 10: home.Home.SearchTop:output_type -> home.SearchTopResponse
	7,  // 11: home.Home.SearchSlideshow:output_type -> home.SearchSlideshowResponse
	10, // 12: home.Home.SearchMainBody:output_type -> home.SearchMainBodyResponse
	13, // 13: home.Home.SearchBottom:output_type -> home.SearchBottomResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_home_proto_init() }
func file_home_proto_init() {
	if File_home_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_home_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_home_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_home_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopInfo); i {
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
		file_home_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTopRequest); i {
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
		file_home_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTopResponse); i {
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
		file_home_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slideshow); i {
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
		file_home_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSlideshowRequest); i {
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
		file_home_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSlideshowResponse); i {
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
		file_home_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainBody); i {
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
		file_home_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMainBodyRequest); i {
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
		file_home_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMainBodyResponse); i {
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
		file_home_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bottom); i {
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
		file_home_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchBottomRequest); i {
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
		file_home_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchBottomResponse); i {
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
			RawDescriptor: file_home_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_home_proto_goTypes,
		DependencyIndexes: file_home_proto_depIdxs,
		MessageInfos:      file_home_proto_msgTypes,
	}.Build()
	File_home_proto = out.File
	file_home_proto_rawDesc = nil
	file_home_proto_goTypes = nil
	file_home_proto_depIdxs = nil
}
