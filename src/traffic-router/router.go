package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FailoverRequest struct {
	FailingProvider string `json:"failing_provider"`
	FailingRegion   string `json:"failing_region"`
	TargetProvider  string `json:"target_provider"`
	Reason          string `json:"reason"`
}

func failoverHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var req FailoverRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("\n==================================================")
	fmt.Println("🚨 PROACTIVE FAILOVER INITIATED 🚨")
	fmt.Printf("Evacuating: %s (%s)\n", req.FailingProvider, req.FailingRegion)
	fmt.Printf("Routing Traffic To: %s\n", req.TargetProvider)
	fmt.Printf("Reason: %s\n", req.Reason)
	fmt.Println("Status: DNS records updated successfully. Traffic shifted.")
	fmt.Println("==================================================\n")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Traffic successfully routed"}`))
}

func main() {
	http.HandleFunc("/failover", failoverHandler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	fmt.Println("Traffic Router listening on port 8002...")
	log.Fatal(http.ListenAndServe(":8002", nil))
}