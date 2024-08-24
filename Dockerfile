FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o book-library ./cmd/book-library

FROM alpine:latest  
WORKDIR /root/
COPY --from=build /app/book-library .

# Set environment variables
ENV PORT=8080
ENV DB_USER=book-bank
ENV DB_NAME=bookbank
ENV DB_HOST=34.143.145.12
ENV DB_PORT=5432
ENV DB_SSLMODE=disable

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the binary
CMD ["./book-library"]