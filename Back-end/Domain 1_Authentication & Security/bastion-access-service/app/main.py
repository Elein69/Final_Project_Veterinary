from fastapi import FastAPI, Depends
from app.auth import verify_jwt
import os

app = FastAPI()

@app.get("/bastion-access")
def bastion_entry(user=Depends(verify_jwt)):
    return {"message": "✅ Acceso autorizado vía Bastion Host"}

if __name__ == "__main__":
    import uvicorn
    if os.getenv("APP_ENV") == "production":
        uvicorn.run("app.main:app", host="0.0.0.0", port=8085)
    else:
        print("🛑 Este microservicio solo se activa en producción.")
