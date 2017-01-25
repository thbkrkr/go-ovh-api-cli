FROM alpine:3.5

RUN apk --no-cache add ca-certificates

COPY /ovhapi /usr/local/bin/ovhapi

ENTRYPOINT ["ovhapi"]
