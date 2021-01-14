package ginwrapper

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/alphaframework/alpha/alog"
)

func defaultHealthHandler(c *gin.Context) {
	c.String(200, "")
}

type Options struct {
	// LivezHandler checks if the app is alive
	LivezHandler func(c *gin.Context)
	// ReadyzHandler checks if the app is ready for incoming requests
	ReadyzHandler func(c *gin.Context)
}

func (o *Options) complete() {
	if o.LivezHandler == nil {
		o.LivezHandler = defaultHealthHandler
	}
	if o.ReadyzHandler == nil {
		o.ReadyzHandler = defaultHealthHandler
	}
}

func New(options *Options) *gin.Engine {
	if options == nil {
		options = &Options{}
	}
	options.complete()

	r := gin.New()
	r.Use(GinResponseBodyLogMiddleware())
	r.Use(Tracer())
	r.Use(Ginzap(alog.Logger, time.RFC3339, true))
	r.Use(RecoveryWithZap(alog.Logger, true))

	r.NoRoute(NoRoute())
	r.NoMethod(NoMethod())

	addHealthEndpoints(r, options)

	return r
}

func addHealthEndpoints(r *gin.Engine, options *Options) {
	r.GET("/livez", options.LivezHandler)
	r.GET("/readyz", options.ReadyzHandler)
}
