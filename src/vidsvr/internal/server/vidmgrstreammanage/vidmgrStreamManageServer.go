// Code generated by goctl. DO NOT EDIT.
// Source: vid.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/vidsvr/internal/logic/vidmgrstreammanage"
	"github.com/i-Things/things/src/vidsvr/internal/svc"
	"github.com/i-Things/things/src/vidsvr/pb/vid"
)

type VidmgrStreamManageServer struct {
	svcCtx *svc.ServiceContext
	vid.UnimplementedVidmgrStreamManageServer
}

func NewVidmgrStreamManageServer(svcCtx *svc.ServiceContext) *VidmgrStreamManageServer {
	return &VidmgrStreamManageServer{
		svcCtx: svcCtx,
	}
}

// 流添加
func (s *VidmgrStreamManageServer) VidmgrStreamCreate(ctx context.Context, in *vid.VidmgrStream) (*vid.Response, error) {
	l := vidmgrstreammanagelogic.NewVidmgrStreamCreateLogic(ctx, s.svcCtx)
	return l.VidmgrStreamCreate(in)
}

// 流更新
func (s *VidmgrStreamManageServer) VidmgrStreamUpdate(ctx context.Context, in *vid.VidmgrStream) (*vid.Response, error) {
	l := vidmgrstreammanagelogic.NewVidmgrStreamUpdateLogic(ctx, s.svcCtx)
	return l.VidmgrStreamUpdate(in)
}

// 删除流
func (s *VidmgrStreamManageServer) VidmgrStreamDelete(ctx context.Context, in *vid.VidmgrStreamDeleteReq) (*vid.Response, error) {
	l := vidmgrstreammanagelogic.NewVidmgrStreamDeleteLogic(ctx, s.svcCtx)
	return l.VidmgrStreamDelete(in)
}

// 获取流列表
func (s *VidmgrStreamManageServer) VidmgrStreamIndex(ctx context.Context, in *vid.VidmgrStreamIndexReq) (*vid.VidmgrStreamIndexResp, error) {
	l := vidmgrstreammanagelogic.NewVidmgrStreamIndexLogic(ctx, s.svcCtx)
	return l.VidmgrStreamIndex(in)
}

// 获取流信息详情
func (s *VidmgrStreamManageServer) VidmgrStreamRead(ctx context.Context, in *vid.VidmgrStreamReadReq) (*vid.VidmgrStream, error) {
	l := vidmgrstreammanagelogic.NewVidmgrStreamReadLogic(ctx, s.svcCtx)
	return l.VidmgrStreamRead(in)
}

// 统计流 在线，离线，未激活
func (s *VidmgrStreamManageServer) VidmgrStreamCount(ctx context.Context, in *vid.VidmgrStreamCountReq) (*vid.VidmgrStreamCountResp, error) {
	l := vidmgrstreammanagelogic.NewVidmgrStreamCountLogic(ctx, s.svcCtx)
	return l.VidmgrStreamCount(in)
}