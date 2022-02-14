
$M59_HOST = "5.45.106.192"
$M59_DEVUSER= "andygeiss"

Write-Output "deploying meridian59 ..."

scp server.zip ${M59_DEVUSER}@${M59_HOST}:/opt/meridian59-dev/

Write-Output "done."
