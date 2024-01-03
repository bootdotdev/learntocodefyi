FROM --platform=linux/amd64 alpine:latest

RUN apk --no-cache add ca-certificates

ADD bin /bin/

CMD ["/bin/app"]
