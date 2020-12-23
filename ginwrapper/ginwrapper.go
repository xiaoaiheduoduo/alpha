package ginwrapper

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/alphaframework/alpha/alog"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(GinResponseBodyLogMiddleware())
	r.Use(Tracer())
	r.Use(Ginzap(alog.Logger, time.RFC3339, true))
	r.Use(RecoveryWithZap(alog.Logger, true))

	return r
}
