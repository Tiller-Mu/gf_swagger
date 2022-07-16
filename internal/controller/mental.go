package controller

import (
	"context"
	"swagger/api"
)

type cMental struct{}

var (
	Mental = cMental{}
)

func (c *cMental) Answer(ctx context.Context, req *api.MentalAnswerReq) (res *api.MentalAnswerRes, err error) {
	res = &api.MentalAnswerRes{}

	return res, nil
}
