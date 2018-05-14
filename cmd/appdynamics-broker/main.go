package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Infra-Red/cf-appdynamics-broker/broker"
	"github.com/Infra-Red/cf-appdynamics-broker/config"
	"github.com/Infra-Red/cf-appdynamics-broker/logger"
	"github.com/pivotal-cf/brokerapi"
)

func loadServices(env config.Specification) ([]brokerapi.Service, error) {
	var service brokerapi.Service
	buf, err := ioutil.ReadFile(env.CatalogPath)
	if err != nil {
		return []brokerapi.Service{}, err
	}
	err = json.Unmarshal(buf, &service)
	if err != nil {
		return []brokerapi.Service{}, err
	}
	return []brokerapi.Service{service}, nil
}

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalln(err)
	}

	logger, err := logger.NewLogger("appdynamics-broker", env)
	if err != nil {
		log.Fatalln(err)
	}

	credentials := brokerapi.BrokerCredentials{
		Username: env.BrokerUsername,
		Password: env.BrokerPassword,
	}

	services, err := loadServices(env)
	if err != nil {
		log.Fatalln(err)
	}

	serviceBroker, err := broker.New(services, logger, env)
	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = env.Port
	}

	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
