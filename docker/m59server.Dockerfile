FROM m59base:latest

ADD server /server
ADD docker/entry.sh /server/

WORKDIR /server

ENTRYPOINT [ "/server/entry.sh" ]
