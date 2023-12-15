package rest

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (a Adapter) Proxy(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Header.Get("X-Tenant-ID")
	customerID := r.Header.Get("X-Customer-ID")

	if tenantID == "" || customerID == "" {
		http.Error(w, "Required headers not found", http.StatusBadRequest)
		return
	}

	shard, err := a.api.GetShard(r.Context(), tenantID, customerID)
	if err != nil {
		http.Error(w, "Unable to identify shard", http.StatusBadRequest)
		return
	}

	address, err := url.Parse(shard.Address)
	if err != nil {
		http.Error(w, "Error parsing shard URL", http.StatusInternalServerError)
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
