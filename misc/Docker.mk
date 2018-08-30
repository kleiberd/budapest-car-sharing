CMD_DOCKER = docker
CMD_DOCKER_BUILD = $(CMD_DOCKER) build
CMD_DOCKER_RUN = $(CMD_DOCKER) run
CMD_DOCKER_LOGIN = $(CMD_DOCKER) login
CMD_DOCKER_PUSH = $(CMD_DOCKER) push

DOCKER = $(shell which $(CMD_DOCKER) 2> /dev/null)

DOCKER_BUILD_PARAMS ?=
DOCKER_USER ?=
DOCKER_PASS ?=
DOCKER_TAG = latest

DOCKER_NAME = kleiberd/budapest-car-sharing-backend-$1
DOCKER_BUILD_FN = $(CMD_DOCKER_BUILD) $(DOCKER_BUILD_PARAMS) --build-arg SUBDIR=$1 -t $(DOCKER_NAME):$(DOCKER_TAG) .
DOCKER_RUN_FN = $(CMD_DOCKER_RUN) $(DOCKER_RUN_PARAMS) $(DOCKER_NAME):$(DOCKER_TAG) $2

define docker
	$(if $(DOCKER),,$(error "Docker is required (https://docs.docker.com/install/)"))

	$(CMD_ENV) $1 $2
endef

docker-build-api:
	$(call docker,$(call DOCKER_BUILD_FN,api))

docker-build-builder-api: DOCKER_BUILD_PARAMS=--target=builder
docker-build-builder-api:
	$(call docker,$(call DOCKER_BUILD_FN,api))

docker-build-collector:
	$(call docker,$(call DOCKER_BUILD_FN,collector))

docker-build-builder-collector: DOCKER_BUILD_PARAMS=--target=builder
docker-build-builder-collector:
	$(call docker,$(call DOCKER_BUILD_FN,collector))

docker-run-tests-api:
	$(call docker,$(call DOCKER_RUN_FN,api,go test))

docker-run-tests-collector:
	$(call docker,$(call DOCKER_RUN_FN,collector,go test))

docker-run-api: DOCKER_RUN_PARAMS=-it -p 8080:8080
docker-run-api:
	$(call docker,$(call DOCKER_RUN_FN,api,/go/bin/api))

docker-run-collector:
	$(call docker,$(call DOCKER_RUN_FN,collector,/go/bin/collector))

docker-copy-artifact-api:
	$(call docker,$(call DOCKER_RUN_FN,api,/bin/sh -c "$(CMD_MKDIR) artifacts && $(CMD_CP) /go/bin/api /artifacts/api"))

docker-login:
	$(call docker, echo '$(DOCKER_PASS)' | $(CMD_DOCKER_LOGIN) -u $(DOCKER_USER) --password-stdin)

docker-push:
	$(call docker,$(CMD_DOCKER_PUSH) $(DOCKER_NAME):$(DOCKER_TAG))

help-docker:
	@echo "$(TEXT_FORMAT_BOLD)docker-build-api$(TEXT_FORMAT_NORMAL)				- Build full API container"
	@echo "$(TEXT_FORMAT_BOLD)docker-build-builder-api$(TEXT_FORMAT_NORMAL)			- Build API container first stage"
	@echo "$(TEXT_FORMAT_BOLD)docker-build-collector$(TEXT_FORMAT_NORMAL)				- Build full Collector container"
	@echo "$(TEXT_FORMAT_BOLD)docker-build-builder-collector$(TEXT_FORMAT_NORMAL)			- Build Collector container first stage"
	@echo "$(TEXT_FORMAT_BOLD)docker-run-tests-api$(TEXT_FORMAT_NORMAL)				- Run tests in API container"
	@echo "$(TEXT_FORMAT_BOLD)docker-run-tests-collector$(TEXT_FORMAT_NORMAL)			- Run tests in Collector container"
