FROM golang:1.17-alpine as builder

# Go env settings
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV CGO_ENABLED=0
ENV GOPROXY="https://proxy.golang.org,direct"

WORKDIR /app

# Cache modules
COPY go.* .
RUN go mod download

# Build a static binary
COPY . .
RUN go build -o ./bin/gopher-translate ./cmd/gopher-translate/main.go

# Binary only container
FROM scratch
WORKDIR /app
COPY --from=builder /app/configs/gopher_rules.json ./configs/
COPY --from=builder /app/bin/gopher-translate .
CMD ["./gopher-translate", "--port", "8080", "--specfile", "./configs/gopher_rules.json"]
EXPOSE 8080