FROM scratch

ADD bin /
ADD client /client
ADD download /download
ADD static /static

ENV M59_HOST=m59server
ENV M59_PORT=59595

CMD [ "./publish.bin" ]
