default: help

THIS_MAKEFILE_PATH = $(word 1,$(MAKEFILE_LIST))
PATH_ROOT = $(shell cd $(dir $(THIS_MAKEFILE_PATH));pwd)
PATH_MISC = $(PATH_ROOT)/misc

include $(PATH_MISC)/Variables.mk
include $(PATH_MISC)/Docker.mk
include $(PATH_MISC)/CircleCI.mk

help: help-circleci help-docker