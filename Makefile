BUILD:=./cmd/social-network

all: clean build run 

build:
	go build -o $(BUILD) $(BUILD)/
	
run: 
	go run $(BUILD)/main.go


.PHONY: compose
compose: compose-down
	docker-compose up -d


.PHONY: compose-down
compose-down:
	docker-compose down


.PHONY: docker-build
docker-build:
	docker build -t social-network .


.PHONY: docker-run 
docker-run:
	docker run social-network 


.PHONY: test
test:
	go test ./...

.PHONY: mod
	go mod download 


.PHONY: clean
clean:
	rm $(BUILD)
