// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/v1/write.proto

package firestore

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A value that is calculated by the server.
type DocumentTransform_FieldTransform_ServerValue int32

const (
	// Unspecified. This value must not be used.
	DocumentTransform_FieldTransform_SERVER_VALUE_UNSPECIFIED DocumentTransform_FieldTransform_ServerValue = 0
	// The time at which the server processed the request, with millisecond
	// precision.
	DocumentTransform_FieldTransform_REQUEST_TIME DocumentTransform_FieldTransform_ServerValue = 1
)

var DocumentTransform_FieldTransform_ServerValue_name = map[int32]string{
	0: "SERVER_VALUE_UNSPECIFIED",
	1: "REQUEST_TIME",
}

var DocumentTransform_FieldTransform_ServerValue_value = map[string]int32{
	"SERVER_VALUE_UNSPECIFIED": 0,
	"REQUEST_TIME":             1,
}

func (x DocumentTransform_FieldTransform_ServerValue) String() string {
	return proto.EnumName(DocumentTransform_FieldTransform_ServerValue_name, int32(x))
}

func (DocumentTransform_FieldTransform_ServerValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{1, 0, 0}
}

// A write on a document.
type Write struct {
	// The operation to execute.
	//
	// Types that are valid to be assigned to Operation:
	//	*Write_Update
	//	*Write_Delete
	//	*Write_Transform
	Operation isWrite_Operation `protobuf_oneof:"operation"`
	// The fields to update in this write.
	//
	// This field can be set only when the operation is `update`.
	// If the mask is not set for an `update` and the document exists, any
	// existing data will be overwritten.
	// If the mask is set and the document on the server has fields not covered by
	// the mask, they are left unchanged.
	// Fields referenced in the mask, but not present in the input document, are
	// deleted from the document on the server.
	// The field paths in this mask must not contain a reserved field name.
	UpdateMask *DocumentMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// An optional precondition on the document.
	//
	// The write will fail if this is set and not met by the target document.
	CurrentDocument      *Precondition `protobuf:"bytes,4,opt,name=current_document,json=currentDocument,proto3" json:"current_document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Write) Reset()         { *m = Write{} }
func (m *Write) String() string { return proto.CompactTextString(m) }
func (*Write) ProtoMessage()    {}
func (*Write) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{0}
}

func (m *Write) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Write.Unmarshal(m, b)
}
func (m *Write) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Write.Marshal(b, m, deterministic)
}
func (m *Write) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Write.Merge(m, src)
}
func (m *Write) XXX_Size() int {
	return xxx_messageInfo_Write.Size(m)
}
func (m *Write) XXX_DiscardUnknown() {
	xxx_messageInfo_Write.DiscardUnknown(m)
}

var xxx_messageInfo_Write proto.InternalMessageInfo

type isWrite_Operation interface {
	isWrite_Operation()
}

type Write_Update struct {
	Update *Document `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type Write_Delete struct {
	Delete string `protobuf:"bytes,2,opt,name=delete,proto3,oneof"`
}

type Write_Transform struct {
	Transform *DocumentTransform `protobuf:"bytes,6,opt,name=transform,proto3,oneof"`
}

func (*Write_Update) isWrite_Operation() {}

func (*Write_Delete) isWrite_Operation() {}

func (*Write_Transform) isWrite_Operation() {}

func (m *Write) GetOperation() isWrite_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *Write) GetUpdate() *Document {
	if x, ok := m.GetOperation().(*Write_Update); ok {
		return x.Update
	}
	return nil
}

func (m *Write) GetDelete() string {
	if x, ok := m.GetOperation().(*Write_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *Write) GetTransform() *DocumentTransform {
	if x, ok := m.GetOperation().(*Write_Transform); ok {
		return x.Transform
	}
	return nil
}

func (m *Write) GetUpdateMask() *DocumentMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *Write) GetCurrentDocument() *Precondition {
	if m != nil {
		return m.CurrentDocument
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Write) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Write_Update)(nil),
		(*Write_Delete)(nil),
		(*Write_Transform)(nil),
	}
}

// A transformation of a document.
type DocumentTransform struct {
	// The name of the document to transform.
	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	// The list of transformations to apply to the fields of the document, in
	// order.
	// This must not be empty.
	FieldTransforms      []*DocumentTransform_FieldTransform `protobuf:"bytes,2,rep,name=field_transforms,json=fieldTransforms,proto3" json:"field_transforms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *DocumentTransform) Reset()         { *m = DocumentTransform{} }
func (m *DocumentTransform) String() string { return proto.CompactTextString(m) }
func (*DocumentTransform) ProtoMessage()    {}
func (*DocumentTransform) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{1}
}

func (m *DocumentTransform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentTransform.Unmarshal(m, b)
}
func (m *DocumentTransform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentTransform.Marshal(b, m, deterministic)
}
func (m *DocumentTransform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentTransform.Merge(m, src)
}
func (m *DocumentTransform) XXX_Size() int {
	return xxx_messageInfo_DocumentTransform.Size(m)
}
func (m *DocumentTransform) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentTransform.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentTransform proto.InternalMessageInfo

func (m *DocumentTransform) GetDocument() string {
	if m != nil {
		return m.Document
	}
	return ""
}

func (m *DocumentTransform) GetFieldTransforms() []*DocumentTransform_FieldTransform {
	if m != nil {
		return m.FieldTransforms
	}
	return nil
}

// A transformation of a field of the document.
type DocumentTransform_FieldTransform struct {
	// The path of the field. See [Document.fields][google.firestore.v1.Document.fields] for the field path syntax
	// reference.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// The transformation to apply on the field.
	//
	// Types that are valid to be assigned to TransformType:
	//	*DocumentTransform_FieldTransform_SetToServerValue
	//	*DocumentTransform_FieldTransform_Increment
	//	*DocumentTransform_FieldTransform_Maximum
	//	*DocumentTransform_FieldTransform_Minimum
	//	*DocumentTransform_FieldTransform_AppendMissingElements
	//	*DocumentTransform_FieldTransform_RemoveAllFromArray
	TransformType        isDocumentTransform_FieldTransform_TransformType `protobuf_oneof:"transform_type"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *DocumentTransform_FieldTransform) Reset()         { *m = DocumentTransform_FieldTransform{} }
func (m *DocumentTransform_FieldTransform) String() string { return proto.CompactTextString(m) }
func (*DocumentTransform_FieldTransform) ProtoMessage()    {}
func (*DocumentTransform_FieldTransform) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{1, 0}
}

func (m *DocumentTransform_FieldTransform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentTransform_FieldTransform.Unmarshal(m, b)
}
func (m *DocumentTransform_FieldTransform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentTransform_FieldTransform.Marshal(b, m, deterministic)
}
func (m *DocumentTransform_FieldTransform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentTransform_FieldTransform.Merge(m, src)
}
func (m *DocumentTransform_FieldTransform) XXX_Size() int {
	return xxx_messageInfo_DocumentTransform_FieldTransform.Size(m)
}
func (m *DocumentTransform_FieldTransform) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentTransform_FieldTransform.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentTransform_FieldTransform proto.InternalMessageInfo

func (m *DocumentTransform_FieldTransform) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

type isDocumentTransform_FieldTransform_TransformType interface {
	isDocumentTransform_FieldTransform_TransformType()
}

type DocumentTransform_FieldTransform_SetToServerValue struct {
	SetToServerValue DocumentTransform_FieldTransform_ServerValue `protobuf:"varint,2,opt,name=set_to_server_value,json=setToServerValue,proto3,enum=google.firestore.v1.DocumentTransform_FieldTransform_ServerValue,oneof"`
}

type DocumentTransform_FieldTransform_Increment struct {
	Increment *Value `protobuf:"bytes,3,opt,name=increment,proto3,oneof"`
}

type DocumentTransform_FieldTransform_Maximum struct {
	Maximum *Value `protobuf:"bytes,4,opt,name=maximum,proto3,oneof"`
}

type DocumentTransform_FieldTransform_Minimum struct {
	Minimum *Value `protobuf:"bytes,5,opt,name=minimum,proto3,oneof"`
}

type DocumentTransform_FieldTransform_AppendMissingElements struct {
	AppendMissingElements *ArrayValue `protobuf:"bytes,6,opt,name=append_missing_elements,json=appendMissingElements,proto3,oneof"`
}

type DocumentTransform_FieldTransform_RemoveAllFromArray struct {
	RemoveAllFromArray *ArrayValue `protobuf:"bytes,7,opt,name=remove_all_from_array,json=removeAllFromArray,proto3,oneof"`
}

func (*DocumentTransform_FieldTransform_SetToServerValue) isDocumentTransform_FieldTransform_TransformType() {
}

func (*DocumentTransform_FieldTransform_Increment) isDocumentTransform_FieldTransform_TransformType() {
}

func (*DocumentTransform_FieldTransform_Maximum) isDocumentTransform_FieldTransform_TransformType() {}

func (*DocumentTransform_FieldTransform_Minimum) isDocumentTransform_FieldTransform_TransformType() {}

func (*DocumentTransform_FieldTransform_AppendMissingElements) isDocumentTransform_FieldTransform_TransformType() {
}

func (*DocumentTransform_FieldTransform_RemoveAllFromArray) isDocumentTransform_FieldTransform_TransformType() {
}

func (m *DocumentTransform_FieldTransform) GetTransformType() isDocumentTransform_FieldTransform_TransformType {
	if m != nil {
		return m.TransformType
	}
	return nil
}

func (m *DocumentTransform_FieldTransform) GetSetToServerValue() DocumentTransform_FieldTransform_ServerValue {
	if x, ok := m.GetTransformType().(*DocumentTransform_FieldTransform_SetToServerValue); ok {
		return x.SetToServerValue
	}
	return DocumentTransform_FieldTransform_SERVER_VALUE_UNSPECIFIED
}

func (m *DocumentTransform_FieldTransform) GetIncrement() *Value {
	if x, ok := m.GetTransformType().(*DocumentTransform_FieldTransform_Increment); ok {
		return x.Increment
	}
	return nil
}

func (m *DocumentTransform_FieldTransform) GetMaximum() *Value {
	if x, ok := m.GetTransformType().(*DocumentTransform_FieldTransform_Maximum); ok {
		return x.Maximum
	}
	return nil
}

func (m *DocumentTransform_FieldTransform) GetMinimum() *Value {
	if x, ok := m.GetTransformType().(*DocumentTransform_FieldTransform_Minimum); ok {
		return x.Minimum
	}
	return nil
}

func (m *DocumentTransform_FieldTransform) GetAppendMissingElements() *ArrayValue {
	if x, ok := m.GetTransformType().(*DocumentTransform_FieldTransform_AppendMissingElements); ok {
		return x.AppendMissingElements
	}
	return nil
}

func (m *DocumentTransform_FieldTransform) GetRemoveAllFromArray() *ArrayValue {
	if x, ok := m.GetTransformType().(*DocumentTransform_FieldTransform_RemoveAllFromArray); ok {
		return x.RemoveAllFromArray
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DocumentTransform_FieldTransform) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DocumentTransform_FieldTransform_SetToServerValue)(nil),
		(*DocumentTransform_FieldTransform_Increment)(nil),
		(*DocumentTransform_FieldTransform_Maximum)(nil),
		(*DocumentTransform_FieldTransform_Minimum)(nil),
		(*DocumentTransform_FieldTransform_AppendMissingElements)(nil),
		(*DocumentTransform_FieldTransform_RemoveAllFromArray)(nil),
	}
}

// The result of applying a write.
type WriteResult struct {
	// The last update time of the document after applying the write. Not set
	// after a `delete`.
	//
	// If the write did not actually change the document, this will be the
	// previous update_time.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The results of applying each [DocumentTransform.FieldTransform][google.firestore.v1.DocumentTransform.FieldTransform], in the
	// same order.
	TransformResults     []*Value `protobuf:"bytes,2,rep,name=transform_results,json=transformResults,proto3" json:"transform_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResult) Reset()         { *m = WriteResult{} }
func (m *WriteResult) String() string { return proto.CompactTextString(m) }
func (*WriteResult) ProtoMessage()    {}
func (*WriteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{2}
}

func (m *WriteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResult.Unmarshal(m, b)
}
func (m *WriteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResult.Marshal(b, m, deterministic)
}
func (m *WriteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResult.Merge(m, src)
}
func (m *WriteResult) XXX_Size() int {
	return xxx_messageInfo_WriteResult.Size(m)
}
func (m *WriteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResult.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResult proto.InternalMessageInfo

func (m *WriteResult) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *WriteResult) GetTransformResults() []*Value {
	if m != nil {
		return m.TransformResults
	}
	return nil
}

// A [Document][google.firestore.v1.Document] has changed.
//
// May be the result of multiple [writes][google.firestore.v1.Write], including deletes, that
// ultimately resulted in a new value for the [Document][google.firestore.v1.Document].
//
// Multiple [DocumentChange][google.firestore.v1.DocumentChange] messages may be returned for the same logical
// change, if multiple targets are affected.
type DocumentChange struct {
	// The new state of the [Document][google.firestore.v1.Document].
	//
	// If `mask` is set, contains only fields that were updated or added.
	Document *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	// A set of target IDs of targets that match this document.
	TargetIds []int32 `protobuf:"varint,5,rep,packed,name=target_ids,json=targetIds,proto3" json:"target_ids,omitempty"`
	// A set of target IDs for targets that no longer match this document.
	RemovedTargetIds     []int32  `protobuf:"varint,6,rep,packed,name=removed_target_ids,json=removedTargetIds,proto3" json:"removed_target_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentChange) Reset()         { *m = DocumentChange{} }
func (m *DocumentChange) String() string { return proto.CompactTextString(m) }
func (*DocumentChange) ProtoMessage()    {}
func (*DocumentChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{3}
}

func (m *DocumentChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentChange.Unmarshal(m, b)
}
func (m *DocumentChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentChange.Marshal(b, m, deterministic)
}
func (m *DocumentChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentChange.Merge(m, src)
}
func (m *DocumentChange) XXX_Size() int {
	return xxx_messageInfo_DocumentChange.Size(m)
}
func (m *DocumentChange) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentChange.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentChange proto.InternalMessageInfo

func (m *DocumentChange) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *DocumentChange) GetTargetIds() []int32 {
	if m != nil {
		return m.TargetIds
	}
	return nil
}

func (m *DocumentChange) GetRemovedTargetIds() []int32 {
	if m != nil {
		return m.RemovedTargetIds
	}
	return nil
}

// A [Document][google.firestore.v1.Document] has been deleted.
//
// May be the result of multiple [writes][google.firestore.v1.Write], including updates, the
// last of which deleted the [Document][google.firestore.v1.Document].
//
// Multiple [DocumentDelete][google.firestore.v1.DocumentDelete] messages may be returned for the same logical
// delete, if multiple targets are affected.
type DocumentDelete struct {
	// The resource name of the [Document][google.firestore.v1.Document] that was deleted.
	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	// A set of target IDs for targets that previously matched this entity.
	RemovedTargetIds []int32 `protobuf:"varint,6,rep,packed,name=removed_target_ids,json=removedTargetIds,proto3" json:"removed_target_ids,omitempty"`
	// The read timestamp at which the delete was observed.
	//
	// Greater or equal to the `commit_time` of the delete.
	ReadTime             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DocumentDelete) Reset()         { *m = DocumentDelete{} }
func (m *DocumentDelete) String() string { return proto.CompactTextString(m) }
func (*DocumentDelete) ProtoMessage()    {}
func (*DocumentDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{4}
}

func (m *DocumentDelete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentDelete.Unmarshal(m, b)
}
func (m *DocumentDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentDelete.Marshal(b, m, deterministic)
}
func (m *DocumentDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentDelete.Merge(m, src)
}
func (m *DocumentDelete) XXX_Size() int {
	return xxx_messageInfo_DocumentDelete.Size(m)
}
func (m *DocumentDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentDelete.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentDelete proto.InternalMessageInfo

func (m *DocumentDelete) GetDocument() string {
	if m != nil {
		return m.Document
	}
	return ""
}

func (m *DocumentDelete) GetRemovedTargetIds() []int32 {
	if m != nil {
		return m.RemovedTargetIds
	}
	return nil
}

func (m *DocumentDelete) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

// A [Document][google.firestore.v1.Document] has been removed from the view of the targets.
//
// Sent if the document is no longer relevant to a target and is out of view.
// Can be sent instead of a DocumentDelete or a DocumentChange if the server
// can not send the new value of the document.
//
// Multiple [DocumentRemove][google.firestore.v1.DocumentRemove] messages may be returned for the same logical
// write or delete, if multiple targets are affected.
type DocumentRemove struct {
	// The resource name of the [Document][google.firestore.v1.Document] that has gone out of view.
	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	// A set of target IDs for targets that previously matched this document.
	RemovedTargetIds []int32 `protobuf:"varint,2,rep,packed,name=removed_target_ids,json=removedTargetIds,proto3" json:"removed_target_ids,omitempty"`
	// The read timestamp at which the remove was observed.
	//
	// Greater or equal to the `commit_time` of the change/delete/remove.
	ReadTime             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DocumentRemove) Reset()         { *m = DocumentRemove{} }
func (m *DocumentRemove) String() string { return proto.CompactTextString(m) }
func (*DocumentRemove) ProtoMessage()    {}
func (*DocumentRemove) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{5}
}

func (m *DocumentRemove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentRemove.Unmarshal(m, b)
}
func (m *DocumentRemove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentRemove.Marshal(b, m, deterministic)
}
func (m *DocumentRemove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentRemove.Merge(m, src)
}
func (m *DocumentRemove) XXX_Size() int {
	return xxx_messageInfo_DocumentRemove.Size(m)
}
func (m *DocumentRemove) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentRemove.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentRemove proto.InternalMessageInfo

func (m *DocumentRemove) GetDocument() string {
	if m != nil {
		return m.Document
	}
	return ""
}

func (m *DocumentRemove) GetRemovedTargetIds() []int32 {
	if m != nil {
		return m.RemovedTargetIds
	}
	return nil
}

func (m *DocumentRemove) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

// A digest of all the documents that match a given target.
type ExistenceFilter struct {
	// The target ID to which this filter applies.
	TargetId int32 `protobuf:"varint,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// The total count of documents that match [target_id][google.firestore.v1.ExistenceFilter.target_id].
	//
	// If different from the count of documents in the client that match, the
	// client must manually determine which documents no longer match the target.
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistenceFilter) Reset()         { *m = ExistenceFilter{} }
func (m *ExistenceFilter) String() string { return proto.CompactTextString(m) }
func (*ExistenceFilter) ProtoMessage()    {}
func (*ExistenceFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77c99aae0973dc, []int{6}
}

func (m *ExistenceFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistenceFilter.Unmarshal(m, b)
}
func (m *ExistenceFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistenceFilter.Marshal(b, m, deterministic)
}
func (m *ExistenceFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistenceFilter.Merge(m, src)
}
func (m *ExistenceFilter) XXX_Size() int {
	return xxx_messageInfo_ExistenceFilter.Size(m)
}
func (m *ExistenceFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistenceFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ExistenceFilter proto.InternalMessageInfo

func (m *ExistenceFilter) GetTargetId() int32 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *ExistenceFilter) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.firestore.v1.DocumentTransform_FieldTransform_ServerValue", DocumentTransform_FieldTransform_ServerValue_name, DocumentTransform_FieldTransform_ServerValue_value)
	proto.RegisterType((*Write)(nil), "google.firestore.v1.Write")
	proto.RegisterType((*DocumentTransform)(nil), "google.firestore.v1.DocumentTransform")
	proto.RegisterType((*DocumentTransform_FieldTransform)(nil), "google.firestore.v1.DocumentTransform.FieldTransform")
	proto.RegisterType((*WriteResult)(nil), "google.firestore.v1.WriteResult")
	proto.RegisterType((*DocumentChange)(nil), "google.firestore.v1.DocumentChange")
	proto.RegisterType((*DocumentDelete)(nil), "google.firestore.v1.DocumentDelete")
	proto.RegisterType((*DocumentRemove)(nil), "google.firestore.v1.DocumentRemove")
	proto.RegisterType((*ExistenceFilter)(nil), "google.firestore.v1.ExistenceFilter")
}

func init() { proto.RegisterFile("google/firestore/v1/write.proto", fileDescriptor_5c77c99aae0973dc) }

var fileDescriptor_5c77c99aae0973dc = []byte{
	// 853 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0x6e, 0xd2, 0x26, 0x5b, 0xbf, 0xa0, 0xd6, 0x3b, 0xcb, 0x6a, 0x4d, 0xd8, 0x6a, 0x43, 0x0e,
	0xa8, 0x07, 0xe4, 0xa8, 0x45, 0xb0, 0x82, 0x85, 0x43, 0xd3, 0x3a, 0x6d, 0xa5, 0x2d, 0x0a, 0x4e,
	0x1a, 0x04, 0xaa, 0x34, 0xcc, 0xda, 0x13, 0xd7, 0x5a, 0x7b, 0xc6, 0x9a, 0x19, 0x87, 0xdd, 0xdf,
	0xc1, 0x85, 0x33, 0xe2, 0xc4, 0xbf, 0xe0, 0xca, 0x8d, 0x1f, 0xc1, 0xff, 0x40, 0x9e, 0xb1, 0xdd,
	0x06, 0xa2, 0x6c, 0x59, 0xed, 0x2d, 0x6f, 0xde, 0xf7, 0x7d, 0xef, 0xf3, 0x7b, 0x33, 0x2f, 0xf0,
	0x24, 0xe2, 0x3c, 0x4a, 0xe8, 0x60, 0x1e, 0x0b, 0x2a, 0x15, 0x17, 0x74, 0xb0, 0x38, 0x18, 0xfc,
	0x24, 0x62, 0x45, 0xdd, 0x4c, 0x70, 0xc5, 0xd1, 0x03, 0x03, 0x70, 0x6b, 0x80, 0xbb, 0x38, 0xe8,
	0xf6, 0x56, 0xb1, 0x02, 0x9e, 0xa6, 0x9c, 0x19, 0x5a, 0xb7, 0xbf, 0x0a, 0x11, 0xf2, 0x20, 0x4f,
	0x29, 0x53, 0x25, 0xa6, 0xaa, 0xad, 0xa3, 0x17, 0xf9, 0x7c, 0xa0, 0xe2, 0x94, 0x4a, 0x45, 0xd2,
	0xac, 0x04, 0x3c, 0x2e, 0x01, 0x24, 0x8b, 0x07, 0x84, 0x31, 0xae, 0x88, 0x8a, 0x39, 0x93, 0x26,
	0xdb, 0xff, 0xa3, 0x09, 0xad, 0xef, 0x0a, 0xa7, 0xe8, 0x29, 0xb4, 0xf3, 0x2c, 0x24, 0x8a, 0x3a,
	0x8d, 0x5e, 0x63, 0xbf, 0x73, 0xb8, 0xe7, 0xae, 0x30, 0xed, 0x9e, 0x94, 0xd5, 0xcf, 0x36, 0xfc,
	0x12, 0x8e, 0x1c, 0x68, 0x87, 0x34, 0xa1, 0x8a, 0x3a, 0xcd, 0x5e, 0x63, 0xdf, 0x2a, 0x32, 0x26,
	0x46, 0x23, 0xb0, 0x94, 0x20, 0x4c, 0xce, 0xb9, 0x48, 0x9d, 0xb6, 0x56, 0xfd, 0x78, 0xad, 0xea,
	0xb4, 0x42, 0x9f, 0x6d, 0xf8, 0x37, 0x54, 0x34, 0x84, 0x8e, 0xa9, 0x85, 0x53, 0x22, 0x5f, 0x3a,
	0x9b, 0x5a, 0xe9, 0xa3, 0xb5, 0x4a, 0x17, 0x44, 0xbe, 0xf4, 0xc1, 0xb0, 0x8a, 0xdf, 0xe8, 0x39,
	0xd8, 0x41, 0x2e, 0x04, 0x65, 0x0a, 0x57, 0x1d, 0x74, 0xb6, 0xd6, 0x08, 0x8d, 0x05, 0x0d, 0x38,
	0x0b, 0xe3, 0xa2, 0x59, 0xfe, 0x6e, 0x49, 0xad, 0xd4, 0x87, 0x1d, 0xb0, 0x78, 0x46, 0x85, 0x6e,
	0x65, 0xff, 0xef, 0x16, 0xdc, 0xff, 0xcf, 0x17, 0xa0, 0x2e, 0x6c, 0xd7, 0x85, 0x8a, 0x8e, 0x5a,
	0x7e, 0x1d, 0xa3, 0x1f, 0xc1, 0x9e, 0xc7, 0x34, 0x09, 0x71, 0xfd, 0x8d, 0xd2, 0x69, 0xf6, 0x36,
	0xf7, 0x3b, 0x87, 0x9f, 0xdd, 0xad, 0x3f, 0xee, 0xa8, 0xa0, 0xd7, 0xa1, 0xbf, 0x3b, 0x5f, 0x8a,
	0x65, 0xf7, 0xaf, 0x2d, 0xd8, 0x59, 0xc6, 0xa0, 0x3d, 0x00, 0x53, 0x34, 0x23, 0xea, 0xba, 0xb4,
	0x64, 0xe9, 0x93, 0x31, 0x51, 0xd7, 0x48, 0xc0, 0x03, 0x49, 0x15, 0x56, 0x1c, 0x4b, 0x2a, 0x16,
	0x54, 0xe0, 0x05, 0x49, 0x72, 0x33, 0xd3, 0x9d, 0xc3, 0xa3, 0xb7, 0xb2, 0xe5, 0x4e, 0xb4, 0xd2,
	0xac, 0x10, 0x3a, 0xdb, 0xf0, 0x6d, 0x49, 0xd5, 0x94, 0xdf, 0x3a, 0x43, 0x5f, 0x82, 0x15, 0xb3,
	0x40, 0x50, 0xdd, 0x24, 0x33, 0xd6, 0xee, 0xca, 0x4a, 0x95, 0xc4, 0x0d, 0x1c, 0x7d, 0x0e, 0xf7,
	0x52, 0xf2, 0x2a, 0x4e, 0xf3, 0xb4, 0x9c, 0xe3, 0x7a, 0x66, 0x05, 0xd6, 0xbc, 0x98, 0x69, 0x5e,
	0xeb, 0x4e, 0x3c, 0x03, 0x46, 0xdf, 0xc3, 0x23, 0x92, 0x65, 0x94, 0x85, 0x38, 0x8d, 0xa5, 0x8c,
	0x59, 0x84, 0x69, 0xa2, 0x9d, 0xc8, 0xf2, 0x6a, 0x3f, 0x59, 0xa9, 0x73, 0x24, 0x04, 0x79, 0x5d,
	0x89, 0x3d, 0x34, 0x0a, 0x17, 0x46, 0xc0, 0x2b, 0xf9, 0x68, 0x0a, 0x0f, 0x05, 0x4d, 0xf9, 0x82,
	0x62, 0x92, 0x24, 0x78, 0x2e, 0x78, 0x8a, 0x49, 0x41, 0x73, 0xee, 0xdd, 0x55, 0x18, 0x19, 0xfe,
	0x51, 0x92, 0x8c, 0x04, 0x4f, 0x75, 0xaa, 0xff, 0x35, 0x74, 0x6e, 0xf7, 0xfa, 0x31, 0x38, 0x13,
	0xcf, 0x9f, 0x79, 0x3e, 0x9e, 0x1d, 0x3d, 0xbf, 0xf4, 0xf0, 0xe5, 0x37, 0x93, 0xb1, 0x77, 0x7c,
	0x3e, 0x3a, 0xf7, 0x4e, 0xec, 0x0d, 0x64, 0xc3, 0x7b, 0xbe, 0xf7, 0xed, 0xa5, 0x37, 0x99, 0xe2,
	0xe9, 0xf9, 0x85, 0x67, 0x37, 0x86, 0x36, 0xec, 0xd4, 0xb7, 0x13, 0xab, 0xd7, 0x19, 0xed, 0xff,
	0xdc, 0x80, 0x8e, 0xde, 0x15, 0x3e, 0x95, 0x79, 0xa2, 0xd0, 0xb3, 0xfa, 0x59, 0x16, 0x3b, 0xa7,
	0x5c, 0x1b, 0x75, 0x37, 0xab, 0x85, 0xe4, 0x4e, 0xab, 0x85, 0x54, 0xbd, 0xc7, 0xe2, 0x00, 0x9d,
	0xc2, 0xfd, 0x1b, 0x79, 0xa1, 0x05, 0xab, 0x37, 0xb0, 0x66, 0x20, 0xbe, 0x5d, 0x93, 0x8c, 0x09,
	0xd9, 0xff, 0xa5, 0x01, 0x3b, 0xd5, 0x45, 0x3c, 0xbe, 0x26, 0x2c, 0xa2, 0xe8, 0x8b, 0x7f, 0x3d,
	0xbd, 0x37, 0x2d, 0xb3, 0x5b, 0x2f, 0x73, 0x0f, 0x40, 0x11, 0x11, 0x51, 0x85, 0xe3, 0x50, 0x3a,
	0xad, 0xde, 0xe6, 0x7e, 0xcb, 0xb7, 0xcc, 0xc9, 0x79, 0x28, 0xd1, 0x27, 0x50, 0x76, 0x3a, 0xc4,
	0xb7, 0x60, 0x6d, 0x0d, 0xb3, 0xcb, 0xcc, 0xb4, 0x42, 0x17, 0x0d, 0xab, 0xad, 0x9d, 0x98, 0x95,
	0xb8, 0x6e, 0x2b, 0xfc, 0x2f, 0x71, 0xf4, 0x14, 0x2c, 0x41, 0x49, 0x68, 0x7a, 0xbf, 0xf5, 0xc6,
	0xde, 0x6f, 0x17, 0xe0, 0x22, 0x5c, 0x72, 0xe5, 0x6b, 0xd5, 0xb7, 0x70, 0xd5, 0x7c, 0xd7, 0xae,
	0x4e, 0x60, 0xd7, 0x7b, 0x15, 0x4b, 0x45, 0x59, 0x40, 0x47, 0x71, 0xa2, 0xa8, 0x40, 0x1f, 0x82,
	0x55, 0x57, 0xd4, 0xb6, 0x5a, 0xfe, 0x76, 0x35, 0x0a, 0xf4, 0x3e, 0xb4, 0x02, 0x9e, 0x33, 0xa5,
	0x17, 0x54, 0xcb, 0x37, 0xc1, 0xf0, 0xb7, 0x06, 0x3c, 0x0a, 0x78, 0xba, 0x6a, 0xda, 0x43, 0xd0,
	0x77, 0x77, 0x5c, 0x98, 0x18, 0x37, 0x7e, 0xf8, 0xaa, 0x84, 0x44, 0x3c, 0x21, 0x2c, 0x72, 0xb9,
	0x88, 0x06, 0x11, 0x65, 0xda, 0xe2, 0xc0, 0xa4, 0x48, 0x16, 0xcb, 0xa5, 0xbf, 0xde, 0x67, 0x75,
	0xf0, 0x6b, 0x73, 0xeb, 0xf4, 0x78, 0x34, 0xf9, 0xbd, 0xf9, 0xc1, 0xa9, 0x51, 0x39, 0x4e, 0x78,
	0x1e, 0xba, 0xa3, 0xba, 0xdc, 0xec, 0xe0, 0xcf, 0x2a, 0x77, 0xa5, 0x73, 0x57, 0x75, 0xee, 0x6a,
	0x76, 0xf0, 0xa2, 0xad, 0xeb, 0x7c, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x45, 0xd0,
	0x55, 0x39, 0x08, 0x00, 0x00,
}
