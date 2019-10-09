// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/gender_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [GenderViewService.GetGenderView][google.ads.googleads.v2.services.GenderViewService.GetGenderView].
type GetGenderViewRequest struct {
	// The resource name of the gender view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGenderViewRequest) Reset()         { *m = GetGenderViewRequest{} }
func (m *GetGenderViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGenderViewRequest) ProtoMessage()    {}
func (*GetGenderViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9903af7aa827f0d3, []int{0}
}

func (m *GetGenderViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGenderViewRequest.Unmarshal(m, b)
}
func (m *GetGenderViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGenderViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGenderViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGenderViewRequest.Merge(m, src)
}
func (m *GetGenderViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGenderViewRequest.Size(m)
}
func (m *GetGenderViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGenderViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGenderViewRequest proto.InternalMessageInfo

func (m *GetGenderViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGenderViewRequest)(nil), "google.ads.googleads.v2.services.GetGenderViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/gender_view_service.proto", fileDescriptor_9903af7aa827f0d3)
}

var fileDescriptor_9903af7aa827f0d3 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x25, 0x79, 0xf0, 0xe0, 0x85, 0xd7, 0x85, 0x41, 0xb4, 0x44, 0x17, 0xa5, 0x76, 0x21, 0x85,
	0xce, 0x40, 0x8a, 0x2e, 0xa6, 0xb8, 0x48, 0x37, 0x71, 0x25, 0xa5, 0x42, 0x16, 0x12, 0x28, 0x31,
	0xb9, 0x84, 0x40, 0x93, 0xa9, 0xb9, 0x69, 0xba, 0x10, 0x37, 0xfe, 0x82, 0x7f, 0xe0, 0x52, 0xf0,
	0x47, 0xba, 0xf5, 0x07, 0x5c, 0xb8, 0xf2, 0x1f, 0x04, 0x49, 0x27, 0x93, 0x58, 0xb5, 0x74, 0x77,
	0xb8, 0xf7, 0x9c, 0x73, 0xcf, 0x1c, 0x46, 0x63, 0x21, 0xe7, 0xe1, 0x14, 0xa8, 0x17, 0x20, 0x15,
	0xb0, 0x40, 0xb9, 0x49, 0x11, 0xd2, 0x3c, 0xf2, 0x01, 0x69, 0x08, 0x49, 0x00, 0xe9, 0x24, 0x8f,
	0x60, 0x31, 0x29, 0x87, 0x64, 0x96, 0xf2, 0x8c, 0xeb, 0x2d, 0x21, 0x20, 0x5e, 0x80, 0xa4, 0xd2,
	0x92, 0xdc, 0x24, 0x52, 0x6b, 0xf4, 0x37, 0xb9, 0xa7, 0x80, 0x7c, 0x9e, 0x7e, 0xb3, 0x17, 0xb6,
	0xc6, 0xa1, 0x14, 0xcd, 0x22, 0xea, 0x25, 0x09, 0xcf, 0xbc, 0x2c, 0xe2, 0x09, 0x96, 0xdb, 0xfd,
	0x2f, 0x5b, 0x7f, 0x1a, 0x41, 0x92, 0x89, 0x45, 0x7b, 0xa0, 0xed, 0xda, 0x90, 0xd9, 0x2b, 0x3b,
	0x27, 0x82, 0xc5, 0x18, 0x6e, 0xe6, 0x80, 0x99, 0x7e, 0xa4, 0x35, 0xe4, 0xb5, 0x49, 0xe2, 0xc5,
	0xd0, 0x54, 0x5a, 0xca, 0xf1, 0xbf, 0xf1, 0x7f, 0x39, 0xbc, 0xf0, 0x62, 0x30, 0x5f, 0x15, 0x6d,
	0xa7, 0x96, 0x5e, 0x8a, 0xfc, 0xfa, 0xb3, 0xa2, 0x35, 0xd6, 0x3c, 0xf5, 0x53, 0xb2, 0xed, 0xcd,
	0xe4, 0xb7, 0x10, 0x46, 0x6f, 0xa3, 0xae, 0x6a, 0x82, 0xd4, 0xaa, 0xf6, 0xc9, 0xfd, 0xcb, 0xdb,
	0x83, 0x4a, 0xf5, 0x5e, 0xd1, 0xd5, 0xed, 0x5a, 0xfc, 0x33, 0x7f, 0x8e, 0x19, 0x8f, 0x21, 0x45,
	0xda, 0x2d, 0xcb, 0x2b, 0x24, 0x48, 0xbb, 0x77, 0xc6, 0xc1, 0xd2, 0x6a, 0xd6, 0xe6, 0x25, 0x9a,
	0x45, 0x48, 0x7c, 0x1e, 0x0f, 0x3f, 0x14, 0xad, 0xe3, 0xf3, 0x78, 0xeb, 0x03, 0x86, 0x7b, 0x3f,
	0x8a, 0x18, 0x15, 0x05, 0x8f, 0x94, 0xab, 0xf3, 0x52, 0x1b, 0xf2, 0xa9, 0x97, 0x84, 0x84, 0xa7,
	0x61, 0x91, 0x60, 0x55, 0x3f, 0xad, 0xaf, 0x6d, 0xfe, 0x4b, 0x03, 0x09, 0x1e, 0xd5, 0x3f, 0xb6,
	0x65, 0x3d, 0xa9, 0x2d, 0x5b, 0x18, 0x5a, 0x01, 0x12, 0x01, 0x0b, 0xe4, 0x98, 0xa4, 0x3c, 0x8c,
	0x4b, 0x49, 0x71, 0xad, 0x00, 0xdd, 0x8a, 0xe2, 0x3a, 0xa6, 0x2b, 0x29, 0xef, 0x6a, 0x47, 0xcc,
	0x19, 0xb3, 0x02, 0x64, 0xac, 0x22, 0x31, 0xe6, 0x98, 0x8c, 0x49, 0xda, 0xf5, 0xdf, 0x55, 0xce,
	0xfe, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xfa, 0x0b, 0x30, 0xf2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GenderViewServiceClient is the client API for GenderViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenderViewServiceClient interface {
	// Returns the requested gender view in full detail.
	GetGenderView(ctx context.Context, in *GetGenderViewRequest, opts ...grpc.CallOption) (*resources.GenderView, error)
}

type genderViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewGenderViewServiceClient(cc *grpc.ClientConn) GenderViewServiceClient {
	return &genderViewServiceClient{cc}
}

func (c *genderViewServiceClient) GetGenderView(ctx context.Context, in *GetGenderViewRequest, opts ...grpc.CallOption) (*resources.GenderView, error) {
	out := new(resources.GenderView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.GenderViewService/GetGenderView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenderViewServiceServer is the server API for GenderViewService service.
type GenderViewServiceServer interface {
	// Returns the requested gender view in full detail.
	GetGenderView(context.Context, *GetGenderViewRequest) (*resources.GenderView, error)
}

// UnimplementedGenderViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGenderViewServiceServer struct {
}

func (*UnimplementedGenderViewServiceServer) GetGenderView(ctx context.Context, req *GetGenderViewRequest) (*resources.GenderView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenderView not implemented")
}

func RegisterGenderViewServiceServer(s *grpc.Server, srv GenderViewServiceServer) {
	s.RegisterService(&_GenderViewService_serviceDesc, srv)
}

func _GenderViewService_GetGenderView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenderViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenderViewServiceServer).GetGenderView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.GenderViewService/GetGenderView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenderViewServiceServer).GetGenderView(ctx, req.(*GetGenderViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GenderViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.GenderViewService",
	HandlerType: (*GenderViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGenderView",
			Handler:    _GenderViewService_GetGenderView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/gender_view_service.proto",
}
