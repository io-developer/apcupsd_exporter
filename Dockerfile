FROM alpine:3.8

RUN apk add --no-cache \
    libc6-compat

ADD bin/apcupsd_exporter /apcupsd_exporter

EXPOSE 9162
CMD ["/apcupsd_exporter"]
