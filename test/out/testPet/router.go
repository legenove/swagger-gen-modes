package testPet

import (
    "github.com/gin-gonic/gin"
    "github.com/legenove/nano-server-sdk/gincore"
)

type ApiBaseHandler func(c *gin.Context) (int, interface{})

func decoratorHandler(handler ApiBaseHandler, decors ...gincore.HandlerDecorator) gin.HandlerFunc {
	apiFunc := func(c *gin.Context) {
		code, obj := handler(c)
		if obj != nil {
			c.JSON(code, obj)
		} else {
			c.Status(code)
		}
	}
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		apiFunc = d(apiFunc)
	}
	return apiFunc
}

