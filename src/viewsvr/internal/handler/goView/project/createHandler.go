package project

import (
	"gitee.com/i-Things/core/shared/errors"
	"gitee.com/i-Things/core/shared/result"
	"github.com/i-Things/things/src/viewsvr/internal/logic/goView/project"
	"github.com/i-Things/things/src/viewsvr/internal/svc"
	"github.com/i-Things/things/src/viewsvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectInfo
		if err := httpx.Parse(r, &req); err != nil {
			result.Http(w, r, nil, errors.Parameter.WithMsg("入参不正确:"+err.Error()))
			return
		}

		l := project.NewCreateLogic(r.Context(), svcCtx)
		err := l.Create(&req)
		result.Http(w, r, nil, err)
	}
}
