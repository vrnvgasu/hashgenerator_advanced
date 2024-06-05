# advanced_final

## SERVICE 1
### Service 1: Stateless сервис, который принимает на вход набор произвольных строк, 
считает от них хэши SHA-3 и возвращает вызывающей стороне
#### generate code by proto
```
mkdir pkg
protoc  api/*.proto --go-grpc_out=pkg --go_out=pkg
```

## SERVICE 2
### Service 2: Statefull сервис, который соответствует спецификации Swagger в api/api.yml
#### generate code by swagger
```
swagger generate server -f api/api.yml  --exclude-main --server-package=internal/handler -A "service2"
```
#### generate code by proto
```
mkdir pkg
protoc  api/*.proto --go-grpc_out=pkg --go_out=pkg
```

-------------------
#### You can set your ENV configs. You can find them in [docker-compose.yml](docker-compose.yml)

-------------------
#### run application
```shell
docker-compose  up
```

#### create hashes
POST localhost:8080/send
```json
[
  "string1","string2","string3"
]
```

#### get hash by id
GET localhost:8080/check?ids=2

-------------------
## Test
### Service1
```shell
cd service1
go generate ./...
go test ./...
```