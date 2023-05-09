package main

import (
	"log"
	"time"

	"github.com/sensu/sensu-go/types"
	"github.com/sensu/sensu-plugin-sdk/sensu"
)

// Config represents the handler plugin config.
type Config struct {
	sensu.PluginConfig
	Wait int
}

const (
	wait = 1
)

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-wait",
			Short:    "Wait",
			Keyspace: "sensu.io/plugins/sensu-wait/config",
		},
	}

	options = []sensu.ConfigOption{
		&sensu.PluginConfigOption[int]{
			Path:      "wait",
			Env:       "SENSU_WAIT",
			Argument:  "wait",
			Shorthand: "-w",
			Default:   int(wait),
			Usage:     "Time to wait in seconds",
			Value:     &plugin.Wait,
		},
	}
)

func main() {
	handler := sensu.NewGoHandler(&plugin.PluginConfig, options, checkArgs, executeHandler)
	handler.Execute()
}

func checkArgs(_ *types.Event) error {
	// nothing really to check, default is 1 second, type is integer
	return nil
}

func executeHandler(event *types.Event) error {
	time.Sleep(time.Duration(plugin.Wait) * time.Second)
	log.Printf("Waiting for %v seconds", plugin.Wait)
	return nil
}
