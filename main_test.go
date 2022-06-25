package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingAPI(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, gin.H{
		"message": "pong",
	}, response)

}
