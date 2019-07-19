package handler

import (
	"github.com/henglory/Demo_A_ServiceStub/spec"
)

func DoA(s interface {
}, req spec.AReq) (res spec.ARes) {
	if req.X > 1000 || req.Y > 1000 {
		res.StatusCode = "01"
		return
	}
	res.Result = req.X + req.Y
	res.StatusCode = "00"
	return
}
