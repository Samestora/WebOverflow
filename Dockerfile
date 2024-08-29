FROM golang:1.23-bookworm AS base
WORKDIR /usr/src/app
COPY . .
RUN go install "github.com/air-verse/air@latest"
RUN go mod tidy
CMD ["air", "-c", ".air.toml"]
