package httpclient

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"github.com/alphaframework/alpha/aconfig"
	"github.com/alphaframework/alpha/alog"
)

const (
	defaultTimeout = 10 * time.Second
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
