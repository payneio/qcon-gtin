FROM phusion/baseimage:0.9.16
MAINTAINER Paul Payne "paul@payne.io"

# RUN apt-get update
# RUN apt-get install -qy ca-certificates
ADD dist/qcon-gtin /bin/qcon-gtin

ENTRYPOINT ["/bin/qcon-gtin"]
