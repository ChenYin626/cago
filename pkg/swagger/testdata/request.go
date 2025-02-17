package testdata

import (
	"github.com/codfrm/cago/pkg/utils/httputils"
	"github.com/codfrm/cago/server/mux"
)

type Item struct {
	ID int64 `json:"id"`
}

// ListRequest 获取脚本列表
type ListRequest struct {
	mux.Meta                     `path:"/script" method:"GET"`
	Type                         int    `form:"type" binding:"oneof=1 2 3 4"` // 1: 脚本 2: 库 3: 后台脚本 4: 定时脚本
	Sort                         string `form:"sort" binding:"oneof=today_download total_download score createtime updatetime"`
	httputils.PageRequest[*Item] `form:",inline"`
}
