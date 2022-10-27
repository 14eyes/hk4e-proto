// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: cmd/cmd_aranara_collection.proto

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

type AranaraCollectionDataNotify_CmdId int32

const (
	AranaraCollectionDataNotify_NONE             AranaraCollectionDataNotify_CmdId = 0
	AranaraCollectionDataNotify_CMD_ID           AranaraCollectionDataNotify_CmdId = 6376
	AranaraCollectionDataNotify_ENET_CHANNEL_ID  AranaraCollectionDataNotify_CmdId = 0
	AranaraCollectionDataNotify_ENET_IS_RELIABLE AranaraCollectionDataNotify_CmdId = 1
)

// Enum value maps for AranaraCollectionDataNotify_CmdId.
var (
	AranaraCollectionDataNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		6376: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	AranaraCollectionDataNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           6376,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x AranaraCollectionDataNotify_CmdId) Enum() *AranaraCollectionDataNotify_CmdId {
	p := new(AranaraCollectionDataNotify_CmdId)
	*p = x
	return p
}

func (x AranaraCollectionDataNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AranaraCollectionDataNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_aranara_collection_proto_enumTypes[0].Descriptor()
}

func (AranaraCollectionDataNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_aranara_collection_proto_enumTypes[0]
}

func (x AranaraCollectionDataNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AranaraCollectionDataNotify_CmdId.Descriptor instead.
func (AranaraCollectionDataNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{1, 0}
}

type AddAranaraCollectionNotify_CmdId int32

const (
	AddAranaraCollectionNotify_NONE             AddAranaraCollectionNotify_CmdId = 0
	AddAranaraCollectionNotify_CMD_ID           AddAranaraCollectionNotify_CmdId = 6368
	AddAranaraCollectionNotify_ENET_CHANNEL_ID  AddAranaraCollectionNotify_CmdId = 0
	AddAranaraCollectionNotify_ENET_IS_RELIABLE AddAranaraCollectionNotify_CmdId = 1
)

// Enum value maps for AddAranaraCollectionNotify_CmdId.
var (
	AddAranaraCollectionNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		6368: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	AddAranaraCollectionNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           6368,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x AddAranaraCollectionNotify_CmdId) Enum() *AddAranaraCollectionNotify_CmdId {
	p := new(AddAranaraCollectionNotify_CmdId)
	*p = x
	return p
}

func (x AddAranaraCollectionNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddAranaraCollectionNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_aranara_collection_proto_enumTypes[1].Descriptor()
}

func (AddAranaraCollectionNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_aranara_collection_proto_enumTypes[1]
}

func (x AddAranaraCollectionNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddAranaraCollectionNotify_CmdId.Descriptor instead.
func (AddAranaraCollectionNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{2, 0}
}

type CataLogFinishedGlobalWatcherAllDataNotify_CmdId int32

const (
	CataLogFinishedGlobalWatcherAllDataNotify_NONE             CataLogFinishedGlobalWatcherAllDataNotify_CmdId = 0
	CataLogFinishedGlobalWatcherAllDataNotify_CMD_ID           CataLogFinishedGlobalWatcherAllDataNotify_CmdId = 6370
	CataLogFinishedGlobalWatcherAllDataNotify_ENET_CHANNEL_ID  CataLogFinishedGlobalWatcherAllDataNotify_CmdId = 0
	CataLogFinishedGlobalWatcherAllDataNotify_ENET_IS_RELIABLE CataLogFinishedGlobalWatcherAllDataNotify_CmdId = 1
)

// Enum value maps for CataLogFinishedGlobalWatcherAllDataNotify_CmdId.
var (
	CataLogFinishedGlobalWatcherAllDataNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		6370: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	CataLogFinishedGlobalWatcherAllDataNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           6370,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x CataLogFinishedGlobalWatcherAllDataNotify_CmdId) Enum() *CataLogFinishedGlobalWatcherAllDataNotify_CmdId {
	p := new(CataLogFinishedGlobalWatcherAllDataNotify_CmdId)
	*p = x
	return p
}

func (x CataLogFinishedGlobalWatcherAllDataNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CataLogFinishedGlobalWatcherAllDataNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_aranara_collection_proto_enumTypes[2].Descriptor()
}

func (CataLogFinishedGlobalWatcherAllDataNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_aranara_collection_proto_enumTypes[2]
}

func (x CataLogFinishedGlobalWatcherAllDataNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CataLogFinishedGlobalWatcherAllDataNotify_CmdId.Descriptor instead.
func (CataLogFinishedGlobalWatcherAllDataNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{4, 0}
}

type CataLogNewFinishedGlobalWatcherNotify_CmdId int32

const (
	CataLogNewFinishedGlobalWatcherNotify_NONE             CataLogNewFinishedGlobalWatcherNotify_CmdId = 0
	CataLogNewFinishedGlobalWatcherNotify_CMD_ID           CataLogNewFinishedGlobalWatcherNotify_CmdId = 6395
	CataLogNewFinishedGlobalWatcherNotify_ENET_CHANNEL_ID  CataLogNewFinishedGlobalWatcherNotify_CmdId = 0
	CataLogNewFinishedGlobalWatcherNotify_ENET_IS_RELIABLE CataLogNewFinishedGlobalWatcherNotify_CmdId = 1
)

// Enum value maps for CataLogNewFinishedGlobalWatcherNotify_CmdId.
var (
	CataLogNewFinishedGlobalWatcherNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		6395: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	CataLogNewFinishedGlobalWatcherNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           6395,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x CataLogNewFinishedGlobalWatcherNotify_CmdId) Enum() *CataLogNewFinishedGlobalWatcherNotify_CmdId {
	p := new(CataLogNewFinishedGlobalWatcherNotify_CmdId)
	*p = x
	return p
}

func (x CataLogNewFinishedGlobalWatcherNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CataLogNewFinishedGlobalWatcherNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_aranara_collection_proto_enumTypes[3].Descriptor()
}

func (CataLogNewFinishedGlobalWatcherNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_aranara_collection_proto_enumTypes[3]
}

func (x CataLogNewFinishedGlobalWatcherNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CataLogNewFinishedGlobalWatcherNotify_CmdId.Descriptor instead.
func (CataLogNewFinishedGlobalWatcherNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{5, 0}
}

type AranaraCollectionSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionIdStateMap map[uint32]AranaraCollectionState `protobuf:"bytes,6,rep,name=collection_id_state_map,json=collectionIdStateMap,proto3" json:"collection_id_state_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=proto.AranaraCollectionState"`
	CollectionType       uint32                            `protobuf:"varint,12,opt,name=collection_type,json=collectionType,proto3" json:"collection_type,omitempty"`
}

func (x *AranaraCollectionSuite) Reset() {
	*x = AranaraCollectionSuite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AranaraCollectionSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AranaraCollectionSuite) ProtoMessage() {}

func (x *AranaraCollectionSuite) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AranaraCollectionSuite.ProtoReflect.Descriptor instead.
func (*AranaraCollectionSuite) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{0}
}

func (x *AranaraCollectionSuite) GetCollectionIdStateMap() map[uint32]AranaraCollectionState {
	if x != nil {
		return x.CollectionIdStateMap
	}
	return nil
}

func (x *AranaraCollectionSuite) GetCollectionType() uint32 {
	if x != nil {
		return x.CollectionType
	}
	return 0
}

type AranaraCollectionDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionSuiteList []*AranaraCollectionSuite `protobuf:"bytes,14,rep,name=collection_suite_list,json=collectionSuiteList,proto3" json:"collection_suite_list,omitempty"`
}

func (x *AranaraCollectionDataNotify) Reset() {
	*x = AranaraCollectionDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AranaraCollectionDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AranaraCollectionDataNotify) ProtoMessage() {}

func (x *AranaraCollectionDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AranaraCollectionDataNotify.ProtoReflect.Descriptor instead.
func (*AranaraCollectionDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{1}
}

func (x *AranaraCollectionDataNotify) GetCollectionSuiteList() []*AranaraCollectionSuite {
	if x != nil {
		return x.CollectionSuiteList
	}
	return nil
}

type AddAranaraCollectionNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionType uint32                 `protobuf:"varint,7,opt,name=collection_type,json=collectionType,proto3" json:"collection_type,omitempty"`
	TargetState    AranaraCollectionState `protobuf:"varint,12,opt,name=target_state,json=targetState,proto3,enum=proto.AranaraCollectionState" json:"target_state,omitempty"`
	FromState      AranaraCollectionState `protobuf:"varint,15,opt,name=from_state,json=fromState,proto3,enum=proto.AranaraCollectionState" json:"from_state,omitempty"`
	CollectionId   uint32                 `protobuf:"varint,8,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
}

func (x *AddAranaraCollectionNotify) Reset() {
	*x = AddAranaraCollectionNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAranaraCollectionNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAranaraCollectionNotify) ProtoMessage() {}

func (x *AddAranaraCollectionNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAranaraCollectionNotify.ProtoReflect.Descriptor instead.
func (*AddAranaraCollectionNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{2}
}

func (x *AddAranaraCollectionNotify) GetCollectionType() uint32 {
	if x != nil {
		return x.CollectionType
	}
	return 0
}

func (x *AddAranaraCollectionNotify) GetTargetState() AranaraCollectionState {
	if x != nil {
		return x.TargetState
	}
	return AranaraCollectionState_ARANARA_COLLECTION_STATE_NONE
}

func (x *AddAranaraCollectionNotify) GetFromState() AranaraCollectionState {
	if x != nil {
		return x.FromState
	}
	return AranaraCollectionState_ARANARA_COLLECTION_STATE_NONE
}

func (x *AddAranaraCollectionNotify) GetCollectionId() uint32 {
	if x != nil {
		return x.CollectionId
	}
	return 0
}

type CataLogGlobalWatcherFinishedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishedGlobalWatcherList []uint32 `protobuf:"varint,8,rep,packed,name=finished_global_watcher_list,json=finishedGlobalWatcherList,proto3" json:"finished_global_watcher_list,omitempty"`
	CatalogType               uint32   `protobuf:"varint,13,opt,name=catalog_type,json=catalogType,proto3" json:"catalog_type,omitempty"`
}

func (x *CataLogGlobalWatcherFinishedData) Reset() {
	*x = CataLogGlobalWatcherFinishedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CataLogGlobalWatcherFinishedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CataLogGlobalWatcherFinishedData) ProtoMessage() {}

func (x *CataLogGlobalWatcherFinishedData) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CataLogGlobalWatcherFinishedData.ProtoReflect.Descriptor instead.
func (*CataLogGlobalWatcherFinishedData) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{3}
}

func (x *CataLogGlobalWatcherFinishedData) GetFinishedGlobalWatcherList() []uint32 {
	if x != nil {
		return x.FinishedGlobalWatcherList
	}
	return nil
}

func (x *CataLogGlobalWatcherFinishedData) GetCatalogType() uint32 {
	if x != nil {
		return x.CatalogType
	}
	return 0
}

type CataLogFinishedGlobalWatcherAllDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishedGlobalWatcherDataList []*CataLogGlobalWatcherFinishedData `protobuf:"bytes,13,rep,name=finished_global_watcher_data_list,json=finishedGlobalWatcherDataList,proto3" json:"finished_global_watcher_data_list,omitempty"`
}

func (x *CataLogFinishedGlobalWatcherAllDataNotify) Reset() {
	*x = CataLogFinishedGlobalWatcherAllDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CataLogFinishedGlobalWatcherAllDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CataLogFinishedGlobalWatcherAllDataNotify) ProtoMessage() {}

func (x *CataLogFinishedGlobalWatcherAllDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CataLogFinishedGlobalWatcherAllDataNotify.ProtoReflect.Descriptor instead.
func (*CataLogFinishedGlobalWatcherAllDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{4}
}

func (x *CataLogFinishedGlobalWatcherAllDataNotify) GetFinishedGlobalWatcherDataList() []*CataLogGlobalWatcherFinishedData {
	if x != nil {
		return x.FinishedGlobalWatcherDataList
	}
	return nil
}

type CataLogNewFinishedGlobalWatcherNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewFinishedGlobalWatcherDataList []*CataLogGlobalWatcherFinishedData `protobuf:"bytes,2,rep,name=new_finished_global_watcher_data_list,json=newFinishedGlobalWatcherDataList,proto3" json:"new_finished_global_watcher_data_list,omitempty"`
}

func (x *CataLogNewFinishedGlobalWatcherNotify) Reset() {
	*x = CataLogNewFinishedGlobalWatcherNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CataLogNewFinishedGlobalWatcherNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CataLogNewFinishedGlobalWatcherNotify) ProtoMessage() {}

func (x *CataLogNewFinishedGlobalWatcherNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_aranara_collection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CataLogNewFinishedGlobalWatcherNotify.ProtoReflect.Descriptor instead.
func (*CataLogNewFinishedGlobalWatcherNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_aranara_collection_proto_rawDescGZIP(), []int{5}
}

func (x *CataLogNewFinishedGlobalWatcherNotify) GetNewFinishedGlobalWatcherDataList() []*CataLogGlobalWatcherFinishedData {
	if x != nil {
		return x.NewFinishedGlobalWatcherDataList
	}
	return nil
}

var File_cmd_cmd_aranara_collection_proto protoreflect.FileDescriptor

var file_cmd_cmd_aranara_collection_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x61, 0x72, 0x61, 0x6e, 0x61, 0x72,
	0x61, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x16, 0x41, 0x72, 0x61, 0x6e,
	0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x12, 0x6e, 0x0a, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x61, 0x6e,
	0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x66, 0x0a, 0x19, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xbf, 0x01, 0x0a, 0x1b, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x51, 0x0a, 0x15, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x61, 0x6e, 0x61,
	0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74,
	0x65, 0x52, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xe8, 0x31, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0xb9, 0x02, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x41, 0x72, 0x61,
	0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a,
	0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x61, 0x6e,
	0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x3c, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x61, 0x6e,
	0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xe0, 0x31, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x22, 0x86, 0x01, 0x0a, 0x20, 0x43, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x1c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x19, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x29, 0x43,
	0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x71, 0x0a, 0x21, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x61,
	0x4c, 0x6f, 0x67, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x1d, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xe2, 0x31, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0xf0, 0x01, 0x0a, 0x25, 0x43,
	0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x78, 0x0a, 0x25, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x61,
	0x4c, 0x6f, 0x67, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x20, 0x6e, 0x65,
	0x77, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xfb, 0x31, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76,
	0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_aranara_collection_proto_rawDescOnce sync.Once
	file_cmd_cmd_aranara_collection_proto_rawDescData = file_cmd_cmd_aranara_collection_proto_rawDesc
)

func file_cmd_cmd_aranara_collection_proto_rawDescGZIP() []byte {
	file_cmd_cmd_aranara_collection_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_aranara_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_aranara_collection_proto_rawDescData)
	})
	return file_cmd_cmd_aranara_collection_proto_rawDescData
}

var file_cmd_cmd_aranara_collection_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_cmd_cmd_aranara_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cmd_cmd_aranara_collection_proto_goTypes = []interface{}{
	(AranaraCollectionDataNotify_CmdId)(0),               // 0: proto.AranaraCollectionDataNotify.CmdId
	(AddAranaraCollectionNotify_CmdId)(0),                // 1: proto.AddAranaraCollectionNotify.CmdId
	(CataLogFinishedGlobalWatcherAllDataNotify_CmdId)(0), // 2: proto.CataLogFinishedGlobalWatcherAllDataNotify.CmdId
	(CataLogNewFinishedGlobalWatcherNotify_CmdId)(0),     // 3: proto.CataLogNewFinishedGlobalWatcherNotify.CmdId
	(*AranaraCollectionSuite)(nil),                       // 4: proto.AranaraCollectionSuite
	(*AranaraCollectionDataNotify)(nil),                  // 5: proto.AranaraCollectionDataNotify
	(*AddAranaraCollectionNotify)(nil),                   // 6: proto.AddAranaraCollectionNotify
	(*CataLogGlobalWatcherFinishedData)(nil),             // 7: proto.CataLogGlobalWatcherFinishedData
	(*CataLogFinishedGlobalWatcherAllDataNotify)(nil),    // 8: proto.CataLogFinishedGlobalWatcherAllDataNotify
	(*CataLogNewFinishedGlobalWatcherNotify)(nil),        // 9: proto.CataLogNewFinishedGlobalWatcherNotify
	nil,                         // 10: proto.AranaraCollectionSuite.CollectionIdStateMapEntry
	(AranaraCollectionState)(0), // 11: proto.AranaraCollectionState
}
var file_cmd_cmd_aranara_collection_proto_depIdxs = []int32{
	10, // 0: proto.AranaraCollectionSuite.collection_id_state_map:type_name -> proto.AranaraCollectionSuite.CollectionIdStateMapEntry
	4,  // 1: proto.AranaraCollectionDataNotify.collection_suite_list:type_name -> proto.AranaraCollectionSuite
	11, // 2: proto.AddAranaraCollectionNotify.target_state:type_name -> proto.AranaraCollectionState
	11, // 3: proto.AddAranaraCollectionNotify.from_state:type_name -> proto.AranaraCollectionState
	7,  // 4: proto.CataLogFinishedGlobalWatcherAllDataNotify.finished_global_watcher_data_list:type_name -> proto.CataLogGlobalWatcherFinishedData
	7,  // 5: proto.CataLogNewFinishedGlobalWatcherNotify.new_finished_global_watcher_data_list:type_name -> proto.CataLogGlobalWatcherFinishedData
	11, // 6: proto.AranaraCollectionSuite.CollectionIdStateMapEntry.value:type_name -> proto.AranaraCollectionState
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_cmd_cmd_aranara_collection_proto_init() }
func file_cmd_cmd_aranara_collection_proto_init() {
	if File_cmd_cmd_aranara_collection_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_cmd_aranara_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AranaraCollectionSuite); i {
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
		file_cmd_cmd_aranara_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AranaraCollectionDataNotify); i {
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
		file_cmd_cmd_aranara_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAranaraCollectionNotify); i {
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
		file_cmd_cmd_aranara_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CataLogGlobalWatcherFinishedData); i {
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
		file_cmd_cmd_aranara_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CataLogFinishedGlobalWatcherAllDataNotify); i {
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
		file_cmd_cmd_aranara_collection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CataLogNewFinishedGlobalWatcherNotify); i {
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
			RawDescriptor: file_cmd_cmd_aranara_collection_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_aranara_collection_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_aranara_collection_proto_depIdxs,
		EnumInfos:         file_cmd_cmd_aranara_collection_proto_enumTypes,
		MessageInfos:      file_cmd_cmd_aranara_collection_proto_msgTypes,
	}.Build()
	File_cmd_cmd_aranara_collection_proto = out.File
	file_cmd_cmd_aranara_collection_proto_rawDesc = nil
	file_cmd_cmd_aranara_collection_proto_goTypes = nil
	file_cmd_cmd_aranara_collection_proto_depIdxs = nil
}
