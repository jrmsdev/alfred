NETNAME='alfrednet'
if ! docker network ls | grep -F ${NETNAME}; then
	docker network create --driver bridge \
		--opt 'com.docker.network.bridge.name=alfrednet0' \
		--subnet '10.0.127.0/24' \
		--gateway '10.0.127.1' ${NETNAME}
fi
