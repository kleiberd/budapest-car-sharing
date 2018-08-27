CMD_PACKER = $(CMD_ENV) packer
CMD_PACKER_BUILD = $(CMD_PACKER) build

PACKER_FOLDER = .packer
PATH_PACKER = $(PATH_ROOT)/$(PACKER_FOLDER)
PATH_BUILD_FILE = $(PATH_PACKER)/bcsb-image.json
PATH_IMAGE_ID = $(PATH_PACKER)/bcsb-image.id

PACKER_IMAGE_ID = $(shell cat ${PATH_IMAGE_ID})

packer-build:
	$(CMD_CD) $(PACKER_FOLDER) && $(CMD_PACKER_BUILD) $(PATH_BUILD_FILE)

help-packer:
	@echo "$(TEXT_FORMAT_BOLD)packer-build$(TEXT_FORMAT_NORMAL)					- Build base image for Terraform with Packer"