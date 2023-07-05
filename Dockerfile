FROM golang:1.20.5-alpine3.18 as build-env

# Copy the source from the current directory to the Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o agency main.go

FROM build-env as test
RUN CGO_ENABLED=0 go test -v ./...

#runtime image
FROM alpine:3.18

COPY --from=build-env /app /app

# for microservice test need bash
RUN apk update && apk add --no-cache bash=5.2.15-r5

CMD ["/app/agency"]
