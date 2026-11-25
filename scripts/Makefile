ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/..)
COMPOSE := docker compose -f "$(ROOT)/docker-compose.yaml"
VAGRANT := cd "$(ROOT)/infra" && env VAGRANT_CWD="$(ROOT)/infra" VAGRANT_DOTFILE_PATH="$(ROOT)/infra/.vagrant" vagrant
IMAGE := iot-webapp:v1.0.0

.DEFAULT_GOAL := up
.PHONY: up build app-up vm-up deploy down

up: app-up deploy

build:
	$(COMPOSE) build

app-up: build
	$(COMPOSE) up -d

vm-up:
	$(VAGRANT) up --provider=virtualbox

deploy: build vm-up
	bash -o pipefail -c 'docker image save "$(IMAGE)" | ($(VAGRANT) ssh -c "sudo k3s ctr images import -")'
	$(VAGRANT) ssh -c "sudo k3s kubectl apply -f -" < "$(ROOT)/infra/webapp.yaml"
	$(VAGRANT) ssh -c "sudo k3s kubectl rollout restart deployment/webapp1 && sudo k3s kubectl rollout status deployment/webapp1 --timeout=120s"

down:
	$(COMPOSE) down
	$(VAGRANT) halt
