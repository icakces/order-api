package config

import "testing"

func setAllEnv(t *testing.T) {
	t.Helper()
	t.Setenv("APP_PORT", "8080")
	t.Setenv("DB_HOST", "127.0.0.1")
	t.Setenv("DB_PORT", "3307")
	t.Setenv("DB_NAME", "order_api")
	t.Setenv("DB_USER", "order_user")
	t.Setenv("DB_PASSWORD", "secret")
}

func TestLoad_Success(t *testing.T) {
	setAllEnv(t)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}
	if cfg.DBHost != "127.0.0.1" {
		t.Errorf("DBHost = %q, want %q", cfg.DBHost, "127.0.0.1")
	}
}

func TestLoad_MissingVar(t *testing.T) {
	setAllEnv(t)
	t.Setenv("DB_HOST", "") // kosong dianggap hilang oleh required()

	_, err := Load()
	if err == nil {
		t.Fatal("Load() should return an error when DB_HOST is empty, got nil")
	}
}
