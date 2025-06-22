# CORS Service

**Language**: Go  
**Port**: 8083  
**Database**: Not required

## 📌 Description

Handles CORS headers configuration for frontend requests.

## ⚙️ Environment Variables

- `ALLOWED_ORIGINS`: Comma-separated list of allowed origins
- `PORT`: Application port (default: 8083)

## 🚀 Running Locally

```bash
# Install dependencies and start the service
go run main.go
```

## 🐳 Docker

Make sure your Dockerfile is present and use the global `docker-compose.yml` to start all services:

```bash
docker compose up --build
```

