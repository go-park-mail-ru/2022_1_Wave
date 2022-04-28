// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: album/albumProto/album.proto

package albumProto

import (
	gatewayProto "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	trackProto "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id" example:"8" db:"id" validate:"min=0,nonnil"`                                                  // @gotags: json:"id" example:"8" db:"id" validate:"min=0,nonnil"
	Title           string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" example:"Mercury" db:"title" validate:"max=256,nonnil"`                                             // @gotags: json:"title" example:"Mercury" db:"title" validate:"max=256,nonnil"
	ArtistId        int64  `protobuf:"varint,3,opt,name=artist_id,json=artistId,proto3" json:"artistId" example:"4" db:"artist_id" validate:"min=0,nonnil"`                      // @gotags: json:"artistId" example:"4" db:"artist_id" validate:"min=0,nonnil"
	CountLikes      int64  `protobuf:"varint,4,opt,name=count_likes,json=countLikes,proto3" json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`                // @gotags: json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"
	CountListenings int64  `protobuf:"varint,5,opt,name=count_listenings,json=countListenings,proto3" json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"` // @gotags: json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"
	Date            int64  `protobuf:"varint,6,opt,name=date,proto3" json:"date" example:"0" db:"date,nonnil"`                                              // @gotags: json:"date" example:"0" db:"date,nonnil"
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_albumProto_album_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_album_albumProto_album_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_album_albumProto_album_proto_rawDescGZIP(), []int{0}
}

func (x *Album) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Album) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Album) GetArtistId() int64 {
	if x != nil {
		return x.ArtistId
	}
	return 0
}

func (x *Album) GetCountLikes() int64 {
	if x != nil {
		return x.CountLikes
	}
	return 0
}

func (x *Album) GetCountListenings() int64 {
	if x != nil {
		return x.CountListenings
	}
	return 0
}

func (x *Album) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type AlbumDataTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64                           `protobuf:"varint,1,opt,name=id,proto3" json:"id" example:"1"`        // @gotags: json:"id" example:"1"
	Title  string                          `protobuf:"bytes,2,opt,name=title,proto3" json:"title" example:"Mercury"`   // @gotags: json:"title" example:"Mercury"
	Artist string                          `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist" example:"Hexed"` // @gotags: json:"artist" example:"Hexed"
	Cover  string                          `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover" example:"assets/album_1.png"`   // @gotags: json:"cover" example:"assets/album_1.png"
	Tracks []*trackProto.TrackDataTransfer `protobuf:"bytes,5,rep,name=Tracks,proto3" json:"tracks"` // @gotags: json:"tracks"
}

func (x *AlbumDataTransfer) Reset() {
	*x = AlbumDataTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_albumProto_album_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumDataTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumDataTransfer) ProtoMessage() {}

func (x *AlbumDataTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_album_albumProto_album_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumDataTransfer.ProtoReflect.Descriptor instead.
func (*AlbumDataTransfer) Descriptor() ([]byte, []int) {
	return file_album_albumProto_album_proto_rawDescGZIP(), []int{1}
}

func (x *AlbumDataTransfer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlbumDataTransfer) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AlbumDataTransfer) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *AlbumDataTransfer) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *AlbumDataTransfer) GetTracks() []*trackProto.TrackDataTransfer {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type AlbumCover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id" example:"1" db:"id" validate:"min=0,nonnil"`                       // @gotags: json:"id" example:"1" db:"id" validate:"min=0,nonnil"
	Quote  string `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote" example:"some phrases" db:"quote" validate:"max=512,nonnil"`                  // @gotags: json:"quote" example:"some phrases" db:"quote" validate:"max=512,nonnil"
	IsDark bool   `protobuf:"varint,3,opt,name=is_dark,json=isDark,proto3" json:"isDark" example:"true" db:"is_dark" validate:"nonnil"` // @gotags: json:"isDark" example:"true" db:"is_dark" validate:"nonnil"
}

func (x *AlbumCover) Reset() {
	*x = AlbumCover{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_albumProto_album_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumCover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumCover) ProtoMessage() {}

func (x *AlbumCover) ProtoReflect() protoreflect.Message {
	mi := &file_album_albumProto_album_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumCover.ProtoReflect.Descriptor instead.
func (*AlbumCover) Descriptor() ([]byte, []int) {
	return file_album_albumProto_album_proto_rawDescGZIP(), []int{2}
}

func (x *AlbumCover) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlbumCover) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *AlbumCover) GetIsDark() bool {
	if x != nil {
		return x.IsDark
	}
	return false
}

type AlbumCoverDataTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote  string `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote" example:"some phrases"`                  // @gotags: json:"quote" example:"some phrases"
	IsDark bool   `protobuf:"varint,2,opt,name=is_dark,json=isDark,proto3" json:"isDark" example:"true"` // @gotags: json:"isDark" example:"true"
}

func (x *AlbumCoverDataTransfer) Reset() {
	*x = AlbumCoverDataTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_albumProto_album_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumCoverDataTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumCoverDataTransfer) ProtoMessage() {}

func (x *AlbumCoverDataTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_album_albumProto_album_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumCoverDataTransfer.ProtoReflect.Descriptor instead.
func (*AlbumCoverDataTransfer) Descriptor() ([]byte, []int) {
	return file_album_albumProto_album_proto_rawDescGZIP(), []int{3}
}

func (x *AlbumCoverDataTransfer) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *AlbumCoverDataTransfer) GetIsDark() bool {
	if x != nil {
		return x.IsDark
	}
	return false
}

type AlbumUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *AlbumDataTransfer `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AlbumUseCaseResponse) Reset() {
	*x = AlbumUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_albumProto_album_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumUseCaseResponse) ProtoMessage() {}

func (x *AlbumUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_album_albumProto_album_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumUseCaseResponse.ProtoReflect.Descriptor instead.
func (*AlbumUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_album_albumProto_album_proto_rawDescGZIP(), []int{4}
}

func (x *AlbumUseCaseResponse) GetData() *AlbumDataTransfer {
	if x != nil {
		return x.Data
	}
	return nil
}

type AlbumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums []*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *AlbumsResponse) Reset() {
	*x = AlbumsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_albumProto_album_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsResponse) ProtoMessage() {}

func (x *AlbumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_album_albumProto_album_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumsResponse.ProtoReflect.Descriptor instead.
func (*AlbumsResponse) Descriptor() ([]byte, []int) {
	return file_album_albumProto_album_proto_rawDescGZIP(), []int{5}
}

func (x *AlbumsResponse) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

type AlbumsCoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Covers []*AlbumCover `protobuf:"bytes,1,rep,name=covers,proto3" json:"covers,omitempty"`
}

func (x *AlbumsCoverResponse) Reset() {
	*x = AlbumsCoverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_albumProto_album_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsCoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsCoverResponse) ProtoMessage() {}

func (x *AlbumsCoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_album_albumProto_album_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumsCoverResponse.ProtoReflect.Descriptor instead.
func (*AlbumsCoverResponse) Descriptor() ([]byte, []int) {
	return file_album_albumProto_album_proto_rawDescGZIP(), []int{6}
}

func (x *AlbumsCoverResponse) GetCovers() []*AlbumCover {
	if x != nil {
		return x.Covers
	}
	return nil
}

var File_album_albumProto_album_proto protoreflect.FileDescriptor

var file_album_albumProto_album_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x06, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x4b, 0x0a,
	0x0a, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x64, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x61, 0x72, 0x6b, 0x22, 0x47, 0x0a, 0x16, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x64, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44,
	0x61, 0x72, 0x6b, 0x22, 0x44, 0x0a, 0x14, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x0e, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x22, 0x40, 0x0a, 0x13, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x73, 0x32, 0xfb, 0x08, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x55, 0x73, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c,
	0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x41, 0x72, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x49, 0x64,
	0x41, 0x72, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x49, 0x64, 0x41, 0x72, 0x67, 0x1a, 0x0c, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x76, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x49, 0x64, 0x41, 0x72, 0x67, 0x1a, 0x11, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x49, 0x64,
	0x41, 0x72, 0x67, 0x1a, 0x15, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x14, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x67, 0x1a, 0x15, 0x2e, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x49, 0x64, 0x41, 0x72, 0x67, 0x1a, 0x15, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73,
	0x12, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x41, 0x72, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x49, 0x64, 0x41, 0x72, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f,
	0x32, 0x30, 0x32, 0x32, 0x5f, 0x31, 0x5f, 0x57, 0x61, 0x76, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_album_albumProto_album_proto_rawDescOnce sync.Once
	file_album_albumProto_album_proto_rawDescData = file_album_albumProto_album_proto_rawDesc
)

func file_album_albumProto_album_proto_rawDescGZIP() []byte {
	file_album_albumProto_album_proto_rawDescOnce.Do(func() {
		file_album_albumProto_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_album_albumProto_album_proto_rawDescData)
	})
	return file_album_albumProto_album_proto_rawDescData
}

var file_album_albumProto_album_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_album_albumProto_album_proto_goTypes = []interface{}{
	(*Album)(nil),                         // 0: album.Album
	(*AlbumDataTransfer)(nil),             // 1: album.AlbumDataTransfer
	(*AlbumCover)(nil),                    // 2: album.AlbumCover
	(*AlbumCoverDataTransfer)(nil),        // 3: album.AlbumCoverDataTransfer
	(*AlbumUseCaseResponse)(nil),          // 4: album.AlbumUseCaseResponse
	(*AlbumsResponse)(nil),                // 5: album.AlbumsResponse
	(*AlbumsCoverResponse)(nil),           // 6: album.AlbumsCoverResponse
	(*trackProto.TrackDataTransfer)(nil),  // 7: track.TrackDataTransfer
	(*empty.Empty)(nil),                   // 8: google.protobuf.Empty
	(*gatewayProto.IdArg)(nil),            // 9: gateway.IdArg
	(*gatewayProto.StringArg)(nil),        // 10: gateway.StringArg
	(*gatewayProto.UserIdAlbumIdArg)(nil), // 11: gateway.UserIdAlbumIdArg
	(*gatewayProto.IntResponse)(nil),      // 12: gateway.IntResponse
}
var file_album_albumProto_album_proto_depIdxs = []int32{
	7,  // 0: album.AlbumDataTransfer.Tracks:type_name -> track.TrackDataTransfer
	1,  // 1: album.AlbumUseCaseResponse.data:type_name -> album.AlbumDataTransfer
	0,  // 2: album.AlbumsResponse.albums:type_name -> album.Album
	2,  // 3: album.AlbumsCoverResponse.covers:type_name -> album.AlbumCover
	8,  // 4: album.AlbumUseCase.GetAll:input_type -> google.protobuf.Empty
	8,  // 5: album.AlbumUseCase.GetAllCovers:input_type -> google.protobuf.Empty
	8,  // 6: album.AlbumUseCase.GetLastId:input_type -> google.protobuf.Empty
	8,  // 7: album.AlbumUseCase.GetLastCoverId:input_type -> google.protobuf.Empty
	0,  // 8: album.AlbumUseCase.Create:input_type -> album.Album
	2,  // 9: album.AlbumUseCase.CreateCover:input_type -> album.AlbumCover
	0,  // 10: album.AlbumUseCase.Update:input_type -> album.Album
	2,  // 11: album.AlbumUseCase.UpdateCover:input_type -> album.AlbumCover
	9,  // 12: album.AlbumUseCase.Delete:input_type -> gateway.IdArg
	9,  // 13: album.AlbumUseCase.DeleteCover:input_type -> gateway.IdArg
	9,  // 14: album.AlbumUseCase.GetById:input_type -> gateway.IdArg
	9,  // 15: album.AlbumUseCase.GetCoverById:input_type -> gateway.IdArg
	8,  // 16: album.AlbumUseCase.GetPopular:input_type -> google.protobuf.Empty
	9,  // 17: album.AlbumUseCase.GetAlbumsFromArtist:input_type -> gateway.IdArg
	8,  // 18: album.AlbumUseCase.GetSize:input_type -> google.protobuf.Empty
	10, // 19: album.AlbumUseCase.SearchByTitle:input_type -> gateway.StringArg
	9,  // 20: album.AlbumUseCase.GetFavorites:input_type -> gateway.IdArg
	11, // 21: album.AlbumUseCase.AddToFavorites:input_type -> gateway.UserIdAlbumIdArg
	11, // 22: album.AlbumUseCase.RemoveFromFavorites:input_type -> gateway.UserIdAlbumIdArg
	5,  // 23: album.AlbumUseCase.GetAll:output_type -> album.AlbumsResponse
	6,  // 24: album.AlbumUseCase.GetAllCovers:output_type -> album.AlbumsCoverResponse
	12, // 25: album.AlbumUseCase.GetLastId:output_type -> gateway.IntResponse
	12, // 26: album.AlbumUseCase.GetLastCoverId:output_type -> gateway.IntResponse
	8,  // 27: album.AlbumUseCase.Create:output_type -> google.protobuf.Empty
	8,  // 28: album.AlbumUseCase.CreateCover:output_type -> google.protobuf.Empty
	8,  // 29: album.AlbumUseCase.Update:output_type -> google.protobuf.Empty
	8,  // 30: album.AlbumUseCase.UpdateCover:output_type -> google.protobuf.Empty
	8,  // 31: album.AlbumUseCase.Delete:output_type -> google.protobuf.Empty
	8,  // 32: album.AlbumUseCase.DeleteCover:output_type -> google.protobuf.Empty
	0,  // 33: album.AlbumUseCase.GetById:output_type -> album.Album
	2,  // 34: album.AlbumUseCase.GetCoverById:output_type -> album.AlbumCover
	5,  // 35: album.AlbumUseCase.GetPopular:output_type -> album.AlbumsResponse
	5,  // 36: album.AlbumUseCase.GetAlbumsFromArtist:output_type -> album.AlbumsResponse
	12, // 37: album.AlbumUseCase.GetSize:output_type -> gateway.IntResponse
	5,  // 38: album.AlbumUseCase.SearchByTitle:output_type -> album.AlbumsResponse
	5,  // 39: album.AlbumUseCase.GetFavorites:output_type -> album.AlbumsResponse
	8,  // 40: album.AlbumUseCase.AddToFavorites:output_type -> google.protobuf.Empty
	8,  // 41: album.AlbumUseCase.RemoveFromFavorites:output_type -> google.protobuf.Empty
	23, // [23:42] is the sub-list for method output_type
	4,  // [4:23] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_album_albumProto_album_proto_init() }
func file_album_albumProto_album_proto_init() {
	if File_album_albumProto_album_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_album_albumProto_album_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
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
		file_album_albumProto_album_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumDataTransfer); i {
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
		file_album_albumProto_album_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumCover); i {
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
		file_album_albumProto_album_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumCoverDataTransfer); i {
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
		file_album_albumProto_album_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumUseCaseResponse); i {
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
		file_album_albumProto_album_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumsResponse); i {
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
		file_album_albumProto_album_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumsCoverResponse); i {
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
			RawDescriptor: file_album_albumProto_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_album_albumProto_album_proto_goTypes,
		DependencyIndexes: file_album_albumProto_album_proto_depIdxs,
		MessageInfos:      file_album_albumProto_album_proto_msgTypes,
	}.Build()
	File_album_albumProto_album_proto = out.File
	file_album_albumProto_album_proto_rawDesc = nil
	file_album_albumProto_album_proto_goTypes = nil
	file_album_albumProto_album_proto_depIdxs = nil
}
