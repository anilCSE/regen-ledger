// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package group

import (
	context "context"
	types "github.com/regen-network/regen-ledger/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateGroup creates a new group with an admin account address, a list of members and an comment.
	CreateGroup(ctx context.Context, in *MsgCreateGroupRequest, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error)
	// UpdateGroupMembers updates the group members with given group id and admin address.
	UpdateGroupMembers(ctx context.Context, in *MsgUpdateGroupMembersRequest, opts ...grpc.CallOption) (*MsgUpdateGroupMembersResponse, error)
	// UpdateGroupAdmin updates the group admin with given group id and previous admin address.
	UpdateGroupAdmin(ctx context.Context, in *MsgUpdateGroupAdminRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAdminResponse, error)
	// UpdateGroupComment updates the group comment with given group id and admin address.
	UpdateGroupComment(ctx context.Context, in *MsgUpdateGroupCommentRequest, opts ...grpc.CallOption) (*MsgUpdateGroupCommentResponse, error)
	// CreateGroupAccount creates a new group account using given DecisionPolicy.
	CreateGroupAccount(ctx context.Context, in *MsgCreateGroupAccountRequest, opts ...grpc.CallOption) (*MsgCreateGroupAccountResponse, error)
	// UpdateGroupAccountAdmin updates a group account admin.
	UpdateGroupAccountAdmin(ctx context.Context, in *MsgUpdateGroupAccountAdminRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAccountAdminResponse, error)
	// UpdateGroupAccountDecisionPolicy allows a group account decision policy to be updated.
	UpdateGroupAccountDecisionPolicy(ctx context.Context, in *MsgUpdateGroupAccountDecisionPolicyRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAccountDecisionPolicyResponse, error)
	// UpdateGroupAccountComment updates a group account comment.
	UpdateGroupAccountComment(ctx context.Context, in *MsgUpdateGroupAccountCommentRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAccountCommentResponse, error)
	// CreateProposal submits a new proposal.
	CreateProposal(ctx context.Context, in *MsgCreateProposalRequest, opts ...grpc.CallOption) (*MsgCreateProposalResponse, error)
	// Vote allows voters to vote on a proposal.
	Vote(ctx context.Context, in *MsgVoteRequest, opts ...grpc.CallOption) (*MsgVoteResponse, error)
	// Exec executes a proposal.
	Exec(ctx context.Context, in *MsgExecRequest, opts ...grpc.CallOption) (*MsgExecResponse, error)
}

type msgClient struct {
	cc                                grpc.ClientConnInterface
	_CreateGroup                      types.Invoker
	_UpdateGroupMembers               types.Invoker
	_UpdateGroupAdmin                 types.Invoker
	_UpdateGroupComment               types.Invoker
	_CreateGroupAccount               types.Invoker
	_UpdateGroupAccountAdmin          types.Invoker
	_UpdateGroupAccountDecisionPolicy types.Invoker
	_UpdateGroupAccountComment        types.Invoker
	_CreateProposal                   types.Invoker
	_Vote                             types.Invoker
	_Exec                             types.Invoker
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc: cc}
}

func (c *msgClient) CreateGroup(ctx context.Context, in *MsgCreateGroupRequest, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error) {
	if invoker := c._CreateGroup; invoker != nil {
		var out MsgCreateGroupResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._CreateGroup, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/CreateGroup")
		if err != nil {
			var out MsgCreateGroupResponse
			err = c._CreateGroup(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgCreateGroupResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupMembers(ctx context.Context, in *MsgUpdateGroupMembersRequest, opts ...grpc.CallOption) (*MsgUpdateGroupMembersResponse, error) {
	if invoker := c._UpdateGroupMembers; invoker != nil {
		var out MsgUpdateGroupMembersResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._UpdateGroupMembers, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/UpdateGroupMembers")
		if err != nil {
			var out MsgUpdateGroupMembersResponse
			err = c._UpdateGroupMembers(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgUpdateGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/UpdateGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupAdmin(ctx context.Context, in *MsgUpdateGroupAdminRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAdminResponse, error) {
	if invoker := c._UpdateGroupAdmin; invoker != nil {
		var out MsgUpdateGroupAdminResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._UpdateGroupAdmin, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/UpdateGroupAdmin")
		if err != nil {
			var out MsgUpdateGroupAdminResponse
			err = c._UpdateGroupAdmin(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgUpdateGroupAdminResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/UpdateGroupAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupComment(ctx context.Context, in *MsgUpdateGroupCommentRequest, opts ...grpc.CallOption) (*MsgUpdateGroupCommentResponse, error) {
	if invoker := c._UpdateGroupComment; invoker != nil {
		var out MsgUpdateGroupCommentResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._UpdateGroupComment, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/UpdateGroupComment")
		if err != nil {
			var out MsgUpdateGroupCommentResponse
			err = c._UpdateGroupComment(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgUpdateGroupCommentResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/UpdateGroupComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateGroupAccount(ctx context.Context, in *MsgCreateGroupAccountRequest, opts ...grpc.CallOption) (*MsgCreateGroupAccountResponse, error) {
	if invoker := c._CreateGroupAccount; invoker != nil {
		var out MsgCreateGroupAccountResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._CreateGroupAccount, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/CreateGroupAccount")
		if err != nil {
			var out MsgCreateGroupAccountResponse
			err = c._CreateGroupAccount(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgCreateGroupAccountResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/CreateGroupAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupAccountAdmin(ctx context.Context, in *MsgUpdateGroupAccountAdminRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAccountAdminResponse, error) {
	if invoker := c._UpdateGroupAccountAdmin; invoker != nil {
		var out MsgUpdateGroupAccountAdminResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._UpdateGroupAccountAdmin, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/UpdateGroupAccountAdmin")
		if err != nil {
			var out MsgUpdateGroupAccountAdminResponse
			err = c._UpdateGroupAccountAdmin(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgUpdateGroupAccountAdminResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/UpdateGroupAccountAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupAccountDecisionPolicy(ctx context.Context, in *MsgUpdateGroupAccountDecisionPolicyRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAccountDecisionPolicyResponse, error) {
	if invoker := c._UpdateGroupAccountDecisionPolicy; invoker != nil {
		var out MsgUpdateGroupAccountDecisionPolicyResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._UpdateGroupAccountDecisionPolicy, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/UpdateGroupAccountDecisionPolicy")
		if err != nil {
			var out MsgUpdateGroupAccountDecisionPolicyResponse
			err = c._UpdateGroupAccountDecisionPolicy(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgUpdateGroupAccountDecisionPolicyResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/UpdateGroupAccountDecisionPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupAccountComment(ctx context.Context, in *MsgUpdateGroupAccountCommentRequest, opts ...grpc.CallOption) (*MsgUpdateGroupAccountCommentResponse, error) {
	if invoker := c._UpdateGroupAccountComment; invoker != nil {
		var out MsgUpdateGroupAccountCommentResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._UpdateGroupAccountComment, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/UpdateGroupAccountComment")
		if err != nil {
			var out MsgUpdateGroupAccountCommentResponse
			err = c._UpdateGroupAccountComment(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgUpdateGroupAccountCommentResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/UpdateGroupAccountComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateProposal(ctx context.Context, in *MsgCreateProposalRequest, opts ...grpc.CallOption) (*MsgCreateProposalResponse, error) {
	if invoker := c._CreateProposal; invoker != nil {
		var out MsgCreateProposalResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._CreateProposal, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/CreateProposal")
		if err != nil {
			var out MsgCreateProposalResponse
			err = c._CreateProposal(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgCreateProposalResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/CreateProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVoteRequest, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	if invoker := c._Vote; invoker != nil {
		var out MsgVoteResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Vote, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/Vote")
		if err != nil {
			var out MsgVoteResponse
			err = c._Vote(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Exec(ctx context.Context, in *MsgExecRequest, opts ...grpc.CallOption) (*MsgExecResponse, error) {
	if invoker := c._Exec; invoker != nil {
		var out MsgExecResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Exec, err = invokerConn.Invoker("/regen.group.v1alpha1.Msg/Exec")
		if err != nil {
			var out MsgExecResponse
			err = c._Exec(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgExecResponse)
	err := c.cc.Invoke(ctx, "/regen.group.v1alpha1.Msg/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateGroup creates a new group with an admin account address, a list of members and an comment.
	CreateGroup(types.Context, *MsgCreateGroupRequest) (*MsgCreateGroupResponse, error)
	// UpdateGroupMembers updates the group members with given group id and admin address.
	UpdateGroupMembers(types.Context, *MsgUpdateGroupMembersRequest) (*MsgUpdateGroupMembersResponse, error)
	// UpdateGroupAdmin updates the group admin with given group id and previous admin address.
	UpdateGroupAdmin(types.Context, *MsgUpdateGroupAdminRequest) (*MsgUpdateGroupAdminResponse, error)
	// UpdateGroupComment updates the group comment with given group id and admin address.
	UpdateGroupComment(types.Context, *MsgUpdateGroupCommentRequest) (*MsgUpdateGroupCommentResponse, error)
	// CreateGroupAccount creates a new group account using given DecisionPolicy.
	CreateGroupAccount(types.Context, *MsgCreateGroupAccountRequest) (*MsgCreateGroupAccountResponse, error)
	// UpdateGroupAccountAdmin updates a group account admin.
	UpdateGroupAccountAdmin(types.Context, *MsgUpdateGroupAccountAdminRequest) (*MsgUpdateGroupAccountAdminResponse, error)
	// UpdateGroupAccountDecisionPolicy allows a group account decision policy to be updated.
	UpdateGroupAccountDecisionPolicy(types.Context, *MsgUpdateGroupAccountDecisionPolicyRequest) (*MsgUpdateGroupAccountDecisionPolicyResponse, error)
	// UpdateGroupAccountComment updates a group account comment.
	UpdateGroupAccountComment(types.Context, *MsgUpdateGroupAccountCommentRequest) (*MsgUpdateGroupAccountCommentResponse, error)
	// CreateProposal submits a new proposal.
	CreateProposal(types.Context, *MsgCreateProposalRequest) (*MsgCreateProposalResponse, error)
	// Vote allows voters to vote on a proposal.
	Vote(types.Context, *MsgVoteRequest) (*MsgVoteResponse, error)
	// Exec executes a proposal.
	Exec(types.Context, *MsgExecRequest) (*MsgExecResponse, error)
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGroup(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGroup(types.UnwrapSDKContext(ctx), req.(*MsgCreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupMembers(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/UpdateGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupMembers(types.UnwrapSDKContext(ctx), req.(*MsgUpdateGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupAdmin(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/UpdateGroupAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupAdmin(types.UnwrapSDKContext(ctx), req.(*MsgUpdateGroupAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupComment(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/UpdateGroupComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupComment(types.UnwrapSDKContext(ctx), req.(*MsgUpdateGroupCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateGroupAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGroupAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGroupAccount(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/CreateGroupAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGroupAccount(types.UnwrapSDKContext(ctx), req.(*MsgCreateGroupAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupAccountAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupAccountAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupAccountAdmin(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/UpdateGroupAccountAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupAccountAdmin(types.UnwrapSDKContext(ctx), req.(*MsgUpdateGroupAccountAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupAccountDecisionPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupAccountDecisionPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupAccountDecisionPolicy(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/UpdateGroupAccountDecisionPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupAccountDecisionPolicy(types.UnwrapSDKContext(ctx), req.(*MsgUpdateGroupAccountDecisionPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupAccountComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupAccountCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupAccountComment(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/UpdateGroupAccountComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupAccountComment(types.UnwrapSDKContext(ctx), req.(*MsgUpdateGroupAccountCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProposal(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/CreateProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProposal(types.UnwrapSDKContext(ctx), req.(*MsgCreateProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(types.UnwrapSDKContext(ctx), req.(*MsgVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Exec(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.group.v1alpha1.Msg/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Exec(types.UnwrapSDKContext(ctx), req.(*MsgExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.group.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _Msg_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroupMembers",
			Handler:    _Msg_UpdateGroupMembers_Handler,
		},
		{
			MethodName: "UpdateGroupAdmin",
			Handler:    _Msg_UpdateGroupAdmin_Handler,
		},
		{
			MethodName: "UpdateGroupComment",
			Handler:    _Msg_UpdateGroupComment_Handler,
		},
		{
			MethodName: "CreateGroupAccount",
			Handler:    _Msg_CreateGroupAccount_Handler,
		},
		{
			MethodName: "UpdateGroupAccountAdmin",
			Handler:    _Msg_UpdateGroupAccountAdmin_Handler,
		},
		{
			MethodName: "UpdateGroupAccountDecisionPolicy",
			Handler:    _Msg_UpdateGroupAccountDecisionPolicy_Handler,
		},
		{
			MethodName: "UpdateGroupAccountComment",
			Handler:    _Msg_UpdateGroupAccountComment_Handler,
		},
		{
			MethodName: "CreateProposal",
			Handler:    _Msg_CreateProposal_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Msg_Exec_Handler,
		},
	},
	Metadata: "regen/group/v1alpha1/tx.proto",
}

const (
	MsgCreateGroupMethod                      = "/regen.group.v1alpha1.Msg/CreateGroup"
	MsgUpdateGroupMembersMethod               = "/regen.group.v1alpha1.Msg/UpdateGroupMembers"
	MsgUpdateGroupAdminMethod                 = "/regen.group.v1alpha1.Msg/UpdateGroupAdmin"
	MsgUpdateGroupCommentMethod               = "/regen.group.v1alpha1.Msg/UpdateGroupComment"
	MsgCreateGroupAccountMethod               = "/regen.group.v1alpha1.Msg/CreateGroupAccount"
	MsgUpdateGroupAccountAdminMethod          = "/regen.group.v1alpha1.Msg/UpdateGroupAccountAdmin"
	MsgUpdateGroupAccountDecisionPolicyMethod = "/regen.group.v1alpha1.Msg/UpdateGroupAccountDecisionPolicy"
	MsgUpdateGroupAccountCommentMethod        = "/regen.group.v1alpha1.Msg/UpdateGroupAccountComment"
	MsgCreateProposalMethod                   = "/regen.group.v1alpha1.Msg/CreateProposal"
	MsgVoteMethod                             = "/regen.group.v1alpha1.Msg/Vote"
	MsgExecMethod                             = "/regen.group.v1alpha1.Msg/Exec"
)
