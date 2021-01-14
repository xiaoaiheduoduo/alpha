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

	r.NoRoute(NoRoute())
	r.NoMethod(NoMethod())

	addHealthEndpoints(r)

	return r
}

func addHealthEndpoints(r *gin.Engine) {
	// For health checking
	r.GET("/livez", func(c *gin.Context) {
		c.String(200, "")
	})
	r.GET("/readyz", func(c *gin.Context) {
		c.String(200, "")
	})
}
