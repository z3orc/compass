FROM golang:1.19.4-alpine3.17 AS build
WORKDIR /app
RUN apk add --no-cache make cmake 
COPY ./ ./
RUN make build

FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/bin/compass /app
RUN chmod 0755 compass

CMD [ "./compass" ]