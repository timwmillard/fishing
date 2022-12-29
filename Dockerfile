FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o fishingd ./cmd/fishingd

# Default API port.
EXPOSE 3000

CMD "/app/fishingd"
