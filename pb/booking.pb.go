// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: booking.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Booking_Status int32

const (
	Booking_Success Booking_Status = 0
	Booking_Fail    Booking_Status = 1
)

// Enum value maps for Booking_Status.
var (
	Booking_Status_name = map[int32]string{
		0: "Success",
		1: "Fail",
	}
	Booking_Status_value = map[string]int32{
		"Success": 0,
		"Fail":    1,
	}
)

func (x Booking_Status) Enum() *Booking_Status {
	p := new(Booking_Status)
	*p = x
	return p
}

func (x Booking_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_booking_proto_enumTypes[0].Descriptor()
}

func (Booking_Status) Type() protoreflect.EnumType {
	return &file_booking_proto_enumTypes[0]
}

func (x Booking_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_Status.Descriptor instead.
func (Booking_Status) EnumDescriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0, 0}
}

type BookingInput_Status int32

const (
	BookingInput_Success BookingInput_Status = 0
	BookingInput_Fail    BookingInput_Status = 1
)

// Enum value maps for BookingInput_Status.
var (
	BookingInput_Status_name = map[int32]string{
		0: "Success",
		1: "Fail",
	}
	BookingInput_Status_value = map[string]int32{
		"Success": 0,
		"Fail":    1,
	}
)

func (x BookingInput_Status) Enum() *BookingInput_Status {
	p := new(BookingInput_Status)
	*p = x
	return p
}

func (x BookingInput_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookingInput_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_booking_proto_enumTypes[1].Descriptor()
}

func (BookingInput_Status) Type() protoreflect.EnumType {
	return &file_booking_proto_enumTypes[1]
}

func (x BookingInput_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookingInput_Status.Descriptor instead.
func (BookingInput_Status) EnumDescriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1, 0}
}

type CreateBookingRequest_Status int32

const (
	CreateBookingRequest_Success CreateBookingRequest_Status = 0
	CreateBookingRequest_Fail    CreateBookingRequest_Status = 1
)

// Enum value maps for CreateBookingRequest_Status.
var (
	CreateBookingRequest_Status_name = map[int32]string{
		0: "Success",
		1: "Fail",
	}
	CreateBookingRequest_Status_value = map[string]int32{
		"Success": 0,
		"Fail":    1,
	}
)

func (x CreateBookingRequest_Status) Enum() *CreateBookingRequest_Status {
	p := new(CreateBookingRequest_Status)
	*p = x
	return p
}

func (x CreateBookingRequest_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateBookingRequest_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_booking_proto_enumTypes[2].Descriptor()
}

func (CreateBookingRequest_Status) Type() protoreflect.EnumType {
	return &file_booking_proto_enumTypes[2]
}

func (x CreateBookingRequest_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateBookingRequest_Status.Descriptor instead.
func (CreateBookingRequest_Status) EnumDescriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3, 0}
}

type UpdateBookingStatusRequest_Status int32

const (
	UpdateBookingStatusRequest_Success  UpdateBookingStatusRequest_Status = 0
	UpdateBookingStatusRequest_Fail     UpdateBookingStatusRequest_Status = 1
	UpdateBookingStatusRequest_Canceled UpdateBookingStatusRequest_Status = 2
)

// Enum value maps for UpdateBookingStatusRequest_Status.
var (
	UpdateBookingStatusRequest_Status_name = map[int32]string{
		0: "Success",
		1: "Fail",
		2: "Canceled",
	}
	UpdateBookingStatusRequest_Status_value = map[string]int32{
		"Success":  0,
		"Fail":     1,
		"Canceled": 2,
	}
)

func (x UpdateBookingStatusRequest_Status) Enum() *UpdateBookingStatusRequest_Status {
	p := new(UpdateBookingStatusRequest_Status)
	*p = x
	return p
}

func (x UpdateBookingStatusRequest_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateBookingStatusRequest_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_booking_proto_enumTypes[3].Descriptor()
}

func (UpdateBookingStatusRequest_Status) Type() protoreflect.EnumType {
	return &file_booking_proto_enumTypes[3]
}

func (x UpdateBookingStatusRequest_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateBookingStatusRequest_Status.Descriptor instead.
func (UpdateBookingStatusRequest_Status) EnumDescriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4, 0}
}

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TotalCost   float32              `protobuf:"fixed32,2,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	TotalTicket int32                `protobuf:"varint,3,opt,name=total_ticket,json=totalTicket,proto3" json:"total_ticket,omitempty"`
	Status      Booking_Status       `protobuf:"varint,4,opt,name=status,proto3,enum=training.Booking_Status" json:"status,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Booking) GetTotalCost() float32 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *Booking) GetTotalTicket() int32 {
	if x != nil {
		return x.TotalTicket
	}
	return 0
}

func (x *Booking) GetStatus() Booking_Status {
	if x != nil {
		return x.Status
	}
	return Booking_Success
}

func (x *Booking) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Booking) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type BookingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCost   float32             `protobuf:"fixed32,1,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	TotalTicket int32               `protobuf:"varint,2,opt,name=total_ticket,json=totalTicket,proto3" json:"total_ticket,omitempty"`
	Status      BookingInput_Status `protobuf:"varint,3,opt,name=status,proto3,enum=training.BookingInput_Status" json:"status,omitempty"`
}

func (x *BookingInput) Reset() {
	*x = BookingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingInput) ProtoMessage() {}

func (x *BookingInput) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingInput.ProtoReflect.Descriptor instead.
func (*BookingInput) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *BookingInput) GetTotalCost() float32 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *BookingInput) GetTotalTicket() int32 {
	if x != nil {
		return x.TotalTicket
	}
	return 0
}

func (x *BookingInput) GetStatus() BookingInput_Status {
	if x != nil {
		return x.Status
	}
	return BookingInput_Success
}

type GetBookingByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookingByIdRequest) Reset() {
	*x = GetBookingByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookingByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingByIdRequest) ProtoMessage() {}

func (x *GetBookingByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingByIdRequest.ProtoReflect.Descriptor instead.
func (*GetBookingByIdRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookingByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalTicket int32                       `protobuf:"varint,1,opt,name=total_ticket,json=totalTicket,proto3" json:"total_ticket,omitempty"`
	Status      CreateBookingRequest_Status `protobuf:"varint,2,opt,name=status,proto3,enum=training.CreateBookingRequest_Status" json:"status,omitempty"`
}

func (x *CreateBookingRequest) Reset() {
	*x = CreateBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingRequest) ProtoMessage() {}

func (x *CreateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingRequest.ProtoReflect.Descriptor instead.
func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookingRequest) GetTotalTicket() int32 {
	if x != nil {
		return x.TotalTicket
	}
	return 0
}

func (x *CreateBookingRequest) GetStatus() CreateBookingRequest_Status {
	if x != nil {
		return x.Status
	}
	return CreateBookingRequest_Success
}

type UpdateBookingStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status UpdateBookingStatusRequest_Status `protobuf:"varint,2,opt,name=status,proto3,enum=training.UpdateBookingStatusRequest_Status" json:"status,omitempty"`
}

func (x *UpdateBookingStatusRequest) Reset() {
	*x = UpdateBookingStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookingStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookingStatusRequest) ProtoMessage() {}

func (x *UpdateBookingStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookingStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookingStatusRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBookingStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBookingStatusRequest) GetStatus() UpdateBookingStatusRequest_Status {
	if x != nil {
		return x.Status
	}
	return UpdateBookingStatusRequest_Success
}

type DeleteBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookingRequest) Reset() {
	*x = DeleteBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookingRequest) ProtoMessage() {}

func (x *DeleteBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookingRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBookingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDeleted bool `protobuf:"varint,1,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *DeleteBookingResponse) Reset() {
	*x = DeleteBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookingResponse) ProtoMessage() {}

func (x *DeleteBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookingResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookingResponse) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x07, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10,
	0x01, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x22, 0x27, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x99, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10,
	0x01, 0x22, 0xa0, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x65, 0x64, 0x10, 0x02, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x32, 0xbf, 0x02, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x50, 0x43, 0x12, 0x44, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x4e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_booking_proto_goTypes = []interface{}{
	(Booking_Status)(0),                    // 0: training.Booking.Status
	(BookingInput_Status)(0),               // 1: training.BookingInput.Status
	(CreateBookingRequest_Status)(0),       // 2: training.CreateBookingRequest.Status
	(UpdateBookingStatusRequest_Status)(0), // 3: training.UpdateBookingStatusRequest.Status
	(*Booking)(nil),                        // 4: training.Booking
	(*BookingInput)(nil),                   // 5: training.BookingInput
	(*GetBookingByIdRequest)(nil),          // 6: training.GetBookingByIdRequest
	(*CreateBookingRequest)(nil),           // 7: training.CreateBookingRequest
	(*UpdateBookingStatusRequest)(nil),     // 8: training.UpdateBookingStatusRequest
	(*DeleteBookingRequest)(nil),           // 9: training.DeleteBookingRequest
	(*DeleteBookingResponse)(nil),          // 10: training.DeleteBookingResponse
	(*timestamp.Timestamp)(nil),            // 11: google.protobuf.Timestamp
}
var file_booking_proto_depIdxs = []int32{
	0,  // 0: training.Booking.status:type_name -> training.Booking.Status
	11, // 1: training.Booking.created_at:type_name -> google.protobuf.Timestamp
	11, // 2: training.Booking.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 3: training.BookingInput.status:type_name -> training.BookingInput.Status
	2,  // 4: training.CreateBookingRequest.status:type_name -> training.CreateBookingRequest.Status
	3,  // 5: training.UpdateBookingStatusRequest.status:type_name -> training.UpdateBookingStatusRequest.Status
	6,  // 6: training.BookingServiceRPC.GetBookingByID:input_type -> training.GetBookingByIdRequest
	7,  // 7: training.BookingServiceRPC.CreateBooking:input_type -> training.CreateBookingRequest
	8,  // 8: training.BookingServiceRPC.UpdateBookingStatus:input_type -> training.UpdateBookingStatusRequest
	9,  // 9: training.BookingServiceRPC.DeleteBooking:input_type -> training.DeleteBookingRequest
	4,  // 10: training.BookingServiceRPC.GetBookingByID:output_type -> training.Booking
	4,  // 11: training.BookingServiceRPC.CreateBooking:output_type -> training.Booking
	4,  // 12: training.BookingServiceRPC.UpdateBookingStatus:output_type -> training.Booking
	10, // 13: training.BookingServiceRPC.DeleteBooking:output_type -> training.DeleteBookingResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingInput); i {
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
		file_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookingByIdRequest); i {
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
		file_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookingRequest); i {
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
		file_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookingStatusRequest); i {
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
		file_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookingRequest); i {
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
		file_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookingResponse); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		EnumInfos:         file_booking_proto_enumTypes,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
