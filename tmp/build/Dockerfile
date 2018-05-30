FROM alpine:3.6

RUN adduser -D artillery-operator
USER artillery-operator

ADD tmp/_output/bin/artillery-operator /usr/local/bin/artillery-operator
