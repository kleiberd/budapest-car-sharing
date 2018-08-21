circleci-local-build:
	$(CMD_CIRCLECI_BUILD) --volume="$(PATH_ROOT)":"/tmp/"

help-circleci:
	@echo "$(TEXT_FORMAT_BOLD)circleci$(TEXT_FORMAT_NORMAL)  		        		- Run CircleCI local build"