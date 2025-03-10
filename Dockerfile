FROM golang:1.24-alpine AS builder

WORKDIR /usr/src/app

# Copy only dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/app ./main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

# Create non-root user
RUN adduser -D appuser
USER appuser

EXPOSE 8080
CMD ["/usr/local/bin/app"]