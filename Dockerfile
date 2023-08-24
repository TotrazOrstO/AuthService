FROM golang:1.19-alpine
WORKDIR $GOPATH/src/medods

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

COPY go.mod go.mod medods.rsa ./
COPY ./ .
RUN go get -d ./...
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./medods ./cmd/main.go


CMD ["./medods"]
EXPOSE 8080