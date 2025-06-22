# Login Service

**Language**: C# (.NET 8)  
**Port**: 5046  
**Database**: MySQL

## 📌 Description

Processes login requests and generates JWT tokens.

## ⚙️ Environment Variables

- `ASPNETCORE_ENVIRONMENT`: Execution environment (e.g., Development, Production)
- `JWT_SECRET`: Secret key used for JWT generation

## 🚀 Running Locally

```bash
# Install dependencies and start the service
dotnet run
```

## 🐳 Docker

Make sure your Dockerfile is present and use the global `docker-compose.yml` to start all services:

```bash
docker compose up --build
```

