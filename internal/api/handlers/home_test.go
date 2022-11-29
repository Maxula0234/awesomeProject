package handlers

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_home(t *testing.T) {
	w := httptest.NewRecorder()
	home(w, nil)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	expectedMessage := "Hello! Your request was processed."
	assert.Equal(t, expectedMessage, string(greeting), "Сообщение не корректно")

}
