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

//frontend : at front-end folder
docker build -f Dockerfile -t ozgurrkarakaya/gm-front-end:1.0.1 .
docker push ozgurrkarakaya/gm-front-end:1.0.1
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

//scale service
docker service scale myapp_listener-service=2
//to scale down
docker service scale myapp_listener-service=1

//update a service with new code, e.g. loggler service, at logger service folder
make build_logger //update binary
docker build -f  Dockerfile -t ozgurrkarakaya/gm-logger-service:1.0.1 .
docker push ozgurrkarakaya/gm-logger-service:1.0.1

//scale service to prevent down time
docker service scale myapp_logger-service=2
//update docker to new version
docker service update --image ozgurrkarakaya/gm-logger-service:1.0.1 myapp_logger-service
//rollback
docker service update --image ozgurrkarakaya/gm-logger-service:1.0.0 myapp_logger-service

//stop docker swarm : scale down individually
docker service scale myapp_broker-service=0
//remove docker swarm project
docker stack rm myapp

//remove swarm from that machine entirely
docker swarm leave
docker swarm leave --force

```

## Web Server
```
docker build -f caddy.dockerfile -t ozgurrkarakaya/gm-micro-caddy:1.0.0 .
docker push ozgurrkarakaya/gm-micro-caddy:1.0.0

//adding backend:80 similar to localhost:80
sudo vi /etc/hosts

docker swarm init
docker stack deploy -c swarm.yml myapp

```

## Kubernetes installment
```
https://minikube.sigs.k8s.io/docs/start/
brew install minikube
brew install kubectl

minikube start --node=2
minikube status
minikube stop
minikube start

minikube dashboard //redirects to k8s dashboard at browser -- ctrl+c to stop

kubectl get pods
kubectl get pods -A

```

## Kubernetes config and run:
```
kubectl create secret docker-registry regcred --docker-server=https://index.docker.io/v1/ --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>
kubectl get secrets
kubectl delete secrets regcred

//created k8s folder and mongo.yml file at project folder
//in project folder run the command
kubectl apply -f k8s
kubectl get pods //mongo pod will be running
minikube dashboard //to check deployments and pods 

kubectl get svc //list services
kubectl get deployments //list deployments

kubectl apply -f k8s/rabbit.yml //run new rabbit.yml file

minikube cache add ozgurrkarakaya/gm-broker-service:1.0.0
kubectl apply -f k8s/broker.yml

```

## Kubernetes troubleshooting:
```
kubectl get pods
kubectl logs broker-service-54ccc98d6b-db5k5
kubectl get deployments
kubectl delete deployments broker-service mongo rabbitmq
kubectl get svc
kubectl delete svc broker-service mongo rabbitmq
```