package proto

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

type BookingCancellationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId int32 `protobuf:"varint,1,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
}

func (x *BookingCancellationRequest) Reset() {
	*x = BookingCancellationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingCancellationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingCancellationRequest) ProtoMessage() {}

func (x *BookingCancellationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingCancellationRequest.ProtoReflect.Descriptor instead.
func (*BookingCancellationRequest) Descriptor() ([]byte, []int) {
	return file_flight_service_proto_rawDescGZIP(), []int{0}
}

func (x *BookingCancellationRequest) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type BookingCancellationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
}

func (x *BookingCancellationResponse) Reset() {
	*x = BookingCancellationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingCancellationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingCancellationResponse) ProtoMessage() {}

func (x *BookingCancellationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingCancellationResponse.ProtoReflect.Descriptor instead.
func (*BookingCancellationResponse) Descriptor() ([]byte, []int) {
	return file_flight_service_proto_rawDescGZIP(), []int{1}
}

func (x *BookingCancellationResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type ChangeBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId    int32 `protobuf:"varint,1,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	FlightNumber int32 `protobuf:"varint,2,opt,name=flightNumber,proto3" json:"flightNumber,omitempty"`
}

func (x *ChangeBookingRequest) Reset() {
	*x = ChangeBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeBookingRequest) ProtoMessage() {}

func (x *ChangeBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeBookingRequest.ProtoReflect.Descriptor instead.
func (*ChangeBookingRequest) Descriptor() ([]byte, []int) {
	return file_flight_service_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeBookingRequest) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *ChangeBookingRequest) GetFlightNumber() int32 {
	if x != nil {
		return x.FlightNumber
	}
	return 0
}

type ChangeBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
}

func (x *ChangeBookingResponse) Reset() {
	*x = ChangeBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeBookingResponse) ProtoMessage() {}

func (x *ChangeBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeBookingResponse.ProtoReflect.Descriptor instead.
func (*ChangeBookingResponse) Descriptor() ([]byte, []int) {
	return file_flight_service_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeBookingResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age            int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	DocumentNumber string `protobuf:"bytes,3,opt,name=documentNumber,proto3" json:"documentNumber,omitempty"`
	Country        string `protobuf:"bytes,4,opt,name=Country,proto3" json:"Country,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_flight_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_flight_service_proto_rawDescGZIP(), []int{4}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetDocumentNumber() string {
	if x != nil {
		return x.DocumentNumber
	}
	return ""
}

func (x *Person) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type BookTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightNumber int32   `protobuf:"varint,1,opt,name=flightNumber,proto3" json:"flightNumber,omitempty"`
	Person       *Person `protobuf:"bytes,2,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *BookTicketRequest) Reset() {
	*x = BookTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookTicketRequest) ProtoMessage() {}

func (x *BookTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookTicketRequest.ProtoReflect.Descriptor instead.
func (*BookTicketRequest) Descriptor() ([]byte, []int) {
	return file_flight_service_proto_rawDescGZIP(), []int{5}
}

func (x *BookTicketRequest) GetFlightNumber() int32 {
	if x != nil {
		return x.FlightNumber
	}
	return 0
}

func (x *BookTicketRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type BookTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool  `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	BookingId int32 `protobuf:"varint,2,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
}

func (x *BookTicketResponse) Reset() {
	*x = BookTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookTicketResponse) ProtoMessage() {}

func (x *BookTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookTicketResponse.ProtoReflect.Descriptor instead.
func (*BookTicketResponse) Descriptor() ([]byte, []int) {
	return file_flight_service_proto_rawDescGZIP(), []int{6}
}

func (x *BookTicketResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *BookTicketResponse) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

var File_flight_service_proto protoreflect.FileDescriptor

var file_flight_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a,
	0x1a, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x1b, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x58, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x35, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x70, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x5e, 0x0a, 0x11, 0x42, 0x6f, 0x6f,
	0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x12, 0x42, 0x6f, 0x6f,
	0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x32, 0xf6, 0x01, 0x0a, 0x0d,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a,
	0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_flight_service_proto_rawDescOnce sync.Once
	file_flight_service_proto_rawDescData = file_flight_service_proto_rawDesc
)

func file_flight_service_proto_rawDescGZIP() []byte {
	file_flight_service_proto_rawDescOnce.Do(func() {
		file_flight_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_service_proto_rawDescData)
	})
	return file_flight_service_proto_rawDescData
}

var file_flight_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_flight_service_proto_goTypes = []interface{}{
	(*BookingCancellationRequest)(nil),  // 0: proto.BookingCancellationRequest
	(*BookingCancellationResponse)(nil), // 1: proto.BookingCancellationResponse
	(*ChangeBookingRequest)(nil),        // 2: proto.ChangeBookingRequest
	(*ChangeBookingResponse)(nil),       // 3: proto.ChangeBookingResponse
	(*Person)(nil),                      // 4: proto.Person
	(*BookTicketRequest)(nil),           // 5: proto.BookTicketRequest
	(*BookTicketResponse)(nil),          // 6: proto.BookTicketResponse
}
var file_flight_service_proto_depIdxs = []int32{
	4, // 0: proto.BookTicketRequest.person:type_name -> proto.Person
	0, // 1: proto.FlightService.CancelBooking:input_type -> proto.BookingCancellationRequest
	5, // 2: proto.FlightService.BookTicket:input_type -> proto.BookTicketRequest
	2, // 3: proto.FlightService.ChangeBooking:input_type -> proto.ChangeBookingRequest
	1, // 4: proto.FlightService.CancelBooking:output_type -> proto.BookingCancellationResponse
	6, // 5: proto.FlightService.BookTicket:output_type -> proto.BookTicketResponse
	3, // 6: proto.FlightService.ChangeBooking:output_type -> proto.ChangeBookingResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flight_service_proto_init() }
func file_flight_service_proto_init() {
	if File_flight_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flight_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingCancellationRequest); i {
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
		file_flight_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingCancellationResponse); i {
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
		file_flight_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeBookingRequest); i {
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
		file_flight_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeBookingResponse); i {
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
		file_flight_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_flight_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookTicketRequest); i {
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
		file_flight_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookTicketResponse); i {
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
			RawDescriptor: file_flight_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_service_proto_goTypes,
		DependencyIndexes: file_flight_service_proto_depIdxs,
		MessageInfos:      file_flight_service_proto_msgTypes,
	}.Build()
	File_flight_service_proto = out.File
	file_flight_service_proto_rawDesc = nil
	file_flight_service_proto_goTypes = nil
	file_flight_service_proto_depIdxs = nil
}
