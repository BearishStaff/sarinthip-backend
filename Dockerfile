# Stage 1: Build the Go application
FROM golang:1.25-bookworm AS builder

WORKDIR /app

# Copy Go module files to leverage Docker caching for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the static binary
RUN CGO_ENABLED=0 go build -o server ./cmd/api/main.go

# Stage 2: Create a minimal runtime image
FROM scratch

# 3. Copy the CA certs from builder (Required for Supabase/Postgres SSL)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 4. Correct the source path (It was /app/server in the builder)
COPY --from=builder /app/server /server

EXPOSE 8080

# Command to run the executable
CMD ["/server"]
