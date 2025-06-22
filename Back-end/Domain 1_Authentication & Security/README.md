# Domain 1: Authentication & Security

This domain is responsible for all authentication, authorization, and security-related logic in the distributed veterinary system.

## 🧩 Microservices Included

### 1. Auth Service (`auth-service`)
- **Language**: Node.js
- **Port**: 3000
- **Database**: PostgreSQL
- **Description**: Handles user registration, authentication, and JWT token issuance.

### 2. Login Service (`login-service`)
- **Language**: C# (.NET 8)
- **Port**: 5046
- **Database**: MySQL
- **Description**: Validates user credentials and returns JWT tokens for frontend sessions.

### 3. CORS Service (`cors-service`)
- **Language**: Go
- **Port**: 8083
- **Database**: None
- **Description**: Manages CORS policy headers for frontend applications.

### 4. Role Manager (`role-manager`)
- **Language**: Java (Spring Boot)
- **Port**: 8084
- **Database**: MySQL
- **Description**: Manages user roles and role assignments for access control.

### 5. Bastion Access Service (`bastion-access-service`)
- **Language**: Python (FastAPI)
- **Port**: 8085
- **Database**: None
- **Description**: Simulates Bastion Host access, available only in production, validates incoming JWT tokens.

## 🔐 Security Stack

- 🔑 **JWT** (JSON Web Tokens) are used for service-to-service authentication and user validation.
- 🧰 **Environment-specific activation**: The `bastion-access-service` only runs in `production` mode.
- 🌐 **CORS control** is handled via an independent microservice (`cors-service`) to centralize origin policies.

## 🐳 Docker & Deployment

All microservices are fully dockerized. A unified `docker-compose.yml` file is located in the root of this domain folder.

To start all services:

```bash
docker compose up --build
```

## 🧪 Testing

Each service includes:
- Unit tests
- Local environment variables
- Independent Dockerfiles

Use `pytest`, `xUnit`, `Moq`, `mvn test`, or the appropriate tool per microservice to run tests locally.

## 📂 Folder Structure

```
Domain 1_Authentication & Security/
│
├── auth-service/
├── login-service/
├── login_service.Tests/
├── cors-service/
├── role-manager/
├── bastion-access-service/
├── docker-compose.yml
└── README.md
```

## 📦 Status

✅ All microservices under Domain 1 are complete, tested, and Docker-ready.
