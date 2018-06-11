#network.go初始化用脚本

DOCKERENV_PATH=../blockchain/network
DOCKER_COMPOSE_CMD=docker-compose
SOURCE_CMD=source

#生成配置文件
.PHONY: generate
generate:
	@cd $(DOCKERENV_PATH) && \
		$(SOURCE_CMD) generateArtifacts.sh

#启动网络
.PHONY: networkup
networkup:
	@cd $(DOCKERENV_PATH) && \
		$(DOCKER_COMPOSE_CMD) -f $COMPOSE_FILE up -d 2>&1

#测试用
.PHONY: test
test:
	@cd $(DOCKERENV_PATH) && \
		ls -a >log.txt
		docker ps