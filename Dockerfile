FROM golang:1.19-alpine
WORKDIR /go/src

COPY . .

RUN go mod tidy
CMD [ "go", "run", "main.go" ]