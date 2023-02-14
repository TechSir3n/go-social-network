FROM golang:latest AS builder 

WORKDIR /social_network/web/app 

COPY go.mod /social_network/web/app/
COPY go.sum /social_network/web/app/

COPY . /social_network/web/app/

RUN go mod download 
RUN go build -o /main ./cmd/social-network/

EXPOSE 3000 3000

ENTRYPOINT ["/main"]