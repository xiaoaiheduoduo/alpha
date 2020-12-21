package main

import (
	"fmt"

	_ "github.com/alphaframework/alpha/aconfig"
	_ "github.com/alphaframework/alpha/aerror"
	_ "github.com/alphaframework/alpha/alog"
	_ "github.com/alphaframework/alpha/alog/gormwrapper"
	_ "github.com/alphaframework/alpha/autil"
	_ "github.com/alphaframework/alpha/autil/ahttp"
	_ "github.com/alphaframework/alpha/autil/ahttp/request"
	_ "github.com/alphaframework/alpha/database"
	_ "github.com/alphaframework/alpha/ginwrapper"
	_ "github.com/alphaframework/alpha/httpclient"
	_ "github.com/alphaframework/alpha/httpserver/rsp"
)

func main() {
	fmt.Println("Hello world")
}
