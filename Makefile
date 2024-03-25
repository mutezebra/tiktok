DIR := $(shell pwd)

BIN := $(DIR)/bin

.PHONY: idl
idl:
	sh $(DIR)/idl/script/idl.sh


.PHONY: run-%
run-%:
	@go mod tidy
	go build -o $(BIN)/$* $(DIR)/cmd/$*/main.go
	@echo "running..."
	@sh ./config/env.sh
	@$(BIN)/$*


.PHONY: env-up
env-up:
	@sh ./config/env.sh
	@docker-compose -f docker-compose.yaml up -d


.PHONY: env-down
env-down:
	@docker-compose -f docker-compose.yaml down

