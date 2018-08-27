CMD_CIRCLECI = $(CMD_ENV) circleci
CMD_CIRCLECI_BUILD = $(CMD_CIRCLECI) build

CIRCLECI_BUILD_FOLDER = _circleci_local_build_repo

circleci-local-build:
	$(CMD_CIRCLECI_BUILD) --volume="$(PATH_ROOT)":"/tmp/"

circleci-clear:
	$(CMD_SUDO) chmod +r $(CIRCLECI_BUILD_FOLDER) && $(CMD_RM) -f -r $(CIRCLECI_BUILD_FOLDER)

help-circleci:
	@echo "$(TEXT_FORMAT_BOLD)circleci-local-build$(TEXT_FORMAT_NORMAL)				- Run CircleCI local build"
	@echo "$(TEXT_FORMAT_BOLD)circleci-clear$(TEXT_FORMAT_NORMAL)					- Remove CircleCI temporary folder"