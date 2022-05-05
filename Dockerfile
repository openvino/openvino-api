# -----------------
# Go Build Stage
# -----------------
FROM golang:1.14-buster AS go-builder

# Move to working directory /build
WORKDIR /app

# Copy and download dependency using go mod
COPY wait-for-it.sh .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

# Copy the code into the container
COPY src src
COPY main.go .
COPY router.go .
COPY languages /languages
COPY static static

# Build the application
RUN go build -o main -v

CMD [ "./main" ]
