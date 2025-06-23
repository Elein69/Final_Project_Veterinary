# Profile Image Upload Microservice

This microservice allows users to upload profile images locally via a RESTful API using FastAPI.

## 📁 Features
- Uploads image files (`.jpeg`, `.png`, `.jpg`, etc.)
- Saves them to a local directory `uploads/`
- Automatically generates a unique filename for each image
- Provides a JSON response upon successful upload

## 🚀 Technologies Used
- **Python 3.11**
- **FastAPI**
- **Docker**
- **Pytest** for functional tests

## 📦 Installation & Usage

### 🐳 Docker (Recommended)
Make sure you are inside the `profile_image_upload` folder.

Build and run with Docker Compose:

```bash
docker-compose up --build
```

The service will be available at: `http://localhost:8087`

### 📫 Upload an Image (via Postman or curl)
**POST** `/upload-image`  
**Form-Data:** `file` (choose image file)

**Sample Response:**
```json
{
  "message": "Image uploaded successfully",
  "filename": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.jpeg"
}
```

### 🧪 Run Functional Test
```bash
pytest tests/test_upload.py
```

## 📁 Project Structure
```
profile_image_upload/
│
├── app/
│   ├── main.py
│   ├── routes/
│   ├── services/
│   └── utils/
│
├── uploads/                # Folder where images are stored
├── tests/
│   └── test_upload.py      # Functional test using requests + pytest
├── Dockerfile
├── docker-compose.yml
└── requirements.txt
```

## 📌 Notes
- This version stores files **locally**. S3 or remote storage can be integrated later.
- Port exposed: `8087`
- Uploaded files are stored in `uploads/` inside the container

---
© 2025 Veterinary Distributed System