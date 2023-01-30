FROM golang:latest

WORKDIR /social_network/web/app 

COPY go.mod /social_network/web/app/
COPY go.sum /social_network/web/app/

COPY . /social_network/web/app/

RUN go mod download 
RUN go build -o main main.go

EXPOSE 8080 8080

CMD [ ". /cmd-social-network/main.go" ]