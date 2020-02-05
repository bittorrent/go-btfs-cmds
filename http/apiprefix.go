package http

import (
	"net/http"
	"strings"
)

type prefixHandler struct {
	prefix   string
	redirect []string
	next     http.Handler
}

func newPrefixHandler(prefix string, redirect []string, next http.Handler) http.Handler {
	return prefixHandler{prefix, redirect, next}
}

func (h prefixHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validPrefixes := append(h.redirect, h.prefix)
	var prefix string
	for _, vp := range validPrefixes {
		if strings.HasPrefix(r.URL.Path, vp) {
			prefix = vp
			break
		}
	}
	if prefix == "" {
		http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
		return
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
	h.next.ServeHTTP(w, r)
}
