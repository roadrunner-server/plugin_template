package plugin_template

import (
	"github.com/roadrunner-server/api/v2/plugins/config"
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
)

// plugin name
const name = "my_plugin"

// Plugin structure should have exactly the `Plugin` name to be found by RR
type Plugin struct {
	log *zap.Logger
	cfg *Config
}

// Init will be called only once
func (p *Plugin) Init(cfg config.Configurer, log *zap.Logger) error {
	const op = errors.Op("my_plugin_init")
	// check for the `my_plugin` key in the configuration
	if !cfg.Has(name) {
		// special type of error, which tells RR to disable this plugin
		return errors.E(errors.Disabled)
	}

	// initialize your configuration
	p.cfg = &Config{}
	err := cfg.UnmarshalKey(name, p.cfg)
	if err != nil {
		return errors.E(op, err)
	}
	// after unmarshal, init values which are not set with the default values
	p.cfg.InitDefaults()

	// init logger
	p.log = new(zap.Logger)
	*p.log = *log

	return nil
}

// Serve called after all Init functions will be resolved (for all plugins)
func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)

	/*
		Your logic is here. Pass errCh channel to your function and send errors in it to notify RR that it should stop execution
	*/

	return errCh
}

// Stop will be called after Serve always. It might be called when you stopping RR or if there is an error returned in Serve
func (p *Plugin) Stop() error {
	// free the resources allocated by your plugin
	return nil
}
