package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/bootdotdev/learn-cicd-starter/internal/database"
)

func TestGetApiAuth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "ApiKey 1234")
	got, err := auth.GetAPIKey(req.Header)
	want := "1234"
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

}

func TestDatabaseUserToUser(t *testing.T) {
	t.Run("Testing if we get correct JSON", func(t *testing.T) {
		// 1. Setup mock database.User
		mockUser := database.User{
			ID:        "aaecf8f6-eaf6-4a26-9979-25e796399923",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
			Name:      "mockUser",
			ApiKey:    "123123123123123123",
		}

		// 2. Call databaseUserToUser
		result, err := databaseUserToUser(mockUser)

		// 3. Assert the outcome (consider both result and error)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		userJSON, err := json.Marshal(result)
		if err != nil {
			t.Errorf("Couldn't marshal user to json: %v", err)
		}

		var backToStructUser User
		err = json.Unmarshal(userJSON, &backToStructUser)
		if err != nil {
			t.Errorf("Couldn't unmarshal user from json: %v", err)
		}
		if !reflect.DeepEqual(backToStructUser, result) {
			t.Errorf("JSON doesn't match. Expected: %v, got: %v", backToStructUser, result)
		}
	})
}
