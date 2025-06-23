def validate_image_format(extension: str) -> bool:
    valid_extensions = [".jpg", ".jpeg", ".png", ".gif"]
    return extension.lower() in valid_extensions
