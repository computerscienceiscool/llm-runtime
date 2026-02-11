package config

import (
	"time"
)

// Config holds the tool configuration
type Config struct {
	RepositoryRoot      string
	MaxFileSize         int64
	MaxWriteSize        int64
	ExcludedPaths       []string
	ExecEnabled         bool
	Interactive         bool
	InputFile           string
	OutputFile          string
	JSONOutput          bool
	Verbose             bool
	RequireConfirmation bool
	BackupBeforeWrite   bool
	AllowedExtensions   []string
	ForceWrite          bool
	ExecWhitelist       []string
	ExecTimeout         time.Duration
	ExecMemoryLimit     string
	ExecCPULimit        int
	ExecContainerImage  string
	ExecNetworkEnabled  bool
	IOContainerImage    string
	IOTimeout           time.Duration
	IOMemoryLimit       string
	IOCPULimit          int
	ContainerPool       PoolConfig
	AuditLogPath        string
}

// PoolConfig holds container pool configuration
type PoolConfig struct {
	Enabled             bool          `yaml:"enabled"`
	Size                int           `yaml:"size"`
	MaxUsesPerContainer int           `yaml:"max_uses_per_container"`
	IdleTimeout         time.Duration `yaml:"idle_timeout"`
	HealthCheckInterval time.Duration `yaml:"health_check_interval"`
	StartupContainers   int           `yaml:"startup_containers"`
}
