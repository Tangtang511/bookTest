package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBookRouter(t *testing.T) {
	var engine = gin.Default()
	BookRouter(engine)

	var book1 = Book{
		Name:      "test",
		Timestamp: 20220331,
	}
	var err error
	var data []byte
	if data, err = json.Marshal(book1); err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(data))
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	req, _ = http.NewRequest(http.MethodGet, "/showBooks", nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var book2 = Book{
		Name: "《小狗钱钱》",
	}
	if data, err = json.Marshal(book2); err != nil {
		t.Error(err)
	}
	req, _ = http.NewRequest(http.MethodPost, "/selectBook", bytes.NewReader(data))
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var book3 = Book{
		Name: "《不存在的书》",
	}
	if data, err = json.Marshal(book3); err != nil {
		t.Error(err)
	}
	req, _ = http.NewRequest(http.MethodPost, "/selectBook", bytes.NewReader(data))
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
