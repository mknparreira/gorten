# Gorten

Gorten is a project for education purposes developed in Golang, designed to implement a checkout service for an e-commerce marketplace platform. This project aims to provide a practical learning experience by covering essential components of an e-commerce system, including:

- **Infrastructure Setup**: Implementing RabbitMQ for messaging and MongoDB for data storage.
- **Core Resources**: Developing several resources such as User, Company, Category, Product, Order, Payment, and Shopping Cart.
- **Checkout Process**: Handling the finalization of purchases.
- **API Enhancements**: Focusing on security, resilience, and scalability improvements.
- **Monitoring and Documentation**: Implementing monitoring and documenting the API with OpenAPI specifications.
- **Future features**: Planning for additional features and improvements.

The Gorten project aims to demonstrate practical applications of Golang in building robust and scalable e-commerce solutions while offering hands-on experience with modern software practices and technologies.

## Architecture

TBD

## Roadmap

The project will be divided into the following phases:

1. Infrastructure Setup **(done)**
2. Implement RabbitMQ
3. Implement MongoDB
4. User resource
5. Company resource
6. Category resource
7. Products resource
8. Orders and Orders Items resource
9. Payments resource
10. Shopping Cart resource
11. Checkout resource
12. Enhance API Security
13. Enhance API Resilience
14. Enhance API Scalability
15. Monitoring
16. The Open API Documentation
17. Future features

# 1. Infrastructure Setup

## Features

| Package                                                                      | Description                                                                                                                                                                                    |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Implement API Gateway using [Kong](https://konghq.com/products/kong-gateway) |                                                                                                                                                                                                |
| Authentication using JWT Kong Plugin                                         | All APIs will only be accessed by a JWT token                                                                                                                                                  |
| Fall-safe using request-termination Kong Plugin                              | Implements rejection of unauthorized requests to implement a security policy that rejects any request not associated with a defined route. This is done by checking if the route is configured |
| The Response Transformer plugin                                              | Add HTTP Headers to avoid XSS and Clickjacking attacks                                                                                                                                         |
| Docker                                                                       |                                                                                                                                                                                                |

# 2. Implements RabbitMQ

TBD

# 3. Implements MongoDB

TBD

# 4. User resource

TBD

# 5. Company resource

TBD

# 6. Category resource

TBD

# 7. Products resource

TBD

# 8. Orders and Orders Items resource

TBD

# 9. Payments resource

TBD

# 10. Shopping Cart resource

TBD

# 11. Checkout resource

TBD

# 12. Enhance API Security

TBD

# 13. Enhance API Resilience

In this phase, I will implement some strategies to improve the resilience of APIs, ensuring they can handle with failures and recover effectively. I will adopting the following strategies and tools:

- Implement Circuit Breaker using Hystrix-Go package.
- Implement retries e backoff using Retry-Go package.
- Configure timeout into the entire services in Kong API Gateway
- Implement Fallback Pattern using Go-Fallback or Resilience4j package (I’m not sure yet).
- Implement Bulkheads.

# 14. Enhance API Scalability

To ensure API Scalability and can handle increasing demand effectively, I will adopting the following strategies and tools:

- Implement Load Balancing with Kubernetes
- Implement Caching with Redis
- Implement Rate Limiting Using the Kong Plugin

# 15. Monitoring

This phase I will integrating monitoring through the entire system to ensure visibility and performance management. The goal is to provide real-time insights, track system health, and identify potential issues early by monitoring various metrics, logs, and events across all services. This will help in maintaining system reliability and performance.

- Implement monitoring across all services

# 16. The Open API Documentation

In this phase, the OpenAPI documentation will be created for every API within the project. This documentation will serve as a detailed reference with all available endpoints, request and response formats, and authentication methods by using the OpenAPI specifications.

- Provide The OpenAPI Documentation for All APIs

# 17. Future features

| Feature    | Description                 |
| ---------- | --------------------------- |
| Kubernetes | Implements Service Register |
| Jenkins    | Implements CI-CD            |
