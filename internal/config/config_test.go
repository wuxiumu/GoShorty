package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("METRICS_ADDR", "")
	t.Setenv("PUBLIC_BASE_URL", "")

	dir := t.TempDir()
	t.Chdir(dir)

	cfg := Load()
	if cfg.HTTPAddr != ":8080" {
		t.Fatalf("HTTPAddr = %q", cfg.HTTPAddr)
	}
	if cfg.MetricsAddr != "127.0.0.1:6060" {
		t.Fatalf("MetricsAddr = %q", cfg.MetricsAddr)
	}
	if cfg.PublicBaseURL != "http://localhost:8080" {
		t.Fatalf("PublicBaseURL = %q", cfg.PublicBaseURL)
	}
	if Current() != cfg {
		t.Fatal("Current() did not match Load()")
	}
}

func TestDotEnvDoesNotOverrideProcessEnv(t *testing.T) {
	dir := t.TempDir()
	contents := "HTTP_ADDR=:9090\nPUBLIC_BASE_URL=http://files.example/\nMETRICS_ADDR=127.0.0.1:16060\n"
	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)

	t.Setenv("HTTP_ADDR", "")
	t.Setenv("METRICS_ADDR", "")
	t.Setenv("PUBLIC_BASE_URL", "")
	fromFile := Load()
	if fromFile.HTTPAddr != ":9090" || fromFile.MetricsAddr != "127.0.0.1:16060" {
		t.Fatalf("file config = %+v", fromFile)
	}
	if fromFile.PublicBaseURL != "http://files.example" {
		t.Fatalf("PublicBaseURL = %q", fromFile.PublicBaseURL)
	}

	t.Setenv("HTTP_ADDR", ":7070")
	overridden := Load()
	if overridden.HTTPAddr != ":7070" {
		t.Fatalf("process env should win, got %q", overridden.HTTPAddr)
	}
	if overridden.MetricsAddr != "127.0.0.1:16060" {
		t.Fatalf("unset metrics should stay on file value, got %q", overridden.MetricsAddr)
	}
}
