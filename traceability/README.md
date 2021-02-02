
# Traceability Agent
The Traceability Agent sends log information about APIs that have been discovered and published to Amplify Central.


# Prerequisite
1. Golang 
2. Make
3. API Central Account
4. API Platform Organization


# Setting Up Amplify Central Access

## Find Organizion ID
<img src="./../img/org.png" width="600"/>

## Create Service Account
<img src="./../img/account.png" width="600"/>

Click the `+Service Account` Button

Add a name and public key

Create a Service Account in Central so that the Agents can connect to the Gateway without exposing client credentials

To generate a public key, you can install OpenSSL and run the commands below:
`openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in private_key.pem -out public_key.pem`




# Docker

## Install
Install WS02 Docker container

`docker run -it -p 8280:8280 -p 8243:8243 -p 9443:9443 --name api-manager wso2/wso2am:3.2.0`

## Running Contianer
` https://localhost:9443/publisher `


# Publish API
Create a basic Pizza API
`https://apim.docs.wso2.com/en/latest/learn/design-api/create-api/create-a-rest-api/`





# Steps to implement traceability agent using this stub
1. Locate the commented tag "CHANGE_HERE" for package import paths in all files and fix them to reference the path correctly.
2. Run "make dep" to resolve the dependencies. This should resolve all dependency packages and vendor them under ./vendor directory
3. Update Makefile to change the name of generated binary image from *apic_traceability_agent* to the desired name. Locate *apic_traceability_agent* string and replace with desired name
4. Update pkg/cmd/root.go to change the name and description of the agent. Locate *apic_traceability_agent* and *Sample Traceability Agent* and replace to desired values
5. Update pkg/config/config.go to define the gateway specific configuration
    - Locate *gateway-section* and replace with the name of your gateway. Same string in pkg/cmd/root.go and sample YAML config file
    - Define gateway specific config properties in *GatewayConfig* struct. Locate the struct variables *ConfigKey1* & struct *config_key_1* and add/replace desired config properties
    - Add config validation. Locate *ValidateCfg()* method and update the implementation to add validation specific to gateway specific config.
    - Update the config binding with command line flags in init(). Locate *gateway-section.config_key_1* and add replace desired config property bindings
    - Update the initialization of gateway specific by parsing the binded properties. Locate *ConfigKey1* & *gateway-section.config_key_1* and add/replace desired config properties
6. Locate pkg/gateway/definition.go to define the structure of the log entry the traceability agent will receive. See pkg/gateway/definition.go for sample definition.
7. Implement the mechanism to read the log entry. Optionally you can wrap the existing beat(for e.g. filebeat) in beater.New() to read events and they setup output event processor to process the events
8. Locate pkg/gateway/eventprocessor.go to perform processing on event to be published. The processing can be performed either on the received event by beat input or before the event is published by transport. See pkg/gateway/eventprocessor.go for example of both type of processing.
9. Locate pkg/gateway/eventmapper.go to map the log entry received by beat to event structure expected for AMPLIFY Central Observer.
10. Run "make build" to build the agent.
11. Rename *apic_traceability_agent.yml* file to the desired agents name and setup the agent config in the file.
12. Execute the agent by running the binary file generated under *bin* directory. The YAML config must be in the current working directory
13. To produce traffic update the ./logs/traffic.log file with a new entry. See ./logs/traffic.log for sample entries

Reference: [SDK Documentation - Building Traceability Agent](https://github.com/Axway/agent-sdk/blob/main/docs/traceability/index.md)