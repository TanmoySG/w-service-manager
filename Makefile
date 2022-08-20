start-soq:
	docker-compose -f onboarding/service-onboarding-q/docker-compose.yml up -d

teardown-soq:
	docker-compose -f onboarding/service-onboarding-q/docker-compose.yml down 

run-kafka-observer:
	node dependencies/extras