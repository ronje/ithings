// Code generated by goctl. DO NOT EDIT.
// Source: vid.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/vidsvr/internal/logic/vidmgrconfigmanage"
	"github.com/i-Things/things/src/vidsvr/internal/svc"
	"github.com/i-Things/things/src/vidsvr/pb/vid"
)

type VidmgrConfigManageServer struct {
	svcCtx *svc.ServiceContext
	vid.UnimplementedVidmgrConfigManageServer
}

func NewVidmgrConfigManageServer(svcCtx *svc.ServiceContext) *VidmgrConfigManageServer {
	return &VidmgrConfigManageServer{
		svcCtx: svcCtx,
	}
}

// 新建配置
func (s *VidmgrConfigManageServer) VidmgrConfigCreate(ctx context.Context, in *vid.VidmgrConfig) (*vid.Response, error) {
	l := vidmgrconfigmanagelogic.NewVidmgrConfigCreateLogic(ctx, s.svcCtx)
	return l.VidmgrConfigCreate(in)
}

// 删除配置
func (s *VidmgrConfigManageServer) VidmgrConfigDelete(ctx context.Context, in *vid.VidmgrConfigDeleteReq) (*vid.Response, error) {
	l := vidmgrconfigmanagelogic.NewVidmgrConfigDeleteLogic(ctx, s.svcCtx)
	return l.VidmgrConfigDelete(in)
}

// 更新配置
func (s *VidmgrConfigManageServer) VidmgrConfigUpdate(ctx context.Context, in *vid.VidmgrConfig) (*vid.Response, error) {
	l := vidmgrconfigmanagelogic.NewVidmgrConfigUpdateLogic(ctx, s.svcCtx)
	return l.VidmgrConfigUpdate(in)
}

// 配置列表
func (s *VidmgrConfigManageServer) VidmgrConfigIndex(ctx context.Context, in *vid.VidmgrConfigIndexReq) (*vid.VidmgrConfigIndexResp, error) {
	l := vidmgrconfigmanagelogic.NewVidmgrConfigIndexLogic(ctx, s.svcCtx)
	return l.VidmgrConfigIndex(in)
}

// 获取配置信息详情
func (s *VidmgrConfigManageServer) VidmgrConfigRead(ctx context.Context, in *vid.VidmgrConfigReadReq) (*vid.VidmgrConfig, error) {
	l := vidmgrconfigmanagelogic.NewVidmgrConfigReadLogic(ctx, s.svcCtx)
	return l.VidmgrConfigRead(in)
}
