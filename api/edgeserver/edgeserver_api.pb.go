// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edgeserver_api.proto

package edgeserver

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Image struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_2122c087cbb44120, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Image) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type NodeInfo struct {
	Ipaddr               string   `protobuf:"bytes,1,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	Camid                string   `protobuf:"bytes,2,opt,name=camid,proto3" json:"camid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2122c087cbb44120, []int{1}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetIpaddr() string {
	if m != nil {
		return m.Ipaddr
	}
	return ""
}

func (m *NodeInfo) GetCamid() string {
	if m != nil {
		return m.Camid
	}
	return ""
}

type CameraParameters struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CameraParameters) Reset()         { *m = CameraParameters{} }
func (m *CameraParameters) String() string { return proto.CompactTextString(m) }
func (*CameraParameters) ProtoMessage()    {}
func (*CameraParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_2122c087cbb44120, []int{2}
}

func (m *CameraParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CameraParameters.Unmarshal(m, b)
}
func (m *CameraParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CameraParameters.Marshal(b, m, deterministic)
}
func (m *CameraParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CameraParameters.Merge(m, src)
}
func (m *CameraParameters) XXX_Size() int {
	return xxx_messageInfo_CameraParameters.Size(m)
}
func (m *CameraParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_CameraParameters.DiscardUnknown(m)
}

var xxx_messageInfo_CameraParameters proto.InternalMessageInfo

type CameraInfo struct {
	Camid                []string `protobuf:"bytes,1,rep,name=camid,proto3" json:"camid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CameraInfo) Reset()         { *m = CameraInfo{} }
func (m *CameraInfo) String() string { return proto.CompactTextString(m) }
func (*CameraInfo) ProtoMessage()    {}
func (*CameraInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2122c087cbb44120, []int{3}
}

func (m *CameraInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CameraInfo.Unmarshal(m, b)
}
func (m *CameraInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CameraInfo.Marshal(b, m, deterministic)
}
func (m *CameraInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CameraInfo.Merge(m, src)
}
func (m *CameraInfo) XXX_Size() int {
	return xxx_messageInfo_CameraInfo.Size(m)
}
func (m *CameraInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CameraInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CameraInfo proto.InternalMessageInfo

func (m *CameraInfo) GetCamid() []string {
	if m != nil {
		return m.Camid
	}
	return nil
}

type ImageStreamParameters struct {
	Appid                string   `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Camid                string   `protobuf:"bytes,2,opt,name=camid,proto3" json:"camid,omitempty"`
	Latency              string   `protobuf:"bytes,3,opt,name=latency,proto3" json:"latency,omitempty"`
	Accuracy             string   `protobuf:"bytes,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	Start                string   `protobuf:"bytes,5,opt,name=start,proto3" json:"start,omitempty"`
	Stop                 string   `protobuf:"bytes,6,opt,name=stop,proto3" json:"stop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageStreamParameters) Reset()         { *m = ImageStreamParameters{} }
func (m *ImageStreamParameters) String() string { return proto.CompactTextString(m) }
func (*ImageStreamParameters) ProtoMessage()    {}
func (*ImageStreamParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_2122c087cbb44120, []int{4}
}

func (m *ImageStreamParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageStreamParameters.Unmarshal(m, b)
}
func (m *ImageStreamParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageStreamParameters.Marshal(b, m, deterministic)
}
func (m *ImageStreamParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageStreamParameters.Merge(m, src)
}
func (m *ImageStreamParameters) XXX_Size() int {
	return xxx_messageInfo_ImageStreamParameters.Size(m)
}
func (m *ImageStreamParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageStreamParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ImageStreamParameters proto.InternalMessageInfo

func (m *ImageStreamParameters) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *ImageStreamParameters) GetCamid() string {
	if m != nil {
		return m.Camid
	}
	return ""
}

func (m *ImageStreamParameters) GetLatency() string {
	if m != nil {
		return m.Latency
	}
	return ""
}

func (m *ImageStreamParameters) GetAccuracy() string {
	if m != nil {
		return m.Accuracy
	}
	return ""
}

func (m *ImageStreamParameters) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ImageStreamParameters) GetStop() string {
	if m != nil {
		return m.Stop
	}
	return ""
}

type AppInfo struct {
	Appid                string   `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Camid                string   `protobuf:"bytes,2,opt,name=camid,proto3" json:"camid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppInfo) Reset()         { *m = AppInfo{} }
func (m *AppInfo) String() string { return proto.CompactTextString(m) }
func (*AppInfo) ProtoMessage()    {}
func (*AppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2122c087cbb44120, []int{5}
}

func (m *AppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfo.Unmarshal(m, b)
}
func (m *AppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfo.Marshal(b, m, deterministic)
}
func (m *AppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfo.Merge(m, src)
}
func (m *AppInfo) XXX_Size() int {
	return xxx_messageInfo_AppInfo.Size(m)
}
func (m *AppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfo proto.InternalMessageInfo

func (m *AppInfo) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *AppInfo) GetCamid() string {
	if m != nil {
		return m.Camid
	}
	return ""
}

type Status struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_2122c087cbb44120, []int{6}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*Image)(nil), "edgeserver.Image")
	proto.RegisterType((*NodeInfo)(nil), "edgeserver.NodeInfo")
	proto.RegisterType((*CameraParameters)(nil), "edgeserver.CameraParameters")
	proto.RegisterType((*CameraInfo)(nil), "edgeserver.CameraInfo")
	proto.RegisterType((*ImageStreamParameters)(nil), "edgeserver.ImageStreamParameters")
	proto.RegisterType((*AppInfo)(nil), "edgeserver.AppInfo")
	proto.RegisterType((*Status)(nil), "edgeserver.Status")
}

func init() { proto.RegisterFile("edgeserver_api.proto", fileDescriptor_2122c087cbb44120) }

var fileDescriptor_2122c087cbb44120 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0xeb, 0xd3, 0x40,
	0x14, 0x6c, 0x7e, 0xbf, 0x36, 0x4d, 0x9e, 0x0a, 0xfa, 0xac, 0x25, 0x84, 0x1e, 0xe2, 0x9e, 0x7a,
	0x2a, 0xa2, 0x28, 0x05, 0x4f, 0x22, 0x52, 0x7a, 0x91, 0x92, 0xd0, 0xb3, 0xbc, 0x24, 0xcf, 0x12,
	0x30, 0xc9, 0xb2, 0xbb, 0x11, 0xfa, 0x79, 0xfc, 0x2a, 0x7e, 0x30, 0xc9, 0x26, 0x6d, 0x82, 0xed,
	0x41, 0xbc, 0xbd, 0x99, 0x9d, 0x99, 0xfd, 0x33, 0x2c, 0x2c, 0x38, 0x3f, 0xb1, 0x66, 0xf5, 0x93,
	0xd5, 0x37, 0x92, 0xc5, 0x46, 0xaa, 0xda, 0xd4, 0x08, 0x03, 0x2b, 0x3e, 0xc2, 0x6c, 0x5f, 0xd2,
	0x89, 0x71, 0x01, 0xb3, 0xa2, 0x1d, 0x02, 0x27, 0x72, 0xd6, 0x4f, 0xe3, 0x0e, 0xe0, 0x0a, 0x7c,
	0x53, 0x94, 0xac, 0x0d, 0x95, 0x32, 0x78, 0x88, 0x9c, 0xb5, 0x1f, 0x0f, 0x84, 0xd8, 0x82, 0xf7,
	0xb5, 0xce, 0x79, 0x5f, 0x7d, 0xaf, 0x71, 0x09, 0x6e, 0x21, 0x29, 0xcf, 0x95, 0x0d, 0xf0, 0xe3,
	0x1e, 0xb5, 0xb9, 0x19, 0x95, 0x45, 0xde, 0xbb, 0x3b, 0x20, 0x10, 0x9e, 0x7f, 0xa6, 0x92, 0x15,
	0x1d, 0x48, 0x51, 0xc9, 0x86, 0x95, 0x16, 0x02, 0xa0, 0xe3, 0x6c, 0xde, 0xd5, 0xe7, 0x44, 0x8f,
	0x83, 0xef, 0x97, 0x03, 0xaf, 0xec, 0x79, 0x13, 0xa3, 0x98, 0xca, 0xc1, 0xdd, 0xea, 0x49, 0x4a,
	0xab, 0xb7, 0xfb, 0x58, 0x70, 0x7f, 0x77, 0x0c, 0x60, 0xfe, 0x83, 0x0c, 0x57, 0xd9, 0x39, 0x78,
	0xb4, 0xfc, 0x05, 0x62, 0x08, 0x1e, 0x65, 0x59, 0xa3, 0x28, 0x3b, 0x07, 0x53, 0xbb, 0x74, 0xc5,
	0x6d, 0x96, 0x36, 0xa4, 0x4c, 0x30, 0xeb, 0xb2, 0x2c, 0x40, 0x84, 0xa9, 0x36, 0xb5, 0x0c, 0x5c,
	0x4b, 0xda, 0x59, 0xbc, 0x87, 0xf9, 0x27, 0x29, 0x2f, 0xd7, 0xf8, 0xd7, 0x63, 0x89, 0x08, 0xdc,
	0xc4, 0x90, 0x69, 0x74, 0xfb, 0x98, 0xda, 0x4e, 0xd6, 0xe6, 0xc5, 0x3d, 0x7a, 0xfb, 0xfb, 0x01,
	0xdc, 0x43, 0x93, 0x26, 0x4d, 0x8a, 0x1f, 0xc0, 0x8b, 0xf9, 0x54, 0x68, 0xc3, 0x0a, 0x17, 0x9b,
	0xa1, 0xd1, 0xcd, 0xa5, 0x91, 0x10, 0xc7, 0x6c, 0x17, 0x2c, 0x26, 0xb8, 0x05, 0x38, 0x56, 0xea,
	0x7f, 0x9c, 0x3b, 0x78, 0xb6, 0x63, 0x33, 0xaa, 0x68, 0x35, 0x96, 0xfd, 0x5d, 0x67, 0xb8, 0xbc,
	0x5d, 0x6d, 0x5d, 0x62, 0x82, 0x5f, 0xc0, 0x4f, 0x9a, 0x54, 0x67, 0xaa, 0x48, 0x19, 0x5f, 0x8f,
	0x65, 0x77, 0xab, 0x0d, 0x5f, 0xdc, 0x48, 0xc4, 0xe4, 0x8d, 0x83, 0x5b, 0x78, 0x72, 0xac, 0xf4,
	0x35, 0xe8, 0xe5, 0x58, 0xd5, 0x3f, 0xff, 0xfd, 0x9b, 0xa4, 0xae, 0xfd, 0x07, 0xef, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x87, 0xd9, 0x1e, 0x75, 0x1f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubSubClient is the client API for PubSub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubSubClient interface {
	Register(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*Status, error)
	Unregister(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*Status, error)
	GetCameraInfo(ctx context.Context, in *CameraParameters, opts ...grpc.CallOption) (*CameraInfo, error)
	Subscribe(ctx context.Context, in *ImageStreamParameters, opts ...grpc.CallOption) (PubSub_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*Status, error)
}

type pubSubClient struct {
	cc *grpc.ClientConn
}

func NewPubSubClient(cc *grpc.ClientConn) PubSubClient {
	return &pubSubClient{cc}
}

func (c *pubSubClient) Register(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/edgeserver.PubSub/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) Unregister(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/edgeserver.PubSub/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) GetCameraInfo(ctx context.Context, in *CameraParameters, opts ...grpc.CallOption) (*CameraInfo, error) {
	out := new(CameraInfo)
	err := c.cc.Invoke(ctx, "/edgeserver.PubSub/GetCameraInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) Subscribe(ctx context.Context, in *ImageStreamParameters, opts ...grpc.CallOption) (PubSub_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSub_serviceDesc.Streams[0], "/edgeserver.PubSub/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSub_SubscribeClient interface {
	Recv() (*Image, error)
	grpc.ClientStream
}

type pubSubSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubSubscribeClient) Recv() (*Image, error) {
	m := new(Image)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubClient) Unsubscribe(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/edgeserver.PubSub/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSubServer is the server API for PubSub service.
type PubSubServer interface {
	Register(context.Context, *NodeInfo) (*Status, error)
	Unregister(context.Context, *NodeInfo) (*Status, error)
	GetCameraInfo(context.Context, *CameraParameters) (*CameraInfo, error)
	Subscribe(*ImageStreamParameters, PubSub_SubscribeServer) error
	Unsubscribe(context.Context, *AppInfo) (*Status, error)
}

// UnimplementedPubSubServer can be embedded to have forward compatible implementations.
type UnimplementedPubSubServer struct {
}

func (*UnimplementedPubSubServer) Register(ctx context.Context, req *NodeInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedPubSubServer) Unregister(ctx context.Context, req *NodeInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (*UnimplementedPubSubServer) GetCameraInfo(ctx context.Context, req *CameraParameters) (*CameraInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCameraInfo not implemented")
}
func (*UnimplementedPubSubServer) Subscribe(req *ImageStreamParameters, srv PubSub_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedPubSubServer) Unsubscribe(ctx context.Context, req *AppInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}

func RegisterPubSubServer(s *grpc.Server, srv PubSubServer) {
	s.RegisterService(&_PubSub_serviceDesc, srv)
}

func _PubSub_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgeserver.PubSub/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Register(ctx, req.(*NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgeserver.PubSub/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Unregister(ctx, req.(*NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_GetCameraInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CameraParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).GetCameraInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgeserver.PubSub/GetCameraInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).GetCameraInfo(ctx, req.(*CameraParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImageStreamParameters)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServer).Subscribe(m, &pubSubSubscribeServer{stream})
}

type PubSub_SubscribeServer interface {
	Send(*Image) error
	grpc.ServerStream
}

type pubSubSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubSubscribeServer) Send(m *Image) error {
	return x.ServerStream.SendMsg(m)
}

func _PubSub_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgeserver.PubSub/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Unsubscribe(ctx, req.(*AppInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _PubSub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "edgeserver.PubSub",
	HandlerType: (*PubSubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PubSub_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PubSub_Unregister_Handler,
		},
		{
			MethodName: "GetCameraInfo",
			Handler:    _PubSub_GetCameraInfo_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _PubSub_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSub_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "edgeserver_api.proto",
}
