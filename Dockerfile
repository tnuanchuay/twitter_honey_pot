FROM golang:1.15

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go

ENTRYPOINT ["./main"]