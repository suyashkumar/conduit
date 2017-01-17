package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/suyashkumar/conduit/server/routes"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setupTests()
	retCode := m.Run()
	os.Exit(retCode)
}

func setupTests() {
	os.Setenv("DEV", "TRUE") // Set dev environment
	//TODO: Bootstrap dev database?
}

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/me", nil)
	w := httptest.NewRecorder()
	sampleEmail := "test@suyash.io"
	samplePrefix := "myPrefix"
	hc := routes.HomeAutoClaims{Email: sampleEmail, Prefix: samplePrefix}
	routes.GetUser(w, req, nil, &hc)
	fmt.Println(w.Body.String())
	var user routes.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, sampleEmail, "Test emails are equal")
}
