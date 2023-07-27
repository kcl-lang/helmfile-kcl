HELM_HOME ?= $(shell helm home)
HELM_3_PLUGINS := $(shell bash -c 'eval $$(helm env); echo $$HELM_PLUGINS')
PKG:= kcl-lang.io/helmfile-kcl

run:
	go run main.go
