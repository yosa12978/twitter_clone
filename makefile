mongo:
	-docker run \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=root \
      	-e MONGO_INITDB_ROOT_PASSWORD=root \
		-v twitter-mcs:/data/db \
		--network twitnet \
		--name twitter-mongo \
		-d mongo
	-docker run \
		-p 8081:8081 \
		-e ME_CONFIG_MONGODB_SERVER=twitter-mongo \
		-e ME_CONFIG_MONGODB_PORT=27017\
		-e ME_CONFIG_BASICAUTH_USERNAME=root \
      	-e ME_CONFIG_BASICAUTH_PASSWORD=root \
		-e ME_CONFIG_MONGODB_URL=mongodb://twitter-mogno:27017 \
		--network twitnet \
		--name twitter-mongoexp \
		-d mongo-express

killmongo:
	-docker kill twitter-mongoexp
	-docker kill twitter-mongo

rmongo:
	-docker rm twitter-mongoexp
	-docker rm twitter-mongo

netcreate:
	docker network create -d bridge twitnet