package broker

import (
	"context"

	"code.cloudfoundry.org/lager"
	"github.com/Infra-Red/cf-appdynamics-broker/appd"
	"github.com/Infra-Red/cf-appdynamics-broker/config"
	"github.com/pivotal-cf/brokerapi"
)

// New returns a new appdynamics service broker instance.
func New(services []brokerapi.Service, logger lager.Logger, env config.Specification) (brokerapi.ServiceBroker, error) {
	conn, err := appd.New(env)
	if err != nil {
		return nil, err
	}
	return &appdynamicsBroker{services: services, logger: logger, env: env, appd: conn}, nil
}

type appdynamicsBroker struct {
	services []brokerapi.Service
	logger   lager.Logger
	env      config.Specification
	appd     *appd.APPD
}

func (sb *appdynamicsBroker) Services(context context.Context) ([]brokerapi.Service, error) {
	return sb.services, nil
}

func (sb *appdynamicsBroker) Provision(context context.Context, instanceID string,
	details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{
		IsAsync:      false,
		DashboardURL: "not implemented",
	}, nil
}

func (sb *appdynamicsBroker) Deprovision(context context.Context, instanceID string,
	details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (sb *appdynamicsBroker) Bind(context context.Context, instanceID,
	bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	creds, err := sb.appd.CreateConnection(context)
	if err != nil {
		return brokerapi.Binding{}, err
	}
	return brokerapi.Binding{
		Credentials: creds,
	}, nil
}

func (sb *appdynamicsBroker) Unbind(context context.Context, instanceID, bindingID string,
	details brokerapi.UnbindDetails) error {
	return nil
}

func (sb *appdynamicsBroker) Update(context context.Context, instanceID string,
	details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

func (sb *appdynamicsBroker) LastOperation(context context.Context, instanceID,
	operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}
