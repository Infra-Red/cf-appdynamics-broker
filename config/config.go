package config

import "github.com/kelseyhightower/envconfig"

type Specification struct {
	AppDynamicsAccountAccessKey string `envconfig:"appdynamics_account_access_key" required:"true"`
	AppDynamicsAccountName      string `envconfig:"appdynamics_account_name" required:"true"`
	AppDynamicsAppName          string `envconfig:"appdynamics_app_name" default:""`
	AppDynamicsHost             string `envconfig:"appdynamics_host" required:"true"`
	AppDynamicsNodeName         string `envconfig:"appdynamics_node_name" default:""`
	AppDynamicsPort             string `envconfig:"appdynamics_port" required:"true"`
	AppDynamicsSSLEnabled       bool   `envconfig:"appdynamics_ssl_enabled" default:"false"`
	AppDynamicsTierName         string `envconfig:"appdynamics_tier_name" default:""`
	BrokerPassword              string `envconfig:"broker_password" required:"true"`
	BrokerUsername              string `envconfig:"broker_username" required:"true"`
	CatalogPath                 string `envconfig:"catalog_path" default:"./catalog.json"`
	LogLevel                    string `envconfig:"log_level" default:"INFO"`
	Port                        string `envconfig:"port" default:"8080"`
}

func LoadEnv() (Specification, error) {
	var env Specification
	err := envconfig.Process("", &env)
	if err != nil {
		return Specification{}, err
	}
	return env, nil
}
