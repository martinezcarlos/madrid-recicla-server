# syntax=docker/dockerfile:1
############################
# STEP 1 build optimized executable binary
############################

FROM golang:1.18-alpine as build

# Creates root directory inside the image.
WORKDIR /server

# Copies go.mod and go.sum into the root directory.
COPY src/go.mod ./
COPY src/go.sum ./

# Downloads GO modules into the image.
RUN go mod download

# Copies source code into the root folder.
COPY src ./

# Compiles the application and produce a static application binary.
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /madrid-recicla-server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /madrid-recicla-server

############################
# STEP 2 build a small image
############################

FROM gcr.io/distroless/base

# Creates root directory inside the image.
WORKDIR /

# Copies the static application binary from the previous build.
COPY --from=build /madrid-recicla-server /madrid-recicla-server

# Executes the server from the static application binary.
ENTRYPOINT ["/madrid-recicla-server"]