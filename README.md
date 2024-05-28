# advanced_final


## SERVICE 2
### Service 2: Statefull сервис, который соответствует спецификации Swagger в api/api.yml
#### generate code by swagger
```
swagger generate server -f api/api.yml  --exclude-main --server-package=internal/handler -A "service2"
```