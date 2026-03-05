from fastapi import FastAPI
from pydantic import BaseModel
import uvicorn
import requests
import datetime

app = FastAPI(title="Multi-Cloud Anomaly Detector")

class CloudMetric(BaseModel):
    provider: str
    region: str
    cpu_usage: float
    latency_ms: int

@app.post("/analyze")
async def analyze_telemetry(metric: CloudMetric):
    print(f"[{datetime.datetime.now()}] Analyzing {metric.provider} {metric.region}...")
    
    # Simulate the Predictive ML Logic
    # If CPU is over 90% and latency is high, predict an imminent crash
    if metric.cpu_usage > 90.0 and metric.latency_ms > 200:
        print(f"CRITICAL ANOMALY DETECTED IN {metric.provider}! Triggering failover...")
        
        failover_payload = {
            "failing_provider": metric.provider,
            "failing_region": metric.region,
            "target_provider": "GCP", # Hardcoded failover target for simulation
            "reason": "Predictive model flagged imminent resource exhaustion."
        }
        
        # Call the Traffic Router service to execute the DNS shift
        try:
            requests.post("http://traffic-router:8002/failover", json=failover_payload)
        except Exception as e:
            print(f"Failed to reach Traffic Router: {e}")
            
        return {"status": "anomaly_detected", "action": "failover_triggered"}
        
    return {"status": "healthy"}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8001)