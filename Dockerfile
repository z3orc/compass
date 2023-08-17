FROM golang:1.20-alpine3.17 AS build
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY ./ ./
RUN apk add --no-cache make cmake git
RUN make build

# Config container
FROM alpine:3.17
WORKDIR /app
RUN apk add --update redis
COPY --from=build /app/bin/compass /app
RUN chmod 0755 compass

CMD [ "./compass" ]