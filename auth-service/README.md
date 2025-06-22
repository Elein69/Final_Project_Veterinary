# Auth Service

**Language**: Node.js  
**Port**: 3000  
**Database**: PostgreSQL

## 📌 Description

Handles user registration, authentication and JWT token generation.

## ⚙️ Environment Variables

- `PORT`: Application port (default: 3000)
- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port (default: 5432)
- `DB_USER`: PostgreSQL username
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name
- `JWT_SECRET`: Secret key for JWT token signing

## 🚀 Running Locally

```bash
# Install dependencies and start the service
npm install && npm run dev
```

## 🐳 Docker

Make sure your Dockerfile is present and use the global `docker-compose.yml` to start all services:

```bash
docker compose up --build
```

