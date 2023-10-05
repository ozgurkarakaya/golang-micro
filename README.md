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