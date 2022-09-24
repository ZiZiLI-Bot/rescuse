FROM golang:1.18-alpine AS builder

# Set HEADER AND ENV FILES
ARG HEADER_FILE
ARG ENV_FILE
ENV HEADER_FILE=header_dev.go
ENV ENV_FILE=.env.dev

RUN apk add bash ca-certificates git gcc g++ libc-dev

# Set working directory for the build
RUN mkdir -p /work/rescues
WORKDIR /work/rescues

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .
RUN ls -la /work/rescues/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# COPY everything else
COPY . /work/rescues/

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rescues .

# EXPOSE 10050

CMD ["./rescues"]