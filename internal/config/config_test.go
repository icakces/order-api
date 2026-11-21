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
	// TODO 1: kalau err != nil, gagalkan test dengan t.Fatalf
	// TODO 2: kalau cfg.DBHost bukan "127.0.0.1", gagalkan dengan t.Errorf
	//         (tulis pesan yang menyebut nilai yang didapat dan yang diharapkan)
	_ = cfg
	_ = err
}

func TestLoad_MissingVar(t *testing.T) {
	setAllEnv(t)
	t.Setenv("DB_HOST", "") // kosong dianggap hilang oleh required()

	_, err := Load()
	// TODO 3: kalau err == nil, gagalkan test (seharusnya error)
	_ = err
}
