// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package UserCenter

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserCenterClient is the client API for UserCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCenterClient interface {
	CreatUser(ctx context.Context, in *Uuser, opts ...grpc.CallOption) (*Uuser, error)
	GetUserInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Uuser, error)
	GetUserInfoByEmail(ctx context.Context, in *S, opts ...grpc.CallOption) (*Uuser, error)
	GetUserPwd(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error)
	GetUserName(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error)
	GetUserEmail(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error)
	GetUserClass(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
	GetUserType(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
	UserIs_Exist(ctx context.Context, in *S, opts ...grpc.CallOption) (*Right, error)
	RefreshingUserData(ctx context.Context, in *Uuser, opts ...grpc.CallOption) (*Uuser, error)
	CreateClass(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error)
	GetClassInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Class, error)
	GetClassTeacher(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
	GetClassName(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error)
	DissolveClass(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	RefreshingClassData(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error)
	FireStudent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Class, error)
}

type userCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCenterClient(cc grpc.ClientConnInterface) UserCenterClient {
	return &userCenterClient{cc}
}

func (c *userCenterClient) CreatUser(ctx context.Context, in *Uuser, opts ...grpc.CallOption) (*Uuser, error) {
	out := new(Uuser)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/creat_user", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Uuser, error) {
	out := new(Uuser)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_user_info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserInfoByEmail(ctx context.Context, in *S, opts ...grpc.CallOption) (*Uuser, error) {
	out := new(Uuser)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_user_info_by_email", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserPwd(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error) {
	out := new(S)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_user_pwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserName(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error) {
	out := new(S)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_user_name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserEmail(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error) {
	out := new(S)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_user_email", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserClass(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_user_class", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserType(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_user_type", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) UserIs_Exist(ctx context.Context, in *S, opts ...grpc.CallOption) (*Right, error) {
	out := new(Right)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/user_is_Exist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) RefreshingUserData(ctx context.Context, in *Uuser, opts ...grpc.CallOption) (*Uuser, error) {
	out := new(Uuser)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/refreshing_user_data", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) CreateClass(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/create_class", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetClassInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_class_info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetClassTeacher(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_class_teacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetClassName(ctx context.Context, in *Id, opts ...grpc.CallOption) (*S, error) {
	out := new(S)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/get_class_name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) DissolveClass(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/dissolve_class", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) RefreshingClassData(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/refreshing_class_data", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) FireStudent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/UserCenter.UserCenter/fire_student", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCenterServer is the server API for UserCenter service.
// All implementations must embed UnimplementedUserCenterServer
// for forward compatibility
type UserCenterServer interface {
	CreatUser(context.Context, *Uuser) (*Uuser, error)
	GetUserInfo(context.Context, *Id) (*Uuser, error)
	GetUserInfoByEmail(context.Context, *S) (*Uuser, error)
	GetUserPwd(context.Context, *Id) (*S, error)
	GetUserName(context.Context, *Id) (*S, error)
	GetUserEmail(context.Context, *Id) (*S, error)
	GetUserClass(context.Context, *Id) (*Id, error)
	GetUserType(context.Context, *Id) (*Id, error)
	UserIs_Exist(context.Context, *S) (*Right, error)
	RefreshingUserData(context.Context, *Uuser) (*Uuser, error)
	CreateClass(context.Context, *Class) (*Class, error)
	GetClassInfo(context.Context, *Id) (*Class, error)
	GetClassTeacher(context.Context, *Id) (*Id, error)
	GetClassName(context.Context, *Id) (*S, error)
	DissolveClass(context.Context, *Id) (*Empty, error)
	RefreshingClassData(context.Context, *Class) (*Class, error)
	FireStudent(context.Context, *Id) (*Class, error)
	mustEmbedUnimplementedUserCenterServer()
}

// UnimplementedUserCenterServer must be embedded to have forward compatible implementations.
type UnimplementedUserCenterServer struct {
}

func (UnimplementedUserCenterServer) CreatUser(context.Context, *Uuser) (*Uuser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatUser not implemented")
}
func (UnimplementedUserCenterServer) GetUserInfo(context.Context, *Id) (*Uuser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserCenterServer) GetUserInfoByEmail(context.Context, *S) (*Uuser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByEmail not implemented")
}
func (UnimplementedUserCenterServer) GetUserPwd(context.Context, *Id) (*S, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPwd not implemented")
}
func (UnimplementedUserCenterServer) GetUserName(context.Context, *Id) (*S, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserName not implemented")
}
func (UnimplementedUserCenterServer) GetUserEmail(context.Context, *Id) (*S, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEmail not implemented")
}
func (UnimplementedUserCenterServer) GetUserClass(context.Context, *Id) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserClass not implemented")
}
func (UnimplementedUserCenterServer) GetUserType(context.Context, *Id) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserType not implemented")
}
func (UnimplementedUserCenterServer) UserIs_Exist(context.Context, *S) (*Right, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserIs_Exist not implemented")
}
func (UnimplementedUserCenterServer) RefreshingUserData(context.Context, *Uuser) (*Uuser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshingUserData not implemented")
}
func (UnimplementedUserCenterServer) CreateClass(context.Context, *Class) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClass not implemented")
}
func (UnimplementedUserCenterServer) GetClassInfo(context.Context, *Id) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassInfo not implemented")
}
func (UnimplementedUserCenterServer) GetClassTeacher(context.Context, *Id) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassTeacher not implemented")
}
func (UnimplementedUserCenterServer) GetClassName(context.Context, *Id) (*S, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassName not implemented")
}
func (UnimplementedUserCenterServer) DissolveClass(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DissolveClass not implemented")
}
func (UnimplementedUserCenterServer) RefreshingClassData(context.Context, *Class) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshingClassData not implemented")
}
func (UnimplementedUserCenterServer) FireStudent(context.Context, *Id) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FireStudent not implemented")
}
func (UnimplementedUserCenterServer) mustEmbedUnimplementedUserCenterServer() {}

// UnsafeUserCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCenterServer will
// result in compilation errors.
type UnsafeUserCenterServer interface {
	mustEmbedUnimplementedUserCenterServer()
}

func RegisterUserCenterServer(s grpc.ServiceRegistrar, srv UserCenterServer) {
	s.RegisterService(&UserCenter_ServiceDesc, srv)
}

func _UserCenter_CreatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).CreatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/creat_user",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).CreatUser(ctx, req.(*Uuser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_user_info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserInfo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserInfoByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserInfoByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_user_info_by_email",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserInfoByEmail(ctx, req.(*S))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_user_pwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserPwd(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_user_name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserName(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_user_email",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserEmail(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_user_class",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserClass(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_user_type",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserType(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_UserIs_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).UserIs_Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/user_is_Exist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).UserIs_Exist(ctx, req.(*S))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_RefreshingUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).RefreshingUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/refreshing_user_data",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).RefreshingUserData(ctx, req.(*Uuser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_CreateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).CreateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/create_class",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).CreateClass(ctx, req.(*Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetClassInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetClassInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_class_info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetClassInfo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetClassTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetClassTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_class_teacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetClassTeacher(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetClassName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetClassName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/get_class_name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetClassName(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_DissolveClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).DissolveClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/dissolve_class",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).DissolveClass(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_RefreshingClassData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).RefreshingClassData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/refreshing_class_data",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).RefreshingClassData(ctx, req.(*Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_FireStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).FireStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCenter.UserCenter/fire_student",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).FireStudent(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCenter_ServiceDesc is the grpc.ServiceDesc for UserCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserCenter.UserCenter",
	HandlerType: (*UserCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "creat_user",
			Handler:    _UserCenter_CreatUser_Handler,
		},
		{
			MethodName: "get_user_info",
			Handler:    _UserCenter_GetUserInfo_Handler,
		},
		{
			MethodName: "get_user_info_by_email",
			Handler:    _UserCenter_GetUserInfoByEmail_Handler,
		},
		{
			MethodName: "get_user_pwd",
			Handler:    _UserCenter_GetUserPwd_Handler,
		},
		{
			MethodName: "get_user_name",
			Handler:    _UserCenter_GetUserName_Handler,
		},
		{
			MethodName: "get_user_email",
			Handler:    _UserCenter_GetUserEmail_Handler,
		},
		{
			MethodName: "get_user_class",
			Handler:    _UserCenter_GetUserClass_Handler,
		},
		{
			MethodName: "get_user_type",
			Handler:    _UserCenter_GetUserType_Handler,
		},
		{
			MethodName: "user_is_Exist",
			Handler:    _UserCenter_UserIs_Exist_Handler,
		},
		{
			MethodName: "refreshing_user_data",
			Handler:    _UserCenter_RefreshingUserData_Handler,
		},
		{
			MethodName: "create_class",
			Handler:    _UserCenter_CreateClass_Handler,
		},
		{
			MethodName: "get_class_info",
			Handler:    _UserCenter_GetClassInfo_Handler,
		},
		{
			MethodName: "get_class_teacher",
			Handler:    _UserCenter_GetClassTeacher_Handler,
		},
		{
			MethodName: "get_class_name",
			Handler:    _UserCenter_GetClassName_Handler,
		},
		{
			MethodName: "dissolve_class",
			Handler:    _UserCenter_DissolveClass_Handler,
		},
		{
			MethodName: "refreshing_class_data",
			Handler:    _UserCenter_RefreshingClassData_Handler,
		},
		{
			MethodName: "fire_student",
			Handler:    _UserCenter_FireStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UserCenter.proto",
}
