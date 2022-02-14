#!/bin/bash

DEVUSER=andygeiss

# install prerequisite packages
apt install -y apt-transport-https ca-certificates curl software-properties-common

# add key for the official Docker repo
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# add official Docker repo
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"

# ensure installation from the Docker repo
apt-cache policy docker-ce

# install Docker
apt install -y docker-ce
usermod -aG docker ${DEVUSER}
