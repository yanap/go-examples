FROM golang:1.13 AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/SHOWROOM-inc/short-server

# Let's cache modules retrieval - those don't change so often
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code necessary to build the application
# You may want to change this to copy only what you actually need.
COPY . .

# Build the application
RUN go build -o /go/bin/peso

# Create the minimal runtime image for cloud deployment
FROM alpine:3 AS cloud
# import https, timezone
RUN apk add --update --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /go/bin/peso .
ENTRYPOINT [ "./peso" ]