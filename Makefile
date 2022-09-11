start-soq:
	docker-compose -f onboarding/docker-compose.yml up -d
	sh onboarding/scripts/monitorStartup.sh


teardown-soq:
	docker-compose -f onboarding/docker-compose.yml down 

run-kafka-observer:
	node dependencies/extras

generate-validity:
	chmod 775 ./schema/codegen.sh;
	./schema/codegen.sh ./schema/service-onboarding/validity.audit.schema.json  ./onboarding/audit/validity/spec/validity/validity.go validity
	./schema/codegen.sh ./schema/service-onboarding/contract.intake.schema.json  ./onboarding/audit/validity/spec/contract/contract.go contract

run-all:
	node onboarding/intake
	cd onboarding/audit/validity && go run main.go
	node onboarding/audit/approval