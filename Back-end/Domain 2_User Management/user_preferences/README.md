# User Preferences Microservice

This microservice stores user interface preferences (theme, language, layout) for the Veterinary Distributed System.

## 📦 Download and Run

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/distributed-system-v2.git
```

### 2. Navigate to the microservice folder

```bash
cd distributed-system-v2/Domain\ 2_User\ Management/user_preferences
```

### 3. Build the project

```bash
mvn clean package
```

### 4. Run with Docker Compose

```bash
docker-compose up --build
```

The application will be available at:

- API: `http://localhost:8086/preferences`
- Database: PostgreSQL at `localhost:5435`
