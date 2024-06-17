# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from the latest golang base image
FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /builder

ADD . /builder

## Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/app

######## Start a new stage from scratch #######
FROM alpine:latest

# Add Maintainer Info
LABEL maintainer="Karpov Sergey <sergi.me@yandex.ru>"

RUN apk --no-cache add ca-certificates
RUN apk add tzdata && rm -rf /var/cache/apk/*

WORKDIR /root/

RUN mkdir ./bin
RUN mkdir ./site
COPY ./site ./site

#Copy the Pre-built binary file from the previous stage
COPY --from=builder /builder/bin/app ./bin/

LABEL Name=vab_bot Version=0.1.0

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
ENTRYPOINT ["/root/bin/app"]
