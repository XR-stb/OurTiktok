package message

import (
	"net/http"

	"OutTiktok/apps/gateway/internal/logic/message"
	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageChatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewMessageChatLogic(r.Context(), svcCtx)
		resp, err := l.MessageChat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
