// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/feed_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for [FeedService.GetFeed][google.ads.googleads.v2.services.FeedService.GetFeed].
type GetFeedRequest struct {
	// The resource name of the feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedRequest) Reset()         { *m = GetFeedRequest{} }
func (m *GetFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedRequest) ProtoMessage()    {}
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86087a3952159448, []int{0}
}

func (m *GetFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedRequest.Unmarshal(m, b)
}
func (m *GetFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedRequest.Merge(m, src)
}
func (m *GetFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedRequest.Size(m)
}
func (m *GetFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedRequest proto.InternalMessageInfo

func (m *GetFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedService.MutateFeeds][google.ads.googleads.v2.services.FeedService.MutateFeeds].
type MutateFeedsRequest struct {
	// The ID of the customer whose feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feeds.
	Operations []*FeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateFeedsRequest) Reset()         { *m = MutateFeedsRequest{} }
func (m *MutateFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsRequest) ProtoMessage()    {}
func (*MutateFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86087a3952159448, []int{1}
}

func (m *MutateFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsRequest.Unmarshal(m, b)
}
func (m *MutateFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsRequest.Merge(m, src)
}
func (m *MutateFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsRequest.Size(m)
}
func (m *MutateFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsRequest proto.InternalMessageInfo

func (m *MutateFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedsRequest) GetOperations() []*FeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an feed.
type FeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedOperation_Create
	//	*FeedOperation_Update
	//	*FeedOperation_Remove
	Operation            isFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *FeedOperation) Reset()         { *m = FeedOperation{} }
func (m *FeedOperation) String() string { return proto.CompactTextString(m) }
func (*FeedOperation) ProtoMessage()    {}
func (*FeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_86087a3952159448, []int{2}
}

func (m *FeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedOperation.Unmarshal(m, b)
}
func (m *FeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedOperation.Marshal(b, m, deterministic)
}
func (m *FeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedOperation.Merge(m, src)
}
func (m *FeedOperation) XXX_Size() int {
	return xxx_messageInfo_FeedOperation.Size(m)
}
func (m *FeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedOperation proto.InternalMessageInfo

func (m *FeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isFeedOperation_Operation interface {
	isFeedOperation_Operation()
}

type FeedOperation_Create struct {
	Create *resources.Feed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedOperation_Update struct {
	Update *resources.Feed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type FeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedOperation_Create) isFeedOperation_Operation() {}

func (*FeedOperation_Update) isFeedOperation_Operation() {}

func (*FeedOperation_Remove) isFeedOperation_Operation() {}

func (m *FeedOperation) GetOperation() isFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedOperation) GetCreate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedOperation) GetUpdate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *FeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedOperation_Create)(nil),
		(*FeedOperation_Update)(nil),
		(*FeedOperation_Remove)(nil),
	}
}

// Response message for an feed mutate.
type MutateFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MutateFeedsResponse) Reset()         { *m = MutateFeedsResponse{} }
func (m *MutateFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsResponse) ProtoMessage()    {}
func (*MutateFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86087a3952159448, []int{3}
}

func (m *MutateFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsResponse.Unmarshal(m, b)
}
func (m *MutateFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsResponse.Merge(m, src)
}
func (m *MutateFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsResponse.Size(m)
}
func (m *MutateFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsResponse proto.InternalMessageInfo

func (m *MutateFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedsResponse) GetResults() []*MutateFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mutate.
type MutateFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedResult) Reset()         { *m = MutateFeedResult{} }
func (m *MutateFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedResult) ProtoMessage()    {}
func (*MutateFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_86087a3952159448, []int{4}
}

func (m *MutateFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedResult.Unmarshal(m, b)
}
func (m *MutateFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedResult.Merge(m, src)
}
func (m *MutateFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedResult.Size(m)
}
func (m *MutateFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedResult proto.InternalMessageInfo

func (m *MutateFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedRequest)(nil), "google.ads.googleads.v2.services.GetFeedRequest")
	proto.RegisterType((*MutateFeedsRequest)(nil), "google.ads.googleads.v2.services.MutateFeedsRequest")
	proto.RegisterType((*FeedOperation)(nil), "google.ads.googleads.v2.services.FeedOperation")
	proto.RegisterType((*MutateFeedsResponse)(nil), "google.ads.googleads.v2.services.MutateFeedsResponse")
	proto.RegisterType((*MutateFeedResult)(nil), "google.ads.googleads.v2.services.MutateFeedResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/feed_service.proto", fileDescriptor_86087a3952159448)
}

var fileDescriptor_86087a3952159448 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x36, 0x5b, 0x69, 0xed, 0xa4, 0xad, 0x65, 0x8a, 0x18, 0x56, 0xc1, 0x25, 0x16, 0x5a, 0x17,
	0xc9, 0x48, 0x6a, 0x11, 0x52, 0x7a, 0x91, 0x82, 0xdb, 0x0a, 0xd6, 0x96, 0x14, 0x7a, 0x21, 0x0b,
	0x61, 0x9a, 0xcc, 0x2e, 0xa1, 0x49, 0x26, 0xce, 0x4c, 0x16, 0x4a, 0xe9, 0x8d, 0xaf, 0x20, 0xbe,
	0x80, 0x97, 0x7a, 0xe5, 0x6b, 0xf4, 0x56, 0x1f, 0xc1, 0x2b, 0x1f, 0x40, 0xf1, 0x4e, 0x26, 0x93,
	0xd9, 0x9f, 0x42, 0xd9, 0xf6, 0xee, 0xe4, 0xcc, 0xf7, 0x7d, 0xe7, 0x9b, 0x73, 0xe6, 0x04, 0x6c,
	0xf4, 0x29, 0xed, 0xa7, 0x04, 0xe1, 0x98, 0x23, 0x15, 0xca, 0x68, 0xe0, 0x22, 0x4e, 0xd8, 0x20,
	0x89, 0x08, 0x47, 0x3d, 0x42, 0xe2, 0xb0, 0xfe, 0x72, 0x0a, 0x46, 0x05, 0x85, 0x2d, 0x85, 0x74,
	0x70, 0xcc, 0x9d, 0x21, 0xc9, 0x19, 0xb8, 0x8e, 0x26, 0x35, 0x9f, 0x5f, 0x27, 0xcb, 0x08, 0xa7,
	0x25, 0xd3, 0xba, 0x4a, 0xaf, 0xf9, 0x58, 0xa3, 0x8b, 0x04, 0xe1, 0x3c, 0xa7, 0x02, 0x8b, 0x84,
	0xe6, 0xbc, 0x3e, 0xad, 0xab, 0xa1, 0xea, 0xeb, 0xa4, 0xec, 0xa1, 0x5e, 0x42, 0xd2, 0x38, 0xcc,
	0x30, 0x3f, 0xad, 0x11, 0x0f, 0x6b, 0x04, 0x2b, 0x22, 0xc4, 0x05, 0x16, 0x25, 0xbf, 0x72, 0x20,
	0x85, 0xa3, 0x34, 0x21, 0xb9, 0x50, 0x07, 0xf6, 0x26, 0x58, 0xda, 0x25, 0xa2, 0x43, 0x48, 0x1c,
	0x90, 0x0f, 0x25, 0xe1, 0x02, 0x3e, 0x05, 0x8b, 0xda, 0x5b, 0x98, 0xe3, 0x8c, 0x58, 0x46, 0xcb,
	0x58, 0x9f, 0x0f, 0x16, 0x74, 0xf2, 0x1d, 0xce, 0x88, 0xfd, 0xd3, 0x00, 0x70, 0xbf, 0x14, 0x58,
	0x10, 0x49, 0xe5, 0x9a, 0xfb, 0x04, 0x98, 0x51, 0xc9, 0x05, 0xcd, 0x08, 0x0b, 0x93, 0xb8, 0x66,
	0x02, 0x9d, 0x7a, 0x13, 0xc3, 0x03, 0x00, 0x68, 0x41, 0x98, 0xba, 0x96, 0xd5, 0x68, 0xcd, 0xac,
	0x9b, 0x2e, 0x72, 0xa6, 0x75, 0xd1, 0x91, 0x45, 0x0e, 0x34, 0x2f, 0x18, 0x93, 0x80, 0x6b, 0xe0,
	0x7e, 0x81, 0x99, 0x48, 0x70, 0x1a, 0xf6, 0x70, 0x92, 0x96, 0x8c, 0x58, 0x33, 0x2d, 0x63, 0xfd,
	0x5e, 0xb0, 0x54, 0xa7, 0x3b, 0x2a, 0x2b, 0xaf, 0x35, 0xc0, 0x69, 0x12, 0x63, 0x41, 0x42, 0x9a,
	0xa7, 0x67, 0xd6, 0xdd, 0x0a, 0xb6, 0xa0, 0x93, 0x07, 0x79, 0x7a, 0x66, 0xff, 0x33, 0xc0, 0xe2,
	0x44, 0x2d, 0xb8, 0x05, 0xcc, 0xb2, 0xa8, 0x48, 0xb2, 0xcd, 0x15, 0xc9, 0x74, 0x9b, 0xda, 0xb1,
	0x9e, 0x84, 0xd3, 0x91, 0x93, 0xd8, 0xc7, 0xfc, 0x34, 0x00, 0x0a, 0x2e, 0x63, 0xe8, 0x83, 0xd9,
	0x88, 0x11, 0x2c, 0x54, 0x0f, 0x4d, 0x77, 0xed, 0xda, 0x9b, 0x0e, 0x5f, 0x43, 0x75, 0xd5, 0xbd,
	0x3b, 0x41, 0x4d, 0x94, 0x12, 0x4a, 0xd0, 0x6a, 0xdc, 0x5a, 0x42, 0x11, 0xa1, 0x05, 0x66, 0x19,
	0xc9, 0xe8, 0x40, 0x75, 0x66, 0x5e, 0x9e, 0xa8, 0xef, 0x1d, 0x13, 0xcc, 0x0f, 0x5b, 0x69, 0x7f,
	0x33, 0xc0, 0xca, 0xc4, 0x48, 0x79, 0x41, 0x73, 0x4e, 0x60, 0x07, 0x3c, 0xb8, 0xd2, 0xe1, 0x90,
	0x30, 0x46, 0x59, 0xa5, 0x66, 0xba, 0x50, 0x1b, 0x62, 0x45, 0xe4, 0x1c, 0x55, 0x6f, 0x2e, 0x58,
	0x99, 0xec, 0xfd, 0x6b, 0x09, 0x87, 0x6f, 0xc1, 0x1c, 0x23, 0xbc, 0x4c, 0x85, 0x9e, 0xbb, 0x3b,
	0x7d, 0xee, 0x23, 0x3f, 0x41, 0x45, 0x0d, 0xb4, 0x84, 0xfd, 0x0a, 0x2c, 0x5f, 0x3d, 0xbc, 0xd1,
	0xcb, 0x75, 0xff, 0x34, 0x80, 0x29, 0x39, 0x47, 0xaa, 0x06, 0xfc, 0x6c, 0x80, 0xb9, 0x7a, 0x03,
	0xe0, 0x8b, 0xe9, 0x8e, 0x26, 0x97, 0xa5, 0x79, 0xd3, 0x71, 0xd8, 0xe8, 0xe3, 0x8f, 0x5f, 0x9f,
	0x1a, 0xcf, 0xe0, 0x9a, 0xdc, 0xfd, 0xf3, 0x09, 0x9b, 0xdb, 0x7a, 0x3f, 0x38, 0x6a, 0x57, 0x3f,
	0x03, 0x8e, 0xda, 0x17, 0xf0, 0xbb, 0x01, 0xcc, 0xb1, 0x71, 0xc0, 0x97, 0xb7, 0xe9, 0x96, 0x5e,
	0xc8, 0xe6, 0xe6, 0x2d, 0x59, 0x6a, 0xe6, 0xf6, 0x66, 0xe5, 0x16, 0xd9, 0x6d, 0xe9, 0x76, 0x64,
	0xef, 0x7c, 0x6c, 0xb9, 0xb7, 0xdb, 0x17, 0xca, 0xac, 0x97, 0x55, 0x02, 0x9e, 0xd1, 0x6e, 0x3e,
	0xba, 0xf4, 0xad, 0x51, 0x91, 0x3a, 0x2a, 0x12, 0xee, 0x44, 0x34, 0xdb, 0xf9, 0x6b, 0x80, 0xd5,
	0x88, 0x66, 0x53, 0x0d, 0xed, 0x2c, 0x8f, 0x8d, 0xe7, 0x50, 0x2e, 0xd8, 0xa1, 0xf1, 0x7e, 0xaf,
	0x66, 0xf5, 0x69, 0x8a, 0xf3, 0xbe, 0x43, 0x59, 0x1f, 0xf5, 0x49, 0x5e, 0xad, 0x1f, 0x1a, 0xd5,
	0xb9, 0xfe, 0xe7, 0xbd, 0xa5, 0x83, 0x2f, 0x8d, 0x99, 0x5d, 0xdf, 0xff, 0xda, 0x68, 0xed, 0x2a,
	0x41, 0x3f, 0xe6, 0x8e, 0x0a, 0x65, 0x74, 0xec, 0x3a, 0x75, 0x61, 0x7e, 0xa9, 0x21, 0x5d, 0x3f,
	0xe6, 0xdd, 0x21, 0xa4, 0x7b, 0xec, 0x76, 0x35, 0xe4, 0x77, 0x63, 0x55, 0xe5, 0x3d, 0xcf, 0x8f,
	0xb9, 0xe7, 0x0d, 0x41, 0x9e, 0x77, 0xec, 0x7a, 0x9e, 0x86, 0x9d, 0xcc, 0x56, 0x3e, 0x37, 0xfe,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x49, 0x64, 0x65, 0xd5, 0x63, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedServiceClient interface {
	// Returns the requested feed in full detail.
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error)
}

type feedServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedServiceClient(cc *grpc.ClientConn) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error) {
	out := new(resources.Feed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.FeedService/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error) {
	out := new(MutateFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.FeedService/MutateFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
type FeedServiceServer interface {
	// Returns the requested feed in full detail.
	GetFeed(context.Context, *GetFeedRequest) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(context.Context, *MutateFeedsRequest) (*MutateFeedsResponse, error)
}

func RegisterFeedServiceServer(s *grpc.Server, srv FeedServiceServer) {
	s.RegisterService(&_FeedService_serviceDesc, srv)
}

func _FeedService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.FeedService/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_MutateFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).MutateFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.FeedService/MutateFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).MutateFeeds(ctx, req.(*MutateFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeed",
			Handler:    _FeedService_GetFeed_Handler,
		},
		{
			MethodName: "MutateFeeds",
			Handler:    _FeedService_MutateFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/feed_service.proto",
}
