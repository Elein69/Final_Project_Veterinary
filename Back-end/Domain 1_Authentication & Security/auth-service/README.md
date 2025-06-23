# Auth Service

This microservice handles user registration, login, authentication, and authorization for the Veterinary Distributed System project.

---

## 🚀 Technologies Used

- Node.js with Express
- TypeScript
- PostgreSQL (via Docker)
- JWT (JSON Web Tokens)
- BcryptJS (password encryption)
- GitHub Actions (CI)
- Jest + Supertest (functional tests)

---

## 📁 Project Structure

auth-service/
├── src/
│ ├── controllers/
│ ├── middleware/
│ ├── routes/
│ ├── utils/
│ ├── app.ts
│ └── db.ts
├── test/
│ └── auth.test.ts
├── Dockerfile
├── docker-compose.yml
├── tsconfig.json
├── jest.config.js
├── .env
└── README.md


---

## ⚙️ Environment Variables

Create a `.env` file with the following content:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=adminpass
DB_NAME=authdb
JWT_SECRET=your_jwt_secret_key

Running with Docker
Start PostgreSQL using Docker:

docker compose up -d

Running Locally in Development

npm install
npm run dev

The service will run at:
📍 http://localhost:3000/api/users

Functional Tests
Run tests with:

npm test

Tests cover registration, login, JWT validation, and protected routes.

Continuous Integration (CI)
GitHub Actions workflow:

.github/workflows/test-auth.yml

Runs automatically when any changes are made to the auth-service/ directory.


# Auth Service

This is the authentication service for the veterinary system.

✅ CI Integrated with GitHub Actions (:

✅ Triggered update to test GitHub Actions (:
