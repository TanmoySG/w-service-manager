# Service Onboarding Namespcae/Publisher Creation
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"publisher":"w-service-onboarding","origin":"wunder.service.onboarding","description":"Namespace for Service-onboarding Microservices"}' \
    http://localhost:8080/publisher

# Service Onboarding Context Creation
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"context":"w-so-intake","origin":"w.so.intake.test","description":"SO-Intake Context","kind":{"logs":[]}}' \
    http://localhost:8080/w-service-onboarding/context
