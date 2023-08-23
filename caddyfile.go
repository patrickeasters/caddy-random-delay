package random_delay

import (
	"fmt"
	"strconv"
	"time"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("random_delay", parseCaddyfileHandler)
}

// parseCaddyfileHandler unmarshals tokens from h into a new middleware handler value.
func parseCaddyfileHandler(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler. Syntax:
//
//	random_delay <name> [<option>]
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	var percent, duration string
	for d.Next() {
		if !d.Args(&percent, &duration) {
			// not enough args
			return d.ArgErr()
		}
		if d.NextArg() {
			// too many args
			return d.ArgErr()
		}
	}

	parsedPercent, err := strconv.ParseFloat(percent, 64)
	if err != nil {
		return fmt.Errorf("failed to parse percentage to be delayed: %w", err)
	}
	m.PercentDelayed = parsedPercent
	parsedDuration, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("failed to parse delay duration: %w", err)
	}
	m.DelayDuration = parsedDuration

	return nil
}
