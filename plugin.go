package plugin_template

import (
	"time"

	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
)

// plugin name
const name = "my_plugin"

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	//
	// func (h *HttpService) Init(cp config.Configurer) error {
	//     h.config := &HttpConfig{}
	//     if err := configProvider.UnmarshalKey("http", h.config); err != nil {
	//         return err
	//     }
	// }
	UnmarshalKey(name string, out interface{}) error

	// Unmarshal unmarshal the config into a Struct. Make sure that the tags
	// on the fields of the structure are properly set.
	Unmarshal(out interface{}) error

	// Get used to get config section
	Get(name string) interface{}

	// Overwrite used to overwrite particular values in the unmarshalled config
	Overwrite(values map[string]interface{}) error

	// Has checks if config section exists.
	Has(name string) bool

	// GracefulTimeout represents timeout for all servers registered in the endure
	GracefulTimeout() time.Duration

	// RRVersion returns running RR version
	RRVersion() string
}

// Plugin structure should have exactly the `Plugin` name to be found by RR
type Plugin struct {
	log *zap.Logger
	cfg *Config
}

// Init will be called only once
func (p *Plugin) Init(cfg Configurer, log *zap.Logger) error {
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
