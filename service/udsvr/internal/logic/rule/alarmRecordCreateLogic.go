package rulelogic

import (
	"context"
	"fmt"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"gitee.com/unitedrhino/things/service/udsvr/internal/domain/scene"
	"gitee.com/unitedrhino/things/service/udsvr/internal/repo/relationDB"
	"time"

	"gitee.com/unitedrhino/things/service/udsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/udsvr/pb/ud"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmRecordCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmRecordCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmRecordCreateLogic {
	return &AlarmRecordCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlarmRecordCreateLogic) AlarmRecordCreate(in *ud.AlarmRecordCreateReq) (*ud.Empty, error) {
	pos, err := relationDB.NewAlarmSceneRepo(l.ctx).FindByFilter(l.ctx, relationDB.AlarmSceneFilter{
		SceneID: in.SceneID, WithAlarmInfo: true, WithSceneInfo: true}, nil)
	if err != nil {
		return nil, err
	}
	if len(pos) == 0 {
		return nil, err
	}
	for _, alarm := range pos {
		if alarm.AlarmInfo.Status == def.False {
			continue
		}
		switch in.Mode {
		case scene.ActionAlarmModeRelieve:
			err = relationDB.NewAlarmRecordRepo(l.ctx).UpdateWithField(l.ctx,
				relationDB.AlarmRecordFilter{AlarmID: alarm.AlarmID, DealStatus: scene.AlarmDealStatusWaring},
				map[string]any{
					"deal_status": scene.AlarmDealStatusIgnore,
					"desc":        fmt.Sprintf("场景:%v(%v)解除告警", in.SceneName, in.SceneID),
				})
			if err != nil {
				return nil, err
			}
		case scene.ActionAlarmModeTrigger:
			po, err := relationDB.NewAlarmRecordRepo(l.ctx).FindOneByFilter(l.ctx, relationDB.AlarmRecordFilter{
				AlarmID:      alarm.AlarmID,
				DealStatuses: []scene.AlarmDealStatus{scene.AlarmDealStatusWaring}, //还处在报警中,新增次数即可
			})
			//err = errors.NotFind //先不开重复
			if err != nil {
				if !errors.Cmp(err, errors.NotFind) { //直接创建并且进行通知
					l.Errorf("NewAlarmRecordFind alarm:%v err:%v", alarm, err)
					return nil, err
				}
				po := relationDB.UdAlarmRecord{
					TenantCode:  alarm.TenantCode,
					ProjectID:   alarm.ProjectID,
					AlarmID:     alarm.AlarmID,
					AlarmName:   alarm.AlarmInfo.Name,
					TriggerType: in.TriggerType,
					ProductID:   in.ProductID,
					DeviceName:  in.DeviceName,
					Level:       alarm.AlarmInfo.Level,
					SceneName:   alarm.SceneInfo.Name,
					SceneID:     alarm.SceneID,
					DealStatus:  scene.AlarmDealStatusWaring,
					Desc:        fmt.Sprintf("自动化触发告警:%v", in.SceneName),
					AlarmCount:  1,
					LastAlarm:   time.Now(),
				}
				err = relationDB.NewAlarmRecordRepo(l.ctx).Insert(l.ctx, &po)
				if err != nil {
					return nil, err
				}
				for _, notify := range alarm.AlarmInfo.Notifies {
					_, err := l.svcCtx.NotifyM.NotifyConfigSend(l.ctx, &sys.NotifyConfigSendReq{
						UserIDs:    alarm.AlarmInfo.UserIDs,
						Accounts:   alarm.AlarmInfo.Accounts,
						NotifyCode: def.NotifyCodeDeviceAlarm,
						TemplateID: notify.TemplateID,
						Type:       notify.Type,
						Params: map[string]string{
							"productID":   in.ProductID,
							"deviceName":  in.DeviceName,
							"sceneName":   in.SceneName,
							"deviceAlias": in.DeviceAlias,
						},
					})
					if err != nil {
						l.Error(err)
						continue
					}
				}
				continue
			}
			po.LastAlarm = time.Now()
			po.AlarmCount++
			err = relationDB.NewAlarmRecordRepo(l.ctx).Update(l.ctx, po)
			if err != nil {
				l.Errorf("NewAlarmRecordUpdate alarm:%v err:%v", alarm, err)
				return nil, err
			}
			continue
		}
	}
	if in.DeviceName != "" && in.ProductID != "" {
		di, err := l.svcCtx.DeviceCache.GetData(l.ctx, devices.Core{ProductID: in.ProductID, DeviceName: in.DeviceName})
		if err != nil {
			l.Error(err)
			return &ud.Empty{}, nil
		}
		total, err := relationDB.NewAlarmRecordRepo(l.ctx).CountByFilter(l.ctx, relationDB.AlarmRecordFilter{
			ProductID:    in.ProductID,
			DeviceName:   in.DeviceName,
			DealStatuses: []int64{scene.AlarmDealStatusWaring, scene.AlarmDealStatusInHand},
		})
		if err != nil {
			l.Error(err)
			return &ud.Empty{}, nil
		}
		if total > 0 && utils.SliceIn(di.Status, def.DeviceStatusOnline, def.DeviceStatusOffline) { //告警中的状态
			_, err := l.svcCtx.DeviceM.DeviceInfoUpdate(l.ctx, &dm.DeviceInfo{ProductID: in.ProductID, DeviceName: in.DeviceName, Status: def.DeviceStatusWarming})
			if err != nil {
				l.Error(err)
				return &ud.Empty{}, nil
			}
		}
		if total == 0 && di.Status == def.DeviceStatusWarming {
			_, err := l.svcCtx.DeviceM.DeviceInfoUpdate(l.ctx, &dm.DeviceInfo{ProductID: in.ProductID, DeviceName: in.DeviceName, Status: di.IsOnline + 1})
			if err != nil {
				l.Error(err)
				return &ud.Empty{}, nil
			}
		}
	}
	return &ud.Empty{}, nil
}
