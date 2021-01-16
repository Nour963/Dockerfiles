FROM golang:1.12.0-alpine3.9
RUN mkdir /app
RUN mkdir -p /app/yes
RUN mkdir -p /app/no
COPY yes.html /app/yes/go.html
COPY no.html /app/no/go.html
ADD main.go /app
WORKDIR /app
RUN apk add git
RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 8080
