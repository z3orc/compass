FROM golang:alpine AS build
WORKDIR /app
RUN apk add --no-cache make cmake 
COPY ./ ./
RUN make build-site

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main /app
COPY --from=build /app/internal/static /app/static
RUN chmod 0755 main
EXPOSE 8080
CMD [ "./main" ]