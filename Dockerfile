FROM  golang:1.14.3
WORKDIR /app/
COPY ./ /app/

RUN GOPROXY="https://gocenter.io" GO111MODULE=on go build cmd/main.go

EXPOSE 8080
ENTRYPOINT ["/app/main"]
