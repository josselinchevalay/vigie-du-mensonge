FROM golang:1.25.0-alpine

WORKDIR /app

# Cache modules
COPY go.mod go.sum ./
RUN go mod download

# Copy sources (including CSVs) with ownership set to 10001
COPY --chown=10001:10001 . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/data-import .

# Ensure readable files and executable binary
RUN chmod -R a+r /app && chmod a+x /app/data-import

# Run as non-root user matching your k8s securityContext
USER 10001:10001

CMD ["/app/data-import"]