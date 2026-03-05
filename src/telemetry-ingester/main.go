package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CloudMetric struct {
	Provider string  `json:"provider"`
	Region   string  `json:"region"`
	CPUUsage float64 `json:"cpu_usage"`
	Latency  int     `json:"latency_ms"`
}

func main() {
	fmt.Println("Starting Multi-Cloud Telemetry Ingester...")

	// Simulate a failing AWS region for testing purposes
	metrics := []CloudMetric{
		{"AWS", "us-east-1", 95.5, 300}, // High CPU and High Latency (Failing)
		{"GCP", "us-central1", 45.0, 40},
		{"Azure", "eastus", 50.2, 45},
	}

	for {
		for _, metric := range metrics {
			jsonData, _ := json.Marshal(metric)
			
			// Send the metric to the Python Anomaly Detector
			resp, err := http.Post("http://anomaly-detector:8001/analyze", "application/json", bytes.NewBuffer(jsonData))
			
			if err != nil {
				fmt.Println("Error sending telemetry:", err)
			} else {
				fmt.Printf("Sent %s telemetry to predictor. Status: %s\n", metric.Provider, resp.Status)
				resp.Body.Close()
			}
		}
		// Pause for 10 seconds before the next scrape
		time.Sleep(10 * time.Second)
	}
}