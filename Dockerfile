FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./

# Download deps
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app ./src

FROM alpine:latest

WORKDIR /app/bin

# Copy the binary from the build stage
COPY --from=build /app/bin/app .

EXPOSE 8080

CMD ["./app"]
