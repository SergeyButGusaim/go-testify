package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func serveHttpRequest(request *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	count := 4
	city := "moscow"
	url := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest("GET", url, nil)

	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"

	responseRecorder := serveHttpRequest(req)
	actual := string(responseRecorder.Body.Bytes()[:])

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, expected, actual)

}
