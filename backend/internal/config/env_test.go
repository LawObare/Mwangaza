package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDotEnvFileDoesNotOverrideExistingEnv(t *testing.T) {
	t.Setenv("MWANGAZA_TEST_EXISTING", "from-env")
	unsetEnv(t, "MWANGAZA_TEST_NEW")

	path := writeEnvFile(t, `
# comment line
MWANGAZA_TEST_EXISTING=from-file
MWANGAZA_TEST_NEW="from-file"
`)

	if err := loadDotEnvFile(path); err != nil {
		t.Fatalf("loadDotEnvFile: %v", err)
	}

	if got := os.Getenv("MWANGAZA_TEST_EXISTING"); got != "from-env" {
		t.Fatalf("existing environment value was overridden: got %q, want %q", got, "from-env")
	}

	if got := os.Getenv("MWANGAZA_TEST_NEW"); got != "from-file" {
		t.Fatalf("new value was not loaded from .env: got %q, want %q", got, "from-file")
	}
}

func TestLoadDotEnvFilesMissingFileIsNotAnError(t *testing.T) {
	if err := loadDotEnvFiles(filepath.Join(t.TempDir(), "does-not-exist")); err != nil {
		t.Fatalf("loadDotEnvFiles: %v", err)
	}
}

func writeEnvFile(t *testing.T, contents string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), ".env")
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("write env file: %v", err)
	}

	return path
}

func unsetEnv(t *testing.T, key string) {
	t.Helper()

	previous, existed := os.LookupEnv(key)
	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("unset %s: %v", key, err)
	}

	t.Cleanup(func() {
		if existed {
			_ = os.Setenv(key, previous)
			return
		}
		_ = os.Unsetenv(key)
	})
}
