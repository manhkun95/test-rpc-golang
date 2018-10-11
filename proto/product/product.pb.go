// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetProductsRequest struct {
	CusAddress           string   `protobuf:"bytes,1,opt,name=cusAddress,proto3" json:"cusAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductsRequest) Reset()         { *m = GetProductsRequest{} }
func (m *GetProductsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductsRequest) ProtoMessage()    {}
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *GetProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsRequest.Unmarshal(m, b)
}
func (m *GetProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsRequest.Marshal(b, m, deterministic)
}
func (m *GetProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsRequest.Merge(m, src)
}
func (m *GetProductsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductsRequest.Size(m)
}
func (m *GetProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsRequest proto.InternalMessageInfo

func (m *GetProductsRequest) GetCusAddress() string {
	if m != nil {
		return m.CusAddress
	}
	return ""
}

type GetProductsResponse struct {
	Result               int32                           `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Products             []*GetProductsResponse_Products `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Count                int32                           `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Page                 int32                           `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetProductsResponse) Reset()         { *m = GetProductsResponse{} }
func (m *GetProductsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductsResponse) ProtoMessage()    {}
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *GetProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsResponse.Unmarshal(m, b)
}
func (m *GetProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsResponse.Marshal(b, m, deterministic)
}
func (m *GetProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsResponse.Merge(m, src)
}
func (m *GetProductsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductsResponse.Size(m)
}
func (m *GetProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsResponse proto.InternalMessageInfo

func (m *GetProductsResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *GetProductsResponse) GetProducts() []*GetProductsResponse_Products {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *GetProductsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetProductsResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type GetProductsResponse_Products struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Detail               string   `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	Imgs                 []string `protobuf:"bytes,5,rep,name=imgs,proto3" json:"imgs,omitempty"`
	Properties           string   `protobuf:"bytes,6,opt,name=properties,proto3" json:"properties,omitempty"`
	Price                int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	TotalNumber          int32    `protobuf:"varint,8,opt,name=totalNumber,proto3" json:"totalNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductsResponse_Products) Reset()         { *m = GetProductsResponse_Products{} }
func (m *GetProductsResponse_Products) String() string { return proto.CompactTextString(m) }
func (*GetProductsResponse_Products) ProtoMessage()    {}
func (*GetProductsResponse_Products) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1, 0}
}

func (m *GetProductsResponse_Products) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsResponse_Products.Unmarshal(m, b)
}
func (m *GetProductsResponse_Products) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsResponse_Products.Marshal(b, m, deterministic)
}
func (m *GetProductsResponse_Products) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsResponse_Products.Merge(m, src)
}
func (m *GetProductsResponse_Products) XXX_Size() int {
	return xxx_messageInfo_GetProductsResponse_Products.Size(m)
}
func (m *GetProductsResponse_Products) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsResponse_Products.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsResponse_Products proto.InternalMessageInfo

func (m *GetProductsResponse_Products) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetProductsResponse_Products) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetProductsResponse_Products) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetProductsResponse_Products) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *GetProductsResponse_Products) GetImgs() []string {
	if m != nil {
		return m.Imgs
	}
	return nil
}

func (m *GetProductsResponse_Products) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *GetProductsResponse_Products) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GetProductsResponse_Products) GetTotalNumber() int32 {
	if m != nil {
		return m.TotalNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*GetProductsRequest)(nil), "proto.GetProductsRequest")
	proto.RegisterType((*GetProductsResponse)(nil), "proto.GetProductsResponse")
	proto.RegisterType((*GetProductsResponse_Products)(nil), "proto.GetProductsResponse.Products")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4d, 0x4f, 0xc3, 0x30,
	0x0c, 0x65, 0xdd, 0xba, 0x0f, 0x4f, 0x70, 0x30, 0x08, 0x85, 0x1d, 0x50, 0x35, 0x2e, 0x3b, 0xed,
	0x30, 0xb8, 0x23, 0x2e, 0x70, 0x43, 0xd0, 0x7f, 0xd0, 0x35, 0xd6, 0x14, 0x69, 0x6b, 0x42, 0x9c,
	0xfc, 0x53, 0xfe, 0x07, 0x7f, 0x01, 0xc5, 0x2d, 0xa3, 0x48, 0xe3, 0x14, 0xfb, 0x59, 0xef, 0xd9,
	0xef, 0x05, 0xce, 0x9d, 0xb7, 0x3a, 0xd6, 0x61, 0xed, 0xbc, 0x0d, 0x16, 0x73, 0x79, 0x96, 0x0f,
	0x80, 0x2f, 0x14, 0xde, 0xda, 0x11, 0x97, 0xf4, 0x11, 0x89, 0x03, 0xde, 0x02, 0xd4, 0x91, 0x9f,
	0xb4, 0xf6, 0xc4, 0xac, 0x06, 0xc5, 0x60, 0x35, 0x2b, 0x7b, 0xc8, 0xf2, 0x2b, 0x83, 0xcb, 0x3f,
	0x34, 0x76, 0xb6, 0x61, 0xc2, 0x6b, 0x18, 0x7b, 0xe2, 0xb8, 0x0f, 0xc2, 0xc9, 0xcb, 0xae, 0xc3,
	0x47, 0x98, 0x76, 0xdb, 0x59, 0x65, 0xc5, 0x70, 0x35, 0xdf, 0xdc, 0xb5, 0x67, 0xac, 0x4f, 0xa8,
	0xac, 0x8f, 0xc0, 0x91, 0x84, 0x57, 0x90, 0xd7, 0x36, 0x36, 0x41, 0x0d, 0x45, 0xb7, 0x6d, 0x10,
	0x61, 0xe4, 0xaa, 0x1d, 0xa9, 0x91, 0x80, 0x52, 0x2f, 0x3e, 0x07, 0x30, 0xfd, 0x11, 0xc0, 0x0b,
	0xc8, 0x8c, 0xee, 0x6e, 0xc9, 0x8c, 0x4e, 0x84, 0xa6, 0x3a, 0x90, 0xca, 0xc4, 0x91, 0xd4, 0x58,
	0xc0, 0x5c, 0x13, 0xd7, 0xde, 0xb8, 0x60, 0x6c, 0x23, 0x0b, 0x66, 0x65, 0x1f, 0x4a, 0xae, 0x34,
	0x85, 0xca, 0xec, 0x65, 0xd1, 0xac, 0xec, 0xba, 0xa4, 0x66, 0x0e, 0x3b, 0x56, 0x79, 0x31, 0x4c,
	0x6a, 0xa9, 0x4e, 0xc9, 0x39, 0x6f, 0x1d, 0xf9, 0x60, 0x88, 0xd5, 0xb8, 0x4d, 0xee, 0x17, 0x49,
	0x46, 0x9c, 0x37, 0x35, 0xa9, 0x49, 0x6b, 0x44, 0x9a, 0x74, 0x43, 0xb0, 0xa1, 0xda, 0xbf, 0xc6,
	0xc3, 0x96, 0xbc, 0x9a, 0xca, 0xac, 0x0f, 0x6d, 0xde, 0x61, 0xd2, 0xb9, 0xc2, 0x67, 0x98, 0xf7,
	0x52, 0xc3, 0x9b, 0x53, 0x49, 0xca, 0x37, 0x2e, 0x16, 0xff, 0x87, 0xbc, 0x3c, 0xdb, 0x8e, 0x65,
	0x78, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x88, 0x10, 0xbb, 0x19, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
}

type productClient struct {
	cc *grpc.ClientConn
}

func NewProductClient(cc *grpc.ClientConn) ProductClient {
	return &productClient{cc}
}

func (c *productClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/proto.Product/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Product/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _Product_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
