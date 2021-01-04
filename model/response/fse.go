package response

import "github.com/xiaoweihong/wolfweb/model"

type FseResponseResult struct {
	model.Repo
}

type FseListResponseResult struct {
	TotalCount int64        `json:"total_count"`
	Repos      []model.Repo `json:"repos"`
}
