// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: wallet.proto

package __

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

type CreateWalletData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletName string `protobuf:"bytes,1,opt,name=walletName,proto3" json:"walletName,omitempty"`
}

func (x *CreateWalletData) Reset() {
	*x = CreateWalletData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletData) ProtoMessage() {}

func (x *CreateWalletData) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletData.ProtoReflect.Descriptor instead.
func (*CreateWalletData) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWalletData) GetWalletName() string {
	if x != nil {
		return x.WalletName
	}
	return ""
}

type CreateMnemonicData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MnemonicLength string `protobuf:"bytes,1,opt,name=mnemonicLength,proto3" json:"mnemonicLength,omitempty"`
	Language       string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *CreateMnemonicData) Reset() {
	*x = CreateMnemonicData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMnemonicData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMnemonicData) ProtoMessage() {}

func (x *CreateMnemonicData) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMnemonicData.ProtoReflect.Descriptor instead.
func (*CreateMnemonicData) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMnemonicData) GetMnemonicLength() string {
	if x != nil {
		return x.MnemonicLength
	}
	return ""
}

func (x *CreateMnemonicData) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic string `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	CoinName string `protobuf:"bytes,2,opt,name=coinName,proto3" json:"coinName,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Private  string `protobuf:"bytes,4,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *Wallet) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *Wallet) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Wallet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Wallet) GetPrivate() string {
	if x != nil {
		return x.Private
	}
	return ""
}

type WalletInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet *Wallet `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
}

func (x *WalletInfo) Reset() {
	*x = WalletInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletInfo) ProtoMessage() {}

func (x *WalletInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletInfo.ProtoReflect.Descriptor instead.
func (*WalletInfo) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *WalletInfo) GetWallet() *Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type MnemonicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic string `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
}

func (x *MnemonicInfo) Reset() {
	*x = MnemonicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MnemonicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MnemonicInfo) ProtoMessage() {}

func (x *MnemonicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MnemonicInfo.ProtoReflect.Descriptor instead.
func (*MnemonicInfo) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *MnemonicInfo) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

var File_wallet_proto protoreflect.FileDescriptor

var file_wallet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x58, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x06,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e,
	0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x22, 0x2d, 0x0a, 0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x22, 0x2a, 0x0a, 0x0c, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x32, 0x79, 0x0a,
	0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x11,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x0b, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e,
	0x69, 0x63, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e,
	0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData = file_wallet_proto_rawDesc
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_proto_rawDescData)
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wallet_proto_goTypes = []interface{}{
	(*CreateWalletData)(nil),   // 0: CreateWalletData
	(*CreateMnemonicData)(nil), // 1: CreateMnemonicData
	(*Wallet)(nil),             // 2: Wallet
	(*WalletInfo)(nil),         // 3: WalletInfo
	(*MnemonicInfo)(nil),       // 4: MnemonicInfo
}
var file_wallet_proto_depIdxs = []int32{
	2, // 0: WalletInfo.wallet:type_name -> Wallet
	0, // 1: WalletService.CreateWallet:input_type -> CreateWalletData
	1, // 2: WalletService.CreateMnemonic:input_type -> CreateMnemonicData
	3, // 3: WalletService.CreateWallet:output_type -> WalletInfo
	4, // 4: WalletService.CreateMnemonic:output_type -> MnemonicInfo
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletData); i {
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
		file_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMnemonicData); i {
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
		file_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
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
		file_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletInfo); i {
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
		file_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MnemonicInfo); i {
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
			RawDescriptor: file_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_rawDesc = nil
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}
