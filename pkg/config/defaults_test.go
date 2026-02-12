package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetDefaultSearchConfig(t *testing.T) {
	cfg := getDefaultSearchConfig()

	if cfg == nil {
		t.Fatal("expected non-nil config")
	}

	// Test default values
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"Enabled", cfg.Enabled, false},
		{"VectorDBPath", cfg.VectorDBPath, "./embeddings.db"},
		{"EmbeddingModel", cfg.EmbeddingModel, "nomic-embed-text"},
		{"MaxResults", cfg.MaxResults, DefaultMaxSearchResults},
		{"MinSimilarityScore", cfg.MinSimilarityScore, DefaultMinSimilarity},
		{"MaxPreviewLength", cfg.MaxPreviewLength, 100},
		{"OllamaURL", cfg.OllamaURL, "http://localhost:11434"},
		{"MaxFileSize", cfg.MaxFileSize, int64(DefaultMaxFileSize)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, tt.got)
			}
		})
	}
}

func TestGetDefaultSearchConfig_IndexExtensions(t *testing.T) {
	cfg := getDefaultSearchConfig()

	expectedExtensions := []string{".go", ".py", ".js", ".md", ".txt", ".yaml", ".json"}

	if len(cfg.IndexExtensions) != len(expectedExtensions) {
		t.Fatalf("expected %d extensions, got %d", len(expectedExtensions), len(cfg.IndexExtensions))
	}

	for i, ext := range expectedExtensions {
		if cfg.IndexExtensions[i] != ext {
			t.Errorf("extension %d: expected %q, got %q", i, ext, cfg.IndexExtensions[i])
		}
	}
}

func TestGetDefaultSearchConfig_ReturnsNewInstance(t *testing.T) {
	cfg1 := getDefaultSearchConfig()
	cfg2 := getDefaultSearchConfig()

	// Modify cfg1
	cfg1.MaxResults = 999
	cfg1.IndexExtensions[0] = ".modified"

	// cfg2 should be unaffected
	if cfg2.MaxResults == 999 {
		t.Error("modifying one config should not affect another")
	}
	if cfg2.IndexExtensions[0] == ".modified" {
		t.Error("modifying one config's slice should not affect another")
	}
}

// Benchmark
func BenchmarkGetDefaultSearchConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDefaultSearchConfig()
	}
}

// TestSetViperDefaults tests that viper defaults are set correctly
func TestSetViperDefaults(t *testing.T) {
	viper.Reset()

	SetViperDefaults()

	// Check repository defaults
	if viper.GetString("repository.root") != "." {
		t.Errorf("repository.root = %q, want '.'", viper.GetString("repository.root"))
	}

	excludedPaths := viper.GetStringSlice("repository.excluded_paths")
	if len(excludedPaths) == 0 {
		t.Error("repository.excluded_paths should not be empty")
	}

	// Check open command defaults
	if !viper.GetBool("commands.open.enabled") {
		t.Error("commands.open.enabled should be true by default")
	}

	if viper.GetInt("commands.open.max_file_size") != DefaultMaxFileSize {
		t.Errorf("commands.open.max_file_size = %d, want 1048576", viper.GetInt("commands.open.max_file_size"))
	}

	// Check write command defaults
	if !viper.GetBool("commands.write.enabled") {
		t.Error("commands.write.enabled should be true by default")
	}

	if viper.GetInt("commands.write.max_file_size") != DefaultMaxWriteSize {
		t.Errorf("commands.write.max_file_size = %d, want 102400", viper.GetInt("commands.write.max_file_size"))
	}

	if !viper.GetBool("commands.write.backup_before_write") {
		t.Error("commands.write.backup_before_write should be true by default")
	}

	// Check exec command defaults
	if viper.GetBool("commands.exec.enabled") {
		t.Error("commands.exec.enabled should be false by default")
	}

	if viper.GetInt("commands.exec.timeout_seconds") != int(DefaultExecTimeout.Seconds()) {
		t.Errorf("commands.exec.timeout_seconds = %d, want 30", viper.GetInt("commands.exec.timeout_seconds"))
	}
}

// TestLoadSearchConfig_Defaults tests LoadSearchConfig with default values
func TestLoadSearchConfig_Defaults(t *testing.T) {
	viper.Reset()

	cfg := LoadSearchConfig()

	if cfg == nil {
		t.Fatal("LoadSearchConfig() returned nil")
	}

	// Should return default config when no viper values are set
	defaultCfg := getDefaultSearchConfig()

	if cfg.Enabled != defaultCfg.Enabled {
		t.Errorf("Enabled = %v, want %v", cfg.Enabled, defaultCfg.Enabled)
	}

	if cfg.MaxResults != defaultCfg.MaxResults {
		t.Errorf("MaxResults = %d, want %d", cfg.MaxResults, defaultCfg.MaxResults)
	}
}

// TestLoadSearchConfig_CustomValues tests LoadSearchConfig with custom viper values
func TestLoadSearchConfig_CustomValues(t *testing.T) {
	viper.Reset()

	// Set custom values
	viper.Set("commands.search.enabled", true)
	viper.Set("commands.search.max_results", 25)
	viper.Set("commands.search.min_similarity_score", 0.8)
	viper.Set("commands.search.ollama_url", "http://custom:11434")

	cfg := LoadSearchConfig()

	if !cfg.Enabled {
		t.Error("Enabled should be true")
	}

	if cfg.MaxResults != 25 {
		t.Errorf("MaxResults = %d, want 25", cfg.MaxResults)
	}

	if cfg.MinSimilarityScore != 0.8 {
		t.Errorf("MinSimilarityScore = %f, want 0.8", cfg.MinSimilarityScore)
	}

	if cfg.OllamaURL != "http://custom:11434" {
		t.Errorf("OllamaURL = %q, want 'http://custom:11434'", cfg.OllamaURL)
	}
}

// TestLoadSearchConfig_PartialOverride tests partial config override
func TestLoadSearchConfig_PartialOverride(t *testing.T) {
	viper.Reset()

	// Override only some values
	viper.Set("commands.search.max_results", 50)

	cfg := LoadSearchConfig()
	defaultCfg := getDefaultSearchConfig()

	// Overridden values
	if cfg.MaxResults != 50 {
		t.Errorf("MaxResults = %d, want 50", cfg.MaxResults)
	}

	// Non-overridden values should still be defaults
	if cfg.Enabled != defaultCfg.Enabled {
		t.Error("Non-overridden Enabled should match default")
	}

	if cfg.EmbeddingModel != defaultCfg.EmbeddingModel {
		t.Error("Non-overridden EmbeddingModel should match default")
	}
}

// TestLoadSearchConfig_IndexExtensions tests loading custom index extensions
func TestLoadSearchConfig_IndexExtensions(t *testing.T) {
	viper.Reset()

	customExtensions := []string{".rs", ".cpp", ".java"}
	viper.Set("commands.search.index_extensions", customExtensions)

	cfg := LoadSearchConfig()

	if len(cfg.IndexExtensions) != 3 {
		t.Errorf("IndexExtensions length = %d, want 3", len(cfg.IndexExtensions))
	}

	for i, ext := range customExtensions {
		if cfg.IndexExtensions[i] != ext {
			t.Errorf("IndexExtensions[%d] = %q, want %q", i, cfg.IndexExtensions[i], ext)
		}
	}
}
