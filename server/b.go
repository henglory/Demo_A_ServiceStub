package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/henglory/Demo_ServiceStubby/handler"
	"github.com/henglory/Demo_ServiceStubby/service"
	"github.com/henglory/Demo_ServiceStubby/spec"
)

func doB(s service.Service, c *gin.Context) {
	var data spec.BReq
	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(500, spec.BRes{
			StatusCode: "99",
		})
	}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		c.JSON(500, spec.BRes{
			StatusCode: "99",
		})
	}
	//
	res := handler.DoB(s, data)
	c.JSON(200, res)
}
