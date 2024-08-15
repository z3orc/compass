FROM golang:1.23-alpine3.20 AS build
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY ./ ./
RUN apk add --no-cache make cmake git
RUN make build

# Config container
FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/build/compass /app
RUN chmod 0755 compass

CMD [ "./compass" ]