CIRCLECI_BUILD_FOLDER = _circleci_local_build_repo

circleci-local-build:
	$(CMD_CIRCLECI_BUILD) --volume="$(PATH_ROOT)":"/tmp/"

circleci-flush:
	$(CMD_SUDO) chmod +r $(CIRCLECI_BUILD_FOLDER) && $(CMD_RM) -R $(CIRCLECI_BUILD_FOLDER)

help-circleci:
	@echo "$(TEXT_FORMAT_BOLD)circleci-local-build$(TEXT_FORMAT_NORMAL)				- Run CircleCI local build"
	@echo "$(TEXT_FORMAT_BOLD)circleci-flush$(TEXT_FORMAT_NORMAL)					- Remove CircleCI temporary folder"