CMD_TERRAFORM = $(CMD_ENV) terraform
CMD_TERRAFORM_INIT = $(CMD_TERRAFORM) init
CMD_TERRAFORM_PLAN = $(CMD_TERRAFORM) plan
CMD_TERRAFORM_APPLY = $(CMD_TERRAFORM) apply
CMD_TERRAFORM_DESTROY = $(CMD_TERRAFORM) destroy

TERRAFORM_FN = $(CMD_CD) $(TERRAFORM_FOLDER) && $1 -var 'image_id=$(PACKER_IMAGE_ID)' $2

TERRAFORM_FOLDER = .terraform

terraform-init:
	$(CMD_CD) $(TERRAFORM_FOLDER) && $(CMD_TERRAFORM_INIT)

terraform-plan: terraform-init
	$(call TERRAFORM_FN,$(CMD_TERRAFORM_PLAN))

terraform-apply: terraform-init
	$(call TERRAFORM_FN,$(CMD_TERRAFORM_APPLY))

terraform-destroy: terraform-init
	$(call TERRAFORM_FN,$(CMD_TERRAFORM_DESTROY))

terraform-apply-auto-approve: terraform-init
	$(call TERRAFORM_FN,$(CMD_TERRAFORM_APPLY),-auto-approve)

terraform-destroy-auto-approve: terraform-init
	$(call TERRAFORM_FN,$(CMD_TERRAFORM_DESTROY),-auto-approve)