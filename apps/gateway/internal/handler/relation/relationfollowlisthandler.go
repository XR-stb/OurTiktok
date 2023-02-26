package relation

import (
	"net/http"
	"strconv"

	"OutTiktok/apps/gateway/internal/logic/relation"
	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowListReq
		req.Token = r.FormValue("token")
		req.UserId, _ = strconv.ParseInt(r.FormValue("user_id"), 10, 64)

		l := relation.NewRelationFollowListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
