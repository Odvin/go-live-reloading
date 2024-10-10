FROM golang:latest

WORKDIR /app

RUN go install github.com/air-verse/air@latest
COPY .air.toml ./

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

ENTRYPOINT ["air"]