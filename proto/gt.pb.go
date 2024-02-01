// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: gt.proto

package proto

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Close bool   `protobuf:"varint,2,opt,name=close,proto3" json:"close,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Data) GetClose() bool {
	if x != nil {
		return x.Close
	}
	return false
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Req:
	//
	//	*Request_Head_
	//	*Request_Data
	//	*Request_ResizeWindow_
	Req isRequest_Req `protobuf_oneof:"req"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{1}
}

func (m *Request) GetReq() isRequest_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *Request) GetHead() *Request_Head {
	if x, ok := x.GetReq().(*Request_Head_); ok {
		return x.Head
	}
	return nil
}

func (x *Request) GetData() *Data {
	if x, ok := x.GetReq().(*Request_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Request) GetResizeWindow() *Request_ResizeWindow {
	if x, ok := x.GetReq().(*Request_ResizeWindow_); ok {
		return x.ResizeWindow
	}
	return nil
}

type isRequest_Req interface {
	isRequest_Req()
}

type Request_Head_ struct {
	Head *Request_Head `protobuf:"bytes,1,opt,name=head,proto3,oneof"`
}

type Request_Data struct {
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type Request_ResizeWindow_ struct {
	ResizeWindow *Request_ResizeWindow `protobuf:"bytes,3,opt,name=resize_window,json=resizeWindow,proto3,oneof"`
}

func (*Request_Head_) isRequest_Req() {}

func (*Request_Data) isRequest_Req() {}

func (*Request_ResizeWindow_) isRequest_Req() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Rsp:
	//
	//	*Response_Head_
	//	*Response_Data
	//	*Response_ExecDone_
	Rsp isResponse_Rsp `protobuf_oneof:"rsp"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{2}
}

func (m *Response) GetRsp() isResponse_Rsp {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func (x *Response) GetHead() *Response_Head {
	if x, ok := x.GetRsp().(*Response_Head_); ok {
		return x.Head
	}
	return nil
}

func (x *Response) GetData() *Data {
	if x, ok := x.GetRsp().(*Response_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Response) GetExecDone() *Response_ExecDone {
	if x, ok := x.GetRsp().(*Response_ExecDone_); ok {
		return x.ExecDone
	}
	return nil
}

type isResponse_Rsp interface {
	isResponse_Rsp()
}

type Response_Head_ struct {
	Head *Response_Head `protobuf:"bytes,1,opt,name=head,proto3,oneof"`
}

type Response_Data struct {
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type Response_ExecDone_ struct {
	ExecDone *Response_ExecDone `protobuf:"bytes,3,opt,name=exec_done,json=execDone,proto3,oneof"`
}

func (*Response_Head_) isResponse_Rsp() {}

func (*Response_Data) isResponse_Rsp() {}

func (*Response_ExecDone_) isResponse_Rsp() {}

type Request_ResizeWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows uint32 `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Cols uint32 `protobuf:"varint,2,opt,name=cols,proto3" json:"cols,omitempty"`
}

func (x *Request_ResizeWindow) Reset() {
	*x = Request_ResizeWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_ResizeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_ResizeWindow) ProtoMessage() {}

func (x *Request_ResizeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_ResizeWindow.ProtoReflect.Descriptor instead.
func (*Request_ResizeWindow) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Request_ResizeWindow) GetRows() uint32 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *Request_ResizeWindow) GetCols() uint32 {
	if x != nil {
		return x.Cols
	}
	return 0
}

type Request_Head struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Head:
	//
	//	*Request_Head_Stfp_
	//	*Request_Head_Pty_
	//	*Request_Head_Exec_
	//	*Request_Head_Dial_
	Head isRequest_Head_Head `protobuf_oneof:"head"`
}

func (x *Request_Head) Reset() {
	*x = Request_Head{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Head) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Head) ProtoMessage() {}

func (x *Request_Head) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Head.ProtoReflect.Descriptor instead.
func (*Request_Head) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{1, 1}
}

func (m *Request_Head) GetHead() isRequest_Head_Head {
	if m != nil {
		return m.Head
	}
	return nil
}

func (x *Request_Head) GetStfp() *Request_Head_Stfp {
	if x, ok := x.GetHead().(*Request_Head_Stfp_); ok {
		return x.Stfp
	}
	return nil
}

func (x *Request_Head) GetPty() *Request_Head_Pty {
	if x, ok := x.GetHead().(*Request_Head_Pty_); ok {
		return x.Pty
	}
	return nil
}

func (x *Request_Head) GetExec() *Request_Head_Exec {
	if x, ok := x.GetHead().(*Request_Head_Exec_); ok {
		return x.Exec
	}
	return nil
}

func (x *Request_Head) GetDial() *Request_Head_Dial {
	if x, ok := x.GetHead().(*Request_Head_Dial_); ok {
		return x.Dial
	}
	return nil
}

type isRequest_Head_Head interface {
	isRequest_Head_Head()
}

type Request_Head_Stfp_ struct {
	Stfp *Request_Head_Stfp `protobuf:"bytes,1,opt,name=stfp,proto3,oneof"`
}

type Request_Head_Pty_ struct {
	Pty *Request_Head_Pty `protobuf:"bytes,2,opt,name=pty,proto3,oneof"`
}

type Request_Head_Exec_ struct {
	Exec *Request_Head_Exec `protobuf:"bytes,3,opt,name=exec,proto3,oneof"`
}

type Request_Head_Dial_ struct {
	Dial *Request_Head_Dial `protobuf:"bytes,4,opt,name=dial,proto3,oneof"`
}

func (*Request_Head_Stfp_) isRequest_Head_Head() {}

func (*Request_Head_Pty_) isRequest_Head_Head() {}

func (*Request_Head_Exec_) isRequest_Head_Head() {}

func (*Request_Head_Dial_) isRequest_Head_Head() {}

type Request_Head_Stfp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request_Head_Stfp) Reset() {
	*x = Request_Head_Stfp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Head_Stfp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Head_Stfp) ProtoMessage() {}

func (x *Request_Head_Stfp) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Head_Stfp.ProtoReflect.Descriptor instead.
func (*Request_Head_Stfp) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{1, 1, 0}
}

type Request_Head_Pty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         string                `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	ResizeWindow *Request_ResizeWindow `protobuf:"bytes,2,opt,name=resize_window,json=resizeWindow,proto3" json:"resize_window,omitempty"`
}

func (x *Request_Head_Pty) Reset() {
	*x = Request_Head_Pty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Head_Pty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Head_Pty) ProtoMessage() {}

func (x *Request_Head_Pty) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Head_Pty.ProtoReflect.Descriptor instead.
func (*Request_Head_Pty) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{1, 1, 1}
}

func (x *Request_Head_Pty) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *Request_Head_Pty) GetResizeWindow() *Request_ResizeWindow {
	if x != nil {
		return x.ResizeWindow
	}
	return nil
}

type Request_Head_Exec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Envs    []string `protobuf:"bytes,2,rep,name=envs,proto3" json:"envs,omitempty"`
}

func (x *Request_Head_Exec) Reset() {
	*x = Request_Head_Exec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Head_Exec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Head_Exec) ProtoMessage() {}

func (x *Request_Head_Exec) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Head_Exec.ProtoReflect.Descriptor instead.
func (*Request_Head_Exec) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{1, 1, 2}
}

func (x *Request_Head_Exec) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Request_Head_Exec) GetEnvs() []string {
	if x != nil {
		return x.Envs
	}
	return nil
}

type Request_Head_Dial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *Request_Head_Dial) Reset() {
	*x = Request_Head_Dial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Head_Dial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Head_Dial) ProtoMessage() {}

func (x *Request_Head_Dial) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Head_Dial.ProtoReflect.Descriptor instead.
func (*Request_Head_Dial) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{1, 1, 3}
}

func (x *Request_Head_Dial) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type Response_Head struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Head:
	//
	//	*Response_Head_Error_
	Head isResponse_Head_Head `protobuf_oneof:"head"`
}

func (x *Response_Head) Reset() {
	*x = Response_Head{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Head) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Head) ProtoMessage() {}

func (x *Response_Head) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Head.ProtoReflect.Descriptor instead.
func (*Response_Head) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{2, 0}
}

func (m *Response_Head) GetHead() isResponse_Head_Head {
	if m != nil {
		return m.Head
	}
	return nil
}

func (x *Response_Head) GetError() *Response_Head_Error {
	if x, ok := x.GetHead().(*Response_Head_Error_); ok {
		return x.Error
	}
	return nil
}

type isResponse_Head_Head interface {
	isResponse_Head_Head()
}

type Response_Head_Error_ struct {
	Error *Response_Head_Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

func (*Response_Head_Error_) isResponse_Head_Head() {}

type Response_ExecDone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	ExitCode uint32 `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *Response_ExecDone) Reset() {
	*x = Response_ExecDone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_ExecDone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_ExecDone) ProtoMessage() {}

func (x *Response_ExecDone) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_ExecDone.ProtoReflect.Descriptor instead.
func (*Response_ExecDone) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Response_ExecDone) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Response_ExecDone) GetExitCode() uint32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type Response_Head_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response_Head_Error) Reset() {
	*x = Response_Head_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gt_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Head_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Head_Error) ProtoMessage() {}

func (x *Response_Head_Error) ProtoReflect() protoreflect.Message {
	mi := &file_gt_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Head_Error.ProtoReflect.Descriptor instead.
func (*Response_Head_Error) Descriptor() ([]byte, []int) {
	return file_gt_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *Response_Head_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gt_proto protoreflect.FileDescriptor

var file_gt_proto_rawDesc = []byte{
	0x0a, 0x08, 0x67, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x30, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x22, 0xdf, 0x04, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x48, 0x00, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x1a, 0x36, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x1a, 0x82, 0x03, 0x0a, 0x04, 0x48, 0x65,
	0x61, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x73, 0x74, 0x66, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x2e, 0x53, 0x74, 0x66, 0x70, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74,
	0x66, 0x70, 0x12, 0x2b, 0x0a, 0x03, 0x70, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x2e, 0x50, 0x74, 0x79, 0x48, 0x00, 0x52, 0x03, 0x70, 0x74, 0x79, 0x12,
	0x2e, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x48, 0x00, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x69, 0x61, 0x6c, 0x1a,
	0x06, 0x0a, 0x04, 0x53, 0x74, 0x66, 0x70, 0x1a, 0x5b, 0x0a, 0x03, 0x50, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x1a, 0x34, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x1a, 0x1a, 0x0a, 0x04, 0x44, 0x69,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x42, 0x05,
	0x0a, 0x03, 0x72, 0x65, 0x71, 0x22, 0xbf, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x48, 0x00, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x21,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x44, 0x6f, 0x6e, 0x65, 0x48, 0x00,
	0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x44, 0x6f, 0x6e, 0x65, 0x1a, 0x65, 0x0a, 0x04, 0x48, 0x65,
	0x61, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x1a, 0x3d, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x42, 0x05, 0x0a, 0x03, 0x72, 0x73, 0x70, 0x32, 0x31, 0x0a, 0x02, 0x47, 0x54, 0x12, 0x2b, 0x0a,
	0x02, 0x47, 0x54, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x79, 0x6f, 0x75, 0x6e, 0x67,
	0x2f, 0x67, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gt_proto_rawDescOnce sync.Once
	file_gt_proto_rawDescData = file_gt_proto_rawDesc
)

func file_gt_proto_rawDescGZIP() []byte {
	file_gt_proto_rawDescOnce.Do(func() {
		file_gt_proto_rawDescData = protoimpl.X.CompressGZIP(file_gt_proto_rawDescData)
	})
	return file_gt_proto_rawDescData
}

var file_gt_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_gt_proto_goTypes = []interface{}{
	(*Data)(nil),                 // 0: proto.Data
	(*Request)(nil),              // 1: proto.Request
	(*Response)(nil),             // 2: proto.Response
	(*Request_ResizeWindow)(nil), // 3: proto.Request.ResizeWindow
	(*Request_Head)(nil),         // 4: proto.Request.Head
	(*Request_Head_Stfp)(nil),    // 5: proto.Request.Head.Stfp
	(*Request_Head_Pty)(nil),     // 6: proto.Request.Head.Pty
	(*Request_Head_Exec)(nil),    // 7: proto.Request.Head.Exec
	(*Request_Head_Dial)(nil),    // 8: proto.Request.Head.Dial
	(*Response_Head)(nil),        // 9: proto.Response.Head
	(*Response_ExecDone)(nil),    // 10: proto.Response.ExecDone
	(*Response_Head_Error)(nil),  // 11: proto.Response.Head.Error
}
var file_gt_proto_depIdxs = []int32{
	4,  // 0: proto.Request.head:type_name -> proto.Request.Head
	0,  // 1: proto.Request.data:type_name -> proto.Data
	3,  // 2: proto.Request.resize_window:type_name -> proto.Request.ResizeWindow
	9,  // 3: proto.Response.head:type_name -> proto.Response.Head
	0,  // 4: proto.Response.data:type_name -> proto.Data
	10, // 5: proto.Response.exec_done:type_name -> proto.Response.ExecDone
	5,  // 6: proto.Request.Head.stfp:type_name -> proto.Request.Head.Stfp
	6,  // 7: proto.Request.Head.pty:type_name -> proto.Request.Head.Pty
	7,  // 8: proto.Request.Head.exec:type_name -> proto.Request.Head.Exec
	8,  // 9: proto.Request.Head.dial:type_name -> proto.Request.Head.Dial
	3,  // 10: proto.Request.Head.Pty.resize_window:type_name -> proto.Request.ResizeWindow
	11, // 11: proto.Response.Head.error:type_name -> proto.Response.Head.Error
	1,  // 12: proto.GT.GT:input_type -> proto.Request
	2,  // 13: proto.GT.GT:output_type -> proto.Response
	13, // [13:14] is the sub-list for method output_type
	12, // [12:13] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_gt_proto_init() }
func file_gt_proto_init() {
	if File_gt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_gt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_gt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_gt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_ResizeWindow); i {
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
		file_gt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Head); i {
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
		file_gt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Head_Stfp); i {
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
		file_gt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Head_Pty); i {
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
		file_gt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Head_Exec); i {
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
		file_gt_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Head_Dial); i {
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
		file_gt_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Head); i {
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
		file_gt_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_ExecDone); i {
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
		file_gt_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Head_Error); i {
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
	file_gt_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Request_Head_)(nil),
		(*Request_Data)(nil),
		(*Request_ResizeWindow_)(nil),
	}
	file_gt_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Response_Head_)(nil),
		(*Response_Data)(nil),
		(*Response_ExecDone_)(nil),
	}
	file_gt_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Request_Head_Stfp_)(nil),
		(*Request_Head_Pty_)(nil),
		(*Request_Head_Exec_)(nil),
		(*Request_Head_Dial_)(nil),
	}
	file_gt_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Response_Head_Error_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gt_proto_goTypes,
		DependencyIndexes: file_gt_proto_depIdxs,
		MessageInfos:      file_gt_proto_msgTypes,
	}.Build()
	File_gt_proto = out.File
	file_gt_proto_rawDesc = nil
	file_gt_proto_goTypes = nil
	file_gt_proto_depIdxs = nil
}
