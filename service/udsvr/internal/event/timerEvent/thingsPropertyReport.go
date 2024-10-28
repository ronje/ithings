package timerEvent

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/domain/application"
	"gitee.com/unitedrhino/share/stores"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/udsvr/internal/domain/scene"
	rulelogic "gitee.com/unitedrhino/things/service/udsvr/internal/logic/rule"
	"gitee.com/unitedrhino/things/service/udsvr/internal/repo/relationDB"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func (l *TimerHandle) SceneThingPropertyReport(in application.PropertyReport) error {
	now := time.Now()
	db := stores.WithNoDebug(l.ctx, relationDB.NewSceneIfTriggerRepo)
	di, err := l.svcCtx.DeviceCache.GetData(l.ctx, in.Device)
	if err != nil {
		return err
	}
	var triggerF = []relationDB.SceneIfTriggerFilter{
		{
			Status:    def.True,
			Type:      scene.TriggerTypeDevice,
			ProjectID: stores.CmpIn(def.RootNode, di.ProjectID),
			AreaID:    stores.CmpIn(def.RootNode, di.AreaID),
			Device:    &in.Device,
			DataID:    in.Identifier,
		},
		{
			Status: def.True,
			Type:   scene.TriggerTypeDevice,
			Device: &in.Device,
			DataID: in.Identifier,
		},
	}
	list, err := db.FindByFilters(l.ctx, triggerF, nil)
	if err != nil {
		l.Error(err)
		return err
	}
	var sceneIDSet = map[int64]struct{}{}
	for _, v := range list {
		var po = v
		if _, ok := sceneIDSet[po.SceneID]; ok {
			continue
		}
		sceneIDSet[po.SceneID] = struct{}{}
		if po.SceneInfo == nil {
			logx.WithContext(l.ctx).Errorf("trigger device not bind scene, trigger:%v", utils.Fmt(po))
			relationDB.NewSceneIfTriggerRepo(l.ctx).Delete(l.ctx, po.ID)
			continue
		}
		po.SceneInfo.Triggers = append(po.SceneInfo.Triggers, po)
		do := rulelogic.PoToSceneInfoDo(po.SceneInfo)
		do.DeviceName = di.DeviceName
		do.DeviceAlias = di.DeviceAlias.GetValue()
		do.ProductID = di.ProductID
		ps, err := l.svcCtx.SchemaCache.GetData(l.ctx, in.Device)
		if err != nil {
			l.Error(err)
			continue
		}
		if !do.If.Triggers[0].Device.IsHit(ps, in.Identifier, in.Param) {
			if po.Device.FirstTriggerTime.Valid { //如果处于触发状态,但是现在不触发了,则需要解除触发
				err := db.UpdateWithField(l.ctx, relationDB.SceneIfTriggerFilter{ID: v.ID}, map[string]any{
					"device_first_trigger_time": nil,
					"last_run_time":             nil,
				})
				if err != nil {
					l.Error(err)
				}
			}
			continue
		}
		if po.Device.FirstTriggerTime.Valid { //如果已经触发过,则忽略(默认边缘触发)
			continue
		}
		if v.Device.StateKeep.Type == scene.StateKeepTypeDuration {
			err := db.UpdateWithField(l.ctx, relationDB.SceneIfTriggerFilter{ID: v.ID}, map[string]any{
				"device_first_trigger_time": now,
				"last_run_time":             nil,
			})
			if err != nil {
				l.Error(err)
			}
			continue
		}
		func() {
			err := db.UpdateWithField(l.ctx, relationDB.SceneIfTriggerFilter{ID: v.ID}, map[string]any{
				"device_first_trigger_time": now,
				"last_run_time":             now,
			})
			if err != nil {
				l.Error(err)
			}
		}()
		ctx := ctxs.BindTenantCode(l.ctx, string(v.SceneInfo.TenantCode), int64(v.SceneInfo.ProjectID))

		if !do.When.IsHit(ctx, now, rulelogic.NewSceneCheckRepo(l.ctx, l.svcCtx, do)) {
			continue
		}
		ctxs.GoNewCtx(ctx, func(ctx context.Context) { //执行任务
			var err error
			if err != nil { //如果失败了下次还可以执行
				logx.WithContext(ctx).Error(err)
				return
			}
			l.SceneExec(ctx, do)
		})
	}
	return nil
}
