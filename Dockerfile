FROM golang:1.21-alpine3.18 AS build
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY ./ ./
RUN apk add --no-cache make cmake git
RUN make build

# Config container
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/bin/compass /app
RUN chmod 0755 compass

CMD [ "./compass" ]
