# ========================================================== #
#   Stage 1: Build the Go application                        #
# ========================================================== #
FROM golang:1.21 AS builder

WORKDIR /app
COPY main.go .

# Initialize module (to avoid "go.mod" errors)
RUN go mod init calculator && go mod tidy

# Compile the Go program as a static binary
RUN CGO_ENABLED=0 go build -o calculator main.go

# ========================================================== #
# Stage 2: Create a small, secure, and efficient final image #
# ========================================================== #

FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /app/calculator .

# Run the Go application
ENTRYPOINT ["/calculator"]
