// Code generated by goctl. DO NOT EDIT!
// Source: add.proto

package server

import (
	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/logic"
	"bookstore/rpc/add/internal/svc"
	"context"
)

type AdderServer struct {
	svcCtx *svc.ServiceContext
}

func NewAdderServer(svcCtx *svc.ServiceContext) *AdderServer {
	return &AdderServer{
		svcCtx: svcCtx,
	}
}

func (s *AdderServer) Add(ctx context.Context, in *add.AddReq) (*add.AddResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}
