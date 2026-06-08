package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTask(t *testing.T) {
	// Создаём фейковый HTTP-запрос
	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/1",
		nil,
	)

	// Создаём фейковый ResponseWriter
	// Он запишет в себя статус, headers и body
	rr := httptest.NewRecorder()

	// Вызываем handler вручную
	GetTask(rr, req)

	// Проверяем статус ответа
	assert.Equal(t, http.StatusOK, rr.Code)

	// Проверяем Content-Type
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Проверяем body
	var task Task

	err := json.NewDecoder(rr.Body).Decode(&task)
	require.NoError(t, err)

	assert.Equal(t, 1, task.ID)
	assert.Equal(t, "Learn Go", task.Title)
}