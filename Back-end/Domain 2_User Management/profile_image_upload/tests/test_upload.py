import requests

def test_upload_image():
    url = "http://localhost:8087/api/upload-profile-image"
    file_path = "tests/test_image.jpeg"  # Imagen de prueba que tú decidas usar
    with open(file_path, "rb") as f:
        files = {"file": ("test_image.jpeg", f, "image/jpeg")}
        response = requests.post(url, files=files)

    assert response.status_code == 200
    json_data = response.json()
    assert "message" in json_data
    assert json_data["message"] == "Image uploaded successfully"
    assert "filename" in json_data
