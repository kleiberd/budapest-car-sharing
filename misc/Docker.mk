CMD_DOCKER = $(CMD_ENV) docker
CMD_DOCKER_BUILD = $(CMD_DOCKER) build
CMD_DOCKER_RUN = $(CMD_DOCKER) run

DOCKER_BUILD_PARAMS ?=
DOCKER_TAG = latest

DOCKER_NAME = kleiberd/budapest-car-sharing-backend-$1
DOCKER_BUILD_FN = $(CMD_DOCKER_BUILD) $(DOCKER_BUILD_PARAMS) --build-arg SUBDIR=$1 -t $(DOCKER_NAME):$(DOCKER_TAG) .
DOCKER_RUN_FN = $(CMD_DOCKER_RUN) $(DOCKER_RUN_PARAMS) $(DOCKER_NAME):$(DOCKER_TAG) $2

docker-build-api:
	$(call DOCKER_BUILD_FN,api)

docker-build-builder-api: DOCKER_BUILD_PARAMS=--target=builder
docker-build-builder-api:
	$(call DOCKER_BUILD_FN,api)

docker-build-collector:
	$(call DOCKER_BUILD_FN,collector)

docker-build-builder-collector: DOCKER_BUILD_PARAMS=--target=builder
docker-build-builder-collector:
	$(call DOCKER_BUILD_FN,collector)

docker-run-tests-api:
	$(call DOCKER_RUN_FN,api,go test)

docker-run-tests-collector:
	$(call DOCKER_RUN_FN,collector,go test)

docker-run-api: DOCKER_RUN_PARAMS=-it -p 8080:8080
docker-run-api:
	$(call DOCKER_RUN_FN,api,/go/bin/api)

docker-run-collector:
	$(call DOCKER_RUN_FN,collector,/go/bin/collector)

docker-copy-artifact-api:
	$(call DOCKER_RUN_FN,api,/bin/sh -c "$(CMD_MKDIR) artifacts && $(CMD_CP) /go/bin/api /artifacts/api")

help-docker:
	@echo "$(TEXT_FORMAT_BOLD)docker-build-api$(TEXT_FORMAT_NORMAL)				- Build full API container"
	@echo "$(TEXT_FORMAT_BOLD)docker-build-builder-api$(TEXT_FORMAT_NORMAL)			- Build API container first stage"
	@echo "$(TEXT_FORMAT_BOLD)docker-build-collector$(TEXT_FORMAT_NORMAL)				- Build full Collector container"
	@echo "$(TEXT_FORMAT_BOLD)docker-build-builder-collector$(TEXT_FORMAT_NORMAL)			- Build Collector container first stage"
	@echo "$(TEXT_FORMAT_BOLD)docker-run-tests-api$(TEXT_FORMAT_NORMAL)				- Run tests in API container"
	@echo "$(TEXT_FORMAT_BOLD)docker-run-tests-collector$(TEXT_FORMAT_NORMAL)			- Run tests in Collector container"
