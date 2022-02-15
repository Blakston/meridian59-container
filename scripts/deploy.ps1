. ("config.ps1")

Write-Output "deploying meridian59 ..."

scp server.zip ${M59_DEVUSER}@${M59_HOST}:~/

Write-Output "done."
