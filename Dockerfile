FROM golang:alpine AS build
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main /app
COPY --from=build /app/static /app/static
RUN chmod 0755 main
EXPOSE 8080
CMD [ "./main" ]