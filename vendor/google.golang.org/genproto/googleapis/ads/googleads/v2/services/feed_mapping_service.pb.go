// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/feed_mapping_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [FeedMappingService.GetFeedMapping][google.ads.googleads.v2.services.FeedMappingService.GetFeedMapping].
type GetFeedMappingRequest struct {
	// The resource name of the feed mapping to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedMappingRequest) Reset()         { *m = GetFeedMappingRequest{} }
func (m *GetFeedMappingRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedMappingRequest) ProtoMessage()    {}
func (*GetFeedMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{0}
}

func (m *GetFeedMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedMappingRequest.Unmarshal(m, b)
}
func (m *GetFeedMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedMappingRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedMappingRequest.Merge(m, src)
}
func (m *GetFeedMappingRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedMappingRequest.Size(m)
}
func (m *GetFeedMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedMappingRequest proto.InternalMessageInfo

func (m *GetFeedMappingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedMappingService.MutateFeedMappings][google.ads.googleads.v2.services.FeedMappingService.MutateFeedMappings].
type MutateFeedMappingsRequest struct {
	// The ID of the customer whose feed mappings are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feed mappings.
	Operations []*FeedMappingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedMappingsRequest) Reset()         { *m = MutateFeedMappingsRequest{} }
func (m *MutateFeedMappingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingsRequest) ProtoMessage()    {}
func (*MutateFeedMappingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{1}
}

func (m *MutateFeedMappingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingsRequest.Unmarshal(m, b)
}
func (m *MutateFeedMappingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedMappingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingsRequest.Merge(m, src)
}
func (m *MutateFeedMappingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingsRequest.Size(m)
}
func (m *MutateFeedMappingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingsRequest proto.InternalMessageInfo

func (m *MutateFeedMappingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedMappingsRequest) GetOperations() []*FeedMappingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedMappingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedMappingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a feed mapping.
type FeedMappingOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedMappingOperation_Create
	//	*FeedMappingOperation_Remove
	Operation            isFeedMappingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *FeedMappingOperation) Reset()         { *m = FeedMappingOperation{} }
func (m *FeedMappingOperation) String() string { return proto.CompactTextString(m) }
func (*FeedMappingOperation) ProtoMessage()    {}
func (*FeedMappingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{2}
}

func (m *FeedMappingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingOperation.Unmarshal(m, b)
}
func (m *FeedMappingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingOperation.Marshal(b, m, deterministic)
}
func (m *FeedMappingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingOperation.Merge(m, src)
}
func (m *FeedMappingOperation) XXX_Size() int {
	return xxx_messageInfo_FeedMappingOperation.Size(m)
}
func (m *FeedMappingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingOperation proto.InternalMessageInfo

type isFeedMappingOperation_Operation interface {
	isFeedMappingOperation_Operation()
}

type FeedMappingOperation_Create struct {
	Create *resources.FeedMapping `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedMappingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedMappingOperation_Create) isFeedMappingOperation_Operation() {}

func (*FeedMappingOperation_Remove) isFeedMappingOperation_Operation() {}

func (m *FeedMappingOperation) GetOperation() isFeedMappingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedMappingOperation) GetCreate() *resources.FeedMapping {
	if x, ok := m.GetOperation().(*FeedMappingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedMappingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedMappingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedMappingOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedMappingOperation_Create)(nil),
		(*FeedMappingOperation_Remove)(nil),
	}
}

// Response message for a feed mapping mutate.
type MutateFeedMappingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedMappingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MutateFeedMappingsResponse) Reset()         { *m = MutateFeedMappingsResponse{} }
func (m *MutateFeedMappingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingsResponse) ProtoMessage()    {}
func (*MutateFeedMappingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{3}
}

func (m *MutateFeedMappingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingsResponse.Unmarshal(m, b)
}
func (m *MutateFeedMappingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedMappingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingsResponse.Merge(m, src)
}
func (m *MutateFeedMappingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingsResponse.Size(m)
}
func (m *MutateFeedMappingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingsResponse proto.InternalMessageInfo

func (m *MutateFeedMappingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedMappingsResponse) GetResults() []*MutateFeedMappingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mapping mutate.
type MutateFeedMappingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedMappingResult) Reset()         { *m = MutateFeedMappingResult{} }
func (m *MutateFeedMappingResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingResult) ProtoMessage()    {}
func (*MutateFeedMappingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{4}
}

func (m *MutateFeedMappingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingResult.Unmarshal(m, b)
}
func (m *MutateFeedMappingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedMappingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingResult.Merge(m, src)
}
func (m *MutateFeedMappingResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingResult.Size(m)
}
func (m *MutateFeedMappingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingResult proto.InternalMessageInfo

func (m *MutateFeedMappingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedMappingRequest)(nil), "google.ads.googleads.v2.services.GetFeedMappingRequest")
	proto.RegisterType((*MutateFeedMappingsRequest)(nil), "google.ads.googleads.v2.services.MutateFeedMappingsRequest")
	proto.RegisterType((*FeedMappingOperation)(nil), "google.ads.googleads.v2.services.FeedMappingOperation")
	proto.RegisterType((*MutateFeedMappingsResponse)(nil), "google.ads.googleads.v2.services.MutateFeedMappingsResponse")
	proto.RegisterType((*MutateFeedMappingResult)(nil), "google.ads.googleads.v2.services.MutateFeedMappingResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/feed_mapping_service.proto", fileDescriptor_378f0144add6067b)
}

var fileDescriptor_378f0144add6067b = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x37, 0xbb, 0x52, 0xed, 0x6c, 0xad, 0x30, 0x5a, 0xba, 0xae, 0x82, 0x4b, 0x2c, 0x58, 0xf6,
	0x30, 0xd1, 0x28, 0x15, 0xd3, 0x56, 0xd8, 0x82, 0x6d, 0x3d, 0xd4, 0x96, 0x14, 0xf6, 0x20, 0x0b,
	0x61, 0x4c, 0x5e, 0x43, 0x20, 0xc9, 0xc4, 0x99, 0xc9, 0x42, 0x29, 0xbd, 0x78, 0xf1, 0x03, 0xf8,
	0x0d, 0x3c, 0x7a, 0xf3, 0xe0, 0xc1, 0xaf, 0x50, 0xf0, 0xe4, 0x37, 0x10, 0x4f, 0x7e, 0x08, 0x91,
	0x64, 0x32, 0xdb, 0xdd, 0xb6, 0xcb, 0x6a, 0x6f, 0x2f, 0xef, 0xbd, 0xdf, 0xef, 0xfd, 0xde, 0x9f,
	0x09, 0x5a, 0x0d, 0x19, 0x0b, 0x63, 0xb0, 0x68, 0x20, 0x2c, 0x65, 0x16, 0xd6, 0xc0, 0xb6, 0x04,
	0xf0, 0x41, 0xe4, 0x83, 0xb0, 0x0e, 0x00, 0x02, 0x2f, 0xa1, 0x59, 0x16, 0xa5, 0xa1, 0x57, 0x79,
	0x49, 0xc6, 0x99, 0x64, 0xb8, 0xad, 0x10, 0x84, 0x06, 0x82, 0x0c, 0xc1, 0x64, 0x60, 0x13, 0x0d,
	0x6e, 0x3d, 0x9d, 0x44, 0xcf, 0x41, 0xb0, 0x9c, 0x9f, 0xe5, 0x57, 0xbc, 0xad, 0x7b, 0x1a, 0x95,
	0x45, 0x16, 0x4d, 0x53, 0x26, 0xa9, 0x8c, 0x58, 0x2a, 0xaa, 0xe8, 0x62, 0x15, 0xe5, 0x99, 0x6f,
	0x09, 0x49, 0x65, 0x7e, 0x36, 0x50, 0xc0, 0xfc, 0x38, 0x82, 0x54, 0xaa, 0x80, 0xb9, 0x86, 0x16,
	0xb6, 0x40, 0x6e, 0x02, 0x04, 0x3b, 0xaa, 0x8e, 0x0b, 0xef, 0x72, 0x10, 0x12, 0x3f, 0x40, 0x37,
	0xb4, 0x10, 0x2f, 0xa5, 0x09, 0x34, 0x8d, 0xb6, 0xb1, 0x3c, 0xeb, 0xce, 0x69, 0xe7, 0x6b, 0x9a,
	0x80, 0xf9, 0xd3, 0x40, 0x77, 0x76, 0x72, 0x49, 0x25, 0x8c, 0x30, 0x08, 0x4d, 0x71, 0x1f, 0x35,
	0xfc, 0x5c, 0x48, 0x96, 0x00, 0xf7, 0xa2, 0xa0, 0x22, 0x40, 0xda, 0xf5, 0x2a, 0xc0, 0x3d, 0x84,
	0x58, 0x06, 0x5c, 0xb5, 0xd0, 0xac, 0xb5, 0xeb, 0xcb, 0x0d, 0x7b, 0x85, 0x4c, 0x9b, 0x1c, 0x19,
	0xa9, 0xb5, 0xab, 0xe1, 0xee, 0x08, 0x13, 0x7e, 0x88, 0x6e, 0x66, 0x94, 0xcb, 0x88, 0xc6, 0xde,
	0x01, 0x8d, 0xe2, 0x9c, 0x43, 0xb3, 0xde, 0x36, 0x96, 0xaf, 0xbb, 0xf3, 0x95, 0x7b, 0x53, 0x79,
	0x8b, 0x26, 0x07, 0x34, 0x8e, 0x02, 0x2a, 0xc1, 0x63, 0x69, 0x7c, 0xd8, 0xbc, 0x5a, 0xa6, 0xcd,
	0x69, 0xe7, 0x6e, 0x1a, 0x1f, 0x9a, 0x1f, 0x0c, 0x74, 0xfb, 0xa2, 0x92, 0x78, 0x1b, 0xcd, 0xf8,
	0x1c, 0xa8, 0x54, 0xb3, 0x69, 0xd8, 0x64, 0xa2, 0xf4, 0xe1, 0x4a, 0x47, 0xb5, 0x6f, 0x5f, 0x71,
	0x2b, 0x3c, 0x6e, 0xa2, 0x19, 0x0e, 0x09, 0x1b, 0x28, 0x9d, 0xb3, 0x45, 0x44, 0x7d, 0x6f, 0x34,
	0xd0, 0xec, 0xb0, 0x31, 0xf3, 0x9b, 0x81, 0x5a, 0x17, 0x8d, 0x5b, 0x64, 0x2c, 0x15, 0x80, 0x37,
	0xd1, 0xc2, 0x99, 0xb6, 0x3d, 0xe0, 0x9c, 0xf1, 0x92, 0xb4, 0x61, 0x63, 0x2d, 0x8f, 0x67, 0x3e,
	0xd9, 0x2f, 0xaf, 0xc3, 0xbd, 0x35, 0x3e, 0x90, 0x97, 0x45, 0x3a, 0xde, 0x47, 0xd7, 0x38, 0x88,
	0x3c, 0x96, 0x7a, 0x27, 0xcf, 0xa7, 0xef, 0xe4, 0x9c, 0x2c, 0xb7, 0x64, 0x70, 0x35, 0x93, 0xf9,
	0x02, 0x2d, 0x4e, 0xc8, 0xf9, 0xa7, 0x53, 0xb3, 0xbf, 0xd6, 0x11, 0x1e, 0x81, 0xee, 0xab, 0xc2,
	0xf8, 0x8b, 0x81, 0xe6, 0xc7, 0x0f, 0x18, 0x3f, 0x9b, 0xae, 0xf6, 0xc2, 0x93, 0x6f, 0xfd, 0xe7,
	0xfe, 0xcc, 0x95, 0xf7, 0x3f, 0x7e, 0x7d, 0xac, 0x3d, 0xc2, 0xa4, 0x78, 0xb5, 0x47, 0x63, 0x2d,
	0xac, 0xeb, 0x2b, 0x17, 0x56, 0xa7, 0x7c, 0xc6, 0x7a, 0x59, 0x56, 0xe7, 0x18, 0x7f, 0x37, 0x10,
	0x3e, 0xbf, 0x46, 0xbc, 0x7a, 0x89, 0x29, 0xeb, 0xb7, 0xd6, 0x5a, 0xbb, 0x1c, 0x58, 0x5d, 0x8e,
	0xb9, 0x56, 0x76, 0xb2, 0x62, 0x3e, 0x2e, 0x3a, 0x39, 0x95, 0x7e, 0x34, 0xf2, 0x7c, 0xd7, 0x3b,
	0xc7, 0x63, 0x8d, 0x38, 0x49, 0x49, 0xe7, 0x18, 0x9d, 0xd6, 0xdd, 0x93, 0x6e, 0xf3, 0xb4, 0x64,
	0x65, 0x65, 0x91, 0x20, 0x3e, 0x4b, 0x36, 0xfe, 0x18, 0x68, 0xc9, 0x67, 0xc9, 0x54, 0x79, 0x1b,
	0x8b, 0xe7, 0xb7, 0xbb, 0x57, 0xfc, 0xa2, 0xf6, 0x8c, 0x37, 0xdb, 0x15, 0x38, 0x64, 0x31, 0x4d,
	0x43, 0xc2, 0x78, 0x68, 0x85, 0x90, 0x96, 0x3f, 0x30, 0xeb, 0xb4, 0xdc, 0xe4, 0x1f, 0xf5, 0xaa,
	0x36, 0x3e, 0xd5, 0xea, 0x5b, 0xdd, 0xee, 0xe7, 0x5a, 0x7b, 0x4b, 0x11, 0x76, 0x03, 0x41, 0x94,
	0x59, 0x58, 0x3d, 0x9b, 0x54, 0x85, 0xc5, 0x89, 0x4e, 0xe9, 0x77, 0x03, 0xd1, 0x1f, 0xa6, 0xf4,
	0x7b, 0x76, 0x5f, 0xa7, 0xfc, 0xae, 0x2d, 0x29, 0xbf, 0xe3, 0x74, 0x03, 0xe1, 0x38, 0xc3, 0x24,
	0xc7, 0xe9, 0xd9, 0x8e, 0xa3, 0xd3, 0xde, 0xce, 0x94, 0x3a, 0x9f, 0xfc, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x24, 0x3b, 0x10, 0x27, 0x4f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedMappingServiceClient is the client API for FeedMappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedMappingServiceClient interface {
	// Returns the requested feed mapping in full detail.
	GetFeedMapping(ctx context.Context, in *GetFeedMappingRequest, opts ...grpc.CallOption) (*resources.FeedMapping, error)
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error)
}

type feedMappingServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedMappingServiceClient(cc *grpc.ClientConn) FeedMappingServiceClient {
	return &feedMappingServiceClient{cc}
}

func (c *feedMappingServiceClient) GetFeedMapping(ctx context.Context, in *GetFeedMappingRequest, opts ...grpc.CallOption) (*resources.FeedMapping, error) {
	out := new(resources.FeedMapping)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.FeedMappingService/GetFeedMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedMappingServiceClient) MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error) {
	out := new(MutateFeedMappingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.FeedMappingService/MutateFeedMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedMappingServiceServer is the server API for FeedMappingService service.
type FeedMappingServiceServer interface {
	// Returns the requested feed mapping in full detail.
	GetFeedMapping(context.Context, *GetFeedMappingRequest) (*resources.FeedMapping, error)
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	MutateFeedMappings(context.Context, *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error)
}

func RegisterFeedMappingServiceServer(s *grpc.Server, srv FeedMappingServiceServer) {
	s.RegisterService(&_FeedMappingService_serviceDesc, srv)
}

func _FeedMappingService_GetFeedMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).GetFeedMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.FeedMappingService/GetFeedMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).GetFeedMapping(ctx, req.(*GetFeedMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedMappingService_MutateFeedMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.FeedMappingService/MutateFeedMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, req.(*MutateFeedMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedMappingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.FeedMappingService",
	HandlerType: (*FeedMappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedMapping",
			Handler:    _FeedMappingService_GetFeedMapping_Handler,
		},
		{
			MethodName: "MutateFeedMappings",
			Handler:    _FeedMappingService_MutateFeedMappings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/feed_mapping_service.proto",
}
