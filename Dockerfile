# Build Stage
FROM golang:1.11-alpine3.9 AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git && go get github.com/golang/dep/cmd/dep

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/app/
WORKDIR /go/src/app/
# Install library dependencies
RUN dep ensure -vendor-only

# Build the application
COPY . /go/src/app
RUN go build -o bin/app

# Final Image
FROM alpine:3.9

COPY --from=build /go/src/app/bin /app/

WORKDIR /app

CMD ["sh", "-c", "./app"]
