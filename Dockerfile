FROM golang:1.21.5-alpine

ENV CGO_ENABLED 0
ENV GOOS linux

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

CMD CompileDaemon --build="go build -o bin/main cmd/main.go" --command=./bin/main