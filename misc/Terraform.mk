CMD_TERRAFORM = terraform
CMD_TERRAFORM_INIT = $(CMD_TERRAFORM) init
CMD_TERRAFORM_PLAN = $(CMD_TERRAFORM) plan
CMD_TERRAFORM_APPLY = $(CMD_TERRAFORM) apply
CMD_TERRAFORM_DESTROY = $(CMD_TERRAFORM) destroy

TERRAFORM = $(shell which $(CMD_TERRAFORM) 2> /dev/null)

TERRAFORM_FN = $1 -var 'image_id=$(PACKER_IMAGE_ID)' $2

TERRAFORM_FOLDER = .terraform

define terraform
	$(if $(TERRAFORM),,$(error "Terraform is required (https://www.terraform.io/downloads.html)"))

	$(CMD_CD) $(TERRAFORM_FOLDER) && $(CMD_ENV) $1
endef

terraform-init:
	$(call terraform,$(CMD_TERRAFORM_INIT))

terraform-plan: terraform-init
	$(call terraform,$(call TERRAFORM_FN,$(CMD_TERRAFORM_PLAN)))

terraform-apply: terraform-init
	$(call terraform,$(call TERRAFORM_FN,$(CMD_TERRAFORM_APPLY)))

terraform-destroy: terraform-init
	$(call terraform,$(call TERRAFORM_FN,$(CMD_TERRAFORM_DESTROY)))

terraform-apply-auto-approve: terraform-init
	$(call terraform,$(call TERRAFORM_FN,$(CMD_TERRAFORM_APPLY),-auto-approve))

terraform-destroy-auto-approve: terraform-init
	$(call terraform,$(call TERRAFORM_FN,$(CMD_TERRAFORM_DESTROY),-auto-approve))

help-terraform:
	@echo "$(TEXT_FORMAT_BOLD)terraform-init$(TEXT_FORMAT_NORMAL)					- Initialize Terraform providers"
	@echo "$(TEXT_FORMAT_BOLD)terraform-plan$(TEXT_FORMAT_NORMAL)					- Check Terraform execution plan"
	@echo "$(TEXT_FORMAT_BOLD)terraform-apply$(TEXT_FORMAT_NORMAL)					- Apply Terraform execution plan"
	@echo "$(TEXT_FORMAT_BOLD)terraform-destroy$(TEXT_FORMAT_NORMAL)				- Destroy all Terraform executed actions"
	@echo "$(TEXT_FORMAT_BOLD)terraform-apply-auto-approve$(TEXT_FORMAT_NORMAL)			- Apply plan with auto approve"
	@echo "$(TEXT_FORMAT_BOLD)terraform-destroy-auto-approve$(TEXT_FORMAT_NORMAL)			- Destroy all actions with auto approve"