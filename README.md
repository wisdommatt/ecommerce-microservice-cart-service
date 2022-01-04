# Ecommerce Cart Service

This is the GRPC cart service that handles all activities in the ecommerce application that interacts with cart (adding to cart, checkout etc).

## Applications

* ## [Jaeger](https://www.jaegertracing.io/)

  * Jaeger is an **open source software for tracing transactions between distributed services**.
* ## [PostgresQL](https://www.mongodb.com/)

  * PostgresQL is a free and open-source relational database management system emphasizing extensibility and SQL compliance.
  * The cart service stores user's cart in PostgresQL database.

### Usage

To install / run the user microservice run the command below:

```bash
docker-compose up
```

## Requirements

The application requires the following:

* Go (v 1.5+)
* Docker (v3+)
* Docker Compose

## Other Micro-Services / Resources

* #### [Product Service](https://github.com/wisdommatt/ecommerce-microservice-product-service)
* #### [Notification Service](https://github.com/wisdommatt/ecommerce-microservice-notification-service)
* #### [User Service](https://github.com/wisdommatt/ecommerce-microservice-user-service)
* #### [Shared](https://github.com/wisdommatt/ecommerce-microservice-shared)

## Public API

The public graphql API that interacts with the microservices internally can be found in [https://github.com/wisdommatt/ecommerce-microservice-public-api](https://github.com/wisdommatt/ecommerce-microservice-public-api).
