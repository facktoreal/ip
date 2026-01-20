.PHONY: dev test
PACKAGES=`go list ./...`

local:
	@go run -mod=vendor app/main.go

test:
	@RICHGO_FORCE_COLOR=1 ENV=test richgo test -v -mod=vendor -cover ./...
