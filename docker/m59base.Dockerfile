FROM ubuntu:latest

RUN apt-get update

RUN dpkg --add-architecture i386
RUN apt-get install -y ca-certificates curl gnupg software-properties-common
RUN curl -L https://dl.winehq.org/wine-builds/winehq.key | apt-key add -
RUN add-apt-repository -y "deb http://dl.winehq.org/wine-builds/ubuntu/ focal main"
RUN apt-get update
RUN apt-get install -y cabextract gnupg2 winbind winehq-stable winetricks xdotool x11-utils xterm xvfb
RUN apt-get clean -y
RUN apt-get autoremove -y

ENV WINEDEBUG=fixme-all
ENV WINEARCH=win32
