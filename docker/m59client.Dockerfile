FROM m59base:latest

RUN mkdir -p /opt/meridian59

ADD bin /opt/meridian59/
ADD client /opt/meridian59/
ADD download /opt/meridian59/
ADD static opt/meridian59/

WORKDIR /opt/meridian59

ENV M59_HOST=m59server
ENV M59_PORT=59595

CMD [ "wine", "bin/publish.exe" ]
