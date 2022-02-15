docker rm -f m59client
docker rm -f m59server

docker network rm m59

docker network create m59

docker run \
	-d \
	-e M59_HOST=m59server \
	-e M59_PORT=59595 \
	--name m59client \
	--net m59 \
	--restart=always \
	--volume $PWD/server/channel:/server/channel:ro \
	-p 80:80 \
	m59client

maintenance=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' m59client)

docker run \
	-d \
	-e M59_MASK=$maintenance \
	--name m59server \
	--net m59 \
	--restart=always \
	--volume $PWD/server/channel:/server/channel:rw \
	--volume $PWD/server/savegame:/server/savegame:rw \
	-p 5959:5959 \
	m59server

