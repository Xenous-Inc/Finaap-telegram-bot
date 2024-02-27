run-dev:
	go run cmd/service/main.go -env-mode=development -config-path=environments/config.yaml
.PHONY: clean