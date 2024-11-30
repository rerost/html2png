FROM golang:1.23

RUN apt-get update && apt-get install -y chromium

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o app .

ENTRYPOINT ["/app"]
