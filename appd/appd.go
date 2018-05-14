package appd

import (
	"context"

	"github.com/Infra-Red/cf-appdynamics-broker/config"
)

type APPD struct {
	accountAccessKey string
	accountName      string
	appName          string
	host             string
	nodeName         string
	port             string
	tierName         string
}

// Credentials contains all information that is needed to connect to AppDynamics Controller
type Credentials struct {
	AccessKey string `json:"account-access-key"`
	AppName   string `json:"application-name"`
	Host      string `json:"host-name"`
	Name      string `json:"account-name"`
	NodeName  string `json:"node-name"`
	Port      string `json:"port"`
	SSL       bool   `json:"ssl-enabled"`
	TierName  string `json:"tier-name"`
}

func New(env config.Specification) (*APPD, error) {
	return &APPD{
		accountAccessKey: env.AppDynamicsAccountAccessKey,
		accountName:      env.AppDynamicsAccountName,
		appName:          env.AppDynamicsAppName,
		host:             env.AppDynamicsHost,
		nodeName:         env.AppDynamicsNodeName,
		port:             env.AppDynamicsPort,
		tierName:         env.AppDynamicsTierName,
	}, nil
}

func (b *APPD) CreateConnection(ctx context.Context) (*Credentials, error) {
	return &Credentials{
		AccessKey: b.accountAccessKey,
		AppName:   b.appName,
		Host:      b.host,
		Name:      b.accountName,
		NodeName:  b.nodeName,
		Port:      b.port,
		SSL:       false,
		TierName:  b.tierName,
	}, nil
}
