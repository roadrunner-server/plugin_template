package plugin_template

// your configuration is here
type Config struct {
	// yaml, sample mapping to your `key` in the .rr.yaml
	keySample string `mapstructure:"key"`
}

// InitDefaults used to initialize default configuration values
func (c *Config) InitDefaults() {
	if c.keySample == "" {
		c.keySample = "some_value"
	}
}
