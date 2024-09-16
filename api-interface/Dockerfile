# Building the binary of the App
FROM golang:latest AS build

# `boilerplate` should be replaced with your project name
WORKDIR /go/src/api-interface

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as release

WORKDIR /app

# Create the `public` dir and copy all the assets into it
RUN mkdir ./static
COPY ./static ./static

# `boilerplate` should be replaced here as well
COPY --from=build /go/src/api-interface/app .
COPY --from=build /go/src/api-interface/.env .

# Add packages
RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/app

# Add packages
RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/app

# Export environment variables from .env file
RUN export $(grep -v '^#' app/app/.env | xargs)

# Expose the port specified in the .env file
ENV S3_PORT=${S3_PORT:-9000}
EXPOSE ${S3_PORT}

# Add the data volume
VOLUME ["/app/data"]

ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/app"]

# Run the application