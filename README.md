# Koal

Koal is a personal goal management and time tracker software. It's approach to task and time management is different than other tools, not too much sophisticated time tracker app and not a useless todo app.

## Backend and Frontend code

Both backends

## Backend Technical Documentation

### Architecture

I followed a modular monolith design approach with Koal, it consists of services with a boundary, no service accesses internal components of other service.

#### API

Each service defines it's API in `/api` directory with `Protobuf` and implement the API servicer in their module, these definitions will create a REST API using grpc-gateway project.

The Gateway is created within the main component and services are given to it as an input.

#### Modules

Each module implements a service API that is defined, and they only implement their own API without any dependency on other modules.

Modules all follow the same architectural pattern. They are using DDD, and this consists of three components, with clear boundaries.

1. API: This layer is responsible to implement the API get the inputs and feed them to Domain layer and return the output. it depends on business layer and infrastructure layer. Tests in this layer focus on input and output given that infrastructure and domain layer are correct.

2. Domain: This layer handles the business logic and domain entities no dependencies on API or infrastructure. Tests in this layer have to test all the possible cases for business rules.

3. Infrastructure: This layer contain components like database, these components are implementing interfaces that are known to Domain. This layer depends on Domain layer. Infrastructure layer implements interfaces so it's easy to swap the implementation so in API tests you can use a completely mocked database implementation that is in memory and fast.

### Components

1. Auth: Implements auth API
2. Todo: Implements the todo features

### Error handling

Errors are monitored with Sentry. If an error cannot be handled it will be annotated and returned, the API layer
will get these errors and return appropriate response.
There also a grpc interceptor which sends these errors to sentry along with the contextual information.

## Development

### Running locally

To run the server install dependencies and run the project

```bash
        go run server/cmd/main.go
```

Then you can interact with API using the Swagger UI at the address `/api-docs/swagger`

### Build

TODO: (glyphack): add complete docs for generating Protobuf messages and
Install buf, [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway#installation)
