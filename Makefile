build:
	@go build -o bin/wheredidmymoneygo

run: build
	@./bin/wheredidmymoneygo

test:
	@go test -v ./..
