FROM golang:1.24-alpine

WORKDIR /app

# Install git & curl
RUN apk add --no-cache git curl

# Set GOPROXY biar cepat
RUN go env -w GOPROXY=https://goproxy.io,direct

# Cache module dependencies before the source tree is mounted at runtime.
COPY go.mod go.sum ./
RUN go mod download

# Install Air
RUN go install github.com/air-verse/air@v1.60.0

# Install gopls di build time (pin versi kompatibel dengan Go 1.24)
RUN go install golang.org/x/tools/gopls@v0.19.1

CMD ["tail", "-f", "/dev/null"]