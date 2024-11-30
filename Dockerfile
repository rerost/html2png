FROM golang:1.23

RUN apt-get update && apt-get install -y chromium

COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o app .

CMD ["./app"]
