// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: store/task.proto

package store

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

// TaskDatabaseCreatePayload is the task payload for creating databases.
type TaskDatabaseCreatePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// common fields
	Skipped       bool   `protobuf:"varint,1,opt,name=skipped,proto3" json:"skipped,omitempty"`
	SkippedReason string `protobuf:"bytes,2,opt,name=skipped_reason,json=skippedReason,proto3" json:"skipped_reason,omitempty"`
	SpecId        string `protobuf:"bytes,3,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	ProjectId     int32  `protobuf:"varint,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DatabaseName  string `protobuf:"bytes,5,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	TableName     string `protobuf:"bytes,6,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	SheetId       int32  `protobuf:"varint,7,opt,name=sheet_id,json=sheetId,proto3" json:"sheet_id,omitempty"`
	CharacterSet  string `protobuf:"bytes,8,opt,name=character_set,json=characterSet,proto3" json:"character_set,omitempty"`
	Collation     string `protobuf:"bytes,9,opt,name=collation,proto3" json:"collation,omitempty"`
	EnvironmentId string `protobuf:"bytes,10,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Labels        string `protobuf:"bytes,11,opt,name=labels,proto3" json:"labels,omitempty"`
}

func (x *TaskDatabaseCreatePayload) Reset() {
	*x = TaskDatabaseCreatePayload{}
	mi := &file_store_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDatabaseCreatePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDatabaseCreatePayload) ProtoMessage() {}

func (x *TaskDatabaseCreatePayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDatabaseCreatePayload.ProtoReflect.Descriptor instead.
func (*TaskDatabaseCreatePayload) Descriptor() ([]byte, []int) {
	return file_store_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskDatabaseCreatePayload) GetSkipped() bool {
	if x != nil {
		return x.Skipped
	}
	return false
}

func (x *TaskDatabaseCreatePayload) GetSkippedReason() string {
	if x != nil {
		return x.SkippedReason
	}
	return ""
}

func (x *TaskDatabaseCreatePayload) GetSpecId() string {
	if x != nil {
		return x.SpecId
	}
	return ""
}

func (x *TaskDatabaseCreatePayload) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *TaskDatabaseCreatePayload) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *TaskDatabaseCreatePayload) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *TaskDatabaseCreatePayload) GetSheetId() int32 {
	if x != nil {
		return x.SheetId
	}
	return 0
}

func (x *TaskDatabaseCreatePayload) GetCharacterSet() string {
	if x != nil {
		return x.CharacterSet
	}
	return ""
}

func (x *TaskDatabaseCreatePayload) GetCollation() string {
	if x != nil {
		return x.Collation
	}
	return ""
}

func (x *TaskDatabaseCreatePayload) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *TaskDatabaseCreatePayload) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

// TaskDatabaseUpdatePayload is the task payload for updating database (DDL & DML).
type TaskDatabaseUpdatePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// common fields
	Skipped               bool                   `protobuf:"varint,1,opt,name=skipped,proto3" json:"skipped,omitempty"`
	SkippedReason         string                 `protobuf:"bytes,2,opt,name=skipped_reason,json=skippedReason,proto3" json:"skipped_reason,omitempty"`
	SpecId                string                 `protobuf:"bytes,3,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	SchemaVersion         string                 `protobuf:"bytes,4,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	SheetId               int32                  `protobuf:"varint,5,opt,name=sheet_id,json=sheetId,proto3" json:"sheet_id,omitempty"`
	PreUpdateBackupDetail *PreUpdateBackupDetail `protobuf:"bytes,6,opt,name=pre_update_backup_detail,json=preUpdateBackupDetail,proto3" json:"pre_update_backup_detail,omitempty"`
	// flags is used for ghost sync
	Flags             map[string]string  `protobuf:"bytes,7,rep,name=flags,proto3" json:"flags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TaskReleaseSource *TaskReleaseSource `protobuf:"bytes,8,opt,name=task_release_source,json=taskReleaseSource,proto3" json:"task_release_source,omitempty"`
}

func (x *TaskDatabaseUpdatePayload) Reset() {
	*x = TaskDatabaseUpdatePayload{}
	mi := &file_store_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDatabaseUpdatePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDatabaseUpdatePayload) ProtoMessage() {}

func (x *TaskDatabaseUpdatePayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDatabaseUpdatePayload.ProtoReflect.Descriptor instead.
func (*TaskDatabaseUpdatePayload) Descriptor() ([]byte, []int) {
	return file_store_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskDatabaseUpdatePayload) GetSkipped() bool {
	if x != nil {
		return x.Skipped
	}
	return false
}

func (x *TaskDatabaseUpdatePayload) GetSkippedReason() string {
	if x != nil {
		return x.SkippedReason
	}
	return ""
}

func (x *TaskDatabaseUpdatePayload) GetSpecId() string {
	if x != nil {
		return x.SpecId
	}
	return ""
}

func (x *TaskDatabaseUpdatePayload) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *TaskDatabaseUpdatePayload) GetSheetId() int32 {
	if x != nil {
		return x.SheetId
	}
	return 0
}

func (x *TaskDatabaseUpdatePayload) GetPreUpdateBackupDetail() *PreUpdateBackupDetail {
	if x != nil {
		return x.PreUpdateBackupDetail
	}
	return nil
}

func (x *TaskDatabaseUpdatePayload) GetFlags() map[string]string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *TaskDatabaseUpdatePayload) GetTaskReleaseSource() *TaskReleaseSource {
	if x != nil {
		return x.TaskReleaseSource
	}
	return nil
}

// TaskDatabaseDataExportPayload is the task payload for database data export.
type TaskDatabaseDataExportPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// common fields
	SpecId   string       `protobuf:"bytes,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	SheetId  int32        `protobuf:"varint,2,opt,name=sheet_id,json=sheetId,proto3" json:"sheet_id,omitempty"`
	Password string       `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Format   ExportFormat `protobuf:"varint,4,opt,name=format,proto3,enum=bytebase.store.ExportFormat" json:"format,omitempty"`
}

func (x *TaskDatabaseDataExportPayload) Reset() {
	*x = TaskDatabaseDataExportPayload{}
	mi := &file_store_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDatabaseDataExportPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDatabaseDataExportPayload) ProtoMessage() {}

func (x *TaskDatabaseDataExportPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDatabaseDataExportPayload.ProtoReflect.Descriptor instead.
func (*TaskDatabaseDataExportPayload) Descriptor() ([]byte, []int) {
	return file_store_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskDatabaseDataExportPayload) GetSpecId() string {
	if x != nil {
		return x.SpecId
	}
	return ""
}

func (x *TaskDatabaseDataExportPayload) GetSheetId() int32 {
	if x != nil {
		return x.SheetId
	}
	return 0
}

func (x *TaskDatabaseDataExportPayload) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *TaskDatabaseDataExportPayload) GetFormat() ExportFormat {
	if x != nil {
		return x.Format
	}
	return ExportFormat_FORMAT_UNSPECIFIED
}

type TaskReleaseSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: projects/{project}/releases/{release}/files/{file}
	// {file} is URL path escaped.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *TaskReleaseSource) Reset() {
	*x = TaskReleaseSource{}
	mi := &file_store_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskReleaseSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReleaseSource) ProtoMessage() {}

func (x *TaskReleaseSource) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReleaseSource.ProtoReflect.Descriptor instead.
func (*TaskReleaseSource) Descriptor() ([]byte, []int) {
	return file_store_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskReleaseSource) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

var File_store_task_proto protoreflect.FileDescriptor

var file_store_task_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf5, 0x02, 0x0a, 0x19, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6b,
	0x69, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x73, 0x68, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0xf0, 0x03, 0x0a, 0x19, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70,
	0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x65, 0x63,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x18, 0x70, 0x72, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x15, 0x70,
	0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x4a, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x51, 0x0a, 0x13, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5, 0x01,
	0x0a, 0x1d, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x34, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_task_proto_rawDescOnce sync.Once
	file_store_task_proto_rawDescData = file_store_task_proto_rawDesc
)

func file_store_task_proto_rawDescGZIP() []byte {
	file_store_task_proto_rawDescOnce.Do(func() {
		file_store_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_task_proto_rawDescData)
	})
	return file_store_task_proto_rawDescData
}

var file_store_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_store_task_proto_goTypes = []any{
	(*TaskDatabaseCreatePayload)(nil),     // 0: bytebase.store.TaskDatabaseCreatePayload
	(*TaskDatabaseUpdatePayload)(nil),     // 1: bytebase.store.TaskDatabaseUpdatePayload
	(*TaskDatabaseDataExportPayload)(nil), // 2: bytebase.store.TaskDatabaseDataExportPayload
	(*TaskReleaseSource)(nil),             // 3: bytebase.store.TaskReleaseSource
	nil,                                   // 4: bytebase.store.TaskDatabaseUpdatePayload.FlagsEntry
	(*PreUpdateBackupDetail)(nil),         // 5: bytebase.store.PreUpdateBackupDetail
	(ExportFormat)(0),                     // 6: bytebase.store.ExportFormat
}
var file_store_task_proto_depIdxs = []int32{
	5, // 0: bytebase.store.TaskDatabaseUpdatePayload.pre_update_backup_detail:type_name -> bytebase.store.PreUpdateBackupDetail
	4, // 1: bytebase.store.TaskDatabaseUpdatePayload.flags:type_name -> bytebase.store.TaskDatabaseUpdatePayload.FlagsEntry
	3, // 2: bytebase.store.TaskDatabaseUpdatePayload.task_release_source:type_name -> bytebase.store.TaskReleaseSource
	6, // 3: bytebase.store.TaskDatabaseDataExportPayload.format:type_name -> bytebase.store.ExportFormat
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_store_task_proto_init() }
func file_store_task_proto_init() {
	if File_store_task_proto != nil {
		return
	}
	file_store_common_proto_init()
	file_store_plan_check_run_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_task_proto_goTypes,
		DependencyIndexes: file_store_task_proto_depIdxs,
		MessageInfos:      file_store_task_proto_msgTypes,
	}.Build()
	File_store_task_proto = out.File
	file_store_task_proto_rawDesc = nil
	file_store_task_proto_goTypes = nil
	file_store_task_proto_depIdxs = nil
}