# agents-wso2

# Synopsis
This repo contains the code for the WS02 Discovery tool and the Traceability tools. 

# Discovery 
The Discovery Agent is used to discover new published APIs. The Discovery Agent pushes both REST and SOAP API definitions to Amplify Central.

If the Discovery Agent discovers an API where the inbound security is not set to PassThrough / API Key / OAuth, the correlating catalog asset will not be created. Discovered APIs that do not have the correct inbound security will only be available in the environment.

The related APIs are published to Amplify Central either as an API Service in environment or an API Service in environment and optionally as Catalog item (default behavior).


# Traceability
The Traceability Agent sends log information about APIs that have been discovered and published to Amplify Central.

<a href="./traceability/README.md"/>Traceability</a> | 
<a href="./discovery/README.md"/>Discovery</a>

