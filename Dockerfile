FROM phusion/baseimage:0.9.16
MAINTAINER Paul Payne "paul@payne.io"

ADD dist/qcon-gtin /bin/qcon-gtin

ENTRYPOINT ["/bin/qcon-gtin"]
