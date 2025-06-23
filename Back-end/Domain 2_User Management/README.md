# Domain 2: User Management Microservices

This folder contains the complete set of microservices developed for **Domain 2: User Management** of the Veterinary System. Each microservice is containerized with Docker and can be orchestrated together using the provided `docker-compose.yml` file.

## ✅ Microservices Included

| Microservice              | Language | Framework     | Database   | Port  |
|--------------------------|----------|---------------|------------|-------|
| user_registry            | Java     | Spring Boot   | PostgreSQL | 8081  |
| user_preferences         | Java     | Spring Boot   | PostgreSQL | 8082  |
| profile_image_upload     | Python   | FastAPI       | Local FS   | 8083  |
| user_notification_linker | Java     | Spring Boot   | MySQL      | 8084  |
| audit_service            | Go       | gRPC          | DynamoDB   | 8085  |

## 🐳 Docker Compose

Use the following command to build and run all services:

```bash
docker compose up --build
```

Each service will start on the port listed above. Make sure the respective databases are available via Docker volumes or external services, as configured.

## 🧪 Testing

Each microservice includes unit tests configured with its respective testing framework (JUnit, pytest, Go test, etc.). Use Maven, pytest, or Go tools depending on the service language.

## 🚀 Deployment & CI/CD

You can extend this setup with GitHub Actions workflows to automate build and test for each service on every push.

## 📁 Folder Structure

Each folder contains:

- `Dockerfile`
- Source code in `/src`
- Test suite (if applicable)
- Config files (like `application.properties`)

## 👤 Author

Developed by Steven Alexander for the Distributed Systems course.

---

For more technical details on each service, refer to their individual README files.
