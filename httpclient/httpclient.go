package httpclient

import (
	"fmt"
	"strings"
	"time"

	"github.com/alphaframework/alpha/aconfig"
	"github.com/alphaframework/alpha/alog"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

const (
	defaultTimeout = 10 * time.Second

	httpProtocol  = "http://"
	httpsProtocol = "https://"
)

func NewResty(sugar *zap.SugaredLogger) *resty.Client {
	client := resty.New()

	client.SetTimeout(defaultTimeout)
	client.SetLogger(NewLogger(sugar))

	client.OnAfterResponse(logRequestMiddleware())

	return client
}

func NewRestyWith(portName aconfig.PortName, appConfig *aconfig.Application, protocol string) (*resty.Client, error) {
	location := appConfig.GetMatchedPrimaryPortLocation(portName)
	if location == nil {
		return nil, fmt.Errorf("missing matched primaryport location (port_name: %s)", portName)
	}

	client := NewResty(alog.Sugar)
	switch {
	case strings.HasPrefix(location.Address, httpProtocol):
		location.Address = location.Address[len(httpProtocol):]
	case strings.HasPrefix(location.Address, httpsProtocol):
		location.Address = location.Address[len(httpsProtocol):]
	}
	hostURL := fmt.Sprintf("%s%s:%d", protocol, location.Address, location.Port)
	client.SetHostURL(hostURL)

	return client, nil
}

func MustNewRestyWith(portName aconfig.PortName, appConfig *aconfig.Application, protocol string) *resty.Client {
	client, err := NewRestyWith(portName, appConfig, protocol)
	if err != nil {
		panic(err)
	}

	return client
}
