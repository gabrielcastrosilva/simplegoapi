FROM alpine:3.13
RUN mkdir /mconf
ADD . /mconf
WORKDIR /mconf