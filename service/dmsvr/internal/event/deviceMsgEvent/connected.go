package deviceMsgEvent

import (
	"context"
	"database/sql"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/domain/deviceAuth"
	"gitee.com/unitedrhino/share/domain/deviceMsg/msgThing"
	"gitee.com/unitedrhino/share/domain/schema"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceStatus"
	devicemanagelogic "gitee.com/unitedrhino/things/service/dmsvr/internal/logic/devicemanage"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ConnectedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	template *schema.Model
	topics   []string
	dreq     msgThing.Req
	DiDB     *relationDB.DeviceInfoRepo
	di       *relationDB.DmDeviceInfo
}

func NewConnectedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectedLogic {
	return &ConnectedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		DiDB:   relationDB.NewDeviceInfoRepo(ctx),
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConnectedLogic) UpdateLoginTime() {
	now := sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}
	if l.di.FirstLogin.Valid == false {
		l.di.FirstLogin = now
	}
	l.di.LastLogin = now
	l.di.IsOnline = def.True
	l.DiDB.Update(l.ctx, l.di)
}

func (l *ConnectedLogic) Handle(msg *deviceStatus.ConnectMsg) error {
	l.Infof("%s req=%+v", utils.FuncName(), utils.Fmt(msg))
	ld, err := deviceAuth.GetClientIDInfo(msg.ClientID)
	if err != nil && msg.Device.DeviceName == "" {
		return err
	}
	if ld != nil && ld.IsNeedRegister {
		return nil
	}
	//di, err := l.svcCtx.DeviceCache.GetData(l.ctx, devices.Core{
	//	ProductID:  ld.ProductID,
	//	DeviceName: ld.DeviceName,
	//})
	//if err != nil {
	//	return err
	//}
	//if di.FirstLogin == 0 {
	//	err := devicemanagelogic.HandleOnlineFix(l.ctx, l.svcCtx, msg)
	//	if err != nil {
	//		l.Error(err)
	//	}
	//	return err
	//}
	err = devicemanagelogic.HandleOnlineFix(l.ctx, l.svcCtx, msg)
	if err != nil {
		l.Error(err)
	}

	//err = l.svcCtx.DeviceStatus.AddDevice(l.ctx, msg)
	return err
	//l.di, err = l.DiDB.FindOneByFilter(l.ctx, relationDB.DeviceFilter{ProductID: ld.ProductID, DeviceNames: []string{ld.DeviceName}})
	//if err != nil {
	//	return err
	//}
	//l.UpdateLoginTime()
	//err = l.svcCtx.StatusRepo.Insert(l.ctx, &deviceLog.Msg{
	//	ProductID:  ld.ProductID,
	//	Msg:     def.ConnectedStatus,
	//	Timestamp:  msg.Timestamp, // 操作时间
	//	DeviceName: ld.DeviceName,
	//})
	//if err != nil {
	//	l.Errorf("%s.HubLogRepo.insert productID:%v deviceName:%v err:%v",
	//		utils.FuncName(), ld.ProductID, ld.DeviceName, err)
	//}
	//err = l.svcCtx.PubApp.DeviceStatusConnected(l.ctx, application.ConnectMsg{
	//	Device: devices.Core{
	//		ProductID:  ld.ProductID,
	//		DeviceName: ld.DeviceName,
	//	},
	//	Timestamp: msg.Timestamp.UnixMilli(),
	//})
	//if err != nil {
	//	l.Errorf("%s.PubApp.DeviceStatusConnected productID:%v deviceName:%v err:%v",
	//		utils.FuncName(), ld.ProductID, ld.DeviceName, err)
	//}
	//return nil
}
