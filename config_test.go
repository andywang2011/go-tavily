package tavily

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Host != DefaultHost {
		t.Errorf("Host = %q, want %q", cfg.Host, DefaultHost)
	}
	if cfg.Timeout <= 0 {
		t.Error("expected positive default timeout")
	}
}

func TestConfig_Validate(t *testing.T) {
	cfg := Config{Secret: "sk"}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() = %v", err)
	}
	if cfg.Host != DefaultHost {
		t.Errorf("Host = %q, want default", cfg.Host)
	}
	if cfg.Timeout <= 0 {
		t.Error("expected default timeout applied")
	}
}

func TestConfig_Validate_MissingSecret(t *testing.T) {
	cfg := Config{Host: "https://a.com"}
	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected error for missing secret")
	}
}

func TestConfig_Validate_CustomHost(t *testing.T) {
	cfg := Config{Host: "https://custom.example.com", Secret: "sk", Timeout: 0}
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
	if cfg.Host != "https://custom.example.com" {
		t.Errorf("Host = %q", cfg.Host)
	}
}
