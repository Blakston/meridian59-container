FROM m59base:latest

ENV WINEDEBUG -all

RUN apt-get install -y --no-install-recommends wine wine64
