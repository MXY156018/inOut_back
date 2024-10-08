// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package auth

import (
	"context"

	"mall-admin/rpc/admin"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiAuthReq  = admin.ApiAuthReq
	ApiAuthResp = admin.ApiAuthResp

	Auth interface {
		// api 授权
		ApiAuth(ctx context.Context, in *ApiAuthReq, opts ...grpc.CallOption) (*ApiAuthResp, error)
	}

	defaultAuth struct {
		cli zrpc.Client
	}
)

func NewAuth(cli zrpc.Client) Auth {
	return &defaultAuth{
		cli: cli,
	}
}

// api 授权
func (m *defaultAuth) ApiAuth(ctx context.Context, in *ApiAuthReq, opts ...grpc.CallOption) (*ApiAuthResp, error) {
	client := admin.NewAuthClient(m.cli.Conn())
	return client.ApiAuth(ctx, in, opts...)
}
