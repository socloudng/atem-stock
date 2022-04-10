FROM alpine

LABEL MAINTAINER="ada223ada@qq.com"
WORKDIR /atem
COPY ./* ./
ENTRYPOINT ["./main"]