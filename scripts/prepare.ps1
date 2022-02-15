. ("config.ps1")

# update server
Write-Output "updating the server ..."
Remove-Item -Force -Path "server" -Recurse
Copy-Item -Destination "server" -Force -Path "../meridian59-base/run/server" -Recurse
Copy-Item -Destination "server/" -Force -Path "../meridian59-base/kod" -Recurse
Copy-Item -Destination "server/" -Force -Path "config/blakserv.cfg"
((Get-Content -path "server/blakserv.cfg" -Raw) -replace '59595',$M59_PORT) | Set-Content -Path "server/blakserv.cfg"
New-Item -Force -ItemType "directory" -Path "server/channel"
New-Item -Force -ItemType "directory" -Path "server/memmap"
New-Item -Force -ItemType "directory" -Path "server/savegame"
Copy-Item -Destination "server/savegame/" -Force -Path "config/savegame/*"

# update client
Write-Output "updating the client ..."
Remove-Item -Force -Path "client" -Recurse
Copy-Item -Destination "client/x64" -Force -Path "../meridian59-dotnet/Meridian59.Ogre.Client/bin/x64/Release" -Recurse
Copy-Item -Destination "client/resources" -Force -Path "../meridian59-dotnet/Resources" -Recurse
Copy-Item -Destination "client/" -Force -Path "../meridian59-dotnet/Meridian59.Patcher/bin/AnyCPU/Release/Meridian59.Patcher.exe"
Copy-Item -Destination "client/" -Force -Path "config/patchurl.txt"
Copy-Item -Destination "client/" -Force -Path "config/configuration.xml"
((Get-Content -path "client/configuration.xml" -Raw) -replace '5.45.106.192',$M59_HOST) | Set-Content -Path "client/configuration.xml"

# sync client/server strings
Write-Output "syncing client and server ..."
Copy-Item -Destination "client/resources/strings/" -Force -Path "server/rsc/rsc0000.rsb"

# create patchinfo.txt
Write-Output "creating patch info ..."
bin/patch.exe client

# create download
Write-Output "creating download ..."
Remove-Item -Force -Path "download" -Recurse
New-Item -Force -ItemType "directory" -Path "download"
Copy-Item -Destination "download/" -Force -Path "client/configuration.xml"
Copy-Item -Destination "download/" -Force -Path "client/patchurl.txt"
Copy-Item -Destination "download/" -Force -Path "client/Meridian59.Patcher.exe"
((Get-Content -path "download/patchurl.txt" -Raw) -replace '5.45.106.192',$M59_HOST) | Set-Content -Path "download/patchurl.txt"
Compress-Archive -DestinationPath "download/patcher.zip" -Force -Path "download"

Write-Output "done."
