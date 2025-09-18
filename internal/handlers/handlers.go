package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

// HelloResponse represents the hello endpoint response
type HelloResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

const appVersion = "v2.1.0"

// HelloHandler handles the main greeting endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message:   "Hello, NDC! Pipeline Test v2.1.0 - Refactored Structure",
		Timestamp: time.Now(),
		Version:   appVersion,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   appVersion,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SimpleHelloHandler provides a simple text response for compatibility
func SimpleHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, NDC! Pipeline Test v2.1.0 - Refactored Structure")
}

// Dashboard serves an HTML page showing the CI/CD pipeline status
func Dashboard(w http.ResponseWriter, r *http.Request) {
	buildNumber := os.Getenv("BUILD_NUMBER")
	if buildNumber == "" {
		buildNumber = "local-dev"
	}
	
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Jenkins CI/CD Demo</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background: #f5f5f5; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .header { text-align: center; color: #2c3e50; margin-bottom: 30px; }
        .pipeline-status { display: flex; justify-content: space-between; margin: 20px 0; }
        .stage { background: #27ae60; color: white; padding: 10px 15px; border-radius: 5px; text-align: center; flex: 1; margin: 0 5px; }
        .stage.active { background: #3498db; animation: pulse 2s infinite; }
        .info-box { background: #ecf0f1; padding: 20px; border-radius: 5px; margin: 15px 0; }
        .success { color: #27ae60; font-weight: bold; }
        .timestamp { color: #7f8c8d; font-size: 0.9em; }
        @keyframes pulse { 0% { opacity: 1; } 50% { opacity: 0.7; } 100% { opacity: 1; } }
        .endpoints { margin-top: 30px; }
        .endpoint { background: #3498db; color: white; padding: 8px 12px; border-radius: 3px; text-decoration: none; margin: 5px; display: inline-block; }
        .endpoint:hover { background: #2980b9; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Jenkins CI/CD Demo</h1>
            <h2>Go Application Successfully Deployed!</h2>
        </div>
        
        <div class="pipeline-status">
            <div class="stage"> Checkout</div>
            <div class="stage"> Dependencies</div>
            <div class="stage"> Tests</div>
            <div class="stage"> Build</div>
            <div class="stage"> Docker</div>
            <div class="stage active">Running</div>
        </div>
        
        <div class="info-box">
            <h3>Deployment Information</h3>
            <p><strong>Build Number:</strong> {{.BuildNumber}}</p>
            <p><strong>Deployment Time:</strong> <span class="timestamp">{{.Timestamp}}</span></p>
            <p><strong>Version:</strong> v2.1.0</p>
            <p><strong>Status:</strong> <span class="success">Healthy</span></p>
        </div>
        
        <div class="info-box">
            <h3>Pipeline Stages Completed</h3>
            <ul>
                <li>Source code checked out from Git</li>
                <li>Go dependencies downloaded</li>
                <li>Unit tests passed</li>
                <li>Application binary built</li>
                <li>Docker image created and pushed</li>
                <li>Application deployed with Docker Compose</li>
            </ul>
        </div>
        
        <div class="endpoints">
            <h3>Available Endpoints</h3>
            <a href="/api/hello" class="endpoint">API Hello</a>
            <a href="/health" class="endpoint">Health Check</a>
            <a href="/metrics" class="endpoint">Metrics</a>
        </div>
    </div>
</body>
</html>`
	
	t, _ := template.New("dashboard").Parse(tmpl)
	data := struct {
		BuildNumber string
		Timestamp   string
	}{
		BuildNumber: buildNumber,
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
	}
	
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, data)
}

// Metrics provides pipeline metrics for students
func Metrics(w http.ResponseWriter, r *http.Request) {
	metrics := map[string]interface{}{
		"build_number":    os.Getenv("BUILD_NUMBER"),
		"deployment_time": time.Now(),
		"go_version":      "1.21",
		"docker_image":    fmt.Sprintf("ogembog/jenkins-cicd-demo:%s", os.Getenv("BUILD_NUMBER")),
		"pipeline_stages": []string{
			"Checkout", "Dependencies", "Tests", "Build", "Docker", "Deploy",
		},
		"status": "running",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}