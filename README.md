# Block-reporter

This is a demo application which demonstrates:  
- How to bootstrap `go-micro`
- How to read the latest blocks from Ethereum
- How to bootstrap the development environment with `Postgres` and `MongoDB`
- How to use `NATS`
- How to create PUB/SUB event architecture with `go-micro`
- How to save all data to `MongoDB`

### References
https://github.com/Dmdv/provision-local
https://github.com/Dmdv/micro-client  
https://github.com/Dmdv/micro-server  


## GO-MICRO: Missing documentation

##### Prerequisites
```shell
make prerequisites
```
- Install helm
- Install minikube
- Install skaffold (`https://skaffold.dev/docs/quickstart`)
- Install micro

# Micro installation
```shell
go install go-micro.dev/v4/cmd/micro@a40f6e8fae19018ead6778a2887c0daf970a9882
```

## How to build

```shell
make proto update tidy
```

## How to run

```shell
micro run
```

or if you want to use config map

```shell
skaffold dev
```

or if you want to deploy to a local minikube cluster (depending on the current context)

```shell
skaffold run
```

## Known issues

If you run with `micro run` then the service name will be default: `go.micro.client`
To assign the service name, use `micro.Name` option

Any explicit server declarations must go before any other options

```go
	// Create service
	srv := micro.NewService(
		micro.Client(grpc.NewClient()), // must come before any other options
		micro.Name(service),
		micro.Version(version))
	srv.Init()
```

If you define transport as grpc in ENV then we have the issue when service name is not set
In that case you'll have to declare grpc transport explicitly

```yaml
---

apiVersion: v1
kind: ConfigMap
metadata:
  name: micro-server-env
data:
  MICRO_REGISTRY: kubernetes
  MICRO_SERVER: grpc

```

Setting `MICRO_SERVER_NAME` explicitly in config map solves this issue

## The list of known ENV variable:

MICRO_CLIENT - Client for go-micro  
MICRO_CLIENT_REQUEST_TIMEOUT - Sets the client request timeout. e.g 500ms, 5s, 1m. Default: 5s  
MICRO_CLIENT_RETRIES - Sets the client retries. Default: 1  
MICRO_CLIENT_POOL_SIZE - Sets the client connection pool size. Default: 1  
MICRO_CLIENT_POOL_TTL - Sets the client connection pool ttl. e.g 500ms, 5s, 1m. Default: 1m  
MICRO_REGISTER_TTL - Register TTL in seconds  
MICRO_REGISTER_INTERVAL - Register interval in seconds  
MICRO_SERVER - Server for go-micro; rpc  
MICRO_SERVER_NAME - Name of the server. go.micro.srv.example  
MICRO_SERVER_VERSION - Version of the server. 1.1.0  
MICRO_SERVER_ID - Id of the server. Auto-generated if not specified  
MICRO_SERVER_ADDRESS - Bind address for the server. 127.0.0.1:8080  
MICRO_SERVER_ADVERTISE - Used instead of the server_address when registering with discovery. 127.0.0.1:8080  
MICRO_SERVER_METADATA - A list of key-value pairs defining metadata. version=1.0.0  
MICRO_BROKER - Broker for pub/sub. http, nats, rabbitmq  
MICRO_BROKER_ADDRESS - Comma-separated list of broker addresses  
MICRO_DEBUG_PROFILE - Debug profiler for cpu and memory stats  
MICRO_REGISTRY - Registry for discovery. etcd, mdns  
MICRO_REGISTRY_ADDRESS - Comma-separated list of registry addresses  
MICRO_RUNTIME - Runtime for building and running services e.g local, kubernetes  
MICRO_RUNTIME_SOURCE - Runtime source for building and running services e.g github.com/micro/service  
MICRO_SELECTOR - Selector used to pick nodes for querying  
MICRO_STORE - Store used for key-value storage  
MICRO_STORE_ADDRESS - Comma-separated list of store addresses  
MICRO_STORE_DATABASE - Database option for the underlying store  
MICRO_STORE_TABLE - Table option for the underlying store  
MICRO_TRANSPORT - Transport mechanism used; http  
MICRO_TRANSPORT_ADDRESS - Comma-separated list of transport addresses  
MICRO_TRACER - "Tracer for distributed tracing, e.g. memory, jaeger  
MICRO_TRACER_ADDRESS - Comma-separated list of tracer addresses  
MICRO_AUTH - Auth for role based access control, e.g. service  
MICRO_AUTH_ID - Account ID used for client authentication  
MICRO_AUTH_SECRET - Account secret used for client authentication  
MICRO_AUTH_NAMESPACE - Namespace for the services auth account, go.micro  
MICRO_AUTH_PUBLIC_KEY - Public key for JWT auth (base64 encoded PEM)  
MICRO_AUTH_PRIVATE_KEY - Private key for JWT auth (base64 encoded PEM)  
MICRO_CONFIG - The source of the config to be used to get configuration
