#setup.go初始化用脚本
DOCKERENV_PATH=../blockchain/network
DOCKER_COMPOSE_CMD=docker-compose
SOURCE_CMD=source

.PHONY: networkup
networkup:
	@cd $(DOCKERENV_PATH) && \
		$(DOCKER_COMPOSE_CMD) -f docker-compose.yaml up --force-recreate

.PHONY: generate
generate:
	@cd $(DOCKERENV_PATH) && \
		$(SOURCE_CMD) generateArtifacts.sh

.PHONY: test
test:
	@cd $(DOCKERENV_PATH) && \
		ls -a >log.txt
		docker ps