FROM python:3.11-slim

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

ENV APP_ENV=production
ENV JWT_SECRET=supersecurekey

CMD ["python", "-m", "app.main"]
