# pull official base image
FROM golang:1.17.1

# set working directory
WORKDIR /go/src/github.com/pmokeev/web-chat

# add golang to PATH
ENV GOPATH=/

# install app dependencies
COPY ./ ./
RUN apt-get update
RUN apt-get -y install postgresql-client
RUN go mod download
RUN go build -o chat-backend ./cmd/main.go

# start app
CMD ["./chat-backend"]