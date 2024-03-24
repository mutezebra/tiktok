#!/bin/bash
# format go imports style
go install golang.org/x/tools/cmd/goimports@latest
goimports  -local github.com/Mutezebra -w .

# format go.mod
go mod tidy