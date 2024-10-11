
.PHONY: generate
generate:
	@echo "Generating code..."
	@go generate -v -x ./main.go
