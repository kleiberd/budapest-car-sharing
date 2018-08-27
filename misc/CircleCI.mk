CMD_CIRCLECI = circleci
CMD_CIRCLECI_BUILD = $(CMD_CIRCLECI) build

CIRCLECI = $(shell which $(CMD_CIRCLECI) 2> /dev/null)
CIRCLECI_BUILD_FOLDER = _circleci_local_build_repo

define circleci
	$(if $(CIRCLECI),,$(error "CircleCI CLI is required (https://circleci.com/docs/2.0/local-cli/#installing-the-circle-downloads-local-cli-on-macos-and-linux-distros)"))

	$(CMD_ENV) $1
endef

circleci-local-build:
	$(call circleci,$(CMD_CIRCLECI_BUILD) --volume="$(PATH_ROOT)":"/tmp/")

circleci-clean:
	$(CMD_SUDO) chmod +r $(CIRCLECI_BUILD_FOLDER) && $(CMD_RM) -f -r $(CIRCLECI_BUILD_FOLDER)

help-circleci:
	@echo "$(TEXT_FORMAT_BOLD)circleci-local-build$(TEXT_FORMAT_NORMAL)				- Run CircleCI local build"
	@echo "$(TEXT_FORMAT_BOLD)circleci-clean$(TEXT_FORMAT_NORMAL)					- Remove CircleCI temporary folder"