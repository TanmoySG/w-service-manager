# Audit-Validation Microservice

The Validation Microservice validates the contract on three parameters
- Service Name
- Repository
- Data Access

If the contract is valid, it is put into the audit topic and if it is invalid it is put into the invalid kafka topic, for further processing

## Running Locally

- Start SOQ
- Start intake microservice 
- Run the audit microservice
```sh
go run main.go
```
- Create a Request and Contract
- To see the messages being created run the observer script using make command