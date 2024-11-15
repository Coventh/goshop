# Go Shop

An example of gin contains many useful features for e-commerce websites

## How to run

### Required Environment

- Postgres
- Redis

### Config

- Copy config file: `cp pkg/config/config.sample.yaml pkg/config/config.yaml`
- You should modify `pkg/config/config.yaml`

```yaml
environment: production
http_port: 8888
grpc_port: 8889
auth_secret: ######
database_uri: postgres://username:password@host:5432/database
redis_uri: localhost:6379
redis_password:
redis_db: 0
```

### Run

```shell script
$ go run cmd/api/main.go
```

```
2023-09-12T15:18:36.684+0700    INFO    http/server.go:58       HTTP server is listening on PORT: 8888
2023-09-12T15:18:36.684+0700    INFO    grpc/server.go:53       GRPC server is listening on PORT: 8889
```

### Test

```shell script
$ go test
```

### Test with Coverage

```shell script
go test -timeout 9000s -a -v -coverprofile=coverage.out -coverpkg=./... ./...
```

**or**

```shell script
make unittest
```

### Document

- API document at: `http://localhost:8888/swagger/index.html`

### Tech stack

- Restful API
- GRPC
- DDD
- Gorm
- Swagger
- Logging
- Jwt-Go
- Gin-gonic
- Redis

### What's next?

- gRPC functions for products and orders
- Push message to notify place order successfully
- Payment with PayPal
- Define error response wrapper

### Donate:

- Ethereum: `0x96b2F7E87786Db12121D66477e87539619aa2c97`
- BNB Chain: `0x96b2F7E87786Db12121D66477e87539619aa2c97`
