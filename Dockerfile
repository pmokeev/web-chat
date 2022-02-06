FROM golang:1.17.1

ENV GOPATH=/
WORKDIR /go/src/github.com/pmokeev/web-chat
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download

RUN go build -o chat-backend ./cmd/main.go

CMD ["./chat-backend"]