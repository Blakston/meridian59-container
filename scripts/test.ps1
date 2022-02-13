# create server package for deployment
Write-Output "creating server package ..."
Compress-Archive -DestinationPath "server.zip" -Path client,server,bin
