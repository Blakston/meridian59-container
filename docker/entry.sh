#!/bin/bash

curl 'http://dl.winehq.org/wine/wine-mono/4.8.3/wine-mono-4.8.3.msi' -o wine-mono-4.8.3.msi
wine msiexec /i wine-mono-4.8.3.msi

Xvfb :0 -screen 0 1024x768x16 &
sleep 4

sed -i "s/::ffff:127.0.0.1/::ffff:$M59_MASK/g" blakserv.cfg

DISPLAY=:0.0 wine blakserv.exe
