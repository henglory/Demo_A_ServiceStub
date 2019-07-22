package handler

import (
	"github.com/henglory/Demo_ServiceStubby/spec"
)

func DoC(s interface {
}, req spec.CReq) (res spec.CRes) {
	if req.X < req.Y {
		res.StatusCode = "01"
		return
	}
	res.Result = req.X - req.Y
	res.StatusCode = "00"
	return
}
