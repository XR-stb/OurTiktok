package publish

import (
	"net/http"

	"OutTiktok/apps/gateway/internal/logic/publish"
	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionReq
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//	return
		//}

		req.Title = r.FormValue("title")
		req.Token = r.FormValue("token")
		file, fileheader, err := r.FormFile("data")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		data := make([]byte, fileheader.Size)
		_, err = file.Read(data)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.Data = data

		l := publish.NewPublishActionLogic(r.Context(), svcCtx)
		resp, err := l.PublishAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
