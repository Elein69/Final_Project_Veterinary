package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Elein69/match-service/controllers"
	"github.com/Elein69/match-service/neo4j"
	"github.com/gin-gonic/gin"
)

func TestFindMatch_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// ✅ Inicializa la conexión con Neo4j
	neo4j.InitDriver()

	router := gin.Default()
	router.POST("/match", controllers.FindMatch)

	// Datos de prueba
	bodyData := map[string]string{
		"pet_id":         "abc123",
		"reason":         "vacunación",
		"preferred_day":  "2025-06-22",
		"preferred_time": "morning",
	}
	body, _ := json.Marshal(bodyData)

	req, _ := http.NewRequest("POST", "/match", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["match_found"] != true {
		t.Errorf("Expected match_found true, got %v", response["match_found"])
	}
}
