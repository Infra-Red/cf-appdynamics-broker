# AppDynamics Broker

This service broker shares a AppDynamics amongst many users via the Open Service Broker API.

## Configuration

There are important environment variables that should be overriden inside the manifest.yml file:

* `BROKER_USERNAME` and `BROKER_PASSWORD` are **required** to setup [HTTP Basic Auth](https://github.com/openservicebrokerapi/servicebroker/blob/v2.12/spec.md#authentication) to the broker API.
* `APPDYNAMICS_HOST` **(required)** : The controller host name.
* `APPDYNAMICS_PORT` **(required)** : The controller port.
* `APPDYNAMICS_ACCOUNT_ACCESS_KEY` **(required)** : The account access key to use when authenticating with the controller.
* `APPDYNAMICS_ACCOUNT_NAME` **(required)** : The account name to use when authenticating with the controller.
* `APPDYNAMICS_SSL_ENABLED` **(optional)** : Whether or not to use an SSL connection to the controller.
* `APPDYNAMICS_APP_NAME` **(optional)** : The application's name.
* `APPDYNAMICS_NODE_NAME` **(optional)** : The application's node name.
* `APPDYNAMICS_TIER_NAME` **(optional)** : The application's tier name.

## Deployment

1. Clone this repository, and `cd` into it.
1. Target the space you want to deploy the broker to.

    ```bash
    $ cf target -o <org> -s <space>
    ```
1. The configuration is entirely read from environment variables. Edit the manifest.yml according to your environment.
1. Deploy the broker as an application.

    ```bash
    $ cf push
    ```
1. [Register a Broker](http://docs.cloudfoundry.org/services/managing-service-brokers.html#register-broker).

    ```bash
    $ cf create-service-broker shared-appdynamics-broker $BROKER_USERNAME $BROKER_PASSWORD $APP-URL
    ```
1. [Enable Access to Service Plans](http://docs.cloudfoundry.org/services/access-control.html#enable-access).

    ```bash
    $ cf enable-service-access a.appdynamics
    ```

## Using

All credentials stored in the `VCAP_SERVICES` environment variable with the JSON key `a.appdynamics`.

1. [Create a Service Instance](https://docs.cloudfoundry.org/devguide/services/managing-services.html#create).

    ```bash
    $ cf create-service a.appdynamics standard my-appdynamics
    ```

1. [Bind a Service Instance](https://docs.cloudfoundry.org/devguide/services/managing-services.html#bind).

    ```bash
    $ cf bind-service my-app my-appdynamics
    ```

## Demo

[![asciicast](https://asciinema.org/a/4ZrRFgWuuluqmXDfpMgIL5xo3.png)](https://asciinema.org/a/4ZrRFgWuuluqmXDfpMgIL5xo3)

## License

```
Copyright 2018 Andrei Krasnitski

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
