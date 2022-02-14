#!/bin/bash

docker rm -f m59client
docker rm -f m59server

docker network rm m59

docker network create m59

docker run \
    -d \
    --name m59client \
    --net m59 \
    --restart=always \
    -p 80:80 \
    m59client

docker run \
    -d \
    --name m59server \
    --net m59 \
    --restart=always \
    -p 5959:5959 \
    m59server
