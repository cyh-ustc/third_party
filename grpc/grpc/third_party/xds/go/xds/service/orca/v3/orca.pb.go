// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: xds/service/orca/v3/orca.proto

package v3

import (
	context "context"
	v3 "github.com/cncf/xds/go/xds/data/orca/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrcaLoadReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportInterval   *duration.Duration `protobuf:"bytes,1,opt,name=report_interval,json=reportInterval,proto3" json:"report_interval,omitempty"`
	RequestCostNames []string           `protobuf:"bytes,2,rep,name=request_cost_names,json=requestCostNames,proto3" json:"request_cost_names,omitempty"`
}

func (x *OrcaLoadReportRequest) Reset() {
	*x = OrcaLoadReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_service_orca_v3_orca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrcaLoadReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrcaLoadReportRequest) ProtoMessage() {}

func (x *OrcaLoadReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xds_service_orca_v3_orca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrcaLoadReportRequest.ProtoReflect.Descriptor instead.
func (*OrcaLoadReportRequest) Descriptor() ([]byte, []int) {
	return file_xds_service_orca_v3_orca_proto_rawDescGZIP(), []int{0}
}

func (x *OrcaLoadReportRequest) GetReportInterval() *duration.Duration {
	if x != nil {
		return x.ReportInterval
	}
	return nil
}

func (x *OrcaLoadReportRequest) GetRequestCostNames() []string {
	if x != nil {
		return x.RequestCostNames
	}
	return nil
}

var File_xds_service_orca_v3_orca_proto protoreflect.FileDescriptor

var file_xds_service_orca_v3_orca_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x78, 0x64, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72,
	0x63, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x78, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x72,
	0x63, 0x61, 0x2e, 0x76, 0x33, 0x1a, 0x27, 0x78, 0x64, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x6f, 0x72, 0x63, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x5f, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x4f, 0x72, 0x63, 0x61,
	0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x42, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x32, 0x75, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x63, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43,
	0x6f, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2a, 0x2e, 0x78, 0x64, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x76, 0x33,
	0x2e, 0x4f, 0x72, 0x63, 0x61, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x63, 0x61, 0x4c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x30, 0x01, 0x42, 0x59, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x4f, 0x72,
	0x63, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72,
	0x63, 0x61, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_service_orca_v3_orca_proto_rawDescOnce sync.Once
	file_xds_service_orca_v3_orca_proto_rawDescData = file_xds_service_orca_v3_orca_proto_rawDesc
)

func file_xds_service_orca_v3_orca_proto_rawDescGZIP() []byte {
	file_xds_service_orca_v3_orca_proto_rawDescOnce.Do(func() {
		file_xds_service_orca_v3_orca_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_service_orca_v3_orca_proto_rawDescData)
	})
	return file_xds_service_orca_v3_orca_proto_rawDescData
}

var file_xds_service_orca_v3_orca_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_xds_service_orca_v3_orca_proto_goTypes = []interface{}{
	(*OrcaLoadReportRequest)(nil), // 0: xds.service.orca.v3.OrcaLoadReportRequest
	(*duration.Duration)(nil),     // 1: google.protobuf.Duration
	(*v3.OrcaLoadReport)(nil),     // 2: xds.data.orca.v3.OrcaLoadReport
}
var file_xds_service_orca_v3_orca_proto_depIdxs = []int32{
	1, // 0: xds.service.orca.v3.OrcaLoadReportRequest.report_interval:type_name -> google.protobuf.Duration
	0, // 1: xds.service.orca.v3.OpenRcaService.StreamCoreMetrics:input_type -> xds.service.orca.v3.OrcaLoadReportRequest
	2, // 2: xds.service.orca.v3.OpenRcaService.StreamCoreMetrics:output_type -> xds.data.orca.v3.OrcaLoadReport
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xds_service_orca_v3_orca_proto_init() }
func file_xds_service_orca_v3_orca_proto_init() {
	if File_xds_service_orca_v3_orca_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_service_orca_v3_orca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrcaLoadReportRequest); i {
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
			RawDescriptor: file_xds_service_orca_v3_orca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xds_service_orca_v3_orca_proto_goTypes,
		DependencyIndexes: file_xds_service_orca_v3_orca_proto_depIdxs,
		MessageInfos:      file_xds_service_orca_v3_orca_proto_msgTypes,
	}.Build()
	File_xds_service_orca_v3_orca_proto = out.File
	file_xds_service_orca_v3_orca_proto_rawDesc = nil
	file_xds_service_orca_v3_orca_proto_goTypes = nil
	file_xds_service_orca_v3_orca_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OpenRcaServiceClient is the client API for OpenRcaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OpenRcaServiceClient interface {
	StreamCoreMetrics(ctx context.Context, in *OrcaLoadReportRequest, opts ...grpc.CallOption) (OpenRcaService_StreamCoreMetricsClient, error)
}

type openRcaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenRcaServiceClient(cc grpc.ClientConnInterface) OpenRcaServiceClient {
	return &openRcaServiceClient{cc}
}

func (c *openRcaServiceClient) StreamCoreMetrics(ctx context.Context, in *OrcaLoadReportRequest, opts ...grpc.CallOption) (OpenRcaService_StreamCoreMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OpenRcaService_serviceDesc.Streams[0], "/xds.service.orca.v3.OpenRcaService/StreamCoreMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &openRcaServiceStreamCoreMetricsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OpenRcaService_StreamCoreMetricsClient interface {
	Recv() (*v3.OrcaLoadReport, error)
	grpc.ClientStream
}

type openRcaServiceStreamCoreMetricsClient struct {
	grpc.ClientStream
}

func (x *openRcaServiceStreamCoreMetricsClient) Recv() (*v3.OrcaLoadReport, error) {
	m := new(v3.OrcaLoadReport)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OpenRcaServiceServer is the server API for OpenRcaService service.
type OpenRcaServiceServer interface {
	StreamCoreMetrics(*OrcaLoadReportRequest, OpenRcaService_StreamCoreMetricsServer) error
}

// UnimplementedOpenRcaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOpenRcaServiceServer struct {
}

func (*UnimplementedOpenRcaServiceServer) StreamCoreMetrics(*OrcaLoadReportRequest, OpenRcaService_StreamCoreMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCoreMetrics not implemented")
}

func RegisterOpenRcaServiceServer(s *grpc.Server, srv OpenRcaServiceServer) {
	s.RegisterService(&_OpenRcaService_serviceDesc, srv)
}

func _OpenRcaService_StreamCoreMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrcaLoadReportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OpenRcaServiceServer).StreamCoreMetrics(m, &openRcaServiceStreamCoreMetricsServer{stream})
}

type OpenRcaService_StreamCoreMetricsServer interface {
	Send(*v3.OrcaLoadReport) error
	grpc.ServerStream
}

type openRcaServiceStreamCoreMetricsServer struct {
	grpc.ServerStream
}

func (x *openRcaServiceStreamCoreMetricsServer) Send(m *v3.OrcaLoadReport) error {
	return x.ServerStream.SendMsg(m)
}

var _OpenRcaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xds.service.orca.v3.OpenRcaService",
	HandlerType: (*OpenRcaServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCoreMetrics",
			Handler:       _OpenRcaService_StreamCoreMetrics_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "xds/service/orca/v3/orca.proto",
}
