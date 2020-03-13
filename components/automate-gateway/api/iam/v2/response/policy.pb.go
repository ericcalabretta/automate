// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/response/policy.proto

package response

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreatePolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreatePolicyResp) Reset()         { *m = CreatePolicyResp{} }
func (m *CreatePolicyResp) String() string { return proto.CompactTextString(m) }
func (*CreatePolicyResp) ProtoMessage()    {}
func (*CreatePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{0}
}

func (m *CreatePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePolicyResp.Unmarshal(m, b)
}
func (m *CreatePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePolicyResp.Marshal(b, m, deterministic)
}
func (m *CreatePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePolicyResp.Merge(m, src)
}
func (m *CreatePolicyResp) XXX_Size() int {
	return xxx_messageInfo_CreatePolicyResp.Size(m)
}
func (m *CreatePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePolicyResp proto.InternalMessageInfo

func (m *CreatePolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type GetPolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPolicyResp) Reset()         { *m = GetPolicyResp{} }
func (m *GetPolicyResp) String() string { return proto.CompactTextString(m) }
func (*GetPolicyResp) ProtoMessage()    {}
func (*GetPolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{1}
}

func (m *GetPolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyResp.Unmarshal(m, b)
}
func (m *GetPolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyResp.Marshal(b, m, deterministic)
}
func (m *GetPolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyResp.Merge(m, src)
}
func (m *GetPolicyResp) XXX_Size() int {
	return xxx_messageInfo_GetPolicyResp.Size(m)
}
func (m *GetPolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyResp proto.InternalMessageInfo

func (m *GetPolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ListPoliciesResp struct {
	Policies             []*common.Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListPoliciesResp) Reset()         { *m = ListPoliciesResp{} }
func (m *ListPoliciesResp) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesResp) ProtoMessage()    {}
func (*ListPoliciesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{2}
}

func (m *ListPoliciesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesResp.Unmarshal(m, b)
}
func (m *ListPoliciesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesResp.Marshal(b, m, deterministic)
}
func (m *ListPoliciesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesResp.Merge(m, src)
}
func (m *ListPoliciesResp) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesResp.Size(m)
}
func (m *ListPoliciesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesResp proto.InternalMessageInfo

func (m *ListPoliciesResp) GetPolicies() []*common.Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type DeletePolicyResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePolicyResp) Reset()         { *m = DeletePolicyResp{} }
func (m *DeletePolicyResp) String() string { return proto.CompactTextString(m) }
func (*DeletePolicyResp) ProtoMessage()    {}
func (*DeletePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{3}
}

func (m *DeletePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePolicyResp.Unmarshal(m, b)
}
func (m *DeletePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePolicyResp.Marshal(b, m, deterministic)
}
func (m *DeletePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePolicyResp.Merge(m, src)
}
func (m *DeletePolicyResp) XXX_Size() int {
	return xxx_messageInfo_DeletePolicyResp.Size(m)
}
func (m *DeletePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePolicyResp proto.InternalMessageInfo

type UpdatePolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdatePolicyResp) Reset()         { *m = UpdatePolicyResp{} }
func (m *UpdatePolicyResp) String() string { return proto.CompactTextString(m) }
func (*UpdatePolicyResp) ProtoMessage()    {}
func (*UpdatePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{4}
}

func (m *UpdatePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePolicyResp.Unmarshal(m, b)
}
func (m *UpdatePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePolicyResp.Marshal(b, m, deterministic)
}
func (m *UpdatePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePolicyResp.Merge(m, src)
}
func (m *UpdatePolicyResp) XXX_Size() int {
	return xxx_messageInfo_UpdatePolicyResp.Size(m)
}
func (m *UpdatePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePolicyResp proto.InternalMessageInfo

func (m *UpdatePolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type GetPolicyVersionResp struct {
	Version              *common.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetPolicyVersionResp) Reset()         { *m = GetPolicyVersionResp{} }
func (m *GetPolicyVersionResp) String() string { return proto.CompactTextString(m) }
func (*GetPolicyVersionResp) ProtoMessage()    {}
func (*GetPolicyVersionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{5}
}

func (m *GetPolicyVersionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyVersionResp.Unmarshal(m, b)
}
func (m *GetPolicyVersionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyVersionResp.Marshal(b, m, deterministic)
}
func (m *GetPolicyVersionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyVersionResp.Merge(m, src)
}
func (m *GetPolicyVersionResp) XXX_Size() int {
	return xxx_messageInfo_GetPolicyVersionResp.Size(m)
}
func (m *GetPolicyVersionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyVersionResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyVersionResp proto.InternalMessageInfo

func (m *GetPolicyVersionResp) GetVersion() *common.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type ListPolicyMembersResp struct {
	// List of policy members.
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPolicyMembersResp) Reset()         { *m = ListPolicyMembersResp{} }
func (m *ListPolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*ListPolicyMembersResp) ProtoMessage()    {}
func (*ListPolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{6}
}

func (m *ListPolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPolicyMembersResp.Unmarshal(m, b)
}
func (m *ListPolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *ListPolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPolicyMembersResp.Merge(m, src)
}
func (m *ListPolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_ListPolicyMembersResp.Size(m)
}
func (m *ListPolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListPolicyMembersResp proto.InternalMessageInfo

func (m *ListPolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ReplacePolicyMembersResp struct {
	// Resulting list of policy members.
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplacePolicyMembersResp) Reset()         { *m = ReplacePolicyMembersResp{} }
func (m *ReplacePolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*ReplacePolicyMembersResp) ProtoMessage()    {}
func (*ReplacePolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{7}
}

func (m *ReplacePolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplacePolicyMembersResp.Unmarshal(m, b)
}
func (m *ReplacePolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplacePolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *ReplacePolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplacePolicyMembersResp.Merge(m, src)
}
func (m *ReplacePolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_ReplacePolicyMembersResp.Size(m)
}
func (m *ReplacePolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplacePolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReplacePolicyMembersResp proto.InternalMessageInfo

func (m *ReplacePolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type RemovePolicyMembersResp struct {
	// Resulting list of policy members.
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePolicyMembersResp) Reset()         { *m = RemovePolicyMembersResp{} }
func (m *RemovePolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*RemovePolicyMembersResp) ProtoMessage()    {}
func (*RemovePolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{8}
}

func (m *RemovePolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePolicyMembersResp.Unmarshal(m, b)
}
func (m *RemovePolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *RemovePolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePolicyMembersResp.Merge(m, src)
}
func (m *RemovePolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_RemovePolicyMembersResp.Size(m)
}
func (m *RemovePolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePolicyMembersResp proto.InternalMessageInfo

func (m *RemovePolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type AddPolicyMembersResp struct {
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPolicyMembersResp) Reset()         { *m = AddPolicyMembersResp{} }
func (m *AddPolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*AddPolicyMembersResp) ProtoMessage()    {}
func (*AddPolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{9}
}

func (m *AddPolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPolicyMembersResp.Unmarshal(m, b)
}
func (m *AddPolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *AddPolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPolicyMembersResp.Merge(m, src)
}
func (m *AddPolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_AddPolicyMembersResp.Size(m)
}
func (m *AddPolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddPolicyMembersResp proto.InternalMessageInfo

func (m *AddPolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type CreateRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateRoleResp) Reset()         { *m = CreateRoleResp{} }
func (m *CreateRoleResp) String() string { return proto.CompactTextString(m) }
func (*CreateRoleResp) ProtoMessage()    {}
func (*CreateRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{10}
}

func (m *CreateRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleResp.Unmarshal(m, b)
}
func (m *CreateRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleResp.Marshal(b, m, deterministic)
}
func (m *CreateRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleResp.Merge(m, src)
}
func (m *CreateRoleResp) XXX_Size() int {
	return xxx_messageInfo_CreateRoleResp.Size(m)
}
func (m *CreateRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleResp proto.InternalMessageInfo

func (m *CreateRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type ListRolesResp struct {
	Roles                []*common.Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListRolesResp) Reset()         { *m = ListRolesResp{} }
func (m *ListRolesResp) String() string { return proto.CompactTextString(m) }
func (*ListRolesResp) ProtoMessage()    {}
func (*ListRolesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{11}
}

func (m *ListRolesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesResp.Unmarshal(m, b)
}
func (m *ListRolesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesResp.Marshal(b, m, deterministic)
}
func (m *ListRolesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesResp.Merge(m, src)
}
func (m *ListRolesResp) XXX_Size() int {
	return xxx_messageInfo_ListRolesResp.Size(m)
}
func (m *ListRolesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesResp proto.InternalMessageInfo

func (m *ListRolesResp) GetRoles() []*common.Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type GetRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetRoleResp) Reset()         { *m = GetRoleResp{} }
func (m *GetRoleResp) String() string { return proto.CompactTextString(m) }
func (*GetRoleResp) ProtoMessage()    {}
func (*GetRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{12}
}

func (m *GetRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleResp.Unmarshal(m, b)
}
func (m *GetRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleResp.Marshal(b, m, deterministic)
}
func (m *GetRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleResp.Merge(m, src)
}
func (m *GetRoleResp) XXX_Size() int {
	return xxx_messageInfo_GetRoleResp.Size(m)
}
func (m *GetRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleResp proto.InternalMessageInfo

func (m *GetRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type DeleteRoleResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoleResp) Reset()         { *m = DeleteRoleResp{} }
func (m *DeleteRoleResp) String() string { return proto.CompactTextString(m) }
func (*DeleteRoleResp) ProtoMessage()    {}
func (*DeleteRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{13}
}

func (m *DeleteRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoleResp.Unmarshal(m, b)
}
func (m *DeleteRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoleResp.Marshal(b, m, deterministic)
}
func (m *DeleteRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoleResp.Merge(m, src)
}
func (m *DeleteRoleResp) XXX_Size() int {
	return xxx_messageInfo_DeleteRoleResp.Size(m)
}
func (m *DeleteRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoleResp proto.InternalMessageInfo

type UpdateRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateRoleResp) Reset()         { *m = UpdateRoleResp{} }
func (m *UpdateRoleResp) String() string { return proto.CompactTextString(m) }
func (*UpdateRoleResp) ProtoMessage()    {}
func (*UpdateRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{14}
}

func (m *UpdateRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRoleResp.Unmarshal(m, b)
}
func (m *UpdateRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRoleResp.Marshal(b, m, deterministic)
}
func (m *UpdateRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRoleResp.Merge(m, src)
}
func (m *UpdateRoleResp) XXX_Size() int {
	return xxx_messageInfo_UpdateRoleResp.Size(m)
}
func (m *UpdateRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRoleResp proto.InternalMessageInfo

func (m *UpdateRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type CreateProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateProjectResp) Reset()         { *m = CreateProjectResp{} }
func (m *CreateProjectResp) String() string { return proto.CompactTextString(m) }
func (*CreateProjectResp) ProtoMessage()    {}
func (*CreateProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{15}
}

func (m *CreateProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectResp.Unmarshal(m, b)
}
func (m *CreateProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectResp.Marshal(b, m, deterministic)
}
func (m *CreateProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectResp.Merge(m, src)
}
func (m *CreateProjectResp) XXX_Size() int {
	return xxx_messageInfo_CreateProjectResp.Size(m)
}
func (m *CreateProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectResp proto.InternalMessageInfo

func (m *CreateProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type UpdateProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateProjectResp) Reset()         { *m = UpdateProjectResp{} }
func (m *UpdateProjectResp) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectResp) ProtoMessage()    {}
func (*UpdateProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{16}
}

func (m *UpdateProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectResp.Unmarshal(m, b)
}
func (m *UpdateProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectResp.Marshal(b, m, deterministic)
}
func (m *UpdateProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectResp.Merge(m, src)
}
func (m *UpdateProjectResp) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectResp.Size(m)
}
func (m *UpdateProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectResp proto.InternalMessageInfo

func (m *UpdateProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type GetProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetProjectResp) Reset()         { *m = GetProjectResp{} }
func (m *GetProjectResp) String() string { return proto.CompactTextString(m) }
func (*GetProjectResp) ProtoMessage()    {}
func (*GetProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{17}
}

func (m *GetProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectResp.Unmarshal(m, b)
}
func (m *GetProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectResp.Marshal(b, m, deterministic)
}
func (m *GetProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectResp.Merge(m, src)
}
func (m *GetProjectResp) XXX_Size() int {
	return xxx_messageInfo_GetProjectResp.Size(m)
}
func (m *GetProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectResp proto.InternalMessageInfo

func (m *GetProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type ListProjectsResp struct {
	Projects             []*common.Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListProjectsResp) Reset()         { *m = ListProjectsResp{} }
func (m *ListProjectsResp) String() string { return proto.CompactTextString(m) }
func (*ListProjectsResp) ProtoMessage()    {}
func (*ListProjectsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{18}
}

func (m *ListProjectsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsResp.Unmarshal(m, b)
}
func (m *ListProjectsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsResp.Marshal(b, m, deterministic)
}
func (m *ListProjectsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsResp.Merge(m, src)
}
func (m *ListProjectsResp) XXX_Size() int {
	return xxx_messageInfo_ListProjectsResp.Size(m)
}
func (m *ListProjectsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsResp proto.InternalMessageInfo

func (m *ListProjectsResp) GetProjects() []*common.Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeleteProjectResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectResp) Reset()         { *m = DeleteProjectResp{} }
func (m *DeleteProjectResp) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectResp) ProtoMessage()    {}
func (*DeleteProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{19}
}

func (m *DeleteProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectResp.Unmarshal(m, b)
}
func (m *DeleteProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectResp.Marshal(b, m, deterministic)
}
func (m *DeleteProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectResp.Merge(m, src)
}
func (m *DeleteProjectResp) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectResp.Size(m)
}
func (m *DeleteProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectResp proto.InternalMessageInfo

type UpgradeToV2Resp struct {
	Reports              []string `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpgradeToV2Resp) Reset()         { *m = UpgradeToV2Resp{} }
func (m *UpgradeToV2Resp) String() string { return proto.CompactTextString(m) }
func (*UpgradeToV2Resp) ProtoMessage()    {}
func (*UpgradeToV2Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{20}
}

func (m *UpgradeToV2Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeToV2Resp.Unmarshal(m, b)
}
func (m *UpgradeToV2Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeToV2Resp.Marshal(b, m, deterministic)
}
func (m *UpgradeToV2Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeToV2Resp.Merge(m, src)
}
func (m *UpgradeToV2Resp) XXX_Size() int {
	return xxx_messageInfo_UpgradeToV2Resp.Size(m)
}
func (m *UpgradeToV2Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeToV2Resp.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeToV2Resp proto.InternalMessageInfo

func (m *UpgradeToV2Resp) GetReports() []string {
	if m != nil {
		return m.Reports
	}
	return nil
}

type ResetToV1Resp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetToV1Resp) Reset()         { *m = ResetToV1Resp{} }
func (m *ResetToV1Resp) String() string { return proto.CompactTextString(m) }
func (*ResetToV1Resp) ProtoMessage()    {}
func (*ResetToV1Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{21}
}

func (m *ResetToV1Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetToV1Resp.Unmarshal(m, b)
}
func (m *ResetToV1Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetToV1Resp.Marshal(b, m, deterministic)
}
func (m *ResetToV1Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetToV1Resp.Merge(m, src)
}
func (m *ResetToV1Resp) XXX_Size() int {
	return xxx_messageInfo_ResetToV1Resp.Size(m)
}
func (m *ResetToV1Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetToV1Resp.DiscardUnknown(m)
}

var xxx_messageInfo_ResetToV1Resp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreatePolicyResp)(nil), "chef.automate.api.iam.v2.CreatePolicyResp")
	proto.RegisterType((*GetPolicyResp)(nil), "chef.automate.api.iam.v2.GetPolicyResp")
	proto.RegisterType((*ListPoliciesResp)(nil), "chef.automate.api.iam.v2.ListPoliciesResp")
	proto.RegisterType((*DeletePolicyResp)(nil), "chef.automate.api.iam.v2.DeletePolicyResp")
	proto.RegisterType((*UpdatePolicyResp)(nil), "chef.automate.api.iam.v2.UpdatePolicyResp")
	proto.RegisterType((*GetPolicyVersionResp)(nil), "chef.automate.api.iam.v2.GetPolicyVersionResp")
	proto.RegisterType((*ListPolicyMembersResp)(nil), "chef.automate.api.iam.v2.ListPolicyMembersResp")
	proto.RegisterType((*ReplacePolicyMembersResp)(nil), "chef.automate.api.iam.v2.ReplacePolicyMembersResp")
	proto.RegisterType((*RemovePolicyMembersResp)(nil), "chef.automate.api.iam.v2.RemovePolicyMembersResp")
	proto.RegisterType((*AddPolicyMembersResp)(nil), "chef.automate.api.iam.v2.AddPolicyMembersResp")
	proto.RegisterType((*CreateRoleResp)(nil), "chef.automate.api.iam.v2.CreateRoleResp")
	proto.RegisterType((*ListRolesResp)(nil), "chef.automate.api.iam.v2.ListRolesResp")
	proto.RegisterType((*GetRoleResp)(nil), "chef.automate.api.iam.v2.GetRoleResp")
	proto.RegisterType((*DeleteRoleResp)(nil), "chef.automate.api.iam.v2.DeleteRoleResp")
	proto.RegisterType((*UpdateRoleResp)(nil), "chef.automate.api.iam.v2.UpdateRoleResp")
	proto.RegisterType((*CreateProjectResp)(nil), "chef.automate.api.iam.v2.CreateProjectResp")
	proto.RegisterType((*UpdateProjectResp)(nil), "chef.automate.api.iam.v2.UpdateProjectResp")
	proto.RegisterType((*GetProjectResp)(nil), "chef.automate.api.iam.v2.GetProjectResp")
	proto.RegisterType((*ListProjectsResp)(nil), "chef.automate.api.iam.v2.ListProjectsResp")
	proto.RegisterType((*DeleteProjectResp)(nil), "chef.automate.api.iam.v2.DeleteProjectResp")
	proto.RegisterType((*UpgradeToV2Resp)(nil), "chef.automate.api.iam.v2.UpgradeToV2Resp")
	proto.RegisterType((*ResetToV1Resp)(nil), "chef.automate.api.iam.v2.ResetToV1Resp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/response/policy.proto", fileDescriptor_b94e203b4b99ed5f)
}

var fileDescriptor_b94e203b4b99ed5f = []byte{
	// 1042 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x98, 0xef, 0x8b, 0x1b, 0xc5,
	0x1b, 0xc0, 0xd9, 0x4d, 0xbf, 0x6d, 0xbf, 0x53, 0xee, 0x7a, 0x5d, 0x2a, 0x86, 0xbe, 0x90, 0x75,
	0x10, 0xaa, 0x6d, 0x92, 0x25, 0xe3, 0xaf, 0x3a, 0x5e, 0x5f, 0xac, 0x15, 0x2a, 0x72, 0x87, 0x35,
	0xb6, 0x11, 0xd4, 0x7b, 0x31, 0xb7, 0x79, 0x2e, 0xdd, 0x92, 0xdd, 0xd9, 0xee, 0x4c, 0x12, 0x8e,
	0x10, 0x10, 0x2a, 0x15, 0xab, 0x54, 0x58, 0xff, 0x01, 0xc1, 0xd7, 0xbe, 0x16, 0xfc, 0x23, 0x44,
	0x51, 0x44, 0x0a, 0xfa, 0xc2, 0xbf, 0x44, 0xe6, 0xc7, 0xe6, 0x2e, 0x31, 0xb9, 0xde, 0xa9, 0x3d,
	0x5b, 0xe9, 0xab, 0x64, 0x76, 0x9e, 0x5f, 0xf3, 0x99, 0x99, 0xe7, 0x79, 0x76, 0xd1, 0x6a, 0xc4,
	0x93, 0x8c, 0xa7, 0x90, 0x4a, 0x11, 0xb0, 0xbe, 0xe4, 0x09, 0x93, 0x50, 0xef, 0x32, 0x09, 0x43,
	0xb6, 0x1d, 0xb0, 0x2c, 0x0e, 0x62, 0x96, 0x04, 0x03, 0x12, 0xe4, 0x20, 0x32, 0x9e, 0x0a, 0x08,
	0x32, 0xde, 0x8b, 0xa3, 0xed, 0x46, 0x96, 0x73, 0xc9, 0xbd, 0x6a, 0x74, 0x1d, 0xb6, 0x1a, 0xa5,
	0x5e, 0x83, 0x65, 0x71, 0x23, 0x66, 0x49, 0x63, 0x40, 0xce, 0xd0, 0x7d, 0xda, 0x8d, 0x78, 0x92,
	0xf0, 0x74, 0xca, 0xea, 0x99, 0x9a, 0xfe, 0x89, 0xea, 0x5d, 0x48, 0xeb, 0x62, 0xc8, 0xba, 0x5d,
	0xc8, 0x03, 0x9e, 0xc9, 0x98, 0xa7, 0x22, 0x60, 0x69, 0xca, 0x25, 0xd3, 0xff, 0x8d, 0x34, 0xfe,
	0xd5, 0x45, 0x2b, 0x97, 0x72, 0x60, 0x12, 0xae, 0x68, 0x23, 0x2d, 0x10, 0x99, 0x77, 0x01, 0x1d,
	0x35, 0x26, 0xab, 0x8e, 0xef, 0x3c, 0x7b, 0x82, 0xf8, 0x8d, 0x45, 0x91, 0x36, 0xac, 0x96, 0x95,
	0xa7, 0x77, 0xdd, 0x22, 0xbc, 0xe3, 0x92, 0x8f, 0x5d, 0xef, 0x23, 0x77, 0x84, 0x53, 0x96, 0x00,
	0xa6, 0x3e, 0x5e, 0xdf, 0xf6, 0xdb, 0x31, 0x0c, 0x21, 0xf7, 0x8d, 0x38, 0xae, 0xe1, 0xb8, 0xa3,
	0x26, 0xa2, 0xbe, 0x90, 0x3c, 0xa9, 0x0f, 0xf4, 0x64, 0x3d, 0x2b, 0x27, 0x13, 0x48, 0x36, 0x21,
	0x17, 0x98, 0xfa, 0xef, 0x63, 0x09, 0x2c, 0xa1, 0xbd, 0x0e, 0xcb, 0xe8, 0x39, 0xbc, 0x51, 0xf3,
	0xb1, 0x90, 0x4c, 0x42, 0xa2, 0xc8, 0xa8, 0xe9, 0x11, 0xce, 0x79, 0x4f, 0xfb, 0x30, 0x36, 0x70,
	0x0d, 0x67, 0x39, 0xbf, 0x01, 0x91, 0x99, 0x2e, 0x07, 0x4d, 0x5c, 0xf3, 0xcb, 0xff, 0x44, 0x9b,
	0x81, 0xad, 0x2d, 0x88, 0xa4, 0x52, 0x0c, 0xd7, 0xd6, 0xde, 0x7a, 0x17, 0x8f, 0x6b, 0xfe, 0x8e,
	0xad, 0x9b, 0x6c, 0xd6, 0x0e, 0x8b, 0x22, 0xc8, 0x24, 0x4b, 0x23, 0xb8, 0x62, 0x1e, 0xcf, 0x37,
	0xb2, 0x31, 0xad, 0xb6, 0x31, 0xc6, 0xf7, 0x5c, 0xb4, 0x74, 0x19, 0xe4, 0x63, 0xb8, 0x0f, 0x02,
	0xee, 0xd7, 0x47, 0xd0, 0xca, 0x5a, 0x2c, 0x0c, 0xdd, 0x18, 0x84, 0xe6, 0xbb, 0x8a, 0x8e, 0x67,
	0x76, 0x5c, 0x75, 0xfc, 0xca, 0xbe, 0x08, 0x4f, 0x34, 0xe8, 0xbd, 0x4a, 0x11, 0xfe, 0x5c, 0x21,
	0x3f, 0x56, 0xbc, 0xef, 0x2b, 0x23, 0x5c, 0x3e, 0x36, 0x34, 0x16, 0x11, 0xf7, 0x9b, 0x7b, 0x32,
	0xaf, 0x37, 0xff, 0x53, 0xd4, 0x95, 0xd5, 0x5d, 0x24, 0x2e, 0xe9, 0x15, 0x97, 0x24, 0xc8, 0x2c,
	0x09, 0x8b, 0x80, 0xcc, 0x45, 0xc0, 0x23, 0xd6, 0xa3, 0x12, 0x84, 0xdc, 0x13, 0x03, 0xeb, 0x77,
	0x62, 0xc9, 0xf3, 0x5d, 0xeb, 0x9d, 0x01, 0xb1, 0xbf, 0xb8, 0x37, 0xc6, 0xf8, 0x2b, 0x17, 0xad,
	0xbc, 0x0e, 0x3d, 0xd8, 0x9d, 0xec, 0x1e, 0xdf, 0xaa, 0xd9, 0x5b, 0xf5, 0xbb, 0x8b, 0x56, 0xae,
	0x65, 0x9d, 0x7f, 0xaa, 0x24, 0x7c, 0xe1, 0x16, 0xe1, 0xe7, 0x2e, 0xf9, 0xcc, 0xf5, 0x3e, 0x99,
	0xe6, 0x6b, 0x7c, 0x74, 0x66, 0x39, 0xfb, 0x53, 0x2c, 0xfb, 0x02, 0x72, 0xc3, 0x32, 0x85, 0xa1,
	0x1a, 0x28, 0x81, 0x1d, 0xc0, 0x29, 0x0c, 0xd5, 0xe0, 0xa0, 0x98, 0x0f, 0x44, 0xf9, 0x81, 0x40,
	0xbe, 0xe5, 0xa0, 0xd3, 0x93, 0xba, 0xd0, 0x86, 0x5c, 0xc4, 0x3c, 0xd5, 0xa0, 0x5f, 0x45, 0xc7,
	0x06, 0x66, 0x68, 0x49, 0x3f, 0xbd, 0x98, 0x74, 0xa9, 0x57, 0x6a, 0xd0, 0x46, 0x11, 0x9e, 0x27,
	0xcf, 0x79, 0x67, 0x47, 0xd8, 0x3e, 0xc1, 0x74, 0x84, 0x13, 0x76, 0x83, 0xe7, 0x98, 0xe2, 0xb6,
	0xbe, 0x84, 0x71, 0x6a, 0x06, 0x4d, 0x3c, 0x1e, 0xe3, 0x0c, 0x3d, 0x31, 0xc9, 0x9f, 0xdb, 0xeb,
	0x06, 0xbb, 0x8e, 0xa2, 0x8a, 0x8e, 0xd9, 0x5d, 0xd0, 0x39, 0xf4, 0xff, 0xad, 0x72, 0x48, 0x2f,
	0x16, 0x21, 0x25, 0x17, 0xbc, 0x97, 0x46, 0x8b, 0xee, 0xb3, 0xc1, 0x2d, 0x14, 0x58, 0xb3, 0x79,
	0xbb, 0x6e, 0xf9, 0x18, 0x0b, 0x54, 0x6d, 0x41, 0xd6, 0x63, 0x11, 0x1c, 0xa2, 0xd3, 0x0f, 0xd0,
	0x93, 0x2d, 0x48, 0xf8, 0xe0, 0x40, 0x3e, 0xcf, 0x16, 0xe1, 0x33, 0x04, 0x7b, 0xfe, 0x68, 0xce,
	0x79, 0x9c, 0xb2, 0xce, 0xd1, 0xe9, 0xb0, 0xd3, 0x39, 0xc4, 0xe5, 0x7c, 0xeb, 0xa0, 0x65, 0xd3,
	0xb3, 0xb5, 0x78, 0x0f, 0xb4, 0x2f, 0x82, 0x8e, 0xa8, 0x43, 0x6a, 0x8f, 0xcc, 0x53, 0x8b, 0x8f,
	0x8c, 0xd6, 0xd0, 0xb2, 0x54, 0x16, 0xe1, 0x4d, 0xc2, 0xbd, 0x64, 0x34, 0x9d, 0xd9, 0xf4, 0x61,
	0xaf, 0xf9, 0x73, 0x72, 0xfc, 0x3b, 0x10, 0xe5, 0x20, 0xfd, 0x75, 0x96, 0xb2, 0x2e, 0xe4, 0x7e,
	0xcb, 0x0a, 0xb2, 0x48, 0x37, 0x94, 0x7a, 0x0d, 0x42, 0x8b, 0x08, 0x7a, 0x4e, 0x4d, 0xc4, 0x2c,
	0xa1, 0xe5, 0xb1, 0xa7, 0xbd, 0xd8, 0x04, 0xff, 0xa5, 0x8b, 0x96, 0xd4, 0x99, 0x53, 0xda, 0x86,
	0xd3, 0x0b, 0xe8, 0x7f, 0xca, 0x67, 0x59, 0xad, 0xef, 0x17, 0xbc, 0x11, 0xa6, 0xbf, 0x38, 0x45,
	0xf8, 0x93, 0x43, 0x7e, 0x70, 0xbc, 0xef, 0x1c, 0x73, 0x43, 0xed, 0xb5, 0x3f, 0xc4, 0xc5, 0xa8,
	0x04, 0xfc, 0x27, 0x77, 0xaa, 0xd6, 0xed, 0xe1, 0x50, 0x39, 0xf2, 0xc9, 0xac, 0xab, 0x38, 0xdd,
	0xca, 0x99, 0x2a, 0x17, 0xaa, 0x4c, 0x7d, 0xe3, 0xa0, 0x13, 0x97, 0x41, 0x3e, 0x82, 0x9b, 0x7b,
	0xdb, 0x41, 0xcb, 0xa6, 0xc0, 0x96, 0xc1, 0xff, 0x4b, 0x81, 0xa8, 0x2b, 0x62, 0xea, 0xcb, 0x23,
	0x48, 0xf1, 0xae, 0x83, 0x4e, 0xd9, 0x77, 0x32, 0x33, 0x55, 0x16, 0x06, 0x2b, 0x79, 0xff, 0xc2,
	0x50, 0xea, 0x95, 0x1a, 0x74, 0xb5, 0x08, 0x5f, 0x21, 0x2f, 0x7b, 0x2f, 0xce, 0x6d, 0xdd, 0x6c,
	0xe5, 0x52, 0xe1, 0x4c, 0xb5, 0x6e, 0xf6, 0xb9, 0x09, 0xc8, 0x76, 0x04, 0x0f, 0x47, 0x40, 0x9f,
	0x3a, 0x68, 0x59, 0x55, 0xcf, 0x87, 0x23, 0x9a, 0xdf, 0x1c, 0xfb, 0x1a, 0x62, 0x37, 0x52, 0xc7,
	0x73, 0x11, 0x1d, 0x2f, 0x37, 0xd6, 0x26, 0xb6, 0x7d, 0x04, 0x34, 0x51, 0xa1, 0x77, 0x9c, 0x22,
	0xbc, 0xed, 0x90, 0x5b, 0x8e, 0xf7, 0xa1, 0x33, 0x9a, 0x6a, 0x1e, 0xfe, 0x52, 0x8c, 0x0b, 0x9a,
	0x76, 0x33, 0x6d, 0xd2, 0xd1, 0x3c, 0xc5, 0x3a, 0xc1, 0x2a, 0x21, 0xbd, 0x8d, 0x4e, 0xd9, 0xb6,
	0x79, 0x07, 0xf8, 0xdf, 0x64, 0x76, 0x1e, 0x9d, 0xbc, 0x96, 0x75, 0x73, 0xd6, 0x81, 0xab, 0xbc,
	0x4d, 0xca, 0x7a, 0x99, 0x43, 0xc6, 0x73, 0x39, 0xa9, 0x97, 0x76, 0x88, 0x4f, 0xa2, 0xa5, 0x16,
	0x08, 0x90, 0x57, 0x79, 0xbb, 0xa9, 0x44, 0x5f, 0x7b, 0xf3, 0xbd, 0x37, 0xba, 0xb1, 0xbc, 0xde,
	0xdf, 0x6c, 0x44, 0x3c, 0x09, 0x14, 0xd6, 0xc9, 0x67, 0x92, 0xe0, 0x80, 0x9f, 0x64, 0x36, 0x8f,
	0xea, 0x0f, 0x21, 0xcf, 0xff, 0x11, 0x00, 0x00, 0xff, 0xff, 0x81, 0x12, 0x3d, 0x5b, 0xcc, 0x11,
	0x00, 0x00,
}
