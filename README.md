# Jenkins CI/CD Demo

A production-ready Go web application demonstrating modern CI/CD practices with Jenkins, Docker, and comprehensive testing.

## 🏗️ Project Structure

```
jenkins-cicd-demo/
├── cmd/server/              # Application entry point
│   └── main.go             # HTTP server with routing
├── internal/handlers/       # HTTP handlers (Go best practices)
│   ├── handlers.go         # Business logic and endpoints
│   └── handlers_test.go    # Comprehensive unit tests
├── deployments/            # Deployment configurations
│   ├── docker-compose.yml  # Container orchestration
│   └── k8s-deployment.yml  # Kubernetes manifests (Deployment, Service, RBAC, HPA, NetworkPolicy)
├── scripts/               # Automation scripts
│   └── deploy.sh          # Deployment automation
├── Dockerfile             # Multi-stage container build
├── Jenkinsfile           # Declarative CI/CD pipeline
├── go.mod               # Go module definition
├── go.sum              # Dependency checksums
└── main               # Compiled binary (gitignored)
```

## ✨ Features

### 🌐 **HTTP Endpoints**
- **`GET /`** - Simple hello message for compatibility
- **`GET /api/hello`** - JSON response with timestamp and version
- **`GET /health`** - Health check endpoint for monitoring
- **`GET /dashboard`** - Interactive CI/CD pipeline visualization
- **`GET /metrics`** - Pipeline metrics and deployment information

### 🧪 **Quality Assurance**
- **5 comprehensive unit tests** covering all handlers
- **HTTP test utilities** for reliable API testing
- **Automated testing** in CI/CD pipeline
- **Health checks** with proper status codes

### 🐳 **Containerization**
- **Multi-stage Dockerfile** with Go build and Alpine runtime
- **Health check integration** for container orchestration
- **Optimized image size** using Alpine Linux
- **Production-ready** container configuration

### 🚀 **CI/CD Pipeline**
- **9-stage Jenkins pipeline** from code to deployment
- **Automated testing** and quality gates
- **Docker image building** and registry push
- **Deployment automation** with health verification
- **Pipeline as Code** using declarative Jenkinsfile

## 🎯 Educational Value

This project demonstrates:
- **Professional Go project structure** (`cmd/`, `internal/`)
- **Test-driven development** practices
- **Infrastructure as Code** principles
- **Automated deployment** strategies
- **Monitoring and observability** basics

## 🚀 Quick Start

### Prerequisites
- [Go](https://golang.org/) 1.21+
- [Docker](https://www.docker.com/)
- [Jenkins](https://www.jenkins.io/) with Docker support
- Docker Hub account for image registry

### Local Development

1. **Clone and setup**
   ```bash
   git clone https://github.com/ZEZE1020/jenkins-cicd-demo.git
   cd jenkins-cicd-demo
   go mod download
   ```

2. **Run locally**
   ```bash
   go run cmd/server/main.go
   # Visit http://localhost:3000/dashboard
   ```

3. **Run tests**
   ```bash
   go test -v ./...
   # All 5 tests should pass
   ```

4. **Build application**
   ```bash
   go build -o main ./cmd/server
   ./main
   ```

### Docker Deployment

1. **Build image**
   ```bash
   docker build -t jenkins-cicd-demo .
   ```

2. **Run container**
   ```bash
   docker run -p 3000:3000 jenkins-cicd-demo
   ```

3. **Using Docker Compose**
   ```bash
   cd deployments
   docker-compose up -d
   ```

### Kubernetes Deployment

Apply all Kubernetes resources from the bundled manifest:

```bash
kubectl apply -f deployments/k8s-deployment.yml
```

This manifest includes:
- Deployment with resource requests/limits and liveness/readiness probes
- ServiceAccount + RBAC (Role/RoleBinding)
- Service and NetworkPolicy
- HorizontalPodAutoscaler (HPA) for CPU and memory-based scaling

## 🔧 Jenkins Pipeline Setup

### Jenkins Configuration
1. **Install required plugins:**
   - Docker Pipeline
   - Docker Commons
   - Pipeline: Stage View

2. **Configure credentials:**
   - Add Docker Hub credentials (ID: `dockerhub-credentials`)
   - Ensure Jenkins can access Docker daemon

3. **Create pipeline job:**
   - New Item → Pipeline
   - Pipeline script from SCM
   - Repository: `https://github.com/ZEZE1020/jenkins-cicd-demo.git`
   - Script Path: `Jenkinsfile`

### Pipeline Stages
The Jenkins pipeline executes these stages automatically:

1. **Checkout** - Clone repository from GitHub
2. **Master Agent Example (Demo Only)** - Demonstrates a stage pinned to Jenkins `master`/`built-in` agent (educational only; runs only when `RUN_MASTER_AGENT_DEMO=true`; use dedicated worker agents in production)
3. **Install Dependencies** - Download Go modules  
4. **Quality Checks (Parallel)** - Run **Lint** and **Test** stages in parallel
5. **Build** - Compile Go binary and build Docker image
6. **Scan Image (Trivy)** - Run container vulnerability scan before push
7. **Push Image** - Upload verified Docker image to Docker Hub registry
8. **Deploy** - Start application with Docker Compose
9. **Verify Deployment** - Health checks and endpoint testing

**Agent roles in this example:**
- **Master/Built-in agent** - Runs only the demo stage when `RUN_MASTER_AGENT_DEMO=true`, to show how a stage can be pinned to the controller node.
- **Worker agents** - Should run regular CI/CD workload (dependency install, lint/test, build, scan, push, deploy, and verification) in normal usage.

**Total pipeline time:** ~5 minutes

## 📊 Monitoring & Health Checks

### Application Endpoints
- **Health Check:** `http://localhost:3000/health`
  ```json
  {
    "status": "healthy",
    "timestamp": "2025-09-18T10:30:00Z",
    "version": "v2.1.0"
  }
  ```

- **Metrics:** `http://localhost:3000/metrics`
  ```json
  {
    "build_number": "123",
    "deployment_time": "2025-09-18T10:30:00Z",
    "go_version": "1.21",
    "pipeline_stages": ["Checkout", "Tests", "Build", "Deploy"],
    "status": "running"
  }
  ```

### Container Health Checks
Docker containers include automated health monitoring:
```bash
docker ps  # Shows health status
docker logs <container-name>  # View application logs
```

## 🎓 Learning Objectives

This project teaches:
- **Modern Go development** patterns and project structure
- **Test-driven development** with comprehensive coverage
- **Containerization** best practices with Docker
- **CI/CD automation** using Jenkins pipelines
- **Infrastructure as Code** principles
- **Monitoring and observability** fundamentals

Perfect for students, developers, and teams learning DevOps practices!

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🔗 Resources

- **Jenkins Documentation:** [jenkins.io/doc](https://jenkins.io/doc)
- **Go Best Practices:** [golang.org/doc/effective_go](https://golang.org/doc/effective_go)
- **Docker Guide:** [docs.docker.com](https://docs.docker.com)
- **CI/CD Patterns:** [martinfowler.com/articles/continuousIntegration.html](https://martinfowler.com/articles/continuousIntegration.html)

---

**Ready to experience CI/CD in action? Clone, build, and deploy!** 🚀
