FROM golang AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine:3.11

COPY --from=builder /dist/main /

COPY .env.yml /
COPY wait-for-it.sh /

RUN apk update && apk add bash
# Script to wait for the database to be initialized.

# Command to run
CMD ./wait-for-it.sh -t 0 database:3306 -- /main