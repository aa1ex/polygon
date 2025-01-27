.PHONY: test-e2e
test-e2e:
	go test ./e2e/elasticsearch -v -count 1

.PHONY: test-e2e-docker-up
test-e2e-docker-up:
	for dc in $(shell find e2e -name 'docker-compose.yml') ; do \
		docker compose -f $$dc up -d ; \
	done

.PHONY: test-e2e-docker-down
test-e2e-docker-down:
	for dc in $(shell find e2e -name 'docker-compose.yml') ; do \
		docker compose -f $$dc down ; \
	done