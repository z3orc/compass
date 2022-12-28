FROM golang:1.19.4-alpine3.17 AS build
WORKDIR /app
RUN apk add --no-cache make cmake 
COPY ./ ./
RUN make build-site

FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/main /app
COPY --from=build /app/internal/static /app/static
RUN chmod 0755 main
CMD [ "./main" ]