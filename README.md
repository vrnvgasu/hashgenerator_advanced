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