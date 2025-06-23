import os
import uuid
from fastapi import UploadFile
from app.utils.file_utils import validate_image_format

UPLOAD_DIR = "uploads"

async def save_profile_image(file: UploadFile) -> str:
    ext = os.path.splitext(file.filename)[1]
    if not validate_image_format(ext):
        raise ValueError("Unsupported image format")

    filename = f"{uuid.uuid4()}{ext}"
    filepath = os.path.join(UPLOAD_DIR, filename)

    with open(filepath, "wb") as buffer:
        buffer.write(await file.read())

    return filename
