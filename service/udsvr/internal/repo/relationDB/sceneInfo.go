package relationDB

import (
	"context"
	"fmt"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/stores"
	"gitee.com/unitedrhino/things/service/udsvr/internal/domain/scene"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
这个是参考样例
使用教程:
1. 将example全局替换为模型的表名
2. 完善todo
*/

type SceneInfoRepo struct {
	db *gorm.DB
}

func NewSceneInfoRepo(in any) *SceneInfoRepo {
	return &SceneInfoRepo{db: stores.GetCommonConn(in)}
}

type SceneInfoFilter struct {
	Name          string `json:"name"`
	IDs           []int64
	Status        int64
	Type          string
	Tag           string
	IsCommon      def.Bool
	AreaID        int64
	DeviceMode    scene.DeviceMode //设备模式
	ProductID     string           //产品id
	DeviceName    string           //设备名
	HasActionType string           //过滤有某个执行类型
	ProjectID     int64
}

func (p SceneInfoRepo) fmtFilter(ctx context.Context, f SceneInfoFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	if len(f.IDs) != 0 {
		db = db.Where("id in ?", f.IDs)
	}
	if f.ProjectID != 0 {
		db = db.Where("project_id = ?", f.ProjectID)
	}
	if f.IsCommon != 0 {
		db = db.Where("is_common = ?", f.IsCommon)
	}
	if f.AreaID != 0 {
		db = db.Where("area_id = ?", f.AreaID)
	}
	if f.DeviceMode != "" {
		//db = db.Where("device_mode=?", f.DeviceMode)
	}
	if f.ProductID != "" {
		db = db.Where("product_id=?", f.ProductID)
	}
	if f.DeviceName != "" {
		db = db.Where("device_name=?", f.DeviceName)
	}
	if f.Tag != "" {
		db = db.Where("tag=?", f.Tag)
	}
	if f.Name != "" {
		db = db.Where("name like ?", "%"+f.Name+"%")
	}
	if f.Type != "" {
		db = db.Where("type = ?", f.Type)
	}
	if f.Status != 0 {
		db = db.Where("status = ?", f.Status)
	}
	if len(f.HasActionType) != 0 {
		subQuery := p.db.Model(&UdSceneThenAction{}).Select("scene_id").
			Where(fmt.Sprintf("%s =?", stores.Col("type")), f.HasActionType)
		db = db.Where("id in (?)", subQuery)
	}
	return db
}

func (p SceneInfoRepo) Insert(ctx context.Context, data *UdSceneInfo) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		triggers := data.Triggers
		actions := data.Actions
		data.Triggers = nil
		data.Actions = nil
		err := p.db.WithContext(ctx).Create(data).Error
		if err != nil {
			return err
		}
		if len(triggers) != 0 {
			for i := range triggers {
				triggers[i].SceneID = data.ID
			}
			err = NewSceneIfTriggerRepo(tx).MultiInsert(ctx, triggers)
			if err != nil {
				return err
			}
		}
		if len(actions) != 0 {
			for i := range actions {
				actions[i].SceneID = data.ID
			}
			err = NewSceneActionRepo(tx).MultiInsert(ctx, actions)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return stores.ErrFmt(err)
}

func (p SceneInfoRepo) FindOneByFilter(ctx context.Context, f SceneInfoFilter) (*UdSceneInfo, error) {
	var result UdSceneInfo
	db := p.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}

func (p SceneInfoRepo) FindByFilter(ctx context.Context, f SceneInfoFilter, page *stores.PageInfo) ([]*UdSceneInfo, error) {
	var results []*UdSceneInfo
	db := p.fmtFilter(ctx, f).Preload("Actions").Preload("Triggers").Model(&UdSceneInfo{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return results, nil
}

func (p SceneInfoRepo) CountByFilter(ctx context.Context, f SceneInfoFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&UdSceneInfo{})
	err = db.Count(&size).Error
	return size, stores.ErrFmt(err)
}

func (p SceneInfoRepo) UpdateHeadImg(ctx context.Context, data *UdSceneInfo) error {
	err := p.db.WithContext(ctx).Select("head_img").Where("id=?", data.ID).Save(data).Error
	return stores.ErrFmt(err)
}

func (d SceneInfoRepo) UpdateWithField(ctx context.Context, f SceneInfoFilter, updates map[string]any) error {
	db := d.fmtFilter(ctx, f)
	err := db.Model(&UdSceneInfo{}).Updates(updates).Error
	return stores.ErrFmt(err)
}

func (p SceneInfoRepo) Update(ctx context.Context, data *UdSceneInfo) error {
	err := p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := NewSceneIfTriggerRepo(tx).DeleteByFilter(ctx, SceneIfTriggerFilter{SceneID: data.ID})
		if err != nil {
			return err
		}
		err = NewSceneActionRepo(tx).DeleteByFilter(ctx, SceneActionFilter{SceneID: data.ID})
		if err != nil {
			return err
		}
		triggers := data.Triggers
		actions := data.Actions
		data.Triggers = nil
		data.Actions = nil
		err = p.db.WithContext(ctx).Where("id = ?", data.ID).Save(data).Error
		if err != nil {
			return err
		}
		if len(triggers) != 0 {
			err = NewSceneIfTriggerRepo(tx).MultiInsert(ctx, triggers)
			if err != nil {
				return err
			}
		}
		if len(actions) != 0 {
			err = NewSceneActionRepo(tx).MultiInsert(ctx, actions)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return stores.ErrFmt(err)
}

func (p SceneInfoRepo) DeleteByFilter(ctx context.Context, f SceneInfoFilter) error {
	db := p.fmtFilter(ctx, f)
	err := db.Delete(&UdSceneInfo{}).Error
	return stores.ErrFmt(err)
}

func (p SceneInfoRepo) Delete(ctx context.Context, id int64) error {
	err := p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).Delete(&UdSceneInfo{}).Error
		if err != nil {
			return err
		}
		err = NewSceneIfTriggerRepo(tx).DeleteByFilter(ctx, SceneIfTriggerFilter{SceneID: id})
		if err != nil {
			return err
		}
		err = NewSceneActionRepo(tx).DeleteByFilter(ctx, SceneActionFilter{SceneID: id})
		if err != nil {
			return err
		}
		err = NewSceneLogRepo(tx).DeleteByFilter(ctx, SceneLogFilter{SceneID: id})
		return err
	})
	return stores.ErrFmt(err)
}
func (p SceneInfoRepo) FindOne(ctx context.Context, id int64) (*UdSceneInfo, error) {
	var result UdSceneInfo
	err := p.db.WithContext(ctx).Preload("Actions").Preload("Triggers").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}

// 批量插入 LightStrategyDevice 记录
func (p SceneInfoRepo) MultiInsert(ctx context.Context, data []*UdSceneInfo) error {
	err := p.db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Model(&UdSceneInfo{}).Create(data).Error
	return stores.ErrFmt(err)
}
