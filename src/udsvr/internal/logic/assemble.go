package logic

import (
	"gitee.com/i-Things/core/shared/def"
	"github.com/i-Things/things/src/udsvr/pb/ud"
)

func ToPageInfo(info *ud.PageInfo) *def.PageInfo {
	if info == nil {
		return nil
	}
	return &def.PageInfo{
		Page: info.GetPage(),
		Size: info.GetSize(),
	}
}
