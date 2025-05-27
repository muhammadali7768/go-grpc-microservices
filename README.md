[E-commerce Microservices App]
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)
![Elasticsearch](https://img.shields.io/badge/Elasticsearch-005571?logo=elasticsearch&logoColor=white)
![GraphQL](https://img.shields.io/badge/GraphQL-E10098?logo=graphql&logoColor=white)

Built by following a [YouTube E-commerce Microservices Course](https://www.youtube.com/watch?v=5UIh1dV7aZ8), this project demonstrates a distributed microservices architecture for an e-commerce platform with the following core components:

- **Account Service** – user registration, authentication, and profile management  
- **Catalog Service** – product listing and search using **Elasticsearch**  
- **Order Service** – order creation, tracking, and status updates  
- **GraphQL API Gateway** – a unified GraphQL interface that aggregates all services

Each service runs independently with its own database:
- **PostgreSQL** is used by the **Account** and **Order** services  
- **Elasticsearch** powers the **Catalog** service

Designed for scalability and modularity, this project reflects real-world practices for cloud-native applications and microservice communication patterns.
