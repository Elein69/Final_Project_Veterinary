# User Registry Microservice

This microservice is part of Domain 2 (User Management) in the distributed veterinary system. It handles user registration and profile retrieval.

## 🚀 Technologies

- Java 17
- Spring Boot 3.1.x
- PostgreSQL (via Docker)
- Spring Data JPA
- Docker & Docker Compose
- Maven
- Lombok

## 📁 Project Structure

```
user-registry/
├── src/
│   └── main/java/com/veterinary/user_registry/
│       ├── model/User.java
│       ├── repository/UserRepository.java
│       ├── service/UserService.java
│       ├── controller/UserController.java
│       └── UserRegistryApplication.java
├── Dockerfile
├── docker-compose.yml
├── pom.xml
└── README.md
```

## ⚙️ Configuration

**application.properties**
```
spring.datasource.url=jdbc:postgresql://user-registry-db:5432/user_registry_db
spring.datasource.username=postgres
spring.datasource.password=postgres
spring.jpa.hibernate.ddl-auto=update
server.port=8085
```

## 🐳 Docker Compose

**docker-compose.yml**
```
version: "3.8"

services:
  user-registry-db:
    image: postgres:15
    container_name: user-registry-db
    environment:
      POSTGRES_DB: user_registry_db
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    ports:
      - "5435:5432"
    volumes:
      - pgdata_user_registry:/var/lib/postgresql/data
    networks:
      - user-registry-net

  user-registry:
    build: .
    container_name: user-registry
    depends_on:
      - user-registry-db
    ports:
      - "8085:8085"
    networks:
      - user-registry-net

volumes:
  pgdata_user_registry:

networks:
  user-registry-net:
```

## 🔄 Build & Run

```
mvn clean package -DskipTests
docker-compose up --build
```

## 🔬 Endpoints

| Method | URL                    | Description           |
|--------|------------------------|-----------------------|
| POST   | /api/users             | Register new user     |
| GET    | /api/users             | Get all users         |
| GET    | /api/users/{id}        | Get user by ID        |

## 🧪 Example Requests

### Create user
```
POST http://localhost:8085/api/users
Content-Type: application/json

{
  "name": "Steven Alexander",
  "email": "steven@email.com",
  "password": "123456"
}
```

### Get all users
```
GET http://localhost:8085/api/users
```

## ✅ Functional Test

**UserControllerTest.java**
```java
package com.veterinary.user_registry.controller;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.server.LocalServerPort;
import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class UserControllerTest {

    @LocalServerPort
    private int port;

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    public void getAllUsers_returnsOk() {
        var response = restTemplate.getForEntity("http://localhost:" + port + "/api/users", String.class);
        assertThat(response.getStatusCode().value()).isEqualTo(200);
    }
}
```

This test checks that the GET `/api/users` endpoint responds successfully.