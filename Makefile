BINARY_NAME=./cmd/social-network

all: clean build run


build:
    GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
    GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
    GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run: 
	./${BINARY_NAME}

dep:
	go mod download

// soon 
test:
	go test

// soon 
bench:
	go test -bench

clean:
	go clean 
	rm ${BINARY_NAME}-darwin 
	rm ${BINARY_NAME}-windows
	rm ${BINARY_NAME}-linux
