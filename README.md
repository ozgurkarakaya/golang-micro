# go-micro

### RUN via Make
```
    make up_build  //builds broker docker image
    make up        //runs broker container
    make start     //starts frontend
    make stop      //stops frontend
    make down      //downs broker container
```

### Protobuf
```
$brew install protobuf
```
```
//at logger-service/logs folder
//at broker-service/logs folder
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto

//at logger-service
//at broker-service
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

### Docker Commands
````
// logger service: at logger service folder
docker build -f Dockerfile -t ozgurrkarakaya/gm-logger-service:1.0.0 .
docker push ozgurrkarakaya/gm-logger-service:1.0.0

//broker service : at broker service folder
docker build -f Dockerfile -t ozgurrkarakaya/gm-broker-service:1.0.0 .
docker push ozgurrkarakaya/gm-broker-service:1.0.0

//auth service : at auth service folder
docker build -f Dockerfile -t ozgurrkarakaya/gm-auth-service:1.0.0 .
docker push ozgurrkarakaya/gm-auth-service:1.0.0

//listener service : at listener service folder
docker build -f Dockerfile -t ozgurrkarakaya/gm-listener-service:1.0.0 .
docker push ozgurrkarakaya/gm-listener-service:1.0.0

//mail service : at mail service folder
docker build -f Dockerfile -t ozgurrkarakaya/gm-mail-service:1.0.0 .
docker push ozgurrkarakaya/gm-mail-service:1.0.0
```

###Docker Swarm
```
//at project folder where swarm.yml exists
docker swarm init
docker swarm join --token SWMTKN-1-5hy3y4700iu44ke64uo93osw3hqf5ndwrlynmognbgqrzbt55w-2ea1vsr7pjua47gitig31kfpu 192.168.65.4:2377

//to get the token again
docker swarm join-token worker

//start docker
docker stack deploy -c swarm.yml myapp

//list services and check replica count
docker service ls
```