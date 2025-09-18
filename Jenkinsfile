// Declarative Jenkinsfile for the CI/CD pipeline
pipeline {
  agent any
  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub-credentials') // Placeholder
    DOCKER_IMAGE = "ogembog/jenkins-cicd-demo:${env.BUILD_NUMBER}"
  }
  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    stage('Install Dependencies') {
      steps {
        sh 'go mod download'
      }
    }
    stage('Run Tests') {
      steps {
        sh 'go test -v ./...'
      }
    }
    stage('Build Application') {
      steps {
        sh 'go build -o main ./cmd/server'
      }
    }
    stage('Build Docker Image') {
      steps {
        sh 'docker build -t $DOCKER_IMAGE .'
      }
    }
    stage('Push Docker Image') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-credentials', usernameVariable: 'DOCKERHUB_USER', passwordVariable: 'DOCKERHUB_PASS')]) {
          sh 'echo $DOCKERHUB_PASS | docker login -u $DOCKERHUB_USER --password-stdin'
          sh 'docker push $DOCKER_IMAGE'
        }
      }
    }
    stage('Deploy') {
      steps {
        sh 'chmod +x scripts/deploy.sh'
        sh './scripts/deploy.sh'
        echo 'Application deployed successfully with Docker Compose!'
      }
    }
    stage('Verification') {
      steps {
        echo 'Running post-deployment verification...'
        sh 'sleep 10'  // Wait for container to start
        sh 'curl -f http://localhost:3000/health || exit 1'
        sh 'curl -f http://localhost:3000/dashboard || exit 1'
        echo 'Application is healthy and responding!'
        echo 'Dashboard: http://your-server:3000/dashboard'
        echo 'Health Check: http://your-server:3000/health'
        echo 'Metrics: http://your-server:3000/metrics'
      }
    }
  }
}
