# Notes

Dumping Notes, Thoughts here.

## Service Intake

Service Intake/Onboarding (part of Service Manager) is supposed to help in onboarding new Services onto the WIP for use as Identity Server. 

![Service Intake Architecture](./diagrams/service-intake.jpg)

For any service onboarded onto wIP, it can use wunder sign-up and log-in for identity, and not worry about authentication and storing identities, while working just on the service logic. 

A service that wants to get onboarded, needs to start an onboarding request to the `service-manager`. On the user/service end,
- Service registers a onboarding request
- Service Manager sends a template contract and policy in return. An UUID for the request is also generated and sent.
- These templates need to be filled by the developer/admin of the service with the required details like 
    - What data is required
    - Purpose of the Data
    - Compliance to data-usage instructions
    - details of app/service, etc.
- The Policy also needs to be filled and signed (some mech. that needs to be formulated, some kind of document signing) 
- These are then sent (verification request) to the service manager.
- Service Manager Cross-checks the UUID and the request and puts the request policy and contract on the Service Onboarding Queues.

The Queue stores all the requests and respective microservices(of service manager) use these to onboard a service onto wunder-platform and wIP.

To-Do [1 Aug] 
- Need to find better names for policy and contract 
- or maybe devide contract into request and contract, 
    - with compliance in contract 
    - and usage and other data in request