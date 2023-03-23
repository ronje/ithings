// Code generated by goctl. DO NOT EDIT.
// Source: rule.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/rulesvr/internal/svc"
	"github.com/i-Things/things/src/rulesvr/pb/rule"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AlarmDeal                = rule.AlarmDeal
	AlarmDealRecordCreateReq = rule.AlarmDealRecordCreateReq
	AlarmDealRecordIndexReq  = rule.AlarmDealRecordIndexReq
	AlarmDealRecordIndexResp = rule.AlarmDealRecordIndexResp
	AlarmInfo                = rule.AlarmInfo
	AlarmInfoDeleteReq       = rule.AlarmInfoDeleteReq
	AlarmInfoIndexReq        = rule.AlarmInfoIndexReq
	AlarmInfoIndexResp       = rule.AlarmInfoIndexResp
	AlarmLog                 = rule.AlarmLog
	AlarmLogIndexReq         = rule.AlarmLogIndexReq
	AlarmLogIndexResp        = rule.AlarmLogIndexResp
	AlarmSceneDeleteReq      = rule.AlarmSceneDeleteReq
	AlarmSceneIndexReq       = rule.AlarmSceneIndexReq
	AlarmSceneIndexResp      = rule.AlarmSceneIndexResp
	AlarmSceneMultiCreateReq = rule.AlarmSceneMultiCreateReq
	FlowInfo                 = rule.FlowInfo
	FlowInfoDeleteReq        = rule.FlowInfoDeleteReq
	FlowInfoIndexReq         = rule.FlowInfoIndexReq
	FlowInfoIndexResp        = rule.FlowInfoIndexResp
	FlowInfoReadReq          = rule.FlowInfoReadReq
	PageInfo                 = rule.PageInfo
	Response                 = rule.Response
	SceneInfo                = rule.SceneInfo
	SceneInfoDeleteReq       = rule.SceneInfoDeleteReq
	SceneInfoIndexReq        = rule.SceneInfoIndexReq
	SceneInfoIndexResp       = rule.SceneInfoIndexResp
	SceneInfoReadReq         = rule.SceneInfoReadReq
	TimeRange                = rule.TimeRange

	SceneLinkage interface {
		SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Response, error)
		SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Response, error)
		SceneInfoDelete(ctx context.Context, in *SceneInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error)
		SceneInfoRead(ctx context.Context, in *SceneInfoReadReq, opts ...grpc.CallOption) (*SceneInfo, error)
	}

	defaultSceneLinkage struct {
		cli zrpc.Client
	}

	directSceneLinkage struct {
		svcCtx *svc.ServiceContext
		svr    rule.SceneLinkageServer
	}
)

func NewSceneLinkage(cli zrpc.Client) SceneLinkage {
	return &defaultSceneLinkage{
		cli: cli,
	}
}

func NewDirectSceneLinkage(svcCtx *svc.ServiceContext, svr rule.SceneLinkageServer) SceneLinkage {
	return &directSceneLinkage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultSceneLinkage) SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Response, error) {
	client := rule.NewSceneLinkageClient(m.cli.Conn())
	return client.SceneInfoCreate(ctx, in, opts...)
}

func (d *directSceneLinkage) SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.SceneInfoCreate(ctx, in)
}

func (m *defaultSceneLinkage) SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Response, error) {
	client := rule.NewSceneLinkageClient(m.cli.Conn())
	return client.SceneInfoUpdate(ctx, in, opts...)
}

func (d *directSceneLinkage) SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.SceneInfoUpdate(ctx, in)
}

func (m *defaultSceneLinkage) SceneInfoDelete(ctx context.Context, in *SceneInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := rule.NewSceneLinkageClient(m.cli.Conn())
	return client.SceneInfoDelete(ctx, in, opts...)
}

func (d *directSceneLinkage) SceneInfoDelete(ctx context.Context, in *SceneInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.SceneInfoDelete(ctx, in)
}

func (m *defaultSceneLinkage) SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error) {
	client := rule.NewSceneLinkageClient(m.cli.Conn())
	return client.SceneInfoIndex(ctx, in, opts...)
}

func (d *directSceneLinkage) SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error) {
	return d.svr.SceneInfoIndex(ctx, in)
}

func (m *defaultSceneLinkage) SceneInfoRead(ctx context.Context, in *SceneInfoReadReq, opts ...grpc.CallOption) (*SceneInfo, error) {
	client := rule.NewSceneLinkageClient(m.cli.Conn())
	return client.SceneInfoRead(ctx, in, opts...)
}

func (d *directSceneLinkage) SceneInfoRead(ctx context.Context, in *SceneInfoReadReq, opts ...grpc.CallOption) (*SceneInfo, error) {
	return d.svr.SceneInfoRead(ctx, in)
}
