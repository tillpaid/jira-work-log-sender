run:
	go run cmd/app/main.go --dev

build:
	go build -o bin/app cmd/app/main.go

test:
	clear
	go test ./internal/...

testf:
	clear
	go test ./internal/... -failfast

build-and-replace:
	go build -o bin/app cmd/app/main.go && sudo mv ./bin/app /usr/local/bin/tt

