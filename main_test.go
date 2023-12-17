package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func serveHttpRequest(request *http.Request) *httptest.ResponseRecorder {
	city := "moscow"
	totalCount := 4
	url := fmt.Sprintf("/cafe?city=%s&count=%d", city, totalCount)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	want := http.StatusOK
	have :=

		assert.Equal(t, want, responseRecorder.Code)
}

func TestMainHandlerWhenCityCorrect(t *testing.T) {

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

}
