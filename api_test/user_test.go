package test

import (
	"encoding/json"
	"golang-web-testing/config"
	"golang-web-testing/services/user"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	t.Run("users/register SUCCESS", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"name": "userregister1",
			"email": "userregister1@mail.com",
			"password": "userregister1"
		}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/users", requestBody)
		routes.ServeHTTP(w, req)

		response := w.Result()
		assert.Equal(t, http.StatusCreated, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, "success register", responseBody["message"])
		assert.Equal(t, "userregister1@mail.com", responseBody["user"].(map[string]interface{})["email"])
		assert.Equal(t, "userregister1", responseBody["user"].(map[string]interface{})["name"])
	})

	t.Run("users/register FAILED EMAIL IS NOT UNIQUE", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"name": "userregister1",
			"email": "userregister1@mail.com",
			"password": "userregister1"
		}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/users", requestBody)
		routes.ServeHTTP(w, req)

		response := w.Result()
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, "email already in use", responseBody["error"])
	})
}

func TestUserLogin(t *testing.T) {
	registeredUser := user.User{
		Email:    "loginuser@mail.com",
		Name:     "loginuser",
		Password: "loginuser",
	}
	if err := config.DB.Create(&registeredUser).Error; err != nil {
		t.Fatal("failed to generate registered user for login test")
	}

	t.Run("users/login SUCCESS", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"email": "loginuser@mail.com",
			"password": "loginuser"
		}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/users/login", requestBody)
		routes.ServeHTTP(w, req)

		response := w.Result()
		assert.Equal(t, http.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, "login success", responseBody["message"])
		assert.Equal(t, "loginuser@mail.com", responseBody["user"].(map[string]interface{})["email"])
		assert.Equal(t, "loginuser", responseBody["user"].(map[string]interface{})["name"])
	})

	t.Run("users/login FAILED email not found", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"email": "notfounduser@mail.com",
			"password": "loginuser"
		}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/users/login", requestBody)
		routes.ServeHTTP(w, req)

		response := w.Result()
		assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, "email/password invalid", responseBody["error"])
	})

	t.Run("users/login FAILED password not match", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"email": "loginuser@mail.com",
			"password": "wrongpassword"
		}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/users/login", requestBody)
		routes.ServeHTTP(w, req)

		response := w.Result()
		assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, "email/password invalid", responseBody["error"])
	})
}
