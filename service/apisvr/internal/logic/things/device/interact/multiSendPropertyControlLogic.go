package interact

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/apisvr/internal/logic/things"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"golang.org/x/sync/errgroup"
	"sync"

	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiSendPropertyControlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	retMsg []*types.DeviceInteractMultiSendPropertyMsg
	err    error
	mutex  sync.Mutex
}

func NewMultiSendPropertyControlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiSendPropertyControlLogic {
	return &MultiSendPropertyControlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctxs.WithDefaultRoot(ctx),
		svcCtx: svcCtx,
	}
}

func (l *MultiSendPropertyControlLogic) MultiSendPropertyControl(req *types.DeviceInteractMultiSendPropertyReq) (resp *types.DeviceInteractMultiSendPropertyResp, err error) {
	if (req.ProductID != "" && len(req.DeviceNames) != 0) || len(req.Devices) != 0 {
		err := l.SendProperty(req.ProductID, req.DeviceNames, req.Devices, req.Data, req.ShadowControl)
		return &types.DeviceInteractMultiSendPropertyResp{List: l.retMsg}, err
	}
	if req.GroupID != 0 || req.AreaID != 0 {
		var ds []*dm.DeviceInfo
		if req.GroupID != 0 {
			dgRet, err := l.svcCtx.DeviceG.GroupDeviceIndex(l.ctx, &dm.GroupDeviceIndexReq{
				GroupID: req.GroupID,
			})
			if err != nil {
				return nil, err
			}
			ds = dgRet.List
		}
		if req.AreaID != 0 {
			ret, err := l.svcCtx.DeviceM.DeviceInfoIndex(l.ctx, &dm.DeviceInfoIndexReq{
				AreaIDs: []int64{req.AreaID},
			})
			if err != nil {
				return nil, err
			}
			ds = ret.List
		}
		var devices = map[string][]string{} //key 是产品id value是设备名列表
		for _, v := range ds {
			if p := devices[v.ProductID]; p != nil {
				devices[v.ProductID] = append(p, v.DeviceName)
				continue
			}
			devices[v.ProductID] = []string{v.DeviceName}
		}
		for p, d := range devices {
			var eg errgroup.Group
			productID := p
			deviceNames := d
			eg.Go(func() error {
				err := l.SendProperty(productID, deviceNames, nil, req.Data, req.ShadowControl)
				if err != nil {
					return err
				}
				return nil
			})
			err := eg.Wait()
			if err != nil {
				return nil, err
			}
		}
		return &types.DeviceInteractMultiSendPropertyResp{List: l.retMsg}, nil
	}
	return nil, errors.Parameter.AddMsg("产品id设备名或分组id或区域id必须填一个")
}
func (l *MultiSendPropertyControlLogic) SendProperty(productID string, deviceNames []string, devices []*types.DeviceCore, data string, shadowControl int64) error {
	list := make([]*types.DeviceInteractMultiSendPropertyMsg, 0)
	dmReq := &dm.PropertyControlMultiSendReq{
		ProductID:     productID,
		DeviceNames:   deviceNames,
		Devices:       things.ToDmDeviceCoresPb(devices),
		Data:          data,
		ShadowControl: shadowControl,
	}
	dmResp, err := l.svcCtx.DeviceInteract.PropertyControlMultiSend(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.MultiSendProperty productID=%v deviceNames=%v data=%v err=%+v", utils.FuncName(), productID, deviceNames, data, er)
		return er
	}
	if len(dmResp.List) > 0 {
		for _, v := range dmResp.List {
			list = append(list, &types.DeviceInteractMultiSendPropertyMsg{
				ProductID:  productID,
				DeviceName: v.DeviceName,
				Code:       v.Code,
				Msg:        v.Msg,
				MsgToken:   v.MsgToken,
				SysMsg:     v.SysMsg,
				SysCode:    v.SysCode,
			})
		}
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.retMsg = append(l.retMsg, list...)
	return nil
}
