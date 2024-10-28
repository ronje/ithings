package otamanagelogic

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/oss/common"
	"gitee.com/unitedrhino/share/stores"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type OtaFirmwareInfoDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PiDB  *relationDB.ProductInfoRepo
	OfDB  *relationDB.OtaFirmwareInfoRepo
	OffDB *relationDB.OtaFirmwareFileRepo
}

func NewOtaFirmwareInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtaFirmwareInfoDeleteLogic {
	return &OtaFirmwareInfoDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PiDB:   relationDB.NewProductInfoRepo(ctx),
		OfDB:   relationDB.NewOtaFirmwareInfoRepo(ctx),
		OffDB:  relationDB.NewOtaFirmwareFileRepo(ctx),
	}
}

// 删除升级包
func (l *OtaFirmwareInfoDeleteLogic) OtaFirmwareInfoDelete(in *dm.WithID) (*dm.Empty, error) {
	//todo debug
	//if err := ctxs.IsRoot(l.ctx); err != nil {
	//	return nil, err
	//}
	l.ctx = ctxs.WithRoot(l.ctx)
	fi, err := l.OfDB.FindOneByFilter(l.ctx, relationDB.OtaFirmwareInfoFilter{ID: in.Id, WithFiles: true})
	if errors.Cmp(err, errors.NotFind) {
		l.Errorf("not find firmware id:" + cast.ToString(in.Id))
		return nil, err
	} else if err != nil {
		return nil, err
	}

	//开启事务
	db := stores.GetCommonConn(l.ctx)
	err = db.Transaction(func(tx *gorm.DB) error {
		//删除升级包文件
		offDb := relationDB.NewOtaFirmwareFileRepo(tx)
		err = offDb.Delete(l.ctx, in.Id)
		if err != nil {
			l.Errorf("%s.DeleteOTAFirmwareFile err=%v", utils.FuncName(), err)
			return err
		}
		ofDb := relationDB.NewOtaFirmwareInfoRepo(tx)
		//删除升级包
		err = ofDb.Delete(l.ctx, in.Id)
		if err != nil {
			l.Errorf("%s.DeleteOTAFirmware err=%v", utils.FuncName(), err)
			return err
		}
		err = relationDB.NewDeviceInfoRepo(tx).UpdateWithField(l.ctx, relationDB.DeviceFilter{ProductID: fi.ProductID, NeedConfirmVersion: fi.Version}, map[string]any{
			"need_confirm_job_id":  0,
			"need_confirm_version": "",
		})
		if err != nil {
			return err
		}
		if len(fi.Jobs) > 0 {
			err = relationDB.NewOtaJobRepo(tx).DeleteByFilter(l.ctx, relationDB.OtaJobFilter{FirmwareID: in.Id})
			if err != nil {
				l.Errorf("%s.DeleteOTAFirmware err=%v", utils.FuncName(), err)
				return err
			}
		}

		err = relationDB.NewOtaFirmwareDeviceRepo(tx).DeleteByFilter(l.ctx, relationDB.OtaFirmwareDeviceFilter{FirmwareID: in.Id})
		// 如果所有操作成功，提交事务
		return err
	})
	if err != nil {
		l.Errorf("failed to commit transaction: %v", err)
		return nil, err
	}
	for _, v := range fi.Files {
		err := l.svcCtx.OssClient.PrivateBucket().Delete(l.ctx, v.FilePath, common.OptionKv{})
		if err != nil {
			l.Error(err)
		}
	}
	return &dm.Empty{}, nil
}
