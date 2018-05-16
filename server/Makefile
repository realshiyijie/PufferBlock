DOCKERENV_PATH=../blockchain/network
DOCKER_COMPOSE_CMD=docker-compose
.PHONY: example-network
example-network:
	@cd $(DOCKERENV_PATH) && \
		$(DOCKER_COMPOSE_CMD) -f docker-compose.yaml up --force-recreate

.PHONY: test
test:
	@cd $(DOCKERENV_PATH) && \
		ls -a >log.txt
		docker ps