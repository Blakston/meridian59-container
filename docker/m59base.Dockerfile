FROM ubuntu:latest

RUN apt-get update && apt-get upgrade -y

ENV WINEDEBUG -all

RUN apt-get install -y --no-install-recommends wine wine64
