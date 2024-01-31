package logic

import (
	"gitee.com/i-Things/core/shared/def"
	"github.com/i-Things/things/src/viewsvr/internal/types"
)

func ToPageInfo(info *types.PageInfo) *def.PageInfo {
	if info == nil {
		return nil
	}
	return &def.PageInfo{
		Page: info.Page,
		Size: info.Size,
	}
}
