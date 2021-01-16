FROM golang:1.12.0-alpine3.9
RUN mkdir /app
RUN mkdir -p /app/static
ADD go.html /app/static
ADD main.go /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 8080
