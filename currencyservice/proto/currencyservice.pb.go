// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: proto/currencyservice.proto

package microshopping

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencyservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_currencyservice_proto_rawDescGZIP(), []int{0}
}

// 当前币种描述
type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 货币code 例如：EUR 欧元 USD 美元
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// 货币单位
	//
	//	例如，如果currencyCode是USD，则1个单位是1美元。
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// 例如，$-1.75表示为 单位＝-1 和 纳米＝-75000000。
	Nanos int32 `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencyservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_proto_currencyservice_proto_rawDescGZIP(), []int{1}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Money) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

// 获得支持的币种响应消息
type GetSupportedCurrenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回字符串切片
	CurrencyCodes []string `protobuf:"bytes,1,rep,name=currency_codes,json=currencyCodes,proto3" json:"currency_codes,omitempty"`
}

func (x *GetSupportedCurrenciesResponse) Reset() {
	*x = GetSupportedCurrenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencyservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedCurrenciesResponse) ProtoMessage() {}

func (x *GetSupportedCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetSupportedCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_proto_currencyservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetSupportedCurrenciesResponse) GetCurrencyCodes() []string {
	if x != nil {
		return x.CurrencyCodes
	}
	return nil
}

// 货币转换请求消息
type CurrencyConversionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   *Money `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ToCode string `protobuf:"bytes,2,opt,name=to_code,json=toCode,proto3" json:"to_code,omitempty"`
}

func (x *CurrencyConversionRequest) Reset() {
	*x = CurrencyConversionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencyservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyConversionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyConversionRequest) ProtoMessage() {}

func (x *CurrencyConversionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyConversionRequest.ProtoReflect.Descriptor instead.
func (*CurrencyConversionRequest) Descriptor() ([]byte, []int) {
	return file_proto_currencyservice_proto_rawDescGZIP(), []int{3}
}

func (x *CurrencyConversionRequest) GetFrom() *Money {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *CurrencyConversionRequest) GetToCode() string {
	if x != nil {
		return x.ToCode
	}
	return ""
}

var File_proto_currencyservice_proto protoreflect.FileDescriptor

var file_proto_currencyservice_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x58, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e,
	0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22,
	0x47, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x19, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xbf, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2d, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x28, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currencyservice_proto_rawDescOnce sync.Once
	file_proto_currencyservice_proto_rawDescData = file_proto_currencyservice_proto_rawDesc
)

func file_proto_currencyservice_proto_rawDescGZIP() []byte {
	file_proto_currencyservice_proto_rawDescOnce.Do(func() {
		file_proto_currencyservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currencyservice_proto_rawDescData)
	})
	return file_proto_currencyservice_proto_rawDescData
}

var file_proto_currencyservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_currencyservice_proto_goTypes = []interface{}{
	(*Empty)(nil),                          // 0: microshopping.Empty
	(*Money)(nil),                          // 1: microshopping.Money
	(*GetSupportedCurrenciesResponse)(nil), // 2: microshopping.GetSupportedCurrenciesResponse
	(*CurrencyConversionRequest)(nil),      // 3: microshopping.CurrencyConversionRequest
}
var file_proto_currencyservice_proto_depIdxs = []int32{
	1, // 0: microshopping.CurrencyConversionRequest.from:type_name -> microshopping.Money
	0, // 1: microshopping.CurrencyService.GetSupportedCurrencies:input_type -> microshopping.Empty
	3, // 2: microshopping.CurrencyService.Convert:input_type -> microshopping.CurrencyConversionRequest
	2, // 3: microshopping.CurrencyService.GetSupportedCurrencies:output_type -> microshopping.GetSupportedCurrenciesResponse
	1, // 4: microshopping.CurrencyService.Convert:output_type -> microshopping.Money
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_currencyservice_proto_init() }
func file_proto_currencyservice_proto_init() {
	if File_proto_currencyservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_currencyservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_currencyservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
		file_proto_currencyservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupportedCurrenciesResponse); i {
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
		file_proto_currencyservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyConversionRequest); i {
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
			RawDescriptor: file_proto_currencyservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currencyservice_proto_goTypes,
		DependencyIndexes: file_proto_currencyservice_proto_depIdxs,
		MessageInfos:      file_proto_currencyservice_proto_msgTypes,
	}.Build()
	File_proto_currencyservice_proto = out.File
	file_proto_currencyservice_proto_rawDesc = nil
	file_proto_currencyservice_proto_goTypes = nil
	file_proto_currencyservice_proto_depIdxs = nil
}
