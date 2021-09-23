.PHONY: run
run:
	go run ./...

.PHONY: test
test: 
	go test -v ./...


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY : audit
audit: vendor
	@echo 'Formatting Code...'
	go fmt ./...
	@echo 'Vetting Code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running Tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verify'
	go mod tidy
	go mod verify
	@echo 'Vendoring'
	go mod vendor
