package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/henglory/Demo_ServiceStubby/handler"
	"github.com/henglory/Demo_ServiceStubby/service"
	"github.com/henglory/Demo_ServiceStubby/spec"
)

func doC(s service.Service, c *gin.Context) {
	var data spec.CReq
	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(500, spec.CRes{
			StatusCode: "99",
		})
	}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		c.JSON(500, spec.BRes{
			StatusCode: "99",
		})
	}

	res := handler.DoC(s, data)
	c.JSON(200, res)
}
