FROM  golang:1.14.3
WORKDIR /app/
COPY ./ /app/

RUN go build cmd/main.go

EXPOSE 8080
ENTRYPOINT ["/app/main"]