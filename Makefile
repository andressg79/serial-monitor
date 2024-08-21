build:
	go build cmd/serial-monitor

lint:
	golangci-lint run

vet:
	go vet ./...

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.txt ./...
