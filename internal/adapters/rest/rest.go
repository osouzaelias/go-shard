package rest

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (a Adapter) Proxy(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Header.Get("X-Tenant-ID")
	shardID := r.Header.Get("X-Shard-ID")
	customerID := r.Header.Get("X-Customer-ID")

	if tenantID == "" || shardID == "" || customerID == "" {
		http.Error(w, "Required headers not found", http.StatusBadRequest)
		return
	}

	cell, err := a.api.GetCell(r.Context(), tenantID, shardID, customerID)
	if err != nil {
		http.Error(w, "Unable to identify cell", http.StatusBadRequest)
		return
	}

	address, err := url.Parse(cell.Address)
	if err != nil {
		http.Error(w, "Error parsing cell URL", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(address)
	proxy.ServeHTTP(w, r)
}

func (a Adapter) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"UP"}`))
}
