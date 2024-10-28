package deviceMsgEvent

import (
	"context"
	"database/sql"
	"encoding/json"
	"gitee.com/unitedrhino/core/service/syssvr/sysExport"
	"gitee.com/unitedrhino/core/service/timed/timedjobsvr/client/timedmanage"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/domain/application"
	"gitee.com/unitedrhino/share/domain/deviceAuth"
	"gitee.com/unitedrhino/share/domain/deviceMsg"
	"gitee.com/unitedrhino/share/domain/deviceMsg/msgOta"
	"gitee.com/unitedrhino/share/domain/deviceMsg/msgThing"
	"gitee.com/unitedrhino/share/domain/schema"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceLog"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceStatus"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/shadow"
	devicemanagelogic "gitee.com/unitedrhino/things/service/dmsvr/internal/logic/devicemanage"
	otamanagelogic "gitee.com/unitedrhino/things/service/dmsvr/internal/logic/otamanage"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/cache"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ThingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	schema *schema.Model
	dreq   msgThing.Req
	repo   msgThing.SchemaDataRepo
}

func NewThingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThingLogic {
	return &ThingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThingLogic) initMsg(msg *deviceMsg.PublishMsg) error {
	var err error
	l.schema, err = l.svcCtx.SchemaRepo.GetData(l.ctx, devices.Core{ProductID: msg.ProductID, DeviceName: msg.DeviceName})
	if err != nil {
		return errors.Database.AddDetail(err)
	}
	err = utils.Unmarshal(msg.Payload, &l.dreq)
	if err != nil {
		return errors.Parameter.AddDetailf("payload unmarshal payload:%v err:%v", string(msg.Payload), err)
	}
	l.repo = l.svcCtx.SchemaManaRepo
	return nil
}

func (l *ThingLogic) DeviceResp(msg *deviceMsg.PublishMsg, err error, data any) *deviceMsg.PublishMsg {
	if !errors.Cmp(err, errors.OK) {
		l.Errorf("%s.DeviceResp err:%v, msg:%v", utils.FuncName(), err, msg)
	}
	resp := &deviceMsg.CommonMsg{
		Method:   deviceMsg.GetRespMethod(l.dreq.Method),
		MsgToken: l.dreq.MsgToken,
		//Timestamp: time.Now().UnixMilli(),
		Data: data,
	}
	if msg.ProtocolCode == "" {
		msg.ProtocolCode = def.ProtocolCodeUnitedRhino
	}
	return &deviceMsg.PublishMsg{
		Handle:       msg.Handle,
		Type:         msg.Type,
		Payload:      resp.AddStatus(err).Bytes(),
		Timestamp:    time.Now().UnixMilli(),
		ProductID:    msg.ProductID,
		DeviceName:   msg.DeviceName,
		ProtocolCode: msg.ProtocolCode,
	}
}

// 设备属性上报
func (l *ThingLogic) HandlePackReport(msg *deviceMsg.PublishMsg, req msgThing.Req) (respMsg *deviceMsg.PublishMsg, err error) {
	pTp, err := msgThing.VerifyProperties(l.schema, req.Properties)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}
	eTp, err := msgThing.VerifyEvents(l.schema, req.Events)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}
	err = l.InsertPackReport(msg, l.schema, devices.Core{
		ProductID:  msg.ProductID,
		DeviceName: msg.DeviceName,
	}, pTp, eTp)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}
	if len(req.SubDevices) != 0 {
		for _, dev := range req.SubDevices {
			d := devices.Core{
				ProductID:  dev.ProductID,
				DeviceName: dev.DeviceName,
			}
			c, err := relationDB.NewGatewayDeviceRepo(l.ctx).CountByFilter(l.ctx, relationDB.GatewayDeviceFilter{
				SubDevice: &d,
				Gateway: &devices.Core{
					ProductID:  msg.ProductID,
					DeviceName: msg.DeviceName,
				},
			})
			if err != nil { //未绑定设备
				return l.DeviceResp(msg, err, nil), err
			}
			if c == 0 {
				err = errors.DeviceNotBound
				return l.DeviceResp(msg, err, nil), err
			}
			di, err := l.svcCtx.DeviceCache.GetData(l.ctx, d)
			if err != nil {
				return l.DeviceResp(msg, err, nil), err
			}
			if di.IsOnline == def.False {
				err := devicemanagelogic.HandleOnlineFix(l.ctx, l.svcCtx, &deviceStatus.ConnectMsg{
					ClientID:  deviceAuth.GenClientID(di.ProductID, di.DeviceName),
					Timestamp: l.dreq.GetTimeStamp(msg.Timestamp),
					Action:    devices.ActionConnected,
					Reason:    "gateway_report_fix",
				})
				if err != nil {
					l.Error(err)
				}
			}
			schema, err := l.svcCtx.SchemaRepo.GetData(l.ctx, devices.Core{ProductID: dev.ProductID, DeviceName: dev.DeviceName})
			if err != nil {
				return l.DeviceResp(msg, err, nil), err
			}
			pTp, err := msgThing.VerifyProperties(schema, dev.Properties)
			if err != nil {
				return l.DeviceResp(msg, err, nil), err
			}
			eTp, err := msgThing.VerifyEvents(schema, dev.Events)
			if err != nil {
				return l.DeviceResp(msg, err, nil), err
			}
			err = l.InsertPackReport(msg, schema, devices.Core{
				ProductID:  dev.ProductID,
				DeviceName: dev.DeviceName,
			}, pTp, eTp)
			if err != nil {
				return l.DeviceResp(msg, err, nil), err
			}
		}
	}
	return l.DeviceResp(msg, errors.OK, nil), nil
}

func (l *ThingLogic) InsertPackReport(msg *deviceMsg.PublishMsg, t *schema.Model, device devices.Core, properties []*msgThing.TimeParam, events []*msgThing.TimeParam) (err error) {
	for _, tp := range properties {
		timeStamp := time.UnixMilli(tp.Timestamp)
		if timeStamp.IsZero() {
			timeStamp = l.dreq.GetTimeStamp(msg.Timestamp)
		}

		paramValues, err := msgThing.ToParamValues(tp.Params)
		if err != nil {
			return err
		}
		ctx := ctxs.CopyCtx(l.ctx)
		utils.Go(ctx, func() {
			startTime := time.Now()
			for identifier, param := range paramValues {
				appMsg := application.PropertyReport{
					Device: device, Timestamp: timeStamp.UnixMilli(),
					Identifier: identifier, Param: param,
				}
				//应用事件通知-设备物模型属性上报通知 ↓↓↓
				err := l.svcCtx.PubApp.DeviceThingPropertyReport(ctx, appMsg)
				if err != nil {
					logx.WithContext(ctx).Errorf("%s.DeviceThingPropertyReport  identifier:%v, param:%v,err:%v", utils.FuncName(), identifier, param, err)
				}
				err = l.svcCtx.WebHook.Publish(l.svcCtx.WithDeviceTenant(ctx, device), sysExport.CodeDmDevicePropertyReport, appMsg)
				if err != nil {
					l.Error(err)
				}
				di, err := l.svcCtx.DeviceCache.GetData(l.ctx, device)
				if err != nil {
					l.Error(err)
				}
				if di != nil {
					err = l.svcCtx.UserSubscribe.Publish(ctx, def.UserSubscribeDevicePropertyReport, appMsg, map[string]any{
						"productID":  device.ProductID,
						"deviceName": device.DeviceName,
						"identifier": identifier,
					}, map[string]any{
						"productID":  device.ProductID,
						"deviceName": device.DeviceName,
					}, map[string]any{
						"projectID": di.ProjectID,
					}, map[string]any{
						"projectID": cast.ToString(di.ProjectID),
						"areaID":    cast.ToString(di.AreaID),
					})
				}
				if err != nil {
					logx.WithContext(ctx).Error(err)
				}
			}
			logx.WithContext(ctx).WithDuration(time.Now().Sub(startTime)).Infof("%s.DeviceThingPropertyReport startTime:%v",
				utils.FuncName(), startTime)
		})

		//插入多条设备物模型属性数据
		err = l.repo.InsertPropertiesData(l.ctx, t, device.ProductID, device.DeviceName, tp.Params, timeStamp, msgThing.Optional{})
		if err != nil {
			l.Errorf("%s.InsertPropertyData err=%+v", utils.FuncName(), err)
			return err
		}
	}
	for _, tp := range events {
		dbData := msgThing.EventData{}
		dbData.Identifier = tp.EventID
		dbData.Type = tp.Type
		dbData.Params, err = msgThing.ToVal(tp.Params)
		if err != nil {
			return err
		}
		dbData.TimeStamp = time.UnixMilli(tp.Timestamp)
		if dbData.TimeStamp.IsZero() {
			dbData.TimeStamp = l.dreq.GetTimeStamp(msg.Timestamp)
		}
		paramValues, err := msgThing.ToParamValues(tp.Params)
		if err != nil {
			return err
		}
		err = l.svcCtx.PubApp.DeviceThingEventReport(l.ctx, application.EventReport{
			Device:     devices.Core{ProductID: device.ProductID, DeviceName: device.DeviceName},
			Timestamp:  dbData.TimeStamp.UnixMilli(),
			Identifier: dbData.Identifier,
			Params:     paramValues,
			Type:       dbData.Type,
		})
		if err != nil {
			l.Errorf("%s.DeviceThingEventReport  err:%v", utils.FuncName(), err)
		}

		err = l.repo.InsertEventData(l.ctx, device.ProductID, device.DeviceName, &dbData)
		if err != nil {
			l.Errorf("%s.InsertEventData err=%+v", utils.FuncName(), err)
			return err
		}
	}
	return nil
}

// 设备属性上报
func (l *ThingLogic) HandlePropertyReport(msg *deviceMsg.PublishMsg, req msgThing.Req) (respMsg *deviceMsg.PublishMsg, err error) {
	tp, err := req.VerifyReqParam(l.schema, schema.ParamProperty)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	} else if len(tp) == 0 {
		err := errors.Parameter.AddMsgf("查不到物模型:%v", req.Params)
		return l.DeviceResp(msg, err, nil), err
	}

	timeStamp := req.GetTimeStamp(msg.Timestamp)
	device := devices.Core{
		ProductID:  msg.ProductID,
		DeviceName: msg.DeviceName,
	}

	paramValues, err := msgThing.ToParamValues(tp)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}
	ctx := ctxs.CopyCtx(l.ctx)
	utils.Go(ctx, func() {
		startTime := time.Now()
		for identifier, param := range paramValues {
			appMsg := application.PropertyReport{
				Device: device, Timestamp: timeStamp.UnixMilli(),
				Identifier: identifier, Param: param,
			}
			//应用事件通知-设备物模型属性上报通知 ↓↓↓
			err := l.svcCtx.PubApp.DeviceThingPropertyReport(ctx, appMsg)
			if err != nil {
				logx.WithContext(ctx).Errorf("%s.DeviceThingPropertyReport  identifier:%v, param:%v,err:%v", utils.FuncName(), identifier, param, err)
			}
			err = l.svcCtx.WebHook.Publish(l.svcCtx.WithDeviceTenant(l.ctx, device), sysExport.CodeDmDevicePropertyReport, appMsg)
			if err != nil {
				l.Error(err)
			}
			di, err := l.svcCtx.DeviceCache.GetData(l.ctx, device)
			if err != nil {
				l.Error(err)
			}
			if di != nil {
				err = l.svcCtx.UserSubscribe.Publish(ctx, def.UserSubscribeDevicePropertyReport, appMsg, map[string]any{
					"productID":  device.ProductID,
					"deviceName": device.DeviceName,
					"identifier": identifier,
				}, map[string]any{
					"productID":  device.ProductID,
					"deviceName": device.DeviceName,
				}, map[string]any{
					"projectID": di.ProjectID,
				}, map[string]any{
					"projectID": cast.ToString(di.ProjectID),
					"areaID":    cast.ToString(di.AreaID),
				})
			}
		}
		logx.WithContext(ctx).WithDuration(time.Now().Sub(startTime)).Infof("%s.DeviceThingPropertyReport startTime:%v",
			utils.FuncName(), startTime)
	})

	//插入多条设备物模型属性数据
	err = l.repo.InsertPropertiesData(l.ctx, l.schema, msg.ProductID, msg.DeviceName, tp, timeStamp, msgThing.Optional{})
	if err != nil {
		l.Errorf("%s.InsertPropertyData err=%+v", utils.FuncName(), err)
		return l.DeviceResp(msg, errors.Database.AddDetail(err), nil), err
	}

	return l.DeviceResp(msg, errors.OK, nil), nil
}

// 设备基础信息上报
func (l *ThingLogic) HandlePropertyReportInfo(msg *deviceMsg.PublishMsg, req msgThing.Req) (respMsg *deviceMsg.PublishMsg, err error) {
	dev := devices.Core{ProductID: msg.ProductID, DeviceName: msg.DeviceName}
	diDeviceBasicInfoDo := &msgThing.DeviceBasicInfo{Core: dev}
	if err = mapstructure.Decode(req.Params, diDeviceBasicInfoDo); err != nil {
		pos, ok := req.Params["position"].(map[string]interface{})
		if !ok {
			return nil, err
		}
		pos["latitude"] = cast.ToFloat64(pos["latitude"])
		pos["longitude"] = cast.ToFloat64(pos["longitude"])
		req.Params["position"] = pos
		err = mapstructure.Decode(req.Params, diDeviceBasicInfoDo)
		if err != nil {
			return nil, err
		}
	}

	dmDeviceInfoReq := ToDmDevicesInfoReq(diDeviceBasicInfoDo)
	if dmDeviceInfoReq.Version != nil {
		ctxs.GoNewCtx(l.ctx, func(ctx context.Context) {
			OtaVersionCheck(l.ctx, l.svcCtx, diDeviceBasicInfoDo.Core, dmDeviceInfoReq.Version.GetValue(), "default")
		})
	}
	_, err = devicemanagelogic.NewDeviceInfoUpdateLogic(l.ctx, l.svcCtx).DeviceInfoUpdate(dmDeviceInfoReq)
	if err != nil {
		l.Errorf("%s.DeviceInfoUpdate productID:%v deviceName:%v err:%v",
			utils.FuncName(), dmDeviceInfoReq.ProductID, dmDeviceInfoReq.DeviceName, err)
		return l.DeviceResp(msg, errors.Database.AddDetail(err), nil), err
	}
	di, err := l.svcCtx.DeviceCache.GetData(l.ctx, dev)
	if err == nil && di.IsOnline == def.False {
		err := devicemanagelogic.HandleOnlineFix(l.ctx, l.svcCtx, &deviceStatus.ConnectMsg{
			ClientID:  deviceAuth.GenClientID(di.ProductID, di.DeviceName),
			Timestamp: l.dreq.GetTimeStamp(msg.Timestamp),
			Action:    devices.ActionConnected,
			Reason:    "baseReportFix",
		})
		if err != nil {
			l.Error(err)
		}
	}
	return l.DeviceResp(msg, errors.OK, nil), nil
}

func OtaVersionCheck(ctx context.Context, svcCtx *svc.ServiceContext, msg devices.Core, version string, module string) {
	log := logx.WithContext(ctx)
	df, err := relationDB.NewOtaFirmwareDeviceRepo(ctx).FindOneByFilter(ctx, relationDB.OtaFirmwareDeviceFilter{
		ProductID:    msg.ProductID,
		DeviceNames:  []string{msg.DeviceName},
		WithFirmware: true,
		WithJob:      true,
		DestVersion:  version,
		Statues:      []int64{msgOta.DeviceStatusQueued},
	})
	if err != nil && !errors.Cmp(err, errors.NotFind) {
		log.Error(err)
		return
	}

	if df == nil || df.Job.IsNeedPush == def.False {
		return
	}
	data, err := otamanagelogic.GenUpgradeParams(ctx, svcCtx, df.Firmware, df.Files)
	if err != nil {
		log.Error(err)
		return
	}
	MsgToken := devices.GenMsgToken(ctx, svcCtx.NodeID)
	upgradeMsg := deviceMsg.CommonMsg{
		MsgToken:  MsgToken,
		Method:    msgOta.TypeUpgrade,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
	payload, _ := json.Marshal(upgradeMsg)
	pi, err := svcCtx.ProductCache.GetData(ctx, df.Firmware.ProductID)
	if err != nil {
		log.Error(err)
		return
	}
	reqMsg := deviceMsg.PublishMsg{
		Handle:       devices.Ota,
		Type:         msgOta.TypeUpgrade,
		Payload:      payload,
		Timestamp:    time.Now().UnixMilli(),
		ProductID:    msg.ProductID,
		DeviceName:   msg.DeviceName,
		ProtocolCode: pi.ProtocolCode,
	}
	err = svcCtx.PubDev.PublishToDev(ctx, &reqMsg)
	if err != nil {
		log.Error(err)
		return
	}
	df.Status = msgOta.DeviceStatusNotified
	df.Detail = "设备上报推送升级包"
	df.PushTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = relationDB.NewOtaFirmwareDeviceRepo(ctx).Update(ctx, df)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (l *ThingLogic) HandlePropertyGetSchema(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	return l.DeviceResp(msg, errors.OK, l.schema.ToSimple()), nil
}

// 设备请求获取 云端记录的最新设备信息
func (l *ThingLogic) HandlePropertyGetStatus(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	respData := make(map[string]any, len(l.schema.Property))
	dataIDs := l.dreq.Identifiers
	{ //设备影子处理
		sr := relationDB.NewShadowRepo(l.ctx)
		shadows, err := sr.FindByFilter(l.ctx, shadow.Filter{
			ProductID:           msg.ProductID,
			DeviceName:          msg.DeviceName,
			UpdatedDeviceStatus: shadow.NotUpdateDevice, //只获取未下发过的
			DataIDs:             dataIDs,
		})
		if err != nil {
			l.Errorf("%s.NewShadowRepo.FindByFilter  err:%v",
				utils.FuncName(), err)
			return nil, err
		}
		if len(shadows) != 0 {
			//插入多条设备物模型属性数据
			err = l.repo.InsertPropertiesData(l.ctx, l.schema, msg.ProductID, msg.DeviceName, shadow.ToValues(shadows, l.schema.Property), time.Now(), msgThing.Optional{Sync: true})
			if err != nil {
				l.Errorf("%s.InsertPropertyData err=%+v", utils.FuncName(), err)
				return l.DeviceResp(msg, errors.Database.AddDetail(err), nil), err
			}
			now := time.Now()
			for _, v := range shadows {
				v.UpdatedDeviceTime = &now
			}
			err = sr.MultiUpdate(l.ctx, shadows)
			if err != nil {
				l.Errorf("%s.MultiUpdate err=%+v", utils.FuncName(), err)
				return l.DeviceResp(msg, errors.Database.AddDetail(err), nil), err
			}
		}
	}
	var propertyMap = schema.PropertyMap{}
	for _, d := range dataIDs {
		p := l.schema.Property[d]
		if p != nil {
			propertyMap[p.Identifier] = p
		}
	}
	if len(propertyMap) == 0 {
		propertyMap = l.schema.Property
	}
	for id, v := range propertyMap {
		data, err := l.repo.GetLatestPropertyDataByID(l.ctx, v, msgThing.LatestFilter{
			ProductID:  msg.ProductID,
			DeviceName: msg.DeviceName,
			DataID:     id,
		})
		if err != nil {
			l.Errorf("%s.GetPropertyDataByID.get id:%s err:%s",
				utils.FuncName(), id, err.Error())
			return nil, err
		}

		if data == nil {
			l.Infof("%s.GetPropertyDataByID not find id:%s", utils.FuncName(), id)
			respData[id], err = v.Define.GetDefaultValue()
			if err != nil {
				l.Errorf("%s.GetDefaultValue id:%s err:%s",
					utils.FuncName(), id, err.Error())
				return nil, err
			}
			continue
		}
		respData[id] = data.Param
	}

	return l.DeviceResp(msg, errors.OK, respData), nil
}

// 属性上报
func (l *ThingLogic) HandleProperty(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Debugf("%s req:%v", utils.FuncName(), msg)
	switch l.dreq.Method { //操作方法
	case deviceMsg.PackReport:
		return l.HandlePackReport(msg, l.dreq)
	case deviceMsg.GetReportReply:
		if l.dreq.Code != errors.OK.Code { //如果不成功,则记录日志即可
			return nil, errors.DeviceResp.AddMsg(l.dreq.Msg).AddDetail(msg.Payload)
		}
		if param, ok := l.dreq.Data.(map[string]any); ok {
			l.dreq.Params = param //新版通过data传递
		}
		_, err = l.HandlePropertyReport(msg, l.dreq)
		return nil, err
	case deviceMsg.Report: //设备属性上报
		return l.HandlePropertyReport(msg, l.dreq)
	case deviceMsg.ReportInfo: //设备基础信息上报
		return l.HandlePropertyReportInfo(msg, l.dreq)
	case deviceMsg.GetStatus: //设备请求获取 云端记录的最新设备信息
		return l.HandlePropertyGetStatus(msg)
	case deviceMsg.GetSchema:
		return l.HandlePropertyGetSchema(msg)
	case deviceMsg.ControlReply: //设备响应的 “云端下发控制指令” 的处理结果
		return l.HandleControl(msg)
	default:
		return nil, errors.Method.AddMsg(l.dreq.Method)
	}
}

func (l *ThingLogic) HandleEvent(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Debugf("%s req:%v", utils.FuncName(), msg)

	dbData := msgThing.EventData{}
	dbData.Identifier = l.dreq.EventID
	dbData.Type = l.dreq.Type

	if l.dreq.Method != deviceMsg.EventPost {
		return nil, errors.Method
	}

	tp, err := l.dreq.VerifyReqParam(l.schema, schema.ParamEvent)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}

	dbData.Params, err = msgThing.ToVal(tp)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}
	dbData.TimeStamp = l.dreq.GetTimeStamp(msg.Timestamp)
	paramValues, err := msgThing.ToParamValues(tp)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}
	err = l.svcCtx.PubApp.DeviceThingEventReport(l.ctx, application.EventReport{
		Device:     devices.Core{ProductID: msg.ProductID, DeviceName: msg.DeviceName},
		Timestamp:  dbData.TimeStamp.UnixMilli(),
		Identifier: dbData.Identifier,
		Params:     paramValues,
		Type:       dbData.Type,
	})
	if err != nil {
		l.Errorf("%s.DeviceThingEventReport  err:%v", utils.FuncName(), err)
	}

	err = l.repo.InsertEventData(l.ctx, msg.ProductID, msg.DeviceName, &dbData)
	if err != nil {
		l.Errorf("%s.InsertEventData err=%+v", utils.FuncName(), err)
		return l.DeviceResp(msg, errors.Database.AddDetail(err), nil), errors.Database.AddDetail(err)
	}
	return l.DeviceResp(msg, errors.OK, nil), nil
}

func (l *ThingLogic) HandleAction(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Debugf("%s req:%v", utils.FuncName(), msg)
	core := devices.Core{
		ProductID:  msg.ProductID,
		DeviceName: msg.DeviceName,
	}
	reqType := deviceMsg.ReqMsg
	timeStamp := l.dreq.GetTimeStamp(msg.Timestamp)
	switch l.dreq.Method {
	case deviceMsg.Action: //设备请求云端
		err = cache.SetDeviceMsg(l.ctx, l.svcCtx.Cache, deviceMsg.ReqMsg, msg, l.dreq.MsgToken)
		if err != nil {
			return nil, err
		}
		ctxs.GoNewCtx(l.ctx, func(ctx context.Context) {
			l.Infof("DeviceThingActionReport.Action device:%v,reqType:%v,req:%v", core, reqType, l.dreq)
			appMsg := application.ActionReport{
				Device: core, Timestamp: timeStamp.UnixMilli(), ReqType: reqType, MsgToken: l.dreq.MsgToken,
				ActionID: l.dreq.ActionID, Params: l.dreq.Params, Dir: schema.ActionDirUp,
				Code: l.dreq.Code, Msg: l.dreq.Msg,
			}
			//应用事件通知-设备物模型事件上报通知 ↓↓↓
			err := l.svcCtx.PubApp.DeviceThingActionReport(ctx, appMsg)
			if err != nil {
				logx.WithContext(ctx).Errorf("%s.DeviceThingActionReport.Action  req:%v,err:%v", utils.FuncName(), utils.Fmt(l.dreq), err)
			}
			di, err := l.svcCtx.DeviceCache.GetData(l.ctx, core)
			if err != nil {
				l.Error(err)
			}
			if di != nil {
				err = l.svcCtx.UserSubscribe.Publish(ctx, def.UserSubscribeDeviceActionReport, appMsg, map[string]any{
					"productID":  core.ProductID,
					"deviceName": core.DeviceName,
					"identifier": l.dreq.ActionID,
				}, map[string]any{
					"productID":  core.ProductID,
					"deviceName": core.DeviceName,
				}, map[string]any{
					"projectID": di.ProjectID,
				}, map[string]any{
					"projectID": cast.ToString(di.ProjectID),
					"areaID":    cast.ToString(di.AreaID),
				})
			}
			if err != nil {
				logx.WithContext(ctx).Error(err)
			}
		})
	case deviceMsg.ActionReply: //云端请求设备的回复
		reqType = deviceMsg.RespMsg
		var resp msgThing.Resp
		err = utils.Unmarshal(msg.Payload, &resp)
		if err != nil {
			return nil, errors.Parameter.AddDetailf("payload unmarshal payload:%v err:%v", string(msg.Payload), err)
		}

		req, err := cache.GetDeviceMsg[msgThing.Req](l.ctx, l.svcCtx.Cache, deviceMsg.ReqMsg, msg.Handle, msg.Type,
			devices.Core{ProductID: msg.ProductID, DeviceName: msg.DeviceName},
			resp.MsgToken)
		if req == nil || err != nil {
			l.Error(req, err)
			return nil, err
		}

		err = cache.SetDeviceMsg(l.ctx, l.svcCtx.Cache, deviceMsg.RespMsg, msg, resp.MsgToken)
		if err != nil {
			l.Error(err)
			return nil, err
		}
		ctxs.GoNewCtx(l.ctx, func(ctx context.Context) {
			l.Infof("DeviceThingActionReport.ActionReply device:%v,reqType:%v,req:%v", core, reqType, l.dreq)
			_, err := l.svcCtx.TimedM.TaskCancel(l.ctx, &timedmanage.TaskWithTaskID{TaskID: resp.MsgToken})
			if err != nil {
				logx.WithContext(ctx).Error(err)
			}
			param, _ := resp.Data.(map[string]any)
			//应用事件通知-设备物模型事件上报通知 ↓↓↓
			appMsg := application.ActionReport{
				Device: core, Timestamp: timeStamp.UnixMilli(), ReqType: reqType, MsgToken: resp.MsgToken,
				ActionID: resp.ActionID, Params: param, Dir: schema.ActionDirUp, Code: resp.Code, Msg: resp.Msg,
			}
			err = l.svcCtx.PubApp.DeviceThingActionReport(ctx, appMsg)
			if err != nil {
				logx.WithContext(ctx).Errorf("%s.DeviceThingActionReport  req:%v,err:%v", utils.FuncName(), utils.Fmt(l.dreq), err)
			}
			di, err := l.svcCtx.DeviceCache.GetData(l.ctx, core)
			if err != nil {
				l.Error(err)
			}
			if di != nil {
				err = l.svcCtx.UserSubscribe.Publish(ctx, def.UserSubscribeDeviceActionReport, appMsg, map[string]any{
					"productID":  core.ProductID,
					"deviceName": core.DeviceName,
					"identifier": resp.ActionID,
				}, map[string]any{
					"productID":  core.ProductID,
					"deviceName": core.DeviceName,
				}, map[string]any{
					"projectID": di.ProjectID,
				}, map[string]any{
					"projectID": cast.ToString(di.ProjectID),
					"areaID":    cast.ToString(di.AreaID),
				})
			}
			if err != nil {
				logx.WithContext(ctx).Error(err)
			}
		})
	}

	return nil, nil
}

func (l *ThingLogic) HandleControl(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Debugf("%s req:%v", utils.FuncName(), msg)

	var resp msgThing.Resp
	err = utils.Unmarshal(msg.Payload, &resp)
	if err != nil {
		return nil, errors.Parameter.AddDetailf("payload unmarshal payload:%v err:%v", string(msg.Payload), err)
	}

	req, err := cache.GetDeviceMsg[msgThing.Req](l.ctx, l.svcCtx.Cache, deviceMsg.ReqMsg, msg.Handle, msg.Type,
		devices.Core{ProductID: msg.ProductID, DeviceName: msg.DeviceName},
		resp.MsgToken)
	if req == nil || err != nil {
		return nil, err
	}

	err = cache.SetDeviceMsg(l.ctx, l.svcCtx.Cache, deviceMsg.RespMsg, msg, resp.MsgToken)
	if err != nil {
		return nil, err
	}

	if resp.Code == errors.OK.GetCode() { //如果设备回复了,且处理成功,需要入库
		_, err = l.HandlePropertyReport(msg, *req)
		return nil, err
	}
	return nil, nil
}

// Handle for topics.DeviceUpThingAll
func (l *ThingLogic) Handle(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Infof("%s req=%v", utils.FuncName(), msg)

	err = l.initMsg(msg)
	if err != nil {
		return nil, err
	}

	var action = devices.Thing
	respMsg, err = func() (respMsg *deviceMsg.PublishMsg, err error) {
		action = msg.Type
		switch msg.Type { //操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为
		case msgThing.TypeProperty: //设备上报的 属性或信息
			return l.HandleProperty(msg)
		case msgThing.TypeEvent: //设备上报的 事件
			return l.HandleEvent(msg)
		case msgThing.TypeAction: //设备响应的 “应用调用设备行为”的执行结果
			return l.HandleAction(msg)
		default:
			action = devices.Thing
			return nil, errors.Parameter.AddDetailf("things types is err:%v", msg.Type)
		}
	}()
	if l.dreq.NoAsk() { //如果不需要回复
		respMsg = nil
	}
	hub := &deviceLog.Hub{
		ProductID:   msg.ProductID,
		Action:      action,
		Timestamp:   time.Now(), // 操作时间
		DeviceName:  msg.DeviceName,
		TraceID:     utils.TraceIdFromContext(l.ctx),
		RequestID:   l.dreq.MsgToken,
		Content:     string(msg.Payload),
		Topic:       msg.Topic,
		ResultCode:  errors.Fmt(err).GetCode(),
		RespPayload: respMsg.GetPayload(),
	}
	_ = l.svcCtx.HubLogRepo.Insert(l.ctx, hub)
	l.svcCtx.UserSubscribe.Publish(l.ctx, def.UserSubscribeDevicePublish, hub.ToApp(), map[string]any{
		"productID":  msg.ProductID,
		"deviceName": msg.DeviceName,
	})

	return
}
