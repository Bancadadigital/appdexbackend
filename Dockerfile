cat > Dockerfile << 'EOF'
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o appdex

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/appdex .
EXPOSE 8081
CMD ["./appdex"]
EOF
