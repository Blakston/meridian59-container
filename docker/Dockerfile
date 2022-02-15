FROM suchja/wine

ADD server /server

USER root

RUN apt install -y xvfb

RUN chown -R xclient /server

USER xclient

ENV DISPLAY=:0.0

WORKDIR /server

ENTRYPOINT [ "/bin/bash" ]

CMD [ "wine", "blakserv.exe" ]
