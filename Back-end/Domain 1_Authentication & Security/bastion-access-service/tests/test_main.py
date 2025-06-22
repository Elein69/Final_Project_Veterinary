import os
import jwt
from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)

# Establecer la misma secret que usa la aplicación
os.environ["JWT_SECRET"] = "supersecurekey"

def create_test_token():
    return jwt.encode({"sub": "test-user", "exp": 9999999999}, "supersecurekey", algorithm="HS256")

def test_access_with_valid_token():
    token = create_test_token()
    response = client.get("/bastion-access", headers={"Authorization": f"Bearer {token}"})
    assert response.status_code == 200
    assert response.json() == {"message": "✅ Acceso autorizado vía Bastion Host"}

def test_access_with_invalid_token():
    token = "invalid.token.here"
    response = client.get("/bastion-access", headers={"Authorization": f"Bearer {token}"})
    assert response.status_code == 401

def test_access_without_token():
    response = client.get("/bastion-access")
    assert response.status_code == 403

