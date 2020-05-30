FROM ubuntu:20.04

ENTRYPOINT [ "/opt/go-experiments/simple-http-server/simple-http-server" ]
WORKDIR /opt/go-experiments/simple-http-server
EXPOSE 80

RUN mkdir -p /opt/go-experiments/simple-http-server
RUN mkdir -p /opt/go-experiments/simple-http-server/files
COPY simple-http-server /opt/go-experiments/simple-http-server/
COPY files /opt/go-experiments/simple-http-server/files/