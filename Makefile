start-soq:
	docker-compose -f onboarding/docker-compose.yml up -d

teardown-soq:
	docker-compose -f onboarding/docker-compose.yml down 

run-kafka-observer:
	node dependencies/extras

generate-validity:
	chmod 775 ./schema/codegen.sh;
	./schema/codegen.sh ./schema/service-onboarding/validity.audit.schema.json  ./onboarding/audit/validity/spec/validity.go