# Service Onboarding - Intake

The Intake Microservice Processes new Onboarding Requests and Contracts. 

## Usage/Testing

Run the Service Onboarding Q. From the project's root directory run
```
make start-soq
```

Then for observability of the Kafka Cluster - to see if events are being produced to the kafka topic run the following command from project's root
```
make run-kafka-observer
```

Then Start the Intake Microservice
```
node onboarding/intake  
``` 

This starts the service and it listens to Port 8080 of localhost.

### Routes

There are two routes for the Intake Microservice

#### `/v1/intake/request` for New Requests

Using any REST Client - POSTMAN, Insomnia, RapidAPI or cURL, you can register a new onboarding request. The Schema of the request can be found [here](../../schema/service-onboarding/request.intake.schema.json). Now using the `POST` method to create/send the request to the server. The Server returns the Contract with some pre-filled fields and instructions to fill in the contract.
```
POST 
    {URL}/v1/intake/request

PAYLOAD
    {
        "kind": "request.intake.service-onboarding",
        "service_name": "test-service-1",
        "request_type": "create",
        "service_owners": [
            "mail@gmail.com"
        ]
    }

RESPONSE
    {
        "request-details": {
            "request_id": "request-6cd2a081-3019-44a0-a2dd-f5b8d926d775",
            "request_status": "queued",
            "timestamp": 1660970878973
        },
        "schema-path": "",
        "documentation": "",
        "instruction": [
            "1. Fill in the contract-template fields",
            "2. Populate the request_id field with the request_id in the request-details section",
            "3. Take the contract-template object and apply/send to the Web API Endpoint",
            "4. Refer to the schema for the contract from the above link in schema-path",
            "5. Make sure the kind is set to - contract.intake.service-onboarding",
            "6. For Additional Details refer to the documentation."
        ],
        "contract-template": {
            "kind": "contract.intake.service-onboarding",
            "request_id": "request-6cd2a081-3019-44a0-a2dd-f5b8d926d775",
            "service": {
            "name": "test-service-1",
            "repository": "<url://link.to.repository>",
            "details": [
                "A short description of the app.",
                "Needs at least to points"
            ]
            },
            "data": [
                {
                    "data" : "field-name",
                    "access": [
                        "read"
                    ],
                    "use": "add usage here for this field"
                }
            ],
            "developer": {
            "admin": [
                "add admin email IDs"
            ],
            "contributor": [
                "add contributor email IDs"
            ]
            }
        }
    }
```

#### `/v1/intake/contract` for Registering Contracts

Once the Request ID and the Contract are filled the same can be registered with the Intake Microservice using this route. The Schema for contract can be found [here](../../schema/service-onboarding/contract.intake.schema.json). Using the `POST` method register the contract. The response from the server is the current status of the request, the request and contract ID.

```
POST 
    {URL}/v1/intake/contract

PAYLOAD
    {
        "kind": "contract.intake.service-onboarding",
        "request_id": "request-6cd2a081-3019-44a0-a2dd-f5b8d926d775",
        "service": {
            "name": "test-service-1",
            "repository": "<url://link.to.repository>",
            "details": [
            "A short description of the app.",
            "Needs at least to points"
            ]
        },
        "data": {
            "sample-field": {
            "access": [
                "read"
            ],
            "use": "add usage here for this field"
            }
        },
        "developer": {
            "admin": [
            "mail@gmail.com"
            ],
            "contributor": []
        }
    }

RESPONSE
    {
        "contract_id": "contract-e6cc7759-fbdf-4fae-8a2f-abc0884d9799",
        "request_id": "request-6cd2a081-3019-44a0-a2dd-f5b8d926d775",
        "request_status": "proccessed",
        "timestamp": 1660971000415
    }
```

### Clean-Up

Stop the services that are running and to teardown/stop the SOQ run
```
make teardown-soq
```