GREEN_COLOR_START:=\033[32m
RED_COLOR_START:=\033[31m
COLOR_END:=\033[0m

# 一键部署
.PHONY: deploy
deploy:
	@echo "$(RED_COLOR_START) deploy may last a long time, don't worry $(COLOR_END)"
	@echo "$(GREEN_COLOR_START) Common modules are being deployed $(COLOR_END)"
	@sudo sh ./scripts/common.sh

	@$(foreach script, dir gateway interaction relation user video, \
		echo "$(GREEN_COLOR_START) $(shell echo $(script) | tr '[:lower:]' '[:upper:]') modules are being deployed $(COLOR_END)"; \
		sudo sh ./scripts/$(script).sh; \
	)

# 格式化代码，我们使用 gofumpt，是 fmt 的严格超集
.PHONY: fmt
fmt:
	gofumpt -l -w .

# 优化 import 顺序结构
.PHONY: import
import:
	goimports -w -local github.com/mutezebra .

# 检查可能的错误
.PHONY: vet
vet:
	go vet ./...

# 代码格式校验
.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml --tests --allow-parallel-runners --sort-results --show-stats --print-resources-usage

# 检查依赖漏洞
.PHONY: vulncheck
vulncheck:
	govulncheck ./...

.PHONY: tidy
tidy:
	go mod tidy

# 一键修正规范并执行代码检查
.PHONY: verify
verify: vet fmt import lint vulncheck tidy

