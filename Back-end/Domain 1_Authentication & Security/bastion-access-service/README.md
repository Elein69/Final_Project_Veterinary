# 🛡️ Bastion Access Service

Microservice of the distributed veterinary system that simulates controlled access via **Bastion Host / JumpBox**, allowing JWT token validation in a secure environment.

---

## 📌 Features

- ✅ Only activates in `production` environment
- ✅ Validates JWT tokens signed with HS256
- ❌ Does not use a database
- 📦 Dockerized and production-ready
- 🔒 Protected endpoint using `Authorization: Bearer <token>`
- 🧪 Automated unit tests with `pytest`

---

## 🚀 Run Locally

### Activate virtual environment

```powershell
.env\Scripts\Activate.ps1
```

### Run the server

```powershell
$env:APP_ENV="production"
$env:JWT_SECRET="supersecurekey"
python -m app.main
```

Access: [http://localhost:8085/bastion-access](http://localhost:8085/bastion-access)

---

## 🧪 JWT Test Token

Generate a valid token for testing:

```python
import jwt
jwt.encode({"sub": "steven", "exp": 9999999999}, "supersecurekey", algorithm="HS256")
```

---

## 🧪 Run Tests

```bash
$env:PYTHONPATH="."
pytest
```

---

## 🐳 Docker

### Dockerfile

```dockerfile
FROM python:3.11-slim

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

ENV APP_ENV=production
ENV JWT_SECRET=supersecurekey

CMD ["python", "-m", "app.main"]
```

### Build Docker image

```bash
docker build -t bastion-access-service .
```

---

## 📁 Project Structure

```
bastion-access-service/
├── app/
│   ├── main.py
│   ├── auth.py
│   └── __init__.py
├── tests/
│   └── test_main.py
├── .env
├── Dockerfile
├── requirements.txt
├── README.md
└── venv/
```
