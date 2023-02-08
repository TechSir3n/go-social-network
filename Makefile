BINARY_NAME=social-network

build:
    GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
    GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
    GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run: build
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
