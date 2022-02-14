#!/bin/bash

chmod +x bin/*
chmod +x docker/*.sh

docker build -t m59base -f docker/m59base.Dockerfile .
docker build -t m59client -f docker/m59client.Dockerfile .
docker build -t m59server -f docker/m59server.Dockerfile .
