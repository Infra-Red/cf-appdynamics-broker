build:
	go build -o "appdynamics-broker" github.com/Infra-Red/cf-appdynamics-broker/cmd/appdynamics-broker

push:
	cf push

register:
	cf create-service-broker shared-appdynamics-broker admin admin http://appdynamics-broker.bosh-lite.com
	cf enable-service-access a.appdynamics

deregister:
	cf purge-service-offering a.appdynamics
	cf delete-service-broker shared-appdynamics-broker
