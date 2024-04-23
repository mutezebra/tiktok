DIR := $(shell pwd)

BIN := $(DIR)/bin

.PHONY: idl
idl:
	sh $(DIR)/idl/script/idl.sh


.PHONY: run-%
run-%:
	@go mod tidy
	go build -o $(BIN)/$* $(DIR)/app/$*/cmd/main.go
	@echo "running..."
	@$(BIN)/$*


.PHONY: env-up
env-up:
	@docker compose -f docker-compose.yaml up -d


.PHONY: env-down
env-down:
	@docker compose -f docker-compose.yaml down

