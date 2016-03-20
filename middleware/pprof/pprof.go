package pprof

import (
	"net/http"
	pp "net/http/pprof"

	"github.com/mholt/caddy/middleware"
)

// BasePath is the base path to match for all pprof requests.
const BasePath = "/debug/pprof"

// Handler is a simple struct whose ServeHTTP will delegate pprof
// endpoints to their equivalent net/http/pprof handlers.
type Handler struct {
	Next middleware.Handler
	Mux  *http.ServeMux
}

// ServeHTTP handles requests to BasePath with pprof, or passes
// all other requests up the chain.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	if middleware.Path(r.URL.Path).Matches(BasePath) {
		h.Mux.ServeHTTP(w, r)
		return 0, nil
	}
	return h.Next.ServeHTTP(w, r)
}

// NewMux returns a new http.ServeMux that routes pprof requests.
// It pretty much copies what the std lib pprof does on init:
// https://golang.org/src/net/http/pprof/pprof.go#L67
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(BasePath+"/", pp.Index)
	mux.HandleFunc(BasePath+"/cmdline", pp.Cmdline)
	mux.HandleFunc(BasePath+"/profile", pp.Profile)
	mux.HandleFunc(BasePath+"/symbol", pp.Symbol)
	mux.HandleFunc(BasePath+"/trace", pp.Trace)
	return mux
}