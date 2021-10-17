FROM  alpine:latest
WORKDIR  /opt
ADD ./http_server ./http_server
EXPOSE 8000
ENTRYPOINT  ["./http_server"]
