# Onboarding Microservices

Start SOQ
```
make start-soq
```

Run Intake Microservice
```
node onboarding/intake
```

Run Audit-Validity
```
cd onboarding/audit/validity && go run main.go
```

Run Audit-Approval (Placeholder)
```
node onboarding/audit/approval
```

Run Integration
```
cd onboarding/integration && node .
```

To create a Onboarding Request and Contract, follow [this](./intake/README.md#routes)