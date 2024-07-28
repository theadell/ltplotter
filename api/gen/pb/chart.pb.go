// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: chart.proto

package pb

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

type PlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X        []float32 `protobuf:"fixed32,1,rep,packed,name=x,proto3" json:"x,omitempty"`
	Y        []*YData  `protobuf:"bytes,2,rep,name=y,proto3" json:"y,omitempty"`
	Metadata *Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *PlotRequest) Reset() {
	*x = PlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlotRequest) ProtoMessage() {}

func (x *PlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlotRequest.ProtoReflect.Descriptor instead.
func (*PlotRequest) Descriptor() ([]byte, []int) {
	return file_chart_proto_rawDescGZIP(), []int{0}
}

func (x *PlotRequest) GetX() []float32 {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *PlotRequest) GetY() []*YData {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *PlotRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type YData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []float32 `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *YData) Reset() {
	*x = YData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YData) ProtoMessage() {}

func (x *YData) ProtoReflect() protoreflect.Message {
	mi := &file_chart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YData.ProtoReflect.Descriptor instead.
func (*YData) Descriptor() ([]byte, []int) {
	return file_chart_proto_rawDescGZIP(), []int{1}
}

func (x *YData) GetValues() []float32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Labels  *Labels  `protobuf:"bytes,2,opt,name=labels,proto3" json:"labels,omitempty"`
	Legends []string `protobuf:"bytes,3,rep,name=legends,proto3" json:"legends,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_chart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_chart_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Metadata) GetLabels() *Labels {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Metadata) GetLegends() []string {
	if x != nil {
		return x.Legends
	}
	return nil
}

type Labels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y string `protobuf:"bytes,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Labels) Reset() {
	*x = Labels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_chart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labels.ProtoReflect.Descriptor instead.
func (*Labels) Descriptor() ([]byte, []int) {
	return file_chart_proto_rawDescGZIP(), []int{3}
}

func (x *Labels) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *Labels) GetY() string {
	if x != nil {
		return x.Y
	}
	return ""
}

type PlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdf []byte `protobuf:"bytes,1,opt,name=pdf,proto3" json:"pdf,omitempty"`
}

func (x *PlotResponse) Reset() {
	*x = PlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlotResponse) ProtoMessage() {}

func (x *PlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlotResponse.ProtoReflect.Descriptor instead.
func (*PlotResponse) Descriptor() ([]byte, []int) {
	return file_chart_proto_rawDescGZIP(), []int{4}
}

func (x *PlotResponse) GetPdf() []byte {
	if x != nil {
		return x.Pdf
	}
	return nil
}

var File_chart_proto protoreflect.FileDescriptor

var file_chart_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70,
	0x6c, 0x6f, 0x74, 0x22, 0x62, 0x0a, 0x0b, 0x50, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x01, 0x78,
	0x12, 0x19, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x6c,
	0x6f, 0x74, 0x2e, 0x59, 0x44, 0x61, 0x74, 0x61, 0x52, 0x01, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x6c, 0x6f, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x0a, 0x05, 0x59, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6c, 0x6f,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x06, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x79,
	0x22, 0x20, 0x0a, 0x0c, 0x50, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x64, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70,
	0x64, 0x66, 0x32, 0x44, 0x0a, 0x0b, 0x50, 0x6c, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x6f,
	0x74, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x6f, 0x74, 0x2e, 0x50, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6c, 0x6f, 0x74, 0x2e, 0x50, 0x6c, 0x6f, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x6c, 0x74, 0x70, 0x6c,
	0x6f, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chart_proto_rawDescOnce sync.Once
	file_chart_proto_rawDescData = file_chart_proto_rawDesc
)

func file_chart_proto_rawDescGZIP() []byte {
	file_chart_proto_rawDescOnce.Do(func() {
		file_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_chart_proto_rawDescData)
	})
	return file_chart_proto_rawDescData
}

var file_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chart_proto_goTypes = []interface{}{
	(*PlotRequest)(nil),  // 0: plot.PlotRequest
	(*YData)(nil),        // 1: plot.YData
	(*Metadata)(nil),     // 2: plot.Metadata
	(*Labels)(nil),       // 3: plot.Labels
	(*PlotResponse)(nil), // 4: plot.PlotResponse
}
var file_chart_proto_depIdxs = []int32{
	1, // 0: plot.PlotRequest.y:type_name -> plot.YData
	2, // 1: plot.PlotRequest.metadata:type_name -> plot.Metadata
	3, // 2: plot.Metadata.labels:type_name -> plot.Labels
	0, // 3: plot.PlotService.GeneratePlot:input_type -> plot.PlotRequest
	4, // 4: plot.PlotService.GeneratePlot:output_type -> plot.PlotResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chart_proto_init() }
func file_chart_proto_init() {
	if File_chart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlotRequest); i {
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
		file_chart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YData); i {
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
		file_chart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_chart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labels); i {
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
		file_chart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlotResponse); i {
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
			RawDescriptor: file_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chart_proto_goTypes,
		DependencyIndexes: file_chart_proto_depIdxs,
		MessageInfos:      file_chart_proto_msgTypes,
	}.Build()
	File_chart_proto = out.File
	file_chart_proto_rawDesc = nil
	file_chart_proto_goTypes = nil
	file_chart_proto_depIdxs = nil
}
