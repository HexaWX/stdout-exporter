package main

import (
	"fmt"

	"github.com/HexaWX/HexaWX/core"
	"github.com/hashicorp/go-plugin"
)

type StdoutExporter struct {
	Prefix string
}

func (e *StdoutExporter) Init(config map[string]string) error {
	// On récupère une option de config, avec une valeur par défaut
	e.Prefix = config["prefix"]
	if e.Prefix == "" {
		e.Prefix = "METEO"
	}
	fmt.Printf("✅ Exporter Stdout initialisé avec le préfixe: %s\n", e.Prefix)
	return nil
}

func (e *StdoutExporter) Export(record core.WeatherRecord) error {
	fmt.Printf("[%s] Temp: %.2f°C | Hum: %.2f%%\n",
		e.Prefix, record.Temperature, record.Humidity)
	return nil
}

func main() {
	exporter := &StdoutExporter{}
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "GOWX_PLUGIN",
		MagicCookieValue: "hello",
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"exporter": &core.ExporterPlugin{Impl: exporter},
		},
	})
}
