
$M59_HOST = "5.45.106.192"
$M59_DEVUSER= "andygeiss"

ssh ${M59_DEVUSER}@${M59_HOST} "rm -rf /opt/meridian59-dev/*"
scp server.zip ${M59_DEVUSER}@${M59_HOST}:/opt/meridian59-dev/
ssh ${M59_DEVUSER}@${M59_HOST} "cd /opt/meridian59-dev && unzip server.zip && rm -f server.zip"
