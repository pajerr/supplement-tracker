run:
	go run main.go	
build:
	go build -o bin/main main.go	
unit-tests:
	go test -v ./...	