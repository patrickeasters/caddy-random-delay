package random_delay

import (
	"net/http"
	"time"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// ServeHTTP implements caddyhttp.MiddlewareHandler
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	if m.randomSource.Float64() < m.PercentDelayed {
		timer := time.NewTimer(m.DelayDuration)
		select {
		case <-r.Context().Done():
			if !timer.Stop() {
				<-timer.C
			}
		case <-timer.C:
		}
	}

	return next.ServeHTTP(w, r)
}
