from fastapi import APIRouter, UploadFile, File, HTTPException, status
from app.services.image_service import save_profile_image

router = APIRouter()

@router.post("/upload-profile-image")
async def upload_profile_image(file: UploadFile = File(...)):
    try:
        filename = await save_profile_image(file)
        return {"message": "Image uploaded successfully", "filename": filename}
    except Exception as e:
        raise HTTPException(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            detail=str(e)
        )
