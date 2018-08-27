CMD_JQ = jq
CMD_PACKER = packer
CMD_PACKER_BUILD = $(CMD_PACKER) build

JQ = $(shell which $(CMD_JQ) 2> /dev/null)
PACKER = $(shell which $(CMD_PACKER) 2> /dev/null)

PACKER_FOLDER = .packer
PATH_PACKER = $(PATH_ROOT)/$(PACKER_FOLDER)
PATH_BUILD_FILE = $(PATH_PACKER)/bcsb-image.json
PATH_IMAGE_ID = $(PATH_PACKER)/bcsb-image.id

PACKER_IMAGE_ID = $(shell cat ${PATH_IMAGE_ID})

define packer
	$(if $(JQ),,$(error "jq JSON processor is required (https://stedolan.github.io/jq/download/)"))
	$(if $(PACKER),,$(error "Packer is required (https://www.packer.io/downloads.html)"))

	$(CMD_CD) $(PACKER_FOLDER) && $(CMD_ENV) $1
endef

packer-build:
	$(call packer, $(CMD_PACKER_BUILD) $(PATH_BUILD_FILE))

help-packer:
	@echo "$(TEXT_FORMAT_BOLD)packer-build$(TEXT_FORMAT_NORMAL)					- Build base image for Terraform with Packer"