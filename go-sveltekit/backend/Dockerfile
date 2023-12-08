# Get Go image from DockerHub.
FROM golang:1.21.1 AS api

WORKDIR /app

COPY go.* .
RUN go mod download
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o strukt-backend ./cmd/main.go

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

WORKDIR /production
COPY --from=api /app/strukt-backend .

EXPOSE 8080
CMD ["./strukt-backend"]