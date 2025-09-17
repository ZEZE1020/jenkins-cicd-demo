# Jenkins CI/CD Demo

This repository demonstrates a simple Go HTTP server with a CI/CD pipeline using Jenkins and Docker.

## Features
- Go HTTP server (`main.go`) responding with "Hello, NDC!"
- Unit tests using Go's built-in testing (`main_test.go`)
- Dockerfile for containerization
- Declarative Jenkinsfile for CI/CD pipeline

## Prerequisites
- [Go](https://golang.org/) 1.21 or later
- [Docker](https://www.docker.com/)
- [Jenkins](https://www.jenkins.io/) with Docker plugin
- Docker Hub account (replace credentials in Jenkinsfile)

## Setup Instructions

### 1. Clone the repository
```sh
git clone https://github.com/yourusername/jenkins-cicd-demo.git
cd jenkins-cicd-demo
```

### 2. Install dependencies
```sh
go mod download
```

### 3. Run the app locally
```sh
go run main.go
```
Visit [http://localhost:3000](http://localhost:3000)

### 4. Run tests
```sh
go test -v ./...
```

### 5. Build the application
```sh
go build -o main .
```

### 6. Build Docker image
```sh
docker build -t yourdockerhubusername/jenkins-cicd-demo .
```

### 7. Run Docker container
```sh
docker run -p 3000:3000 yourdockerhubusername/jenkins-cicd-demo
```

### 8. Jenkins Pipeline Setup
- Create a new Jenkins Pipeline project.
- Add your Docker Hub credentials in Jenkins (ID: `dockerhub-credentials`).
- Use the provided `Jenkinsfile` for the pipeline script.
- The pipeline stages:
  1. **Checkout**: Clones the repo
  2. **Install Dependencies**: Runs `go mod download`
  3. **Run Tests**: Runs `go test -v ./...`
  4. **Build Application**: Compiles the Go binary
  5. **Build Docker Image**: Builds the Docker image
  6. **Push Docker Image**: Pushes to Docker Hub (replace with your username)
  7. **Deploy**: Echoes a message (replace with real deployment logic)

## Notes
- Replace all placeholder values (e.g., Docker Hub username, credentials) with your actual details.
- For real deployment, update the `Deploy` stage in the Jenkinsfile.

---

For questions, open an issue or contact the maintainer.
