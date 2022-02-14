
# create server package for deployment
Write-Output "creating server package ..."
Compress-Archive -DestinationPath "server.zip" -Force -Path bin,client,docker,download,static,server
