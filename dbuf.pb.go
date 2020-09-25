// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dbuf.proto

package dbuf

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

type GetQueueStateResponse_QueuesState int32

const (
	GetQueueStateResponse_QUEUE_STATE_INVALID     GetQueueStateResponse_QueuesState = 0
	GetQueueStateResponse_QUEUE_STATE_BUFFERING   GetQueueStateResponse_QueuesState = 1
	GetQueueStateResponse_QUEUE_STATE_DRAINING    GetQueueStateResponse_QueuesState = 2
	GetQueueStateResponse_QUEUE_STATE_PASSTHROUGH GetQueueStateResponse_QueuesState = 3
)

var GetQueueStateResponse_QueuesState_name = map[int32]string{
	0: "QUEUE_STATE_INVALID",
	1: "QUEUE_STATE_BUFFERING",
	2: "QUEUE_STATE_DRAINING",
	3: "QUEUE_STATE_PASSTHROUGH",
}

var GetQueueStateResponse_QueuesState_value = map[string]int32{
	"QUEUE_STATE_INVALID":     0,
	"QUEUE_STATE_BUFFERING":   1,
	"QUEUE_STATE_DRAINING":    2,
	"QUEUE_STATE_PASSTHROUGH": 3,
}

func (x GetQueueStateResponse_QueuesState) String() string {
	return proto.EnumName(GetQueueStateResponse_QueuesState_name, int32(x))
}

func (GetQueueStateResponse_QueuesState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{4, 0}
}

type ModifyQueueRequest_QueueAction int32

const (
	ModifyQueueRequest_QUEUE_ACTION_INVALID                 ModifyQueueRequest_QueueAction = 0
	ModifyQueueRequest_QUEUE_ACTION_RELEASE                 ModifyQueueRequest_QueueAction = 1
	ModifyQueueRequest_QUEUE_ACTION_RELEASE_AND_PASSTHROUGH ModifyQueueRequest_QueueAction = 2
	ModifyQueueRequest_QUEUE_ACTION_DROP                    ModifyQueueRequest_QueueAction = 3
	ModifyQueueRequest_QUEUE_ACTION_RESIZE                  ModifyQueueRequest_QueueAction = 4
)

var ModifyQueueRequest_QueueAction_name = map[int32]string{
	0: "QUEUE_ACTION_INVALID",
	1: "QUEUE_ACTION_RELEASE",
	2: "QUEUE_ACTION_RELEASE_AND_PASSTHROUGH",
	3: "QUEUE_ACTION_DROP",
	4: "QUEUE_ACTION_RESIZE",
}

var ModifyQueueRequest_QueueAction_value = map[string]int32{
	"QUEUE_ACTION_INVALID":                 0,
	"QUEUE_ACTION_RELEASE":                 1,
	"QUEUE_ACTION_RELEASE_AND_PASSTHROUGH": 2,
	"QUEUE_ACTION_DROP":                    3,
	"QUEUE_ACTION_RESIZE":                  4,
}

func (x ModifyQueueRequest_QueueAction) String() string {
	return proto.EnumName(ModifyQueueRequest_QueueAction_name, int32(x))
}

func (ModifyQueueRequest_QueueAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{5, 0}
}

// TODO(max): move from cli flags to config file?
type DbufConfig struct {
	DataPlaneUrls         []string `protobuf:"bytes,1,rep,name=data_plane_urls,json=dataPlaneUrls,proto3" json:"data_plane_urls,omitempty"`
	MaximumNumberOfQueues uint64   `protobuf:"varint,2,opt,name=maximum_number_of_queues,json=maximumNumberOfQueues,proto3" json:"maximum_number_of_queues,omitempty"`
	MaximumBufferedBytes  uint64   `protobuf:"varint,3,opt,name=maximum_buffered_bytes,json=maximumBufferedBytes,proto3" json:"maximum_buffered_bytes,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *DbufConfig) Reset()         { *m = DbufConfig{} }
func (m *DbufConfig) String() string { return proto.CompactTextString(m) }
func (*DbufConfig) ProtoMessage()    {}
func (*DbufConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{0}
}

func (m *DbufConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DbufConfig.Unmarshal(m, b)
}
func (m *DbufConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DbufConfig.Marshal(b, m, deterministic)
}
func (m *DbufConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DbufConfig.Merge(m, src)
}
func (m *DbufConfig) XXX_Size() int {
	return xxx_messageInfo_DbufConfig.Size(m)
}
func (m *DbufConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DbufConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DbufConfig proto.InternalMessageInfo

func (m *DbufConfig) GetDataPlaneUrls() []string {
	if m != nil {
		return m.DataPlaneUrls
	}
	return nil
}

func (m *DbufConfig) GetMaximumNumberOfQueues() uint64 {
	if m != nil {
		return m.MaximumNumberOfQueues
	}
	return 0
}

func (m *DbufConfig) GetMaximumBufferedBytes() uint64 {
	if m != nil {
		return m.MaximumBufferedBytes
	}
	return 0
}

type GetDbufStateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDbufStateRequest) Reset()         { *m = GetDbufStateRequest{} }
func (m *GetDbufStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetDbufStateRequest) ProtoMessage()    {}
func (*GetDbufStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{1}
}

func (m *GetDbufStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDbufStateRequest.Unmarshal(m, b)
}
func (m *GetDbufStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDbufStateRequest.Marshal(b, m, deterministic)
}
func (m *GetDbufStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDbufStateRequest.Merge(m, src)
}
func (m *GetDbufStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetDbufStateRequest.Size(m)
}
func (m *GetDbufStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDbufStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDbufStateRequest proto.InternalMessageInfo

type GetDbufStateResponse struct {
	QueueIdLow           uint64   `protobuf:"varint,1,opt,name=queue_id_low,json=queueIdLow,proto3" json:"queue_id_low,omitempty"`
	QueueIdHigh          uint64   `protobuf:"varint,2,opt,name=queue_id_high,json=queueIdHigh,proto3" json:"queue_id_high,omitempty"`
	FreeQueues           uint64   `protobuf:"varint,3,opt,name=free_queues,json=freeQueues,proto3" json:"free_queues,omitempty"`
	MaximumMemory        uint64   `protobuf:"varint,4,opt,name=maximum_memory,json=maximumMemory,proto3" json:"maximum_memory,omitempty"`
	FreeMemory           uint64   `protobuf:"varint,5,opt,name=free_memory,json=freeMemory,proto3" json:"free_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDbufStateResponse) Reset()         { *m = GetDbufStateResponse{} }
func (m *GetDbufStateResponse) String() string { return proto.CompactTextString(m) }
func (*GetDbufStateResponse) ProtoMessage()    {}
func (*GetDbufStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{2}
}

func (m *GetDbufStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDbufStateResponse.Unmarshal(m, b)
}
func (m *GetDbufStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDbufStateResponse.Marshal(b, m, deterministic)
}
func (m *GetDbufStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDbufStateResponse.Merge(m, src)
}
func (m *GetDbufStateResponse) XXX_Size() int {
	return xxx_messageInfo_GetDbufStateResponse.Size(m)
}
func (m *GetDbufStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDbufStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDbufStateResponse proto.InternalMessageInfo

func (m *GetDbufStateResponse) GetQueueIdLow() uint64 {
	if m != nil {
		return m.QueueIdLow
	}
	return 0
}

func (m *GetDbufStateResponse) GetQueueIdHigh() uint64 {
	if m != nil {
		return m.QueueIdHigh
	}
	return 0
}

func (m *GetDbufStateResponse) GetFreeQueues() uint64 {
	if m != nil {
		return m.FreeQueues
	}
	return 0
}

func (m *GetDbufStateResponse) GetMaximumMemory() uint64 {
	if m != nil {
		return m.MaximumMemory
	}
	return 0
}

func (m *GetDbufStateResponse) GetFreeMemory() uint64 {
	if m != nil {
		return m.FreeMemory
	}
	return 0
}

type GetQueueStateRequest struct {
	QueueId              uint64   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQueueStateRequest) Reset()         { *m = GetQueueStateRequest{} }
func (m *GetQueueStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetQueueStateRequest) ProtoMessage()    {}
func (*GetQueueStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{3}
}

func (m *GetQueueStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQueueStateRequest.Unmarshal(m, b)
}
func (m *GetQueueStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQueueStateRequest.Marshal(b, m, deterministic)
}
func (m *GetQueueStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueueStateRequest.Merge(m, src)
}
func (m *GetQueueStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetQueueStateRequest.Size(m)
}
func (m *GetQueueStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueueStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueueStateRequest proto.InternalMessageInfo

func (m *GetQueueStateRequest) GetQueueId() uint64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

type GetQueueStateResponse struct {
	MaximumBuffers       uint64                            `protobuf:"varint,1,opt,name=maximum_buffers,json=maximumBuffers,proto3" json:"maximum_buffers,omitempty"`
	FreeBuffers          uint64                            `protobuf:"varint,2,opt,name=free_buffers,json=freeBuffers,proto3" json:"free_buffers,omitempty"`
	MaximumMemory        uint64                            `protobuf:"varint,3,opt,name=maximum_memory,json=maximumMemory,proto3" json:"maximum_memory,omitempty"`
	FreeMemory           uint64                            `protobuf:"varint,4,opt,name=free_memory,json=freeMemory,proto3" json:"free_memory,omitempty"`
	State                GetQueueStateResponse_QueuesState `protobuf:"varint,5,opt,name=state,proto3,enum=dbuf.GetQueueStateResponse_QueuesState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *GetQueueStateResponse) Reset()         { *m = GetQueueStateResponse{} }
func (m *GetQueueStateResponse) String() string { return proto.CompactTextString(m) }
func (*GetQueueStateResponse) ProtoMessage()    {}
func (*GetQueueStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{4}
}

func (m *GetQueueStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQueueStateResponse.Unmarshal(m, b)
}
func (m *GetQueueStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQueueStateResponse.Marshal(b, m, deterministic)
}
func (m *GetQueueStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueueStateResponse.Merge(m, src)
}
func (m *GetQueueStateResponse) XXX_Size() int {
	return xxx_messageInfo_GetQueueStateResponse.Size(m)
}
func (m *GetQueueStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueueStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueueStateResponse proto.InternalMessageInfo

func (m *GetQueueStateResponse) GetMaximumBuffers() uint64 {
	if m != nil {
		return m.MaximumBuffers
	}
	return 0
}

func (m *GetQueueStateResponse) GetFreeBuffers() uint64 {
	if m != nil {
		return m.FreeBuffers
	}
	return 0
}

func (m *GetQueueStateResponse) GetMaximumMemory() uint64 {
	if m != nil {
		return m.MaximumMemory
	}
	return 0
}

func (m *GetQueueStateResponse) GetFreeMemory() uint64 {
	if m != nil {
		return m.FreeMemory
	}
	return 0
}

func (m *GetQueueStateResponse) GetState() GetQueueStateResponse_QueuesState {
	if m != nil {
		return m.State
	}
	return GetQueueStateResponse_QUEUE_STATE_INVALID
}

type ModifyQueueRequest struct {
	Action               ModifyQueueRequest_QueueAction `protobuf:"varint,1,opt,name=action,proto3,enum=dbuf.ModifyQueueRequest_QueueAction" json:"action,omitempty"`
	QueueId              uint64                         `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	NewQueueSize         uint64                         `protobuf:"varint,3,opt,name=new_queue_size,json=newQueueSize,proto3" json:"new_queue_size,omitempty"`
	DestinationAddress   string                         `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ModifyQueueRequest) Reset()         { *m = ModifyQueueRequest{} }
func (m *ModifyQueueRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyQueueRequest) ProtoMessage()    {}
func (*ModifyQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{5}
}

func (m *ModifyQueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyQueueRequest.Unmarshal(m, b)
}
func (m *ModifyQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyQueueRequest.Marshal(b, m, deterministic)
}
func (m *ModifyQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyQueueRequest.Merge(m, src)
}
func (m *ModifyQueueRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyQueueRequest.Size(m)
}
func (m *ModifyQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyQueueRequest proto.InternalMessageInfo

func (m *ModifyQueueRequest) GetAction() ModifyQueueRequest_QueueAction {
	if m != nil {
		return m.Action
	}
	return ModifyQueueRequest_QUEUE_ACTION_INVALID
}

func (m *ModifyQueueRequest) GetQueueId() uint64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *ModifyQueueRequest) GetNewQueueSize() uint64 {
	if m != nil {
		return m.NewQueueSize
	}
	return 0
}

func (m *ModifyQueueRequest) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

type ModifyQueueResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyQueueResponse) Reset()         { *m = ModifyQueueResponse{} }
func (m *ModifyQueueResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyQueueResponse) ProtoMessage()    {}
func (*ModifyQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{6}
}

func (m *ModifyQueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyQueueResponse.Unmarshal(m, b)
}
func (m *ModifyQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyQueueResponse.Marshal(b, m, deterministic)
}
func (m *ModifyQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyQueueResponse.Merge(m, src)
}
func (m *ModifyQueueResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyQueueResponse.Size(m)
}
func (m *ModifyQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyQueueResponse proto.InternalMessageInfo

type SubscribeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{7}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

type Notification struct {
	// Types that are valid to be assigned to MessageType:
	//	*Notification_Ready_
	//	*Notification_FirstBuffer_
	//	*Notification_DroppedPacket_
	MessageType          isNotification_MessageType `protobuf_oneof:"message_type"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{8}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

type isNotification_MessageType interface {
	isNotification_MessageType()
}

type Notification_Ready_ struct {
	Ready *Notification_Ready `protobuf:"bytes,1,opt,name=ready,proto3,oneof"`
}

type Notification_FirstBuffer_ struct {
	FirstBuffer *Notification_FirstBuffer `protobuf:"bytes,2,opt,name=first_buffer,json=firstBuffer,proto3,oneof"`
}

type Notification_DroppedPacket_ struct {
	DroppedPacket *Notification_DroppedPacket `protobuf:"bytes,3,opt,name=dropped_packet,json=droppedPacket,proto3,oneof"`
}

func (*Notification_Ready_) isNotification_MessageType() {}

func (*Notification_FirstBuffer_) isNotification_MessageType() {}

func (*Notification_DroppedPacket_) isNotification_MessageType() {}

func (m *Notification) GetMessageType() isNotification_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *Notification) GetReady() *Notification_Ready {
	if x, ok := m.GetMessageType().(*Notification_Ready_); ok {
		return x.Ready
	}
	return nil
}

func (m *Notification) GetFirstBuffer() *Notification_FirstBuffer {
	if x, ok := m.GetMessageType().(*Notification_FirstBuffer_); ok {
		return x.FirstBuffer
	}
	return nil
}

func (m *Notification) GetDroppedPacket() *Notification_DroppedPacket {
	if x, ok := m.GetMessageType().(*Notification_DroppedPacket_); ok {
		return x.DroppedPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Notification) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Notification_Ready_)(nil),
		(*Notification_FirstBuffer_)(nil),
		(*Notification_DroppedPacket_)(nil),
	}
}

type Notification_Ready struct {
	Config               *DbufConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Notification_Ready) Reset()         { *m = Notification_Ready{} }
func (m *Notification_Ready) String() string { return proto.CompactTextString(m) }
func (*Notification_Ready) ProtoMessage()    {}
func (*Notification_Ready) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{8, 0}
}

func (m *Notification_Ready) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification_Ready.Unmarshal(m, b)
}
func (m *Notification_Ready) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification_Ready.Marshal(b, m, deterministic)
}
func (m *Notification_Ready) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification_Ready.Merge(m, src)
}
func (m *Notification_Ready) XXX_Size() int {
	return xxx_messageInfo_Notification_Ready.Size(m)
}
func (m *Notification_Ready) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification_Ready.DiscardUnknown(m)
}

var xxx_messageInfo_Notification_Ready proto.InternalMessageInfo

func (m *Notification_Ready) GetConfig() *DbufConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type Notification_FirstBuffer struct {
	NewBufferId          uint32   `protobuf:"varint,1,opt,name=new_buffer_id,json=newBufferId,proto3" json:"new_buffer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification_FirstBuffer) Reset()         { *m = Notification_FirstBuffer{} }
func (m *Notification_FirstBuffer) String() string { return proto.CompactTextString(m) }
func (*Notification_FirstBuffer) ProtoMessage()    {}
func (*Notification_FirstBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{8, 1}
}

func (m *Notification_FirstBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification_FirstBuffer.Unmarshal(m, b)
}
func (m *Notification_FirstBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification_FirstBuffer.Marshal(b, m, deterministic)
}
func (m *Notification_FirstBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification_FirstBuffer.Merge(m, src)
}
func (m *Notification_FirstBuffer) XXX_Size() int {
	return xxx_messageInfo_Notification_FirstBuffer.Size(m)
}
func (m *Notification_FirstBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification_FirstBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_Notification_FirstBuffer proto.InternalMessageInfo

func (m *Notification_FirstBuffer) GetNewBufferId() uint32 {
	if m != nil {
		return m.NewBufferId
	}
	return 0
}

type Notification_DroppedPacket struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification_DroppedPacket) Reset()         { *m = Notification_DroppedPacket{} }
func (m *Notification_DroppedPacket) String() string { return proto.CompactTextString(m) }
func (*Notification_DroppedPacket) ProtoMessage()    {}
func (*Notification_DroppedPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b6dc1b0941c986, []int{8, 2}
}

func (m *Notification_DroppedPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification_DroppedPacket.Unmarshal(m, b)
}
func (m *Notification_DroppedPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification_DroppedPacket.Marshal(b, m, deterministic)
}
func (m *Notification_DroppedPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification_DroppedPacket.Merge(m, src)
}
func (m *Notification_DroppedPacket) XXX_Size() int {
	return xxx_messageInfo_Notification_DroppedPacket.Size(m)
}
func (m *Notification_DroppedPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification_DroppedPacket.DiscardUnknown(m)
}

var xxx_messageInfo_Notification_DroppedPacket proto.InternalMessageInfo

func (m *Notification_DroppedPacket) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func init() {
	proto.RegisterEnum("dbuf.GetQueueStateResponse_QueuesState", GetQueueStateResponse_QueuesState_name, GetQueueStateResponse_QueuesState_value)
	proto.RegisterEnum("dbuf.ModifyQueueRequest_QueueAction", ModifyQueueRequest_QueueAction_name, ModifyQueueRequest_QueueAction_value)
	proto.RegisterType((*DbufConfig)(nil), "dbuf.DbufConfig")
	proto.RegisterType((*GetDbufStateRequest)(nil), "dbuf.GetDbufStateRequest")
	proto.RegisterType((*GetDbufStateResponse)(nil), "dbuf.GetDbufStateResponse")
	proto.RegisterType((*GetQueueStateRequest)(nil), "dbuf.GetQueueStateRequest")
	proto.RegisterType((*GetQueueStateResponse)(nil), "dbuf.GetQueueStateResponse")
	proto.RegisterType((*ModifyQueueRequest)(nil), "dbuf.ModifyQueueRequest")
	proto.RegisterType((*ModifyQueueResponse)(nil), "dbuf.ModifyQueueResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "dbuf.SubscribeRequest")
	proto.RegisterType((*Notification)(nil), "dbuf.Notification")
	proto.RegisterType((*Notification_Ready)(nil), "dbuf.Notification.Ready")
	proto.RegisterType((*Notification_FirstBuffer)(nil), "dbuf.Notification.FirstBuffer")
	proto.RegisterType((*Notification_DroppedPacket)(nil), "dbuf.Notification.DroppedPacket")
}

func init() { proto.RegisterFile("dbuf.proto", fileDescriptor_60b6dc1b0941c986) }

var fileDescriptor_60b6dc1b0941c986 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xe3, 0x54,
	0x10, 0xae, 0x93, 0xb4, 0xd0, 0x71, 0x92, 0x0d, 0xa7, 0xcd, 0xae, 0x93, 0x45, 0x50, 0xac, 0xc2,
	0x56, 0x48, 0xa4, 0xdb, 0x80, 0x84, 0x56, 0x82, 0x0b, 0xa7, 0x71, 0x1b, 0x4b, 0xdd, 0xb4, 0xeb,
	0x34, 0x5c, 0xec, 0x8d, 0xe5, 0x9f, 0xe3, 0xd4, 0x50, 0xff, 0xac, 0x8f, 0x4d, 0x36, 0x7d, 0x15,
	0xc4, 0x15, 0xcf, 0xc2, 0x8b, 0x20, 0x9e, 0x80, 0x27, 0x40, 0xe7, 0x27, 0x89, 0x9d, 0x06, 0xb4,
	0x77, 0xc9, 0x7c, 0xdf, 0x8c, 0xbe, 0x6f, 0x66, 0xce, 0x18, 0xc0, 0x73, 0x72, 0xbf, 0x97, 0xa4,
	0x71, 0x16, 0xa3, 0x1a, 0xfd, 0xad, 0xfe, 0x21, 0x01, 0x0c, 0x9d, 0xdc, 0x3f, 0x8f, 0x23, 0x3f,
	0x98, 0xa1, 0xaf, 0xe0, 0x89, 0x67, 0x67, 0xb6, 0x95, 0xdc, 0xdb, 0x11, 0xb6, 0xf2, 0xf4, 0x9e,
	0x28, 0xd2, 0x51, 0xf5, 0x64, 0xdf, 0x6c, 0xd0, 0xf0, 0x0d, 0x8d, 0x4e, 0xd3, 0x7b, 0x82, 0xbe,
	0x07, 0x25, 0xb4, 0xdf, 0x07, 0x61, 0x1e, 0x5a, 0x51, 0x1e, 0x3a, 0x38, 0xb5, 0x62, 0xdf, 0x7a,
	0x97, 0xe3, 0x1c, 0x13, 0xa5, 0x72, 0x24, 0x9d, 0xd4, 0xcc, 0xb6, 0xc0, 0xc7, 0x0c, 0xbe, 0xf6,
	0xdf, 0x30, 0x10, 0x7d, 0x07, 0x4f, 0x97, 0x89, 0x4e, 0xee, 0xfb, 0x38, 0xc5, 0x9e, 0xe5, 0x2c,
	0x32, 0x4c, 0x94, 0x2a, 0x4b, 0x3b, 0x14, 0xe8, 0x40, 0x80, 0x03, 0x8a, 0xa9, 0x6d, 0x38, 0xb8,
	0xc4, 0x19, 0xd5, 0x39, 0xc9, 0xec, 0x0c, 0x9b, 0xf8, 0x5d, 0x8e, 0x49, 0xa6, 0xfe, 0x29, 0xc1,
	0x61, 0x39, 0x4e, 0x92, 0x38, 0x22, 0x18, 0x1d, 0x41, 0x9d, 0x89, 0xb1, 0x02, 0xcf, 0xba, 0x8f,
	0xe7, 0x8a, 0xc4, 0x6a, 0x03, 0x8b, 0x19, 0xde, 0x55, 0x3c, 0x47, 0x2a, 0x34, 0x56, 0x8c, 0xbb,
	0x60, 0x76, 0x27, 0x54, 0xcb, 0x82, 0x32, 0x0a, 0x66, 0x77, 0xe8, 0x73, 0x90, 0xfd, 0x14, 0xe3,
	0xa5, 0x2f, 0x2e, 0x10, 0x68, 0x48, 0x98, 0xf9, 0x12, 0x9a, 0x4b, 0x33, 0x21, 0x0e, 0xe3, 0x74,
	0xa1, 0xd4, 0x18, 0xa7, 0x21, 0xa2, 0xaf, 0x59, 0x70, 0x55, 0x47, 0x70, 0x76, 0xd7, 0x75, 0x38,
	0x41, 0x3d, 0x63, 0x36, 0x58, 0xd1, 0xa2, 0x3f, 0xd4, 0x81, 0x8f, 0x97, 0x22, 0x85, 0x85, 0x8f,
	0x84, 0x3e, 0xf5, 0xef, 0x0a, 0xb4, 0x37, 0x72, 0x84, 0xf7, 0x17, 0xf0, 0xa4, 0xdc, 0x61, 0x22,
	0x72, 0x9b, 0xa5, 0xd6, 0x12, 0xf4, 0x05, 0xd4, 0x99, 0xac, 0x25, 0x4b, 0x74, 0x80, 0xc6, 0x96,
	0x94, 0xc7, 0x06, 0xab, 0x1f, 0x60, 0xb0, 0xb6, 0x69, 0x10, 0xfd, 0x08, 0xbb, 0x84, 0x8a, 0x64,
	0xde, 0x9b, 0xfd, 0x17, 0x3d, 0xb6, 0x87, 0x5b, 0xf5, 0xf7, 0x78, 0x6f, 0x79, 0x8c, 0x67, 0xa9,
	0xef, 0x41, 0x2e, 0x44, 0xd1, 0x33, 0x38, 0x78, 0x33, 0xd5, 0xa7, 0xba, 0x35, 0xb9, 0xd5, 0x6e,
	0x75, 0xcb, 0x18, 0xff, 0xa4, 0x5d, 0x19, 0xc3, 0xd6, 0x0e, 0xea, 0x40, 0xbb, 0x08, 0x0c, 0xa6,
	0x17, 0x17, 0xba, 0x69, 0x8c, 0x2f, 0x5b, 0x12, 0x52, 0xe0, 0xb0, 0x08, 0x0d, 0x4d, 0xcd, 0x18,
	0x53, 0xa4, 0x82, 0x9e, 0xc3, 0xb3, 0x22, 0x72, 0xa3, 0x4d, 0x26, 0xb7, 0x23, 0xf3, 0x7a, 0x7a,
	0x39, 0x6a, 0x55, 0xd5, 0xbf, 0x2a, 0x80, 0x5e, 0xc7, 0x5e, 0xe0, 0x2f, 0x98, 0x80, 0xe5, 0x60,
	0x7e, 0x80, 0x3d, 0xdb, 0xcd, 0x82, 0x38, 0x62, 0xad, 0x6d, 0xf6, 0x8f, 0xb9, 0xa1, 0xc7, 0x4c,
	0xee, 0x46, 0x63, 0x5c, 0x53, 0xe4, 0x94, 0xc6, 0x5a, 0x29, 0x8d, 0x15, 0x1d, 0x43, 0x33, 0xc2,
	0x73, 0xbe, 0x71, 0x16, 0x09, 0x1e, 0xb0, 0x68, 0x78, 0x3d, 0xc2, 0x73, 0xde, 0xab, 0xe0, 0x01,
	0xa3, 0x53, 0x38, 0xf0, 0x30, 0xc9, 0x82, 0xc8, 0xa6, 0xf5, 0x2c, 0xdb, 0xf3, 0x52, 0x4c, 0x08,
	0xeb, 0xfb, 0xbe, 0x89, 0x0a, 0x90, 0xc6, 0x11, 0xf5, 0x37, 0x49, 0x74, 0x90, 0x2b, 0x59, 0x77,
	0x43, 0x3b, 0xbf, 0x35, 0xae, 0xc7, 0x85, 0x16, 0x6e, 0x22, 0xa6, 0x7e, 0xa5, 0x6b, 0x13, 0xbd,
	0x25, 0xa1, 0x13, 0x38, 0xde, 0x86, 0x58, 0xda, 0x78, 0x58, 0x6a, 0x5a, 0x05, 0xb5, 0xe1, 0x93,
	0x12, 0x73, 0x68, 0x5e, 0xdf, 0xb4, 0xaa, 0xeb, 0xb1, 0xad, 0x0a, 0x4c, 0x8c, 0xb7, 0x7a, 0xab,
	0x46, 0x5f, 0x77, 0xa9, 0x73, 0x7c, 0x11, 0x54, 0x04, 0xad, 0x49, 0xee, 0x10, 0x37, 0x0d, 0x9c,
	0xd5, 0x8b, 0xff, 0xa7, 0x02, 0xf5, 0x71, 0x9c, 0x05, 0x7e, 0xe0, 0x32, 0x83, 0xe8, 0x25, 0xec,
	0xa6, 0xd8, 0xf6, 0x16, 0x6c, 0x10, 0x72, 0x5f, 0xe1, 0x83, 0x28, 0x52, 0x7a, 0x26, 0xc5, 0x47,
	0x3b, 0x26, 0x27, 0xa2, 0x73, 0xa8, 0xfb, 0x41, 0x4a, 0x32, 0xb1, 0xf7, 0x6c, 0x02, 0x72, 0xff,
	0xb3, 0x2d, 0x89, 0x17, 0x94, 0xc6, 0x9f, 0xc2, 0x68, 0xc7, 0x94, 0xfd, 0xf5, 0x5f, 0x64, 0x40,
	0xd3, 0x4b, 0xe3, 0x24, 0xc1, 0x9e, 0x95, 0xd8, 0xee, 0x2f, 0x38, 0x63, 0x73, 0x92, 0xfb, 0x47,
	0x5b, 0xca, 0x0c, 0x39, 0xf1, 0x86, 0xf1, 0x46, 0x3b, 0x66, 0xc3, 0x2b, 0x06, 0xba, 0x67, 0xb0,
	0xcb, 0x14, 0xa2, 0x13, 0xd8, 0x73, 0xd9, 0x15, 0x16, 0x5e, 0x5a, 0xbc, 0xd6, 0xfa, 0x3a, 0x9b,
	0x02, 0xef, 0x9e, 0x81, 0x5c, 0xd0, 0x46, 0x6f, 0x19, 0x5d, 0x1a, 0xee, 0x67, 0x79, 0x2b, 0x1a,
	0xa6, 0x1c, 0xe1, 0x39, 0x67, 0x18, 0x5e, 0xf7, 0x6b, 0x68, 0x94, 0x74, 0x3c, 0xba, 0x2d, 0x8d,
	0xd5, 0x12, 0x0e, 0x9a, 0x50, 0x0f, 0x31, 0x21, 0xf6, 0x0c, 0x5b, 0xd9, 0x22, 0xc1, 0xfd, 0xdf,
	0x2b, 0x20, 0xb3, 0x1b, 0x8b, 0xd3, 0x5f, 0x03, 0x17, 0x23, 0x1d, 0xea, 0xc5, 0xab, 0x8b, 0x3a,
	0xab, 0xe7, 0xbc, 0x79, 0xa1, 0xbb, 0xdd, 0x6d, 0x90, 0x38, 0x54, 0x23, 0x68, 0x94, 0x2e, 0x00,
	0xea, 0x6e, 0x3d, 0x0b, 0xbc, 0xd0, 0xf3, 0xff, 0x39, 0x19, 0x68, 0x00, 0x72, 0x61, 0x81, 0x90,
	0xf2, 0x5f, 0xaf, 0xb1, 0xdb, 0xd9, 0x82, 0x88, 0x1a, 0xaf, 0x60, 0x7f, 0xb5, 0x6d, 0xe8, 0x29,
	0xe7, 0x6d, 0xae, 0x5f, 0x17, 0x3d, 0x1e, 0xef, 0x4b, 0x69, 0xf0, 0x0a, 0x3a, 0x71, 0x3a, 0xeb,
	0xc5, 0x21, 0x76, 0x93, 0x34, 0xfe, 0x19, 0xbb, 0x19, 0xa7, 0xcd, 0xd2, 0xc4, 0x7d, 0xfb, 0xe9,
	0x2c, 0xc8, 0xee, 0x72, 0xa7, 0xe7, 0xc6, 0xe1, 0x29, 0x65, 0x7c, 0x23, 0x28, 0xa7, 0x94, 0xe2,
	0xec, 0xb1, 0x6f, 0xf1, 0xb7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xad, 0x1a, 0x3f, 0xdd, 0x99,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DbufServiceClient is the client API for DbufService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DbufServiceClient interface {
	GetDbufState(ctx context.Context, in *GetDbufStateRequest, opts ...grpc.CallOption) (*GetDbufStateResponse, error)
	GetQueueState(ctx context.Context, in *GetQueueStateRequest, opts ...grpc.CallOption) (*GetQueueStateResponse, error)
	ModifyQueue(ctx context.Context, in *ModifyQueueRequest, opts ...grpc.CallOption) (*ModifyQueueResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DbufService_SubscribeClient, error)
}

type dbufServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDbufServiceClient(cc grpc.ClientConnInterface) DbufServiceClient {
	return &dbufServiceClient{cc}
}

func (c *dbufServiceClient) GetDbufState(ctx context.Context, in *GetDbufStateRequest, opts ...grpc.CallOption) (*GetDbufStateResponse, error) {
	out := new(GetDbufStateResponse)
	err := c.cc.Invoke(ctx, "/dbuf.DbufService/GetDbufState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbufServiceClient) GetQueueState(ctx context.Context, in *GetQueueStateRequest, opts ...grpc.CallOption) (*GetQueueStateResponse, error) {
	out := new(GetQueueStateResponse)
	err := c.cc.Invoke(ctx, "/dbuf.DbufService/GetQueueState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbufServiceClient) ModifyQueue(ctx context.Context, in *ModifyQueueRequest, opts ...grpc.CallOption) (*ModifyQueueResponse, error) {
	out := new(ModifyQueueResponse)
	err := c.cc.Invoke(ctx, "/dbuf.DbufService/ModifyQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbufServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DbufService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DbufService_serviceDesc.Streams[0], "/dbuf.DbufService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dbufServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DbufService_SubscribeClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type dbufServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *dbufServiceSubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DbufServiceServer is the server API for DbufService service.
type DbufServiceServer interface {
	GetDbufState(context.Context, *GetDbufStateRequest) (*GetDbufStateResponse, error)
	GetQueueState(context.Context, *GetQueueStateRequest) (*GetQueueStateResponse, error)
	ModifyQueue(context.Context, *ModifyQueueRequest) (*ModifyQueueResponse, error)
	Subscribe(*SubscribeRequest, DbufService_SubscribeServer) error
}

// UnimplementedDbufServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDbufServiceServer struct {
}

func (*UnimplementedDbufServiceServer) GetDbufState(ctx context.Context, req *GetDbufStateRequest) (*GetDbufStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDbufState not implemented")
}
func (*UnimplementedDbufServiceServer) GetQueueState(ctx context.Context, req *GetQueueStateRequest) (*GetQueueStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueueState not implemented")
}
func (*UnimplementedDbufServiceServer) ModifyQueue(ctx context.Context, req *ModifyQueueRequest) (*ModifyQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyQueue not implemented")
}
func (*UnimplementedDbufServiceServer) Subscribe(req *SubscribeRequest, srv DbufService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterDbufServiceServer(s *grpc.Server, srv DbufServiceServer) {
	s.RegisterService(&_DbufService_serviceDesc, srv)
}

func _DbufService_GetDbufState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDbufStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbufServiceServer).GetDbufState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbuf.DbufService/GetDbufState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbufServiceServer).GetDbufState(ctx, req.(*GetDbufStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbufService_GetQueueState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueueStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbufServiceServer).GetQueueState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbuf.DbufService/GetQueueState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbufServiceServer).GetQueueState(ctx, req.(*GetQueueStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbufService_ModifyQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbufServiceServer).ModifyQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbuf.DbufService/ModifyQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbufServiceServer).ModifyQueue(ctx, req.(*ModifyQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbufService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DbufServiceServer).Subscribe(m, &dbufServiceSubscribeServer{stream})
}

type DbufService_SubscribeServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type dbufServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *dbufServiceSubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

var _DbufService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbuf.DbufService",
	HandlerType: (*DbufServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDbufState",
			Handler:    _DbufService_GetDbufState_Handler,
		},
		{
			MethodName: "GetQueueState",
			Handler:    _DbufService_GetQueueState_Handler,
		},
		{
			MethodName: "ModifyQueue",
			Handler:    _DbufService_ModifyQueue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _DbufService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dbuf.proto",
}
