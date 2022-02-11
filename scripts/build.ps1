# update server
Remove-Item -Force -Path "server" -Recurse
Copy-Item -Destination "server" -Force -Path "../meridian59-base/run/server" -Recurse
Copy-Item -Destination "config/" -Force -Path "../meridian59-base/run/server/blakserv.cfg"

# update client
Remove-Item -Force -Path "client" -Recurse
Copy-Item -Destination "client/x64" -Force -Path "../meridian59-dotnet/Meridian59.Ogre.Client/bin/x64/Release" -Recurse
Copy-Item -Destination "client/resources" -Force -Path "../meridian59-dotnet/Resources" -Recurse
Copy-Item -Destination "client/" -Force -Path "../meridian59-dotnet/Meridian59.Patcher/bin/AnyCPU/Release/Meridian59.Patcher.exe"
Copy-Item -Destination "client/" -Force -Path "../meridian59-dotnet/Meridian59.Patcher/patchurl.txt"
Copy-Item -Destination "config/" -Force -Path "../meridian59-dotnet/Meridian59.Patcher/patchurl.txt"

# sync client/server strings
Copy-Item -Destination "client/resources/strings/" -Force -Path "server/rsc/rsc0000.rsb"
