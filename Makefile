start-soq:
	docker-compose -f onboarding/docker-compose.yml up -d

teardown-soq:
	docker-compose -f onboarding/docker-compose.yml down 

run-kafka-observer:
	node dependencies/extras