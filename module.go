package random_delay

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Middleware{})
}

// Middleware implements an HTTP handler that delays approximately PercentDelayed requests
// by DelayDuration
type Middleware struct {
	DelayDuration  time.Duration `json:"delay_duration"`
	PercentDelayed float64       `json:"percent_delayed"`

	randomSource *rand.Rand
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.random_delay",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
	src := rand.NewSource(time.Now().UnixNano())
	m.randomSource = rand.New(src)
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	if m.PercentDelayed < 0 || m.PercentDelayed > 1 {
		return fmt.Errorf("PercentDelayed must be between 0 and 100")
	}
	return nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
