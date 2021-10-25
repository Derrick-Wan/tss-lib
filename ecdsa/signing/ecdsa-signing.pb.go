// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: protob/ecdsa-signing.proto

package signing

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

//
// Represents a P2P message sent to each party during Round 1 of the ECDSA TSS signing protocol.
type SignRound1Message1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C               []byte   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	RangeProofAlice [][]byte `protobuf:"bytes,2,rep,name=range_proof_alice,json=rangeProofAlice,proto3" json:"range_proof_alice,omitempty"`
}

func (x *SignRound1Message1) Reset() {
	*x = SignRound1Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound1Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound1Message1) ProtoMessage() {}

func (x *SignRound1Message1) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound1Message1.ProtoReflect.Descriptor instead.
func (*SignRound1Message1) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{0}
}

func (x *SignRound1Message1) GetC() []byte {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *SignRound1Message1) GetRangeProofAlice() [][]byte {
	if x != nil {
		return x.RangeProofAlice
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 1 of the ECDSA TSS signing protocol.
type SignRound1Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
}

func (x *SignRound1Message2) Reset() {
	*x = SignRound1Message2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound1Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound1Message2) ProtoMessage() {}

func (x *SignRound1Message2) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound1Message2.ProtoReflect.Descriptor instead.
func (*SignRound1Message2) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{1}
}

func (x *SignRound1Message2) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

//
// Represents a P2P message sent to each party during Round 2 of the ECDSA TSS signing protocol.
type SignRound2Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C1         []byte   `protobuf:"bytes,1,opt,name=c1,proto3" json:"c1,omitempty"`
	C2         []byte   `protobuf:"bytes,2,opt,name=c2,proto3" json:"c2,omitempty"`
	ProofBob   [][]byte `protobuf:"bytes,3,rep,name=proof_bob,json=proofBob,proto3" json:"proof_bob,omitempty"`
	ProofBobWc [][]byte `protobuf:"bytes,4,rep,name=proof_bob_wc,json=proofBobWc,proto3" json:"proof_bob_wc,omitempty"`
}

func (x *SignRound2Message) Reset() {
	*x = SignRound2Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound2Message) ProtoMessage() {}

func (x *SignRound2Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound2Message.ProtoReflect.Descriptor instead.
func (*SignRound2Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{2}
}

func (x *SignRound2Message) GetC1() []byte {
	if x != nil {
		return x.C1
	}
	return nil
}

func (x *SignRound2Message) GetC2() []byte {
	if x != nil {
		return x.C2
	}
	return nil
}

func (x *SignRound2Message) GetProofBob() [][]byte {
	if x != nil {
		return x.ProofBob
	}
	return nil
}

func (x *SignRound2Message) GetProofBobWc() [][]byte {
	if x != nil {
		return x.ProofBobWc
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 3 of the ECDSA TSS signing protocol.
type SignRound3Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Theta []byte `protobuf:"bytes,1,opt,name=theta,proto3" json:"theta,omitempty"`
}

func (x *SignRound3Message) Reset() {
	*x = SignRound3Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound3Message) ProtoMessage() {}

func (x *SignRound3Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound3Message.ProtoReflect.Descriptor instead.
func (*SignRound3Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{3}
}

func (x *SignRound3Message) GetTheta() []byte {
	if x != nil {
		return x.Theta
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 4 of the ECDSA TSS signing protocol.
type SignRound4Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeCommitment [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	ProofAlphaX  []byte   `protobuf:"bytes,2,opt,name=proof_alpha_x,json=proofAlphaX,proto3" json:"proof_alpha_x,omitempty"`
	ProofAlphaY  []byte   `protobuf:"bytes,3,opt,name=proof_alpha_y,json=proofAlphaY,proto3" json:"proof_alpha_y,omitempty"`
	ProofT       []byte   `protobuf:"bytes,4,opt,name=proof_t,json=proofT,proto3" json:"proof_t,omitempty"`
}

func (x *SignRound4Message) Reset() {
	*x = SignRound4Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound4Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound4Message) ProtoMessage() {}

func (x *SignRound4Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound4Message.ProtoReflect.Descriptor instead.
func (*SignRound4Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{4}
}

func (x *SignRound4Message) GetDeCommitment() [][]byte {
	if x != nil {
		return x.DeCommitment
	}
	return nil
}

func (x *SignRound4Message) GetProofAlphaX() []byte {
	if x != nil {
		return x.ProofAlphaX
	}
	return nil
}

func (x *SignRound4Message) GetProofAlphaY() []byte {
	if x != nil {
		return x.ProofAlphaY
	}
	return nil
}

func (x *SignRound4Message) GetProofT() []byte {
	if x != nil {
		return x.ProofT
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 5 of the ECDSA TSS signing protocol.
type SignRound5Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
}

func (x *SignRound5Message) Reset() {
	*x = SignRound5Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound5Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound5Message) ProtoMessage() {}

func (x *SignRound5Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound5Message.ProtoReflect.Descriptor instead.
func (*SignRound5Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{5}
}

func (x *SignRound5Message) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 6 of the ECDSA TSS signing protocol.
type SignRound6Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeCommitment [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	ProofAlphaX  []byte   `protobuf:"bytes,2,opt,name=proof_alpha_x,json=proofAlphaX,proto3" json:"proof_alpha_x,omitempty"`
	ProofAlphaY  []byte   `protobuf:"bytes,3,opt,name=proof_alpha_y,json=proofAlphaY,proto3" json:"proof_alpha_y,omitempty"`
	ProofT       []byte   `protobuf:"bytes,4,opt,name=proof_t,json=proofT,proto3" json:"proof_t,omitempty"`
	VProofAlphaX []byte   `protobuf:"bytes,5,opt,name=v_proof_alpha_x,json=vProofAlphaX,proto3" json:"v_proof_alpha_x,omitempty"`
	VProofAlphaY []byte   `protobuf:"bytes,6,opt,name=v_proof_alpha_y,json=vProofAlphaY,proto3" json:"v_proof_alpha_y,omitempty"`
	VProofT      []byte   `protobuf:"bytes,7,opt,name=v_proof_t,json=vProofT,proto3" json:"v_proof_t,omitempty"`
	VProofU      []byte   `protobuf:"bytes,8,opt,name=v_proof_u,json=vProofU,proto3" json:"v_proof_u,omitempty"`
}

func (x *SignRound6Message) Reset() {
	*x = SignRound6Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound6Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound6Message) ProtoMessage() {}

func (x *SignRound6Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound6Message.ProtoReflect.Descriptor instead.
func (*SignRound6Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{6}
}

func (x *SignRound6Message) GetDeCommitment() [][]byte {
	if x != nil {
		return x.DeCommitment
	}
	return nil
}

func (x *SignRound6Message) GetProofAlphaX() []byte {
	if x != nil {
		return x.ProofAlphaX
	}
	return nil
}

func (x *SignRound6Message) GetProofAlphaY() []byte {
	if x != nil {
		return x.ProofAlphaY
	}
	return nil
}

func (x *SignRound6Message) GetProofT() []byte {
	if x != nil {
		return x.ProofT
	}
	return nil
}

func (x *SignRound6Message) GetVProofAlphaX() []byte {
	if x != nil {
		return x.VProofAlphaX
	}
	return nil
}

func (x *SignRound6Message) GetVProofAlphaY() []byte {
	if x != nil {
		return x.VProofAlphaY
	}
	return nil
}

func (x *SignRound6Message) GetVProofT() []byte {
	if x != nil {
		return x.VProofT
	}
	return nil
}

func (x *SignRound6Message) GetVProofU() []byte {
	if x != nil {
		return x.VProofU
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 7 of the ECDSA TSS signing protocol.
type SignRound7Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
}

func (x *SignRound7Message) Reset() {
	*x = SignRound7Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound7Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound7Message) ProtoMessage() {}

func (x *SignRound7Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound7Message.ProtoReflect.Descriptor instead.
func (*SignRound7Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{7}
}

func (x *SignRound7Message) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 8 of the ECDSA TSS signing protocol.
type SignRound8Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeCommitment [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
}

func (x *SignRound8Message) Reset() {
	*x = SignRound8Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound8Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound8Message) ProtoMessage() {}

func (x *SignRound8Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound8Message.ProtoReflect.Descriptor instead.
func (*SignRound8Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{8}
}

func (x *SignRound8Message) GetDeCommitment() [][]byte {
	if x != nil {
		return x.DeCommitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 9 of the ECDSA TSS signing protocol.
type SignRound9Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S []byte `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *SignRound9Message) Reset() {
	*x = SignRound9Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound9Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound9Message) ProtoMessage() {}

func (x *SignRound9Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound9Message.ProtoReflect.Descriptor instead.
func (*SignRound9Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{9}
}

func (x *SignRound9Message) GetS() []byte {
	if x != nil {
		return x.S
	}
	return nil
}

var File_protob_ecdsa_signing_proto protoreflect.FileDescriptor

var file_protob_ecdsa_signing_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2d, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x62, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x73, 0x73, 0x6c, 0x69, 0x62, 0x2e, 0x65, 0x63, 0x64,
	0x73, 0x61, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x4e, 0x0a, 0x12, 0x53, 0x69,
	0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31,
	0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x63, 0x12, 0x2a,
	0x0a, 0x11, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x34, 0x0a, 0x12, 0x53, 0x69,
	0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x72, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x63, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x63, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x62,
	0x6f, 0x62, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x42,
	0x6f, 0x62, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x62, 0x6f, 0x62, 0x5f,
	0x77, 0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x42,
	0x6f, 0x62, 0x57, 0x63, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x68, 0x65, 0x74, 0x61, 0x22,
	0x99, 0x01, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x34, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x58, 0x12, 0x22,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70, 0x68,
	0x61, 0x59, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x54, 0x22, 0x33, 0x0a, 0x11, 0x53,
	0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x35, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x9f, 0x02, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x36, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x64,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x58, 0x12,
	0x22, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x59, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x54, 0x12, 0x25, 0x0a, 0x0f,
	0x76, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x76, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x58, 0x12, 0x25, 0x0a, 0x0f, 0x76, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x5f, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x76, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x59, 0x12, 0x1a, 0x0a, 0x09, 0x76, 0x5f,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x76,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x54, 0x12, 0x1a, 0x0a, 0x09, 0x76, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x5f, 0x75, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x76, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x55, 0x22, 0x33, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x37,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x38, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x21, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x39, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x01, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2f, 0x73, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protob_ecdsa_signing_proto_rawDescOnce sync.Once
	file_protob_ecdsa_signing_proto_rawDescData = file_protob_ecdsa_signing_proto_rawDesc
)

func file_protob_ecdsa_signing_proto_rawDescGZIP() []byte {
	file_protob_ecdsa_signing_proto_rawDescOnce.Do(func() {
		file_protob_ecdsa_signing_proto_rawDescData = protoimpl.X.CompressGZIP(file_protob_ecdsa_signing_proto_rawDescData)
	})
	return file_protob_ecdsa_signing_proto_rawDescData
}

var file_protob_ecdsa_signing_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protob_ecdsa_signing_proto_goTypes = []interface{}{
	(*SignRound1Message1)(nil), // 0: binance.tsslib.ecdsa.signing.SignRound1Message1
	(*SignRound1Message2)(nil), // 1: binance.tsslib.ecdsa.signing.SignRound1Message2
	(*SignRound2Message)(nil),  // 2: binance.tsslib.ecdsa.signing.SignRound2Message
	(*SignRound3Message)(nil),  // 3: binance.tsslib.ecdsa.signing.SignRound3Message
	(*SignRound4Message)(nil),  // 4: binance.tsslib.ecdsa.signing.SignRound4Message
	(*SignRound5Message)(nil),  // 5: binance.tsslib.ecdsa.signing.SignRound5Message
	(*SignRound6Message)(nil),  // 6: binance.tsslib.ecdsa.signing.SignRound6Message
	(*SignRound7Message)(nil),  // 7: binance.tsslib.ecdsa.signing.SignRound7Message
	(*SignRound8Message)(nil),  // 8: binance.tsslib.ecdsa.signing.SignRound8Message
	(*SignRound9Message)(nil),  // 9: binance.tsslib.ecdsa.signing.SignRound9Message
}
var file_protob_ecdsa_signing_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protob_ecdsa_signing_proto_init() }
func file_protob_ecdsa_signing_proto_init() {
	if File_protob_ecdsa_signing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protob_ecdsa_signing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound1Message1); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound1Message2); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound2Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound3Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound4Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound5Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound6Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound7Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound8Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound9Message); i {
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
			RawDescriptor: file_protob_ecdsa_signing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protob_ecdsa_signing_proto_goTypes,
		DependencyIndexes: file_protob_ecdsa_signing_proto_depIdxs,
		MessageInfos:      file_protob_ecdsa_signing_proto_msgTypes,
	}.Build()
	File_protob_ecdsa_signing_proto = out.File
	file_protob_ecdsa_signing_proto_rawDesc = nil
	file_protob_ecdsa_signing_proto_goTypes = nil
	file_protob_ecdsa_signing_proto_depIdxs = nil
}
