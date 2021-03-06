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
	Result               int32       `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Products             []*Products `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Count                int32       `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Page                 int32       `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
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

func (m *GetProductsResponse) GetProducts() []*Products {
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

type Products struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Detail               string   `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	Imgs                 string   `protobuf:"bytes,5,opt,name=imgs,proto3" json:"imgs,omitempty"`
	Properties           string   `protobuf:"bytes,6,opt,name=properties,proto3" json:"properties,omitempty"`
	Price                int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	TotalNumber          int32    `protobuf:"varint,8,opt,name=totalNumber,proto3" json:"totalNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Products) Reset()         { *m = Products{} }
func (m *Products) String() string { return proto.CompactTextString(m) }
func (*Products) ProtoMessage()    {}
func (*Products) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *Products) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Products.Unmarshal(m, b)
}
func (m *Products) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Products.Marshal(b, m, deterministic)
}
func (m *Products) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Products.Merge(m, src)
}
func (m *Products) XXX_Size() int {
	return xxx_messageInfo_Products.Size(m)
}
func (m *Products) XXX_DiscardUnknown() {
	xxx_messageInfo_Products.DiscardUnknown(m)
}

var xxx_messageInfo_Products proto.InternalMessageInfo

func (m *Products) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Products) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Products) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Products) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Products) GetImgs() string {
	if m != nil {
		return m.Imgs
	}
	return ""
}

func (m *Products) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *Products) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Products) GetTotalNumber() int32 {
	if m != nil {
		return m.TotalNumber
	}
	return 0
}

type CreateProductRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductRequest) Reset()         { *m = CreateProductRequest{} }
func (m *CreateProductRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProductRequest) ProtoMessage()    {}
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{3}
}

func (m *CreateProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductRequest.Unmarshal(m, b)
}
func (m *CreateProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductRequest.Marshal(b, m, deterministic)
}
func (m *CreateProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductRequest.Merge(m, src)
}
func (m *CreateProductRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProductRequest.Size(m)
}
func (m *CreateProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductRequest proto.InternalMessageInfo

func (m *CreateProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateProductRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateProductResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductResponse) Reset()         { *m = CreateProductResponse{} }
func (m *CreateProductResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProductResponse) ProtoMessage()    {}
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{4}
}

func (m *CreateProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductResponse.Unmarshal(m, b)
}
func (m *CreateProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductResponse.Marshal(b, m, deterministic)
}
func (m *CreateProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductResponse.Merge(m, src)
}
func (m *CreateProductResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProductResponse.Size(m)
}
func (m *CreateProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductResponse proto.InternalMessageInfo

func (m *CreateProductResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*GetProductsRequest)(nil), "proto.GetProductsRequest")
	proto.RegisterType((*GetProductsResponse)(nil), "proto.GetProductsResponse")
	proto.RegisterType((*Products)(nil), "proto.Products")
	proto.RegisterType((*CreateProductRequest)(nil), "proto.CreateProductRequest")
	proto.RegisterType((*CreateProductResponse)(nil), "proto.CreateProductResponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xed, 0x4a, 0xeb, 0x40,
	0x10, 0xbd, 0x49, 0x9b, 0x7e, 0x4c, 0xe9, 0xbd, 0x30, 0xb7, 0xf7, 0xb2, 0x56, 0x91, 0x92, 0x5f,
	0x05, 0xa1, 0x42, 0xf5, 0x05, 0x44, 0xd0, 0x3f, 0x45, 0x24, 0x6f, 0x90, 0x26, 0x43, 0x59, 0x68,
	0xb3, 0xeb, 0xee, 0xe6, 0x0d, 0x7c, 0x0f, 0x5f, 0xca, 0x07, 0x92, 0xfd, 0x48, 0x4d, 0x6d, 0xf5,
	0x57, 0x66, 0xe6, 0xcc, 0x9c, 0x33, 0x27, 0xb3, 0x30, 0x96, 0x4a, 0x94, 0x75, 0x61, 0x16, 0x52,
	0x09, 0x23, 0x30, 0x71, 0x9f, 0xf4, 0x16, 0xf0, 0x91, 0xcc, 0xb3, 0x87, 0x74, 0x46, 0x2f, 0x35,
	0x69, 0x83, 0x97, 0x00, 0x45, 0xad, 0xef, 0xca, 0x52, 0x91, 0xd6, 0x2c, 0x9a, 0x45, 0xf3, 0x61,
	0xd6, 0xaa, 0xa4, 0xaf, 0x11, 0xfc, 0x3d, 0x18, 0xd3, 0x52, 0x54, 0x9a, 0xf0, 0x3f, 0xf4, 0x14,
	0xe9, 0x7a, 0x6b, 0xdc, 0x4c, 0x92, 0x85, 0x0c, 0xaf, 0x60, 0x10, 0xd4, 0x35, 0x8b, 0x67, 0x9d,
	0xf9, 0x68, 0xf9, 0xc7, 0xaf, 0xb1, 0xd8, 0x53, 0xec, 0x1b, 0x70, 0x02, 0x49, 0x21, 0xea, 0xca,
	0xb0, 0x8e, 0xe3, 0xf0, 0x09, 0x22, 0x74, 0x65, 0xbe, 0x21, 0xd6, 0x75, 0x45, 0x17, 0xa7, 0xef,
	0x11, 0x0c, 0x1a, 0x02, 0xfc, 0x0d, 0x31, 0x2f, 0x83, 0x6e, 0xcc, 0x4b, 0x3b, 0x50, 0xe5, 0x3b,
	0x62, 0xb1, 0xdb, 0xde, 0xc5, 0x38, 0x83, 0x51, 0x49, 0xba, 0x50, 0x5c, 0x1a, 0x2e, 0x2a, 0x27,
	0x30, 0xcc, 0xda, 0x25, 0xeb, 0xa0, 0x24, 0x93, 0xf3, 0xad, 0x13, 0x1a, 0x66, 0x21, 0xb3, 0x6c,
	0x7c, 0xb7, 0xd1, 0x2c, 0xf1, 0x6c, 0x36, 0xb6, 0x7f, 0x49, 0x2a, 0x21, 0x49, 0x19, 0x4e, 0x9a,
	0xf5, 0xfc, 0x5f, 0xfa, 0xac, 0x58, 0x23, 0x52, 0xf1, 0x82, 0x58, 0xdf, 0x1b, 0x71, 0x89, 0xdd,
	0xc1, 0x08, 0x93, 0x6f, 0x9f, 0xea, 0xdd, 0x9a, 0x14, 0x1b, 0x38, 0xac, 0x5d, 0x4a, 0x57, 0x30,
	0xb9, 0x57, 0x94, 0x1b, 0x0a, 0xde, 0x9a, 0xab, 0x34, 0x8e, 0xa2, 0xef, 0x1d, 0xc5, 0x47, 0x8e,
	0xd2, 0x6b, 0xf8, 0xf7, 0x85, 0xed, 0xe7, 0x63, 0x2d, 0xdf, 0x22, 0xe8, 0x87, 0x5e, 0x5c, 0xc1,
	0xf8, 0x60, 0x18, 0xcf, 0xc3, 0xdd, 0x4e, 0x2d, 0x38, 0xbd, 0x38, 0x0d, 0x7a, 0xbd, 0xf4, 0x17,
	0x3e, 0xc0, 0xa8, 0xf5, 0x6a, 0xf0, 0x2c, 0xb4, 0x1f, 0x3f, 0xc0, 0xe9, 0xf4, 0x14, 0xd4, 0xf0,
	0xac, 0x7b, 0x0e, 0xbc, 0xf9, 0x08, 0x00, 0x00, 0xff, 0xff, 0x22, 0xa5, 0x80, 0xee, 0xd3, 0x02,
	0x00, 0x00,
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
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
}

type productClient struct {
	cc *grpc.ClientConn
}

func NewProductClient(cc *grpc.ClientConn) ProductClient {
	return &productClient{cc}
}

func (c *productClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/proto.Product/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Product/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _Product_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
