package relation

import (
	"net/http"

	"OutTiktok/apps/gateway/internal/logic/relation"
	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFriendListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := relation.NewRelationFriendListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFriendList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
