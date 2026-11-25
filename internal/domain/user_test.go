package domain

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUser_JSONOmitPasswordHash(t *testing.T) {
	user := User{
		ID:           1,
		Name:         "John Doe",
		Email:        "test@email.com",
		PasswordHash: "super-secret-hash",
		Role:         RoleCustomer,
	}

	b, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("marshal user: %v", err)
	}

	if strings.Contains(string(b), "super-secret-hash") {
		t.Errorf("json output must be not containt super-secret-hash")
	}

	if strings.Contains(string(b), "password_hash") {
		t.Errorf("json output must be not containt password_hash")
	}
}
