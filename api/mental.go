package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"swagger/internal/model"
)

// MentalAnswerReq is the request struct for mental.Answer
type MentalAnswerReq struct {
	g.Meta `path:"/answer" tags:"MentalSrv" method:"post" summary:"问题回答请求结构体"`
	model.MentalAnswerInput
}

type MentalAnswerRes struct {
	model.MentalAnswerOutput
}
