package fnversioning

import (
  "net/http"
  
	"github.com/fnproject/fn/fnext"
	_ "github.com/fnproject/fn/api/models"
	"github.com/fnproject/fn/api/server"
)

type VersioningMiddleware struct {
}

func (m *VersioningMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if h, ok := r.Header["X-Function-Versioning"]; ok == true {
      version := h[0]
      r.URL.Path = r.URL.Path + "/" + version
    }
    next.ServeHTTP(w, r)
  }
}

func init() {
    server.RegisterExtension(&fnext.Extension{
        Name:  "github.com/chanwit/fnversioning",
        Setup: setup,
    })
}

func setup(s *fnext.ExtServer) error {
    s.AddAPIMiddleware(&VersioningMiddleware{})
    return nil
}