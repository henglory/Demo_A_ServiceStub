package handler

import (
	"github.com/henglory/Demo_ServiceStubby/spec"
)

func DoB(s interface {
}, req spec.BReq) (res spec.BRes) {
	res.Result = req.X * 2
	res.StatusCode = "00"
	return
}
