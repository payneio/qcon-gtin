FROM nordstrom/baseimage-ubuntu:14.04
MAINTAINER Paul Payne "paul@payne.io"

RUN apt-get update
RUN apt-get install -qy ca-certificates
ADD dist/ic-api-gtin /bin/ic-api-gtin

ENTRYPOINT ["/bin/ic-api-gtin"]
