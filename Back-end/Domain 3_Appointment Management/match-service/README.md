# 🧠 Match Service

The Match Service is responsible for finding the best available veterinarian for a pet based on service type and availability. It uses **Neo4j** as the graph database and is written in **Go** using the **Gin** web framework.

## 🚀 Features

- Exposes a `POST /match` endpoint
- Connects to Neo4j via `bolt://`
- Real-time matchmaking using Cypher queries
- Containerized using Docker

## 🔧 Environment Variables

The service requires the following variables, loaded from `.env`:

```env
NEO4J_URI=bolt://localhost:7687
NEO4J_USER=neo4j
NEO4J_PASSWORD=testpass

Running Locally

go run ./cmd/main.go


Run with Docker
To build and run:

docker build -t match-service .
docker run -p 3002:3002 --env-file .env match-service






