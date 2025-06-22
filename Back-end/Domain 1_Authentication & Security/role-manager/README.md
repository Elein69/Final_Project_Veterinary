# Role Manager Service

**Language**: Java (Spring Boot)  
**Port**: 8084  
**Database**: MySQL

## 📌 Description

Manages user roles and their assignments.

## ⚙️ Environment Variables

- `SPRING_DATASOURCE_URL`: MySQL connection string
- `SPRING_DATASOURCE_USERNAME`: MySQL username
- `SPRING_DATASOURCE_PASSWORD`: MySQL password

## 🚀 Running Locally

```bash
# Install dependencies and start the service
mvn spring-boot:run
```

## 🐳 Docker

Make sure your Dockerfile is present and use the global `docker-compose.yml` to start all services:

```bash
docker compose up --build
```

